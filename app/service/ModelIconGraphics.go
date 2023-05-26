package service

import (
	"encoding/base64"
	"log"
	"os"
	"strings"
	"yssim-go/library/omc"
	"yssim-go/library/stringOperation"
)

func GetIconNew(modelName string, icon bool) map[string]interface{} {
	data := make(map[string]interface{}, 0)
	iconData := omc.OMC.GetIconAnnotation(modelName)
	if len(iconData) > 8 {
		bitmapData := iconData[8].([]interface{})
		Bitmap := getBitmapImage(bitmapData)
		if Bitmap != "" {
			data = map[string]interface{}{
				"type":     "base64",
				"graphics": Bitmap,
			}
			return data
		}
	}
	graphics := map[string]interface{}{}
	nameType := omc.OMC.GetClassRestriction(modelName)
	if nameType != "connector" || (icon && nameType == "connector") {
		graphics = getIconAnnotationGraphics(modelName, nameType)
	} else {
		graphics = getDiagramAnnotationGraphics(modelName, nameType)
	}

	data = map[string]interface{}{
		"type":     "graphics",
		"graphics": graphics,
	}
	return data
}

func getIconAnnotationGraphics(modelName, nameType string) map[string]interface{} {
	modelNameList := GetICList(modelName)
	modelIconAnnotation := getIconAnnotation(modelNameList)
	AnnotationConfig := []interface{}{}
	data := map[string]interface{}{"extent1Diagram": "-,-", "extent2Diagram": "-,-", "initialScale": "0.1"}
	componentsData, componentAnnotationsData := getElementsAndModelName(modelNameList)
	subShapes := iconSubShapes(modelIconAnnotation, modelName)
	inputOutputs := iconInputOutputs(componentsData, componentAnnotationsData, modelName)
	if len(subShapes) == 0 && len(inputOutputs) == 0 && len(modelIconAnnotation) == 0 {
		return nil
	}
	if len(modelIconAnnotation) > 8 {
		AnnotationConfig = modelIconAnnotation[:8]
		x1, y1, x2, y2 := AnnotationConfig[0], AnnotationConfig[1], AnnotationConfig[2], AnnotationConfig[3]
		data["extent1Diagram"] = strings.Replace(strings.Join([]string{x1.(string), y1.(string)}, ","), "-,-", "-100.0,-100.0", 1)
		data["extent2Diagram"] = strings.Replace(strings.Join([]string{x2.(string), y2.(string)}, ","), "-,-", "100.0,100.0", 1)
		if AnnotationConfig[5].(string) != "-" {
			data["initialScale"] = AnnotationConfig[5].(string)
		}
	}
	data["output_type"] = "[]"
	data["graphType"] = nameType
	data["classname"] = modelName
	data["parent"] = ""
	data["visible"] = "true"
	data["mobility"] = true
	data["rotation"] = "0"
	data["inputOutputs"] = inputOutputs
	data["subShapes"] = subShapes
	return data
}

func getDiagramAnnotationGraphics(modelName, nameType string) map[string]interface{} {
	modelIconAnnotation := getDiagramAnnotation(modelName)
	AnnotationConfig := []interface{}{}
	data := map[string]interface{}{"extent1Diagram": "-,-", "extent2Diagram": "-,-", "initialScale": "0.1"}

	subShapes := iconSubShapes(modelIconAnnotation, modelName)
	if len(subShapes) == 0 && len(modelIconAnnotation) == 0 {
		return nil
	}
	if len(modelIconAnnotation) > 8 {
		AnnotationConfig = modelIconAnnotation[:8]
		x1, y1, x2, y2 := AnnotationConfig[0], AnnotationConfig[1], AnnotationConfig[2], AnnotationConfig[3]
		data["extent1Diagram"] = strings.Replace(strings.Join([]string{x1.(string), y1.(string)}, ","), "-,-", "-100.0,-100.0", 1)
		data["extent2Diagram"] = strings.Replace(strings.Join([]string{x2.(string), y2.(string)}, ","), "-,-", "100.0,100.0", 1)
		if AnnotationConfig[5].(string) != "-" {
			data["initialScale"] = AnnotationConfig[5].(string)
		}
	}
	data["output_type"] = "[]"
	data["graphType"] = nameType
	data["classname"] = modelName
	data["parent"] = ""
	data["visible"] = "true"
	data["mobility"] = true
	data["rotation"] = "0"
	data["inputOutputs"] = make([]interface{}, 0)
	data["subShapes"] = subShapes
	return data
}

func iconSubShapes(cData []interface{}, modelName string) []map[string]interface{} {
	dataList := make([]map[string]interface{}, 0, 1)
	if len(cData) < 8 {
		return dataList
	}
	d := cData[8].([]interface{})
	for i := 0; i < len(d); i += 2 {
		data := map[string]interface{}{}
		drawingDataList := d[i+1].([]interface{})
		if drawingDataList[0].(string) != "true" {
			continue
		}
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
		dataType := d[i]
		if len(drawingDataList) < 9 {
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
				originalTextString := drawingDataList[9].(string)
				data["textType"] = "var"
				if strings.Contains(originalTextString, "%") {
					data["textType"] = "text"
				}
				textList := stringOperation.PluralSplit(originalTextString, []string{"/", ",", "\t", "\n", "\r", " "})
				for _, t := range textList {
					pSignIndex := strings.Index(t, "%")
					if pSignIndex != -1 {
						classNameAll := omc.OMC.GetInheritedClassesListAll([]string{modelName})
						varName := t[pSignIndex+1:]
						varValue := ""
						if varName != "name" {
							varName = strings.TrimSuffix(varName, "%")
							for _, name := range classNameAll {
								varValue = omc.OMC.GetParameterValue(name, varName)
								if varValue != "" {
									break
								}

								if varValue == "" {
									varValue = varName
								}

								if len(varValue) > 20 && (strings.Contains(varValue, ".") || strings.Contains(varValue, " ")) {
									varValueList := strings.Split(varValue, ".") // 某些值是模型全称的需要取最后一部分。所以分割一下
									varValue = varValueList[len(varValueList)-1]
								}
							}
							Unit := ""
							classNameList := append(classNameAll, modelName)
							for n := 0; n < len(classNameList); n++ {
								Unit = omc.OMC.GetElementModifierValue(classNameList[n], varName+"."+"unit")
								if Unit != "" {
									Unit = " " + Unit
									break
								}
							}
							if Unit == "" {
								for n := 0; n < len(classNameList); n++ {
									classnameData := omc.OMC.GetElements(classNameList[n])
									for p := 0; p < len(classnameData); p++ {
										name := classnameData[p].([]interface{})[3].(string)
										varClassName := classnameData[p].([]interface{})[2].(string)
										if name != varName {
											continue
										}
										Unit = " " + getDerivedClassModifierValueALL(varClassName)
										break
									}
								}
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

func iconInputOutputs(cData [][]interface{}, caData [][]interface{}, modelName string) []map[string]interface{} {
	dataList := make([]map[string]interface{}, 0, 1)
	var cDataFilter [][]interface{}
	var caDataFilter [][]interface{}
	dataLen := func() int {
		if len(cData) > len(caData) {
			return len(caData)
		}
		return len(cData)
	}()
	for i := 0; i < dataLen; i++ {
		nameType := omc.OMC.GetClassRestriction(cData[i][2].(string))
		if nameType == "connector" {
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

			//initialScale := "1"
			//if len(modelIconAnnotationAll) > 0 {
			//	initialScale = modelIconAnnotationAll[5].(string)
			//}
			if len(caDataFilter[i]) < 1 {
				continue
			}
			caf := caDataFilter[i][placementIndex+1].([]interface{})
			if len(caf) < 7 || caf[0].(string) != "true" {
				// 出现错误会使数据不可用， 长度小于预期，弃用
				continue
			}

			data := map[string]interface{}{}

			data["graphType"] = "connecter"
			//data["ID"] = strconv.Itoa(i)
			data["classname"] = classname
			data["name"] = cDataFilter[i][3]
			//data["original_name"] = cDataFilter[i][3]
			data["parent"] = modelName
			data["visible"] = caf[0]
			data["mobility"] = false
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
			IconAnnotationData := getIconAnnotation([]string{classname})
			data["subShapes"] = iconSubShapes(IconAnnotationData, modelName)
			dataList = append(dataList, data)
		}
	}
	return dataList
}

func getBitmapImage(bitmapData []interface{}) string {
	for i := 0; i < len(bitmapData); i += 2 {
		imageData := bitmapData[i]
		if imageData == "Bitmap" {
			image := bitmapData[i+1].([]interface{})[5].(string)
			imageUri := bitmapData[i+1].([]interface{})[4].(string)
			if strings.HasPrefix(imageUri, "modelica://") {
				imageFile := omc.OMC.UriToFilename(imageUri)
				file, err := os.ReadFile(imageFile)
				if err != nil {
					log.Println("获取模型图表文件信息失败: ", err)
					return ""
				}
				fileBase64Str := base64.StdEncoding.EncodeToString(file)
				return "data:image/png;base64," + fileBase64Str
			}
			return "data:image/png;base64," + image
		}
	}
	return ""
}

func getIconAnnotation(modelNameList []string) []interface{} {
	modelIconAnnotation := []interface{}{}
	for i := 0; i < len(modelNameList); i++ {
		IconAnnotation := omc.OMC.GetIconAnnotations(modelNameList[i])
		if len(modelIconAnnotation) > 8 && len(IconAnnotation) > 8 {
			for p := 0; p < len(IconAnnotation[8:]); p++ {
				modelIconAnnotation[8] = append(modelIconAnnotation[8].([]interface{}), IconAnnotation[8:][p].([]interface{})...)
			}
		} else {
			modelIconAnnotation = append(modelIconAnnotation, IconAnnotation...)
		}
	}
	return modelIconAnnotation
}

func getDiagramAnnotation(modelName string) []interface{} {
	modelIconAnnotation := []interface{}{}
	IconAnnotation := omc.OMC.GetDiagramAnnotation(modelName)
	if len(modelIconAnnotation) > 8 && len(IconAnnotation) > 8 {
		for p := 0; p < len(IconAnnotation[8:]); p++ {
			modelIconAnnotation[8] = append(modelIconAnnotation[8].([]interface{}), IconAnnotation[8:][p].([]interface{})...)
		}
	} else {
		modelIconAnnotation = append(modelIconAnnotation, IconAnnotation...)
	}
	return modelIconAnnotation
}
