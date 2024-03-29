package service

import (
	"encoding/base64"
	"log"
	"os"
	"strconv"
	"strings"

	"yssim-go/library/omc"
	"yssim-go/library/stringOperation"
)

func GetIconNew(modelName, componentName string, icon bool) map[string]any {
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
	if (modelType != "connector" && modelType != "expandable connector") || (icon && (modelType == "connector" || modelType == "expandable connector")) {
		graphics = getIconAnnotationGraphics(modelName, modelType, componentName)
	} else {
		graphics = getDiagramAnnotationGraphics(modelName, modelType)
	}

	data = map[string]any{
		"type":     "graphics",
		"graphics": graphics,
	}
	return data
}

type coordinateSystemData struct {
	Extent1Diagram      []float64 `json:"extent1_diagram"`
	Extent2Diagram      []float64 `json:"extent2_diagram"`
	PreserveAspectRatio bool      `json:"preserve_aspect_ratio"`
	InitialScale        float64   `json:"initial_scale"`
}

func GetModelExtent(coordinateSystem any) (float64, float64, float64, float64) {
	coordinate := coordinateSystem.(coordinateSystemData)
	initialScale := coordinate.InitialScale
	x1 := coordinate.Extent1Diagram[0]
	y1 := coordinate.Extent1Diagram[1]
	x2 := coordinate.Extent2Diagram[0]
	y2 := coordinate.Extent2Diagram[1]
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

// getCoordinateSystemRecursion 会根据提供的模型列表直到找到有数据为止
func getCoordinateSystemRecursion(modelNameList []string, isIcon bool) coordinateSystemData {
	data := coordinateSystemData{
		[]float64{-100, -100},
		[]float64{100, 100},
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
			// initialScale, _ := strconv.ParseFloat(data["initialScale"], 64)
			// x1, _ := strconv.ParseFloat(coordinateSystem[0].(string), 64)
			// y1, _ := strconv.ParseFloat(coordinateSystem[1].(string), 64)
			// x2, _ := strconv.ParseFloat(coordinateSystem[2].(string), 64)
			// y2, _ := strconv.ParseFloat(coordinateSystem[3].(string), 64)
			// x1, y1, x2, y2 = x1*initialScale, y1*initialScale, x2*initialScale, y2*initialScale
			// x1Str := strconv.FormatFloat(x1, 'f', 1, 64)
			// y1Str := strconv.FormatFloat(y1, 'f', 1, 64)
			// x2Str := strconv.FormatFloat(x2, 'f', 1, 64)
			// y2Str := strconv.FormatFloat(y2, 'f', 1, 64)
			// data["extent1Diagram"] = strings.Join([]string{x1Str, y1Str}, ",")
			x1, _ := strconv.ParseFloat(coordinateSystem[0].(string), 64)
			y1, _ := strconv.ParseFloat(coordinateSystem[1].(string), 64)
			x2, _ := strconv.ParseFloat(coordinateSystem[2].(string), 64)
			y2, _ := strconv.ParseFloat(coordinateSystem[3].(string), 64)
			data.Extent1Diagram = []float64{x1, y1}
			// data["extent2Diagram"] = strings.Join([]string{x2Str, y2Str}, ",")
			data.Extent2Diagram = []float64{x2, y2}
			return data
		}
	}
	return data
}

func getIconAnnotationGraphics(modelName, modelType, parentName string) map[string]any {
	data := map[string]any{}
	modelNameList := GetICList(modelName)
	modelIconAnnotation := getIconAnnotation(modelNameList)
	coordinateSystem := getCoordinateSystemRecursion(modelNameList, false)
	componentsData, componentAnnotationsData := getElementsAndModelName(modelNameList)
	subShapes := iconSubShapes(modelIconAnnotation, modelName)
	inputOutputs := iconInputOutputs(componentsData, componentAnnotationsData, modelName, parentName)
	if len(subShapes) == 0 && len(inputOutputs) == 0 && len(modelIconAnnotation) == 0 {
		return nil
	}
	data["output_type"] = "[]"
	data["graphType"] = modelType
	data["classname"] = modelName
	data["parent"] = ""
	data["visible"] = "true"
	data["mobility"] = true
	data["rotation"] = "0"
	data["inputOutputs"] = inputOutputs
	data["subShapes"] = subShapes
	data["extent1Diagram"] = func() string {
		d := []string{}
		for _, p := range coordinateSystem.Extent1Diagram {
			f := strconv.FormatFloat(p*coordinateSystem.InitialScale, 'f', -1, 64)
			d = append(d, f)
		}
		return strings.Join(d, ",")
	}()
	data["extent2Diagram"] = func() string {
		d := []string{}
		for _, p := range coordinateSystem.Extent2Diagram {
			f := strconv.FormatFloat(p*coordinateSystem.InitialScale, 'f', -1, 64)
			d = append(d, f)
		}
		return strings.Join(d, ",")
	}()
	data["coordinate_system"] = coordinateSystem
	return data
}

func getDiagramAnnotationGraphics(modelName, modelType string) map[string]any {
	data := map[string]any{}
	nameList := GetICList(modelName)
	modelIconAnnotation := getDiagramAnnotation(nameList)
	coordinateSystem := getCoordinateSystemRecursion(nameList, false)
	subShapes := make([]map[string]any, 0)
	subShapes = append(subShapes, iconSubShapes(modelIconAnnotation, modelName)...)
	if len(subShapes) == 0 && len(modelIconAnnotation) == 0 {
		return nil
	}

	data["output_type"] = "[]"
	data["graphType"] = modelType
	data["classname"] = modelName
	data["parent"] = ""
	data["visible"] = "true"
	data["mobility"] = true
	data["rotation"] = "0"
	data["visibleList"] = GetConnectionOption(modelName, modelType)
	data["inputOutputs"] = make([]any, 0)
	data["subShapes"] = subShapes
	data["extent1Diagram"] = func() string {
		d := []string{}
		for _, p := range coordinateSystem.Extent1Diagram {
			f := strconv.FormatFloat(p*coordinateSystem.InitialScale, 'f', -1, 64)
			d = append(d, f)
		}
		return strings.Join(d, ",")
	}()
	data["extent2Diagram"] = func() string {
		d := []string{}
		for _, p := range coordinateSystem.Extent2Diagram {
			f := strconv.FormatFloat(p*coordinateSystem.InitialScale, 'f', -1, 64)
			d = append(d, f)
		}
		return strings.Join(d, ",")
	}()
	// data["extent1Diagram"] = func() []float64 {
	//	d := []float64{}
	//	for _, p := range coordinateSystem.Extent2Diagram {
	//		d = append(d, p*coordinateSystem.InitialScale)
	//	}
	//	return d
	// }()
	// data["extent2Diagram"] = func() []float64 {
	//	d := []float64{}
	//	for _, p := range coordinateSystem.Extent2Diagram {
	//		d = append(d, p*coordinateSystem.InitialScale)
	//	}
	//	return d
	// }()
	data["coordinate_system"] = coordinateSystem
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
								c := GetModelParameters(modelName, "", "", varName)
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

func iconInputOutputs(cData [][]any, caData [][]any, modelName, parentName string) []map[string]any {
	dataList := make([]map[string]any, 0, 1)
	var cDataFilter [][]any
	var caDataFilter [][]any

	for i := 0; i < len(cData); i++ {
		modelType := omc.OMC.GetClassRestriction(cData[i][2].(string))
		if modelType == "connector" || modelType == "expandable connector" {
			cData[i] = append(cData[i], modelType)
			cDataFilter = append(cDataFilter, cData[i])
			caDataFilter = append(caDataFilter, caData[i])
		}
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

			// initialScale := "1"
			// if len(modelIconAnnotationAll) > 0 {
			//	initialScale = modelIconAnnotationAll[5].(string)
			// }
			if len(caDataFilter[i]) < 1 {
				continue
			}
			caf := caDataFilter[i][placementIndex+1].([]any)
			if len(caf) < 7 || caf[0].(string) != "true" {
				// 出现错误会使数据不可用， 长度小于预期，弃用
				continue
			}

			data := map[string]any{}

			data["graphType"] = cDataFilter[i][17]
			data["connector_sizing"] = cDataFilter[i][16]
			// data["ID"] = strconv.Itoa(i)
			data["classname"] = classname
			data["name"] = cDataFilter[i][3]
			data["extend_name"] = modelName
			data["visible"] = caf[0]
			data["mobility"] = false
			data["parent"] = parentName
			data["comment"] = cDataFilter[i][4]
			// data["initialScale"] = initialScale
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

			if caf[10].(string) != "-" {
				extentX1, _ := caf[10].(string)
				extentY1, _ := caf[11].(string)
				extentX2, _ := caf[12].(string)
				extentY2, _ := caf[13].(string)
				data["originDiagram"] = strings.Join([]string{caf[8].(string), caf[9].(string)}, ",")
				data["extent1Diagram"] = strings.Join([]string{extentX1, extentY1}, ",")
				data["extent2Diagram"] = strings.Join([]string{extentX2, extentY2}, ",")
			} else {
				data["extent1Diagram"] = strings.Join([]string{caf[3].(string), caf[4].(string)}, ",")
				data["extent2Diagram"] = strings.Join([]string{caf[5].(string), caf[6].(string)}, ",")
				data["originDiagram"] = strings.Join([]string{caf[1].(string), caf[2].(string)}, ",")
			}
			data["extent1Diagram"] = strings.Replace(data["extent1Diagram"].(string), "-,-", "-100.0,-100.0", 1)
			data["extent2Diagram"] = strings.Replace(data["extent2Diagram"].(string), "-,-", "100.0,100.0", 1)
			data["rotateAngle"] = rotateAngle
			data["rotation"] = rotateAngle
			data["output_type"] = func() string {
				t := cDataFilter[i][14].(string)
				return t
			}()
			data["inputOutputs"] = make([]string, 0)
			nameList := GetICList(classname)
			IconAnnotationData := getIconAnnotation(nameList)
			data["subShapes"] = iconSubShapes(IconAnnotationData, modelName)
			coordinateSystem := getCoordinateSystemRecursion(nameList, false)
			data["coordinate_system"] = coordinateSystem
			dataList = append(dataList, data)
		}
	}
	return dataList
}

func getBitmapImage(bitmapData []any, modelName, modelType string) map[string]any {
	modelNameList := GetICList(modelName)
	modelIconAnnotation := getIconAnnotation(modelNameList)
	data := map[string]any{"extent1Diagram": "-,-", "extent2Diagram": "-,-", "initialScale": "0.1"}
	componentsData, componentAnnotationsData := getElementsAndModelName(modelNameList)
	subShapes := iconSubShapes(modelIconAnnotation, modelName)
	inputOutputs := iconInputOutputs(componentsData, componentAnnotationsData, modelName, "")
	if len(subShapes) == 0 && len(inputOutputs) == 0 && len(modelIconAnnotation) == 0 {
		return nil
	}
	coordinateSystem := getCoordinateSystemRecursion(modelNameList, false)
	data["extent1Diagram"] = func() string {
		d := []string{}
		for _, p := range coordinateSystem.Extent1Diagram {
			f := strconv.FormatFloat(p*coordinateSystem.InitialScale, 'f', -1, 64)
			d = append(d, f)
		}
		return strings.Join(d, ",")
	}()
	data["extent2Diagram"] = func() string {
		d := []string{}
		for _, p := range coordinateSystem.Extent2Diagram {
			f := strconv.FormatFloat(p*coordinateSystem.InitialScale, 'f', -1, 64)
			d = append(d, f)
		}
		return strings.Join(d, ",")
	}()
	data["coordinate_system"] = coordinateSystem
	data["output_type"] = "[]"
	data["graphType"] = modelType
	data["classname"] = modelName
	data["parent"] = ""
	data["visible"] = "true"
	data["mobility"] = true
	data["rotation"] = "0"
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
