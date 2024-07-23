package serviceV2

import (
	"encoding/base64"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	serviceV1 "yssim-go/app/v1/service"
	"yssim-go/library/omc"
	modelComponent "yssim-go/library/omc/component"
	"yssim-go/library/stringOperation"
)

func GetIcon(modelName, componentName string, icon bool, displayAllConnector bool) map[string]any {
	data := make(map[string]any)
	iconData := omc.OMC.GetIconAnnotation(modelName)
	modelType := omc.OMC.GetClassRestriction(modelName)
	if len(iconData) > 8 {
		bitmapData := iconData[8].([]any)
		Bitmap := getBitmapImage(bitmapData, modelName, modelType)
		if Bitmap != nil {
			data = map[string]any{
				"type":     "base64",
				"graphics": Bitmap,
			}
			return data
		}
	}

	graphics := map[string]any{}
	if icon {
		graphics = getIconAnnotationGraphics(modelName, modelType, componentName)
	} else {
		if modelType == "connector" || modelType == "expandable connector" {
			graphics = getDiagramAnnotationGraphics(modelName, modelType)
			if graphics != nil {
				iconInstance := getModelInstance(modelName)
				iconInstance.DataPreprocessing()
				graphics["direction"] = iconInstance.Prefixes.Direction
				graphics["restriction"] = iconInstance.Restriction
				graphics["type"] = getConnectorType(componentName, iconInstance)
				graphics["visibleList"] = GetConnectionOption(componentName, iconInstance, nil)
				graphics["connectors"] = getElementsConnectorList(iconInstance, componentName, displayAllConnector)
			}
		} else {
			graphics = getIconAnnotationGraphics(modelName, modelType, componentName)
			if modelType != "model" && graphics != nil {
				iconInstance := getModelInstance(modelName)
				iconInstance.DataPreprocessing()
				graphics["connectors"] = getElementsConnectorList(iconInstance, componentName, displayAllConnector)
			}
		}
	}

	data = map[string]any{
		"type":     "graphics",
		"graphics": graphics,
	}
	return data
}

type coordinateSystemData struct {
	Extent              [][]float64 `json:"extent"`
	PreserveAspectRatio bool        `json:"preserveAspectRatio"`
	InitialScale        float64     `json:"initialScale"`
}

func GetModelExtent(coordinateSystem any) (float64, float64, float64, float64) {
	if coordinateSystem == nil {
		return 0.0, 0.0, 0.0, 0.0
	}
	coordinate := coordinateSystem.(coordinateSystemData)
	initialScale := coordinate.InitialScale
	x1 := coordinate.Extent[0][0]
	y1 := coordinate.Extent[0][1]
	x2 := coordinate.Extent[1][0]
	y2 := coordinate.Extent[1][1]
	x1, y1, x2, y2 = x1*initialScale, y1*initialScale, x2*initialScale, y2*initialScale
	return x1, y1, x2, y2
}

func GetModelExtentToString(coordinateSystem any) []string {
	x1, y1, x2, y2 := GetModelExtent(coordinateSystem)
	x1Str := strconv.FormatFloat(x1, 'f', 1, 64)
	y1Str := strconv.FormatFloat(y1, 'f', 1, 64)
	x2Str := strconv.FormatFloat(x2, 'f', 1, 64)
	y2Str := strconv.FormatFloat(y2, 'f', 1, 64)

	extent1Diagram := strings.Join([]string{x1Str, y1Str}, ",")
	extent2Diagram := strings.Join([]string{x2Str, y2Str}, ",")
	return []string{extent1Diagram, extent2Diagram}
}

func getCoordinateSystem(name string, isIcon bool) []any {
	res := omc.OMC.GetCoordinateSystem(name, isIcon)
	return res
}

// getCoordinateSystemRecursion 会根据提供的模型列表直到找到有数据为止
func getCoordinateSystemRecursion(modelNameList []string, isIcon bool) coordinateSystemData {
	data := coordinateSystemData{
		[][]float64{{-100, -100}, {100, 100}},
		true,
		0.1,
	}
	for i := len(modelNameList) - 1; i >= 0; i-- {
		coordinateSystem := getCoordinateSystem(modelNameList[i], isIcon)
		if len(coordinateSystem) == 8 && coordinateSystem[0] != "-" {
			data.PreserveAspectRatio, _ = strconv.ParseBool(coordinateSystem[4].(string))
			data.InitialScale = func() float64 {
				if coordinateSystem[5] == "-" {
					return 0.1
				}
				InitialScale, _ := strconv.ParseFloat(coordinateSystem[5].(string), 64)
				return InitialScale
			}()
			x1, _ := strconv.ParseFloat(coordinateSystem[0].(string), 64)
			y1, _ := strconv.ParseFloat(coordinateSystem[1].(string), 64)
			x2, _ := strconv.ParseFloat(coordinateSystem[2].(string), 64)
			y2, _ := strconv.ParseFloat(coordinateSystem[3].(string), 64)
			data.Extent[0] = []float64{x1, y1}
			data.Extent[1] = []float64{x2, y2}
			return data
		}
	}
	return data
}

func getIconAnnotationGraphics(modelName, modelType, parentName string) map[string]any {
	data := map[string]any{}
	modelNameList := modelComponent.GetICList(modelName)
	modelIconAnnotation := getIconAnnotation(modelNameList)
	coordinateSystem := getCoordinateSystemRecursion(modelNameList, false)
	componentsData, componentAnnotationsData := getElementsAndModelName(modelNameList)
	subShapes := iconSubShapes(modelIconAnnotation, modelName)
	inputOutputs := iconInputOutputs(componentsData, componentAnnotationsData, modelName, parentName)
	if len(subShapes) == 0 && len(inputOutputs) == 0 && len(modelIconAnnotation) == 0 {
		return nil
	}
	data["outputType"] = make(map[string]any, 0)
	data["classname"] = modelName
	data["direction"] = ""
	data["restriction"] = modelType
	data["type"] = ""
	data["visible"] = true
	data["rotation"] = 0
	data["connectors"] = inputOutputs
	data["subShapes"] = subShapes
	data["extents"] = [][]float64{
		{coordinateSystem.Extent[0][0] * coordinateSystem.InitialScale, coordinateSystem.Extent[0][1] * coordinateSystem.InitialScale},
		{coordinateSystem.Extent[1][0] * coordinateSystem.InitialScale, coordinateSystem.Extent[1][1] * coordinateSystem.InitialScale},
	}
	data["coordinateSystem"] = coordinateSystem
	return data
}

func getDiagramAnnotationGraphics(modelName, modelType string) map[string]any {
	data := map[string]any{}
	nameList := modelComponent.GetICList(modelName)
	modelIconAnnotation := getDiagramAnnotation(nameList)
	coordinateSystem := getCoordinateSystemRecursion(nameList, false)
	subShapes := make([]map[string]any, 0)
	subShapes = append(subShapes, iconSubShapes(modelIconAnnotation, modelName)...)
	if len(subShapes) == 0 && len(modelIconAnnotation) == 0 {
		return nil
	}

	data["outputType"] = make(map[string]any, 0)
	data["classname"] = modelName
	data["parentName"] = ""
	data["visible"] = true
	data["rotation"] = 0
	data["subShapes"] = subShapes
	data["extents"] = [][]float64{
		{coordinateSystem.Extent[0][0] * coordinateSystem.InitialScale, coordinateSystem.Extent[0][1] * coordinateSystem.InitialScale},
		{coordinateSystem.Extent[1][0] * coordinateSystem.InitialScale, coordinateSystem.Extent[1][1] * coordinateSystem.InitialScale},
	}

	data["coordinateSystem"] = coordinateSystem
	return data
}

func iconSubShapes(cData []any, modelName string) []map[string]any {
	dataList := make([]map[string]any, 0, 1)
	if len(cData) < 8 {
		return dataList
	}
	d := cData[8].([]any)
	for i := 0; i < len(d); i += 2 {
		data := map[string]any{}
		drawingDataList := d[i+1].([]any)
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
		dataType := d[i]
		if len(drawingDataList) < 9 {
			if dataType == "Bitmap" {
				data["type"] = dataType
				data["visible"] = true
				data["offset"] = oneDimensionalProcessing(drawingDataList[1])
				rotation, _ := strconv.ParseFloat(drawingDataList[2].(string), 64)
				data["rotation"] = rotation
				data["extentsPoints"] = twoDimensionalProcessing(drawingDataList[3].([]any))
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

		data["type"] = dataType
		data["visible"] = true
		data["offset"] = oneDimensionalProcessing(drawingDataList[1])
		rotation, _ := strconv.ParseFloat(drawingDataList[2].(string), 64)
		data["rotation"] = rotation
		data["lineColor"] = oneDimensionalProcessing(drawingDataList[3])
		data["fillColor"] = oneDimensionalProcessing(drawingDataList[4])
		data["linePattern"] = map[string]any{"name": drawingDataList[5]}
		data["fillPattern"] = map[string]any{"name": drawingDataList[6]}

		switch dataType {
		case "Polygon":
			ppData := twoDimensionalProcessing(drawingDataList[8].([]any))
			data["extentsPoints"] = ppData
			if len(ppData) < 4 {
				data["extentsPoints"] = append(ppData, ppData[0])
			}
			data["smooth"] = map[string]any{"name": drawingDataList[9]}
			thickness, _ := strconv.ParseFloat(drawingDataList[7].(string), 64)
			data["thickness"] = thickness
		case "Line":
			delete(data, "fillColor")
			delete(data, "fillPattern")
			data["extentsPoints"] = twoDimensionalProcessing(drawingDataList[3].([]any))
			data["lineColor"] = oneDimensionalProcessing(drawingDataList[4])
			thickness, _ := strconv.ParseFloat(drawingDataList[6].(string), 64)
			data["thickness"] = thickness
			arrow := drawingDataList[7].([]any)
			data["arrow"] = []map[string]any{{"name": arrow[0]}, {"name": arrow[1]}}
			arrowSize, _ := strconv.ParseFloat(drawingDataList[8].(string), 64)
			data["arrowSize"] = arrowSize
			data["smooth"] = map[string]any{"name": drawingDataList[9]}
		case "Text":
			data["lineColor"] = oneDimensionalProcessing(drawingDataList[3])
			thickness, _ := strconv.ParseFloat(drawingDataList[7].(string), 64)
			data["thickness"] = thickness
			data["fillPattern"] = map[string]any{"name": drawingDataList[6]}
			data["extentsPoints"] = twoDimensionalProcessing(drawingDataList[8].([]any))
			data["varName"] = ""
			typeOriginalTextString, ok := drawingDataList[9].([]any)
			if ok {
				data["textString"] = typeOriginalTextString[0]
			} else {
				originalTextString := drawingDataList[9].(string)
				varName := ""
				data["textType"] = "var"
				if strings.Contains(originalTextString, "%") {
					data["textType"] = "text"
				}
				textList := stringOperation.PluralSplit(originalTextString, []string{"/", ",", "\t", "\n", "\r", " "})
				for _, t := range textList {
					pSignIndex := strings.Index(t, "%")
					if pSignIndex != -1 {
						varName = t[pSignIndex+1:]
						varValue := ""
						Unit := ""
						if varName != "name" {
							varName = strings.TrimSuffix(varName, "%")
							if varName != "" {
								c := serviceV1.GetModelParameters(modelName, "", "", varName)
								varValue, Unit = c[1].(string), c[0].(string)
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
				data["varName"] = varName
				data["textString"] = originalTextString
			}
			fontSize, _ := strconv.ParseFloat(drawingDataList[10].(string), 64)
			data["fontSize"] = fontSize
			data["textColor"] = oneDimensionalProcessing(drawingDataList[11])
			data["fontName"] = drawingDataList[12]
			data["glyph"] = drawingDataList[13]
			data["horizontalAlignment"] = drawingDataList[14]
		case "Rectangle":
			thickness, _ := strconv.ParseFloat(drawingDataList[7].(string), 64)
			data["thickness"] = thickness
			data["fillPattern"] = map[string]any{"name": drawingDataList[6]}
			data["borderPattern"] = map[string]any{"name": drawingDataList[8]}
			data["extentsPoints"] = twoDimensionalProcessing(drawingDataList[9].([]any))
			radius, _ := strconv.ParseFloat(drawingDataList[10].(string), 64)
			data["borderRadius"] = radius
		case "Ellipse":
			thickness, _ := strconv.ParseFloat(drawingDataList[7].(string), 64)
			data["thickness"] = thickness
			data["fillPattern"] = map[string]any{"name": drawingDataList[6]}
			data["extentsPoints"] = twoDimensionalProcessing(drawingDataList[8].([]any))
			startingAngle, _ := strconv.ParseFloat(drawingDataList[9].(string), 64)
			data["startingAngle"] = startingAngle
			endingAngle, _ := strconv.ParseFloat(drawingDataList[10].(string), 64)
			data["endingAngle"] = endingAngle
		}
		dataList = append(dataList, data)
	}
	return dataList
}

func iconInputOutputs(cData [][]any, caData [][]any, modelName, parentName string) []map[string]any {
	dataList := make([]map[string]any, 0, 1)
	var cDataFilter [][]any
	var caDataFilter [][]any

	for i := 0; i < len(cData); i++ {
		cDataFilter = append(cDataFilter, cData[i])
		caDataFilter = append(caDataFilter, caData[i])
	}
	if cDataFilter == nil || caDataFilter == nil {
		return dataList
	}

	for i := 0; i < len(cDataFilter); i++ {
		classname := cDataFilter[i][2].(string)
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
			data := map[string]any{}
			data["classname"] = classname
			data["name"] = cDataFilter[i][3]
			data["modelName"] = modelName
			data["visible"] = true

			data["parentName"] = parentName
			data["comment"] = cDataFilter[i][4]
			rotateAngle := func() float64 {
				r := 0.0
				if caf[14] != "" {
					r, _ = strconv.ParseFloat(caf[14].(string), 64)
					return r
				}
				if caf[7] != "-" {
					r, _ = strconv.ParseFloat(caf[7].(string), 64)
				}
				return r
			}()
			if caf[10].(string) != "-" {
				originX, _ := strconv.ParseFloat(caf[8].(string), 64)
				originY, _ := strconv.ParseFloat(caf[9].(string), 64)
				extentX1, _ := strconv.ParseFloat(caf[10].(string), 64)
				extentY1, _ := strconv.ParseFloat(caf[11].(string), 64)
				extentX2, _ := strconv.ParseFloat(caf[12].(string), 64)
				extentY2, _ := strconv.ParseFloat(caf[13].(string), 64)
				data["origin"] = []float64{originX, originY}
				data["extents"] = [][]float64{{extentX1, extentY1}, {extentX2, extentY2}}
			} else {
				originX, _ := strconv.ParseFloat(caf[1].(string), 64)
				originY, _ := strconv.ParseFloat(caf[2].(string), 64)
				extentX1, _ := strconv.ParseFloat(caf[3].(string), 64)
				extentY1, _ := strconv.ParseFloat(caf[4].(string), 64)
				extentX2, _ := strconv.ParseFloat(caf[5].(string), 64)
				extentY2, _ := strconv.ParseFloat(caf[6].(string), 64)
				data["origin"] = []float64{originX, originY}
				data["extents"] = [][]float64{{extentX1, extentY1}, {extentX2, extentY2}}
			}
			data["rotation"] = rotateAngle
			if cDataFilter[i][14].(string) != "[]" {
				data["outputType"] = map[string]any{"name": cDataFilter[i][14].(string)}
			}
			if cDataFilter[i][16] != "" {
				data["outputType"] = map[string]any{"name": cDataFilter[i][16], "connectorSizing": true}
			}
			data["connectors"] = make([]map[string]any, 0)
			iconInstance := getModelInstance(classname)
			iconInstance.DataPreprocessing()
			data["direction"] = iconInstance.Prefixes.Direction
			data["restriction"] = iconInstance.Restriction

			data["type"] = ""
			if len(iconInstance.Elements) > 0 && iconInstance.Elements[0].BaseClass != nil && iconInstance.Elements[0].BaseClass.BasicType {
				data["type"] = iconInstance.Elements[0].BaseClass.Name
			}
			data["type"] = getConnectorType(data["name"].(string), iconInstance)

			nameList := modelComponent.GetICList(classname)
			IconAnnotationData := getIconAnnotation(nameList)
			data["subShapes"] = iconSubShapes(IconAnnotationData, modelName)
			coordinateSystem := getCoordinateSystemRecursion(nameList, false)
			data["coordinateSystem"] = coordinateSystem
			dataList = append(dataList, data)
		}
	}
	return dataList
}

func getBitmapImage(bitmapData []any, modelName, modelType string) map[string]any {
	modelNameList := modelComponent.GetICList(modelName)
	modelIconAnnotation := getIconAnnotation(modelNameList)
	data := map[string]any{}
	componentsData, componentAnnotationsData := getElementsAndModelName(modelNameList)
	subShapes := iconSubShapes(modelIconAnnotation, modelName)
	inputOutputs := iconInputOutputs(componentsData, componentAnnotationsData, modelName, "")
	if len(subShapes) == 0 && len(inputOutputs) == 0 && len(modelIconAnnotation) == 0 {
		return nil
	}
	coordinateSystem := getCoordinateSystemRecursion(modelNameList, false)
	data["extents"] = [][]float64{
		{coordinateSystem.Extent[0][0] * coordinateSystem.InitialScale, coordinateSystem.Extent[0][1] * coordinateSystem.InitialScale},
		{coordinateSystem.Extent[1][0] * coordinateSystem.InitialScale, coordinateSystem.Extent[1][1] * coordinateSystem.InitialScale},
	}
	data["coordinateSystem"] = coordinateSystem
	data["outputType"] = make(map[string]any, 0)
	data["restriction"] = modelType
	data["classname"] = modelName
	data["parentName"] = ""
	data["visible"] = true
	data["rotation"] = 0
	data["inputOutputs"] = inputOutputs
	data["subShapes"] = subShapes
	for i := 0; i < len(bitmapData); i += 2 {
		imageData := bitmapData[i]
		if imageData == "Bitmap" {
			image := bitmapData[i+1].([]any)[5].(string)
			imageUri := bitmapData[i+1].([]any)[4].(string)
			if strings.HasPrefix(imageUri, "modelica://") {
				imageFile := omc.OMC.UriToFilename(imageUri)
				file, err := os.ReadFile(imageFile)
				if err != nil {
					log.Println("获取模型图表文件信息失败: ", err)
					return nil
				}
				fileBase64Str := base64.StdEncoding.EncodeToString(file)
				data["base64"] = "data:image/png;base64," + fileBase64Str
				return data
			}
			data["base64"] = "data:image/png;base64," + image
			return data
		}
	}
	return nil
}

func getIconAnnotation(modelNameList []string) []any {
	modelIconAnnotation := []any{}
	for i := 0; i < len(modelNameList); i++ {
		IconAnnotation := omc.OMC.GetIconAnnotations(modelNameList[i])
		if len(modelIconAnnotation) > 8 && len(IconAnnotation) > 8 {
			for p := 0; p < len(IconAnnotation[8:]); p++ {
				modelIconAnnotation[8] = append(modelIconAnnotation[8].([]any), IconAnnotation[8:][p].([]any)...)
			}
		} else {
			modelIconAnnotation = append(modelIconAnnotation, IconAnnotation...)
		}
	}
	return modelIconAnnotation
}

func getDiagramAnnotation(nameList []string) []any {

	modelIconAnnotation := []any{}
	for _, name := range nameList {
		IconAnnotation := omc.OMC.GetDiagramAnnotation(name)
		if len(modelIconAnnotation) > 8 && len(IconAnnotation) > 8 {
			for p := 0; p < len(IconAnnotation[8:]); p++ {
				modelIconAnnotation[8] = append(modelIconAnnotation[8].([]any), IconAnnotation[8:][p].([]any)...)
			}
		} else {
			modelIconAnnotation = append(modelIconAnnotation, IconAnnotation...)
		}
	}

	return modelIconAnnotation
}
func getElementsAndModelName(nameList []string) ([][]any, [][]any) {
	var ComponentData [][]any
	var AnnotationData [][]any

	for _, name := range nameList {
		connectorSizingList := []string{}
		componentsData := omc.OMC.GetElements(name)

		componentsList := [][]any{}
		annotationsData := omc.OMC.GetElementAnnotations(name)
		// {Dialog("General","Parameters",true,false,false,-,-,-,-,"",true)}
		for index, c := range componentsData {
			modelType := omc.OMC.GetClassRestriction(c.([]any)[2].(string))
			if modelType != "connector" && modelType != "expandable connector" {
				continue
			}
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
				// if len(annotation) <= 1 || dListTab[len(dListTab)-1] == "true" {
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

func find(data []any, str string) bool {
	for i := 0; i < len(data); i++ {
		if reflect.TypeOf(data[i]).String() == "string" && data[i].(string) == str {
			return true
		}
	}
	return false
}

func oneDimensionalProcessing(Data any) []float64 {
	var data []float64
	if reflect.TypeOf(Data).String() == "[]interface {}" {
		drawingData := Data.([]any)
		defer func() {
			if err := recover(); err != nil {
				data = []float64{0, 0}
			}
		}()
		if len(drawingData) >= 1 && reflect.TypeOf(drawingData[0]).String() != "string" {
			return []float64{0, 0}
		}
		for i := 0; i < len(drawingData); i++ {
			d, _ := strconv.ParseFloat(drawingData[i].(string), 64)
			data = append(data, d)
		}
		return data
	}
	return []float64{0, 0}
}

func twoDimensionalProcessing(drawingData []any) [][]float64 {
	var data [][]float64
	dataType := reflect.TypeOf(drawingData[0].([]any)[0]).String()

	if dataType == "string" {
		for i := 0; i < len(drawingData); i++ {
			var piList []float64
			for pi := 0; pi < len(drawingData[i].([]any)); pi++ {
				d, _ := strconv.ParseFloat(drawingData[i].([]any)[pi].(string), 64)
				piList = append(piList, d)
			}
			data = append(data, piList)
		}
		return data
	}
	if dataType == "[]interface {}" {
		for i := 0; i < len(drawingData); {
			for pi := 0; pi < len(drawingData[i].([]any)); pi++ {
				var piList []float64
				for pii := 0; pii < len(drawingData[i].([]any)[pi].([]any)); pii++ {
					d, _ := strconv.ParseFloat(drawingData[i].([]any)[pi].([]any)[pii].(string), 64)
					piList = append(piList, d)
				}
				data = append(data, piList)
			}
			// break
		}
		return data
	}
	return data
}
