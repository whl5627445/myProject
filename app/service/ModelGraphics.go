package service

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"yssim-go/library/omc"
)

type graphicsData struct {
	data [][]map[string]interface{}
}

func GetGraphicsData(modelName string) [][]map[string]interface{} {
	var g = graphicsData{}
	g.data = [][]map[string]interface{}{{}, {}}
	nameType := omc.OMC.GetClassRestriction(modelName)
	if nameType == "connector" || nameType == "expandable connector" {
		interfaceDiagramAnnotationData := omc.OMC.GetDiagramAnnotation(modelName)
		interfaceGraphicsData := interfaceDiagramAnnotationData[8].([]interface{})
		data := make(map[string]interface{}, 0)
		data["ID"] = "0"
		data["classname"] = modelName
		data["extent1Diagram"] = interfaceDiagramAnnotationData[0].(string) + "," + interfaceDiagramAnnotationData[1].(string)
		data["extent2Diagram"] = interfaceDiagramAnnotationData[2].(string) + "," + interfaceDiagramAnnotationData[3].(string)
		data["graphType"] = ""
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
		data["inputOutputs"] = make([]map[string]interface{}, 0, 1)
		data["subShapes"] = make([]map[string]interface{}, 0, 1)
		data1 := g.data01(interfaceGraphicsData)
		data["subShapes"] = data1
		g.data[1] = append(g.data[1], data)
		log.Println("g.data: ", g.data)
		return g.data
	}
	nameList := g.getICList(modelName)
	diagramAnnotationData := omc.OMC.GetDiagramAnnotationList(nameList)
	if len(diagramAnnotationData) >= 8 {
		dData := diagramAnnotationData[len(diagramAnnotationData)-1]
		data1 := g.data01(dData.([]interface{}))
		g.data[0] = append(g.data[0], data1...)
	}
	g.getnthconnectionData(nameList)

	//nameList第一个名字是模型自身的名字，先获取模型自身的视图数据
	componentsData := omc.OMC.GetElementsList([]string{modelName})
	componentAnnotationsData := omc.OMC.GetComponentAnnotationsList([]string{modelName})
	//componentAnnotationsData := getElementAndDiagramAnnotations([]string{modelName})
	data2 := g.data02(componentsData, componentAnnotationsData, false, "")
	for i := 0; i < len(data2); i++ {
		data2[i]["mobility"] = true //模型自身的组件是可以移动的，设置字段"mobility"为true
	}
	g.data[1] = append(g.data[1], data2...)

	//nameList第二个名字开始是继承模型的名字，获取继承模型的视图数据
	componentsData = omc.OMC.GetElementsList(nameList[:len(nameList)-1])
	componentAnnotationsData = getElementAndDiagramAnnotations(nameList[:len(nameList)-1])
	data2 = g.data02(componentsData, componentAnnotationsData, false, "")
	for i := 0; i < len(data2); i++ {
		data2[i]["mobility"] = false //继承模型的组件是不可以移动的，设置字段"mobility"为false
	}

	g.data[1] = append(g.data[1], data2...)
	return g.data
}

func GetComponentGraphicsData(modelName string, componentName string) [][]map[string]interface{} {
	var g = graphicsData{}
	g.data = [][]map[string]interface{}{{}, {}}

	components := omc.OMC.GetElementsList([]string{modelName})
	componentAnnotations := getElementAndDiagramAnnotations([]string{modelName})
	var componentsData [][]interface{}
	var componentAnnotationsData [][]interface{}
	for i := 0; i < len(components); i++ {
		if components[i] != nil {
			if components[i][3] == componentName {
				componentsData = append(componentsData, components[i])
				componentAnnotationsData = append(componentAnnotationsData, componentAnnotations[i])
			}
		}
	}
	data2 := g.data02(componentsData, componentAnnotationsData, false, "")
	g.data[1] = append(g.data[1], data2...)
	return g.data
}

func oneDimensionalProcessing(Data interface{}) string {
	if reflect.TypeOf(Data).String() == "[]interface {}" {
		drawingData := Data.([]interface{})
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

func twoDimensionalProcessing(drawingData []interface{}) []string {
	var data []string
	dataType := reflect.TypeOf(drawingData[0].([]interface{})[0]).String()

	if dataType == "string" {
		for i := 0; i < len(drawingData); i++ {
			var piList []string
			for pi := 0; pi < len(drawingData[i].([]interface{})); pi++ {
				piList = append(piList, drawingData[i].([]interface{})[pi].(string))
			}
			data = append(data, strings.Join(piList, ","))
		}
		return data
	}
	if dataType == "[]interface {}" {
		for i := 0; i < len(drawingData); {
			for pi := 0; pi < len(drawingData[i].([]interface{})); pi++ {
				var piList []string
				for pii := 0; pii < len(drawingData[i].([]interface{})[pi].([]interface{})); pii++ {
					piList = append(piList, drawingData[i].([]interface{})[pi].([]interface{})[pii].(string))
				}
				data = append(data, strings.Join(piList, ","))
			}
			break
		}
		return data
	}
	return data
}

func (g *graphicsData) getICList(name string) []string {
	dataList := GetICList(name)
	return dataList
}

func find(data []interface{}, str string) bool {
	for i := 0; i < len(data); i++ {
		if reflect.TypeOf(data[i]).String() == "string" && data[i].(string) == str {
			return true
		}
	}
	return false
}

func (g *graphicsData) data01(cData []interface{}) []map[string]interface{} {
	dataList := make([]map[string]interface{}, 0, 1)
	for i := 0; i < len(cData); i += 2 {
		data := map[string]interface{}{}
		drawingDataList := cData[i+1].([]interface{})

		DynamicSelect := find(drawingDataList, "DynamicSelect")
		if DynamicSelect {
			var drawingDataListFilter []interface{}
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
				data["points"] = twoDimensionalProcessing(drawingDataList[3].([]interface{}))
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
			data["polygonPoints"] = twoDimensionalProcessing(drawingDataList[8].([]interface{}))
			ppData := data["polygonPoints"].([]string)
			if len(ppData) < 4 {
				data["polygonPoints"] = append(ppData, ppData[0])
			}
			data["smooth"] = drawingDataList[9]
		case "Line":
			delete(data, "fillColor")
			delete(data, "fillPattern")
			data["points"] = twoDimensionalProcessing(drawingDataList[3].([]interface{}))
			data["color"] = oneDimensionalProcessing(drawingDataList[4])
			data["lineThickness"] = drawingDataList[6]
			data["arrow"] = oneDimensionalProcessing(drawingDataList[7])
			data["arrowSize"] = drawingDataList[8]
			data["smooth"] = drawingDataList[9]
		case "Text":
			data["fillPattern"] = drawingDataList[6]
			data["extentsPoints"] = twoDimensionalProcessing(drawingDataList[8].([]interface{}))
			typeOriginalTextString, ok := drawingDataList[9].([]interface{})
			if ok {
				data["originalTextString"] = typeOriginalTextString[0]
			} else {
				data["originalTextString"] = drawingDataList[9]
			}
			data["fontSize"] = drawingDataList[10]
			data["textColor"] = oneDimensionalProcessing(drawingDataList[11])
			data["fontName"] = drawingDataList[12]
			data["textStyles"] = drawingDataList[13]
			data["horizontalAlignment"] = drawingDataList[14]
		case "Rectangle":
			data["fillPattern"] = drawingDataList[6]
			data["borderPattern"] = drawingDataList[8]
			data["extentsPoints"] = twoDimensionalProcessing(drawingDataList[9].([]interface{}))
			data["radius"] = drawingDataList[10]
		case "Ellipse":
			data["fillPattern"] = drawingDataList[6]
			data["extentsPoints"] = twoDimensionalProcessing(drawingDataList[8].([]interface{}))
			data["startAngle"] = drawingDataList[9]
			data["endAngle"] = drawingDataList[10]
		}
		dataList = append(dataList, data)
	}
	return dataList
}

func (g *graphicsData) data02(cData [][]interface{}, caData [][]interface{}, isIcon bool, parent string) []map[string]interface{} {
	dataList := make([]map[string]interface{}, 0, 1)
	var cDataFilter [][]interface{}
	var caDataFilter [][]interface{}
	dataLen := func() int {
		if len(cData) > len(caData) {
			return len(cDataFilter)
		}
		return len(caData)
	}()
	if isIcon == true && cData != nil && caData != nil {
		for i := 0; i < dataLen; i++ {
			//cDataSplit := strings.Split(cData[i][2].(string), ".")
			//for ii := 0; ii < len(cDataSplit); ii++ {
			//	if "Interfaces" == cDataSplit[ii] {
			//		cDataFilter = append(cDataFilter, cData[i])
			//		caDataFilter = append(caDataFilter, caData[i])
			//	}
			//}
			nameType := omc.OMC.GetClassRestriction(cData[i][2].(string))
			if nameType == "connector" || nameType == "expandable connector" {
				cDataFilter = append(cDataFilter, cData[i])
				caDataFilter = append(caDataFilter, caData[i])
			}
		}
	} else {
		cDataFilter = cData
		caDataFilter = caData
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
		if len(caDataFilter[i]) > 2 {
			caDataFilter[i] = caDataFilter[i][len(caDataFilter[i])-2:]
		}
		name := cDataFilter[i][2].(string)
		nameList := g.getICList(name)
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
		if placementIndex != -1 || cDataFilter[i][9] == "true" {

			componentsData := omc.OMC.GetElementsList(nameList)
			componentAnnotationsData := omc.OMC.GetElementAnnotationsList(nameList)
			IconAnnotationData := getIconAndDiagramAnnotations(nameList, isIcon)
			if len(caDataFilter[i]) < 1 {
				continue
			}
			caf := caDataFilter[i][placementIndex+1].([]interface{})
			if len(caf) < 7 {
				// 出现错误会使数据不可用， 长度小于预期，弃用
				continue
			}
			// TODO 默认填充色修改
			rotateAngle := func() string {
				if caf[7] == "-" {
					return "0"
				} else {
					return caf[7].(string)
				}
			}()
			data := map[string]interface{}{"type": "Transformation"}

			data["graphType"] = func() string {
				nameType := omc.OMC.GetClassRestriction(name)
				if nameType == "connector" || nameType == "expandable connector" {
					return "connecter"
				}
				return ""
			}()
			data["ID"] = strconv.Itoa(i)
			data["classname"] = name
			data["name"] = cDataFilter[i][3]
			data["original_name"] = cDataFilter[i][3]
			data["parent"] = parent
			data["visible"] = caf[0]
			data["rotateAngle"] = rotateAngle
			data["originDiagram"] = strings.Join([]string{caf[1].(string), caf[2].(string)}, ",")
			data["extent1Diagram"] = strings.Join([]string{caf[3].(string), caf[4].(string)}, ",")
			data["extent2Diagram"] = strings.Join([]string{caf[5].(string), caf[6].(string)}, ",")
			data["rotation"] = rotateAngle
			data["output_type"] = func() string {
				t := cDataFilter[i][len(cDataFilter[i])-1].([]interface{})
				str := fmt.Sprintf("%s", t)
				return str
			}()
			data["inputOutputs"] = g.data02(componentsData, componentAnnotationsData, true, data["name"].(string))
			data["subShapes"] = g.data01(IconAnnotationData)
			dataList = append(dataList, data)
		}
	}
	return dataList
}

func (g *graphicsData) getnthconnectionData(nameList []string) {
	ConnectionCount := omc.OMC.GetConnectionCountList(nameList)
	for i := 0; i < len(ConnectionCount); i++ {
		for c := 0; c < ConnectionCount[i]; c++ {
			ncData := omc.OMC.GetNthConnection(nameList[i], c+1)
			ncaData := omc.OMC.GetNthConnectionAnnotation(nameList[i], c+1) //
			d1Data := g.data01(ncaData)
			if len(ncData) != 0 && len(ncaData) != 0 && len(d1Data) != 0 {
				daData := d1Data[0]
				if i == 0 { //i==0的时候，表示目前遍历的是模型自身的组件，模型自身的组件可以移动，设在"mobility"为true
					daData["mobility"] = true
				} else {
					daData["mobility"] = false
				}
				daData["connectionfrom_original_name"] = ncData[0]
				daData["connectionto_original_name"] = ncData[1]
				re1, _ := regexp.Compile("[[0-9]+]$")
				re2, _ := regexp.Compile("[[0-9]+].")
				connectionfrom := re1.ReplaceAll([]byte(ncData[0]), []byte(""))
				connectionto := re1.ReplaceAll([]byte(ncData[1]), []byte(""))
				daData["connectionfrom"] = string(re2.ReplaceAll(connectionfrom, []byte(".")))
				daData["connectionto"] = string(re2.ReplaceAll(connectionto, []byte(".")))
				g.data[0] = append(g.data[0], daData)
			}
		}
	}
}

func getElementAndDiagramAnnotations(nameList []string) [][]interface{} {
	var data [][]interface{}

	for _, name := range nameList {
		var result []interface{}
		//if strings.Index(name, "Interfaces") == -1 {
		nameType := omc.OMC.GetClassRestriction(name)
		if nameType == "connector" || nameType == "expandable connector" {
			result = omc.OMC.GetDiagramAnnotation(name)
			if len(result) > 8 {
				result = result[8].([]interface{})
			}
		} else {
			result = omc.OMC.GetElementAnnotations(name)
		}
		for _, i := range result {
			data = append(data, i.([]interface{}))
		}
	}
	return data
}

func getIconAndDiagramAnnotations(nameList []string, isIcon bool) []interface{} {
	var data []interface{}
	for _, name := range nameList {
		var result []interface{}
		nameType := omc.OMC.GetClassRestriction(name)
		if (nameType == "connector" || nameType == "expandable connector") && isIcon == false {
			result = omc.OMC.GetDiagramAnnotation(name)
			if len(result) > 8 {
				result = result[8].([]interface{})
			}
		} else {
			result = omc.OMC.GetIconAnnotation(name)
		}
		for _, i := range result {
			data = append(data, i)
		}
	}
	return data
}

//Buildings.ThermalZones.EnergyPlus.Examples.SmallOffice.Unconditioned
