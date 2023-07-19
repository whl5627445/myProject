package service

import (
	"encoding/base64"
	"log"
	"os"
	"reflect"
	"regexp"
	"strings"
	"yssim-go/config"
	"yssim-go/library/omc"
	"yssim-go/library/stringOperation"
)

type graphicsData struct {
	data          []any
	modelName     string
	modelNameList []string
	permissions   string
}

var allModelCache = config.R

func GetGraphicsData(modelName, permissions string) []any {
	var g = graphicsData{}
	//g.data = [][]map[string]any{{}, {}}
	g.data = make([]any, 3)
	g.data[0] = make([]map[string]any, 0)
	g.data[1] = make([]map[string]any, 0)
	g.data[2] = make(map[string]any, 0)
	//ctx := context.Background()
	//g.permissions = permissions
	//if permissions == "sys" {
	//	msg, _ := allModelCache.HGet(ctx, config.USERNAME+"-yssim-modelGraphicsData", modelName).Bytes()
	//	msg := []byte{}
	//if len(msg) > 0 {
	//	err := sonic.Unmarshal(msg, &g.data)
	//	if err != nil {
	//		log.Println("GetGraphicsData 反序列化错误 err", err)
	//		return nil
	//	}
	//	return g.data
	//}
	//}
	g.modelName = modelName
	nameType := omc.OMC.GetClassRestriction(modelName)
	if nameType == "connector" || nameType == "expandable connector" {
		g.data = g.getConnectorModelDiagram(modelName)
	} else {
		g.modelNameList = g.getICList(modelName)
		g.getDiagramAnnotationData()
		g.getnthconnectionData()
		g.getData02()
	}
	//if permissions == "sys" {
	//	redisData, _ := sonic.Marshal(g.data)
	//	allModelCache.HSet(ctx, config.USERNAME+"-yssim-modelGraphicsData", modelName, redisData)
	//}
	return g.data
}

func GetComponentGraphicsData(modelName, componentName string) []any {
	var g = graphicsData{}
	//g.data = [][]map[string]any{{}, {}}
	g.data = make([]any, 3)
	g.data[0] = make([]map[string]any, 0)
	g.data[1] = make([]map[string]any, 0)
	g.data[2] = make(map[string]any, 0)
	g.modelName = modelName
	components, componentAnnotations := getElementsAndModelName([]string{modelName})
	var componentsData [][]any
	var componentAnnotationsData [][]any
	for i := 0; i < len(components); i++ {
		if components[i] != nil {
			if components[i][3] == componentName {
				nameType := omc.OMC.GetClassRestriction(components[i][2].(string))
				if nameType == "connector" || nameType == "expandable connector" {
					data := g.getConnectorComponentDiagram(components[i], componentAnnotations[i])
					return data
				}
				componentsData = append(componentsData, components[i])
				componentAnnotationsData = append(componentAnnotationsData, componentAnnotations[i])
				break
			}
		}
	}
	data2 := g.data02(componentsData, componentAnnotationsData, false, "", true)
	for i := 0; i < len(data2); i++ {
		if len(data2[i]["subShapes"].([]map[string]any)) == 0 {
			data2[i]["subShapes"] = append(data2[i]["subShapes"].([]map[string]any), rectangleDefault)
		}
	}
	g.data[1] = append(g.data[1].([]map[string]any), data2...)
	return g.data
}

func (g *graphicsData) getData02() {
	// nameList第一个名字是模型自身的名字，之后是继承的模型名字
	for i := len(g.modelNameList) - 1; i >= 0; i-- {
		var data2 []map[string]any
		mobility := false
		if i == len(g.modelNameList)-1 && g.permissions != "sys" {
			mobility = true
		}
		componentsData, componentAnnotationsData := getElementsAndModelName([]string{g.modelNameList[i]})
		data := g.data02(componentsData, componentAnnotationsData, false, "", mobility)
		data2 = append(data2, data...)
		g.data[1] = append(g.data[1].([]map[string]any), data2...)
	}
	for i := 0; i < len(g.data[1].([]map[string]any)); i++ {
		if len(g.data[1].([]map[string]any)[i]["subShapes"].([]map[string]any)) == 0 {
			g.data[1].([]map[string]any)[i]["subShapes"] = append(g.data[1].([]map[string]any)[i]["subShapes"].([]map[string]any), rectangleDefault)
		}
	}
}

func oneDimensionalProcessing(Data any) string {
	if reflect.TypeOf(Data).String() == "[]interface {}" {
		drawingData := Data.([]any)
		var data []string
		defer func() {
			if err := recover(); err != nil {
				data = []string{}
			}
		}()
		if len(drawingData) >= 1 && reflect.TypeOf(drawingData[0]).String() != "string" {
			return ""
		}
		for i := 0; i < len(drawingData); i++ {
			data = append(data, drawingData[i].(string))
		}

		return strings.Join(data, ",")
	}
	return ""
}

func twoDimensionalProcessing(drawingData []any) []string {
	var data []string
	dataType := reflect.TypeOf(drawingData[0].([]any)[0]).String()

	if dataType == "string" {
		for i := 0; i < len(drawingData); i++ {
			var piList []string
			for pi := 0; pi < len(drawingData[i].([]any)); pi++ {
				piList = append(piList, drawingData[i].([]any)[pi].(string))
			}
			data = append(data, strings.Join(piList, ","))
		}
		return data
	}
	if dataType == "[]interface {}" {
		for i := 0; i < len(drawingData); {
			for pi := 0; pi < len(drawingData[i].([]any)); pi++ {
				var piList []string
				for pii := 0; pii < len(drawingData[i].([]any)[pi].([]any)); pii++ {
					piList = append(piList, drawingData[i].([]any)[pi].([]any)[pii].(string))
				}
				data = append(data, strings.Join(piList, ","))
			}
			//break
		}
		return data
	}
	return data
}

func (g *graphicsData) getICList(name string) []string {
	dataList := GetICList(name)
	return dataList
}

func find(data []any, str string) bool {
	for i := 0; i < len(data); i++ {
		if reflect.TypeOf(data[i]).String() == "string" && data[i].(string) == str {
			return true
		}
	}
	return false
}

func (g *graphicsData) data01(cData []any, className, component, modelName string) []map[string]any {
	dataList := make([]map[string]any, 0, 1)

	for i := 0; i < len(cData); i += 2 {
		data := map[string]any{}
		drawingDataList := cData[i+1].([]any)
		if drawingDataList[0].(string) != "true" {
			continue
		}
		DynamicSelect := find(drawingDataList, "DynamicSelect")
		if DynamicSelect {
			var drawingDataListFilter []any
			for _, i2 := range drawingDataList {
				if i2 == "DynamicSelect" {
					continue
				}
				drawingDataListFilter = append(drawingDataListFilter, i2)
			}
			drawingDataList = drawingDataListFilter
		}

		if len(drawingDataList) < 9 {
			dataType := cData[i]
			if dataType == "Bitmap" {
				data["type"] = dataType
				data["visible"] = drawingDataList[0]
				data["originalPoint"] = oneDimensionalProcessing(drawingDataList[1])
				data["rotation"] = drawingDataList[2]
				data["points"] = twoDimensionalProcessing(drawingDataList[3].([]any))
				dataImagePath := drawingDataList[4]
				if drawingDataList[5] == "" {
					imagePath := omc.OMC.UriToFilename(dataImagePath.(string))
					bytes, err := os.ReadFile(imagePath)
					if err != nil {
						log.Println("dataImagePath 错误：", dataImagePath)
						log.Println("err：", err)
					}
					data["imageBase64"] = base64.StdEncoding.EncodeToString(bytes)
				} else {
					data["imageBase64"] = drawingDataList[5]
				}
				dataList = append(dataList, data)
				continue
			}
			continue
		}
		dataType := cData[i]
		data["type"] = dataType
		data["type"] = dataType
		data["visible"] = drawingDataList[0]
		data["originalPoint"] = oneDimensionalProcessing(drawingDataList[1])
		data["rotation"] = drawingDataList[2]
		data["color"] = oneDimensionalProcessing(drawingDataList[3])
		data["fillColor"] = oneDimensionalProcessing(drawingDataList[4])
		data["linePattern"] = drawingDataList[5]
		data["fillPattern"] = drawingDataList[6]
		data["lineThickness"] = drawingDataList[7]
		switch dataType {
		case "Polygon":
			data["polygonPoints"] = twoDimensionalProcessing(drawingDataList[8].([]any))
			ppData := data["polygonPoints"].([]string)
			if len(ppData) < 4 {
				data["polygonPoints"] = append(ppData, ppData[0])
			}
			data["smooth"] = drawingDataList[9]
		case "Line":
			delete(data, "fillColor")
			delete(data, "fillPattern")
			data["points"] = twoDimensionalProcessing(drawingDataList[3].([]any))
			data["color"] = oneDimensionalProcessing(drawingDataList[4])
			data["lineThickness"] = drawingDataList[6]
			data["arrow"] = oneDimensionalProcessing(drawingDataList[7])
			data["arrowSize"] = drawingDataList[8]
			data["smooth"] = drawingDataList[9]
		case "Text":
			data["fillPattern"] = drawingDataList[6]
			data["extentsPoints"] = twoDimensionalProcessing(drawingDataList[8].([]any))
			typeOriginalTextString, ok := drawingDataList[9].([]any)
			if ok {
				data["originalTextString"] = typeOriginalTextString[0]
			} else {
				originalTextString := drawingDataList[9].(string)
				data["textType"] = "var"
				if strings.Contains(originalTextString, "%") {
					data["textType"] = "text"
				}
				textList := stringOperation.PluralSplit(originalTextString, []string{"/", ",", "\t", "\n", "\r", " "})
				for _, t := range textList {
					pSignIndex := strings.Index(t, "%")
					if pSignIndex != -1 {
						varName := t[pSignIndex+1:]
						varValue := ""
						Unit := ""
						if varName != "name" {
							varName = strings.TrimSuffix(varName, "%")
							if varName != "" {
								d := GetModelParameters(g.modelName, component, className, varName)
								Unit, varValue = d[0].(string), d[1].(string)
								if varValue == "" {
									varValue = varName
								}
							}

							if Unit != "" {
								Unit = " " + Unit
							}
							oldVarName := "%" + varName
							varValueUnit := varName + Unit

							varValueUnit = strings.Replace(varValueUnit, varName, varValue, 1)
							originalTextString = strings.Replace(originalTextString, oldVarName, varValueUnit, 1)
						}
					}
				}
				data["originalTextString"] = originalTextString
			}
			data["fontSize"] = drawingDataList[10]
			data["textColor"] = oneDimensionalProcessing(drawingDataList[11])
			data["fontName"] = drawingDataList[12]
			data["textStyles"] = drawingDataList[13]
			data["horizontalAlignment"] = drawingDataList[14]
		case "Rectangle":
			data["fillPattern"] = drawingDataList[6]
			data["borderPattern"] = drawingDataList[8]
			data["extentsPoints"] = twoDimensionalProcessing(drawingDataList[9].([]any))
			data["radius"] = drawingDataList[10]
		case "Ellipse":
			data["fillPattern"] = drawingDataList[6]
			data["extentsPoints"] = twoDimensionalProcessing(drawingDataList[8].([]any))
			data["startAngle"] = drawingDataList[9]
			data["endAngle"] = drawingDataList[10]
		}
		dataList = append(dataList, data)
	}
	return dataList
}

func (g *graphicsData) data02(cData [][]any, caData [][]any, isIcon bool, parent string, mobility bool) []map[string]any {
	dataList := make([]map[string]any, 0, 1)
	var cDataFilter [][]any
	var caDataFilter [][]any

	for i := 0; i < len(cData); i++ {

		nameType := omc.OMC.GetClassRestriction(cData[i][2].(string))
		if isIcon && cData != nil && caData != nil {
			if nameType == "connector" || nameType == "expandable connector" {
				cData[i] = append(cData[i], nameType)
				cDataFilter = append(cDataFilter, cData[i])
				caDataFilter = append(caDataFilter, caData[i])
			}
		} else {
			cData[i] = append(cData[i], nameType)
			cDataFilter = append(cDataFilter, cData[i])
			caDataFilter = append(caDataFilter, caData[i])
		}
	}
	if cDataFilter == nil || caDataFilter == nil {
		return dataList
	}
	dataLen2 := func() int {
		if len(caDataFilter) > len(cDataFilter) {
			return len(cDataFilter)
		}
		return len(caDataFilter)
	}()
	for i := 0; i < dataLen2; i++ {
		modelName := cDataFilter[i][15].(string)

		classname := cDataFilter[i][2].(string)

		nameList := g.getICList(classname)
		DynamicSelect := find(caDataFilter[i], "DynamicSelect")
		if DynamicSelect {
			continue
		}
		placementIndex := func() int {
			for index, p := range caDataFilter[i] {
				if p == "Placement" {
					return index
				}
			}
			return -1
		}()
		if placementIndex != -1 {

			if len(caDataFilter[i]) < 1 {
				continue
			}
			caf := caDataFilter[i][placementIndex+1].([]any)
			if len(caf) < 7 || caf[0].(string) != "true" {
				// 出现错误会使数据不可用， 长度小于预期，弃用
				continue
			}

			data := map[string]any{"type": "Transformation"}

			data["graphType"] = cDataFilter[i][17]
			//data["ID"] = strconv.Itoa(i)
			data["classname"] = classname
			data["name"] = cDataFilter[i][3]
			data["original_name"] = cDataFilter[i][3]
			data["extend_name"] = modelName
			data["is_extend"] = func() bool {
				if g.modelName == modelName {
					return false
				}
				return true
			}()

			data["parent"] = parent
			data["connector_sizing"] = cDataFilter[i][16]
			data["visible"] = caf[0]
			data["mobility"] = mobility
			//data["initialScale"] = initialScale
			rotateAngle := func() string {
				if caf[14] != "" {
					return caf[14].(string)
				}
				if caf[7] != "-" {
					return caf[7].(string)
				} else {
					return "0"
				}
			}()
			if parent != "" && caf[10].(string) != "-" {
				extentX1, extentY1, extentX2, extentY2 := caf[10].(string), caf[11].(string), caf[12].(string), caf[13].(string)
				data["extent1Diagram"] = strings.Join([]string{extentX1, extentY1}, ",")
				data["extent2Diagram"] = strings.Join([]string{extentX2, extentY2}, ",")
				data["originDiagram"] = strings.Join([]string{caf[8].(string), caf[9].(string)}, ",")
			} else {
				data["extent1Diagram"] = strings.Join([]string{caf[3].(string), caf[4].(string)}, ",")
				data["extent2Diagram"] = strings.Join([]string{caf[5].(string), caf[6].(string)}, ",")
				data["originDiagram"] = strings.Join([]string{caf[1].(string), caf[2].(string)}, ",")
			}
			data["rotateAngle"] = rotateAngle
			data["rotation"] = rotateAngle
			data["output_type"] = func() string {
				t := cDataFilter[i][14].(string)
				return t
			}()
			coordinateSystem := getCoordinateSystem(classname, isIcon)
			data["coordinate_system"] = map[string]any{
				"extent1":               []any{coordinateSystem[0], coordinateSystem[1]},
				"extent2":               []any{coordinateSystem[2], coordinateSystem[3]},
				"preserve_aspect_ratio": coordinateSystem[4],
				"initialScale":          coordinateSystem[5],
			}
			componentsData, componentAnnotationsData := getElementsAndModelName(nameList)
			IconAnnotationData := getIconAndDiagramAnnotations(nameList, isIcon)
			data["inputOutputs"] = g.data02(componentsData, componentAnnotationsData, true, data["name"].(string), false)
			data["subShapes"] = g.data01(IconAnnotationData, classname, cDataFilter[i][3].(string), modelName)
			dataList = append(dataList, data)
		}
	}
	return dataList
}

func (g *graphicsData) getnthconnectionData() {
	ConnectionCount := omc.OMC.GetConnectionCountList(g.modelNameList)
	ncDistinct := map[string]bool{}
	for i := 0; i < len(ConnectionCount); i++ {
		for c := 0; c < ConnectionCount[i]; c++ {
			ncData := omc.OMC.GetNthConnection(g.modelNameList[i], c+1)
			if len(ncData) > 1 {
				if ncDistinct[ncData[0]+ncData[1]] {
					continue
				}
				ncDistinct[ncData[0]+ncData[1]] = true
			}
			ncaData := omc.OMC.GetNthConnectionAnnotation(g.modelNameList[i], c+1)
			d1Data := g.data01(ncaData, g.modelNameList[i], g.modelNameList[i], g.modelNameList[i])
			if len(ncData) != 0 && len(ncaData) != 0 && len(d1Data) != 0 {
				for _, daData := range d1Data {
					if daData["type"] == "Line" {
						if i == len(g.modelNameList)-1 && g.permissions != "sys" { // i==0的时候，表示目前遍历的是模型自身的组件，模型自身的组件可以移动，设在"mobility"为true
							daData["mobility"] = true
						} else {
							daData["mobility"] = false
						}
						daData["connectionfrom_original_name"] = ncData[0]
						daData["connectionto_original_name"] = ncData[1]
						//re1, _ := regexp.Compile("\\[[0-9a-zA-Z]+\\]$")
						//re2, _ := regexp.Compile("\\[[0-9a-zA-Z]+\\].")
						re1, _ := regexp.Compile("\\[[0-9a-zA-Z,，]+\\]$")
						re2, _ := regexp.Compile("\\[[0-9a-zA-Z,，]+\\].")
						connectionfrom := re1.ReplaceAll([]byte(ncData[0]), []byte(""))
						connectionto := re1.ReplaceAll([]byte(ncData[1]), []byte(""))
						daData["connectionfrom"] = string(re2.ReplaceAll(connectionfrom, []byte(".")))
						daData["connectionto"] = string(re2.ReplaceAll(connectionto, []byte(".")))
					}
					g.data[0] = append(g.data[0].([]map[string]any), daData)
				}
			}
		}
	}
}

func getElementsAndModelName(nameList []string) ([][]any, [][]any) {
	var ComponentData [][]any
	var AnnotationData [][]any

	for _, name := range nameList {
		connectorSizingList := []string{}
		componentsData := omc.OMC.GetElements(name)
		annotationsData := omc.OMC.GetElementAnnotations(name)
		componentsList := [][]any{}
		// {Dialog("General","Parameters",true,false,false,-,-,-,-,"",true)}
		for index, c := range componentsData {

			component := c.([]any)
			component = append(component, name)
			componentsList = append(componentsList, component)
			annotation := annotationsData[index].([]any)

			connectorSizing := getDialogConnectorSizing(annotation)
			if connectorSizing {
				connectorSizingList = append(connectorSizingList, component[3].(string))
			}
			AnnotationData = append(AnnotationData, annotation)
		}
		for _, c := range componentsList {
			connectorSizingName := ""
			for _, s := range connectorSizingList {
				if c[14] == "["+s+"]" {
					connectorSizingName = s
				}
			}
			c = append(c, connectorSizingName)
			ComponentData = append(ComponentData, c)
		}
	}

	return ComponentData, AnnotationData
}

func getDialogConnectorSizing(annotation []any) bool {
	for n := 0; n < len(annotation); n++ {
		if annotation[n] == "Dialog" {
			tabIndex := n + 1
			dListTab := annotation[tabIndex].([]any)
			if tabIndex > 0 && len(dListTab) > 3 {
				//if len(annotation) <= 1 || dListTab[len(dListTab)-1] == "true" {
				if len(annotation) <= 1 || dListTab[2] == "false" {
					continue
				}
				if dListTab[10] == "true" {
					return true
				}
			}
		}
	}
	return false
}

func getIconAndDiagramAnnotations(nameList []string, isIcon bool) []any {
	res := omc.OMC.GetIconAndDiagramAnnotations(nameList, isIcon)
	return res
}

func getCoordinateSystem(name string, isIcon bool) []any {
	res := omc.OMC.GetCoordinateSystem(name, isIcon)
	return res
}

func (g *graphicsData) getDiagramAnnotationData() {
	for index, name := range g.modelNameList {
		modelNameDiagramAnnotationData := omc.OMC.GetDiagramAnnotation(name)
		if len(modelNameDiagramAnnotationData) > 8 && modelNameDiagramAnnotationData[len(modelNameDiagramAnnotationData)-1] != "" {
			dData := modelNameDiagramAnnotationData[len(modelNameDiagramAnnotationData)-1]
			coordinateSystem := modelNameDiagramAnnotationData[:8]
			g.data[2] = map[string]any{
				"extent1":               []any{coordinateSystem[0], coordinateSystem[1]},
				"extent2":               []any{coordinateSystem[2], coordinateSystem[3]},
				"preserve_aspect_ratio": coordinateSystem[4],
			}
			g.data = append(g.data)
			data1 := g.data01(dData.([]any), g.modelName, g.modelName, g.modelName)
			for _, d := range data1 {
				if g.permissions != "sys" && index == len(g.modelNameList)-1 {
					d["mobility"] = true
				} else {
					d["mobility"] = false
				}
				d["diagram"] = true
				g.data[0] = append(g.data[0].([]map[string]any), d)
			}
		}
	}
}

func (g *graphicsData) getConnectorComponentDiagram(components, componentAnnotationsData []any) []any {
	className := components[2].(string)
	componentName := components[3].(string)
	if componentAnnotationsData[0] == "Placement" {
		data := make(map[string]any, 0)
		interfaceGraphicsData := getIconAndDiagramAnnotations([]string{className}, false)
		coordinateSystem := interfaceGraphicsData[:8]
		data["coordinate_system"] = map[string]any{
			"extent1":               []any{coordinateSystem[0], coordinateSystem[1]},
			"extent2":               []any{coordinateSystem[2], coordinateSystem[3]},
			"preserve_aspect_ratio": coordinateSystem[4],
		}
		caf := componentAnnotationsData[1].([]any)
		data["ID"] = "0"
		data["classname"] = className
		data["extent1Diagram"] = strings.Replace(caf[3].(string)+","+caf[4].(string), "-,-", "-100.0,-100.0", 1)
		data["extent2Diagram"] = strings.Replace(caf[5].(string)+","+caf[6].(string), "-,-", "100.0,100.0", 1)
		data["graphType"] = "connector"
		data["mobility"] = true
		data["name"] = componentName
		data["originDiagram"] = strings.Join([]string{caf[1].(string), caf[2].(string)}, ",")
		data["original_name"] = componentName
		data["output_type"] = "[]"
		data["parent"] = ""
		data["rotateAngle"] = "0.0"
		data["rotation"] = "0.0"
		data["type"] = "Transformation"
		data["visible"] = "true"
		data["inputOutputs"] = make([]map[string]any, 0, 1)
		data["subShapes"] = g.data01(interfaceGraphicsData, className, className, className)
		g.data[1] = append(g.data[1].([]map[string]any), data)
		return g.data
	}
	return nil
}

func (g *graphicsData) getConnectorModelDiagram(modelName string) []any {
	interfaceDiagramAnnotationData := omc.OMC.GetDiagramAnnotation(modelName)
	if len(interfaceDiagramAnnotationData) > 8 {
		interfaceGraphicsData := interfaceDiagramAnnotationData[8].([]any)
		data := make(map[string]any, 0)
		coordinateSystem := interfaceDiagramAnnotationData[:8]
		data["coordinate_system"] = map[string]any{
			"extent1":               []any{coordinateSystem[0], coordinateSystem[1]},
			"extent2":               []any{coordinateSystem[2], coordinateSystem[3]},
			"preserve_aspect_ratio": coordinateSystem[4],
		}
		data["ID"] = "0"
		data["classname"] = modelName
		data["extent1Diagram"] = strings.Replace(interfaceDiagramAnnotationData[0].(string)+","+interfaceDiagramAnnotationData[1].(string), "-,-", "-100.0,-100.0", 1)
		data["extent2Diagram"] = strings.Replace(interfaceDiagramAnnotationData[2].(string)+","+interfaceDiagramAnnotationData[3].(string), "-,-", "100.0,100.0", 1)
		data["graphType"] = "model"
		data["mobility"] = false
		data["name"] = ""
		data["originDiagram"] = "0.0,0.0"
		data["original_name"] = ""
		data["output_type"] = "[]"
		data["parent"] = ""
		data["rotateAngle"] = "0.0"
		data["rotation"] = "0.0"
		data["type"] = "Transformation"
		data["visible"] = "true"
		data["inputOutputs"] = make([]map[string]any, 0, 1)
		data["subShapes"] = make([]map[string]any, 0, 1)
		data1 := g.data01(interfaceGraphicsData, modelName, modelName, modelName)
		data["subShapes"] = data1
		g.data[1] = append(g.data[1].([]map[string]any), data)
	}
	return g.data
}

var rectangleDefault = map[string]any{"borderPattern": "BorderPattern.None", "color": "0,0,127", "extentsPoints": []string{"-100.0,-100.0", "100.0,100.0"},
	"fillColor":     "255,255,255",
	"fillPattern":   "FillPattern.Solid",
	"linePattern":   "LinePattern.Solid",
	"lineThickness": "0.25",
	"originalPoint": "0.0,0.0",
	"radius":        "0.0",
	"rotation":      "0.0",
	"type":          "Rectangle",
	"visible":       "true",
}
