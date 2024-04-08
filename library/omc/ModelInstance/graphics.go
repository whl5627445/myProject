package instance

import (
	"encoding/base64"
	"log"
	"os"
	"strconv"
	"strings"

	"yssim-go/library/omc"
	"yssim-go/library/stringOperation"
)

// getRectangle 获取Rectangle类型图形数据
func getRectangle(data []any, graphics map[string]any) map[string]any {
	graphics["visible"] = data[0]       // 是否显示
	graphics["offset"] = data[1]        // 偏移量
	graphics["rotation"] = data[2]      // 旋转角度
	graphics["lineColor"] = data[3]     // 线的颜色
	graphics["fillColor"] = data[4]     // 填充颜色
	graphics["linePattern"] = data[5]   // 线的样式
	graphics["fillPattern"] = data[6]   // 填充样式
	graphics["lineThickness"] = data[7] // 线的粗细
	graphics["borderPattern"] = data[8] // 边框图案
	graphics["extentsPoints"] = data[9] // 范围坐标
	graphics["borderRadius"] = data[10] // 边框半径
	graphics["type"] = "Rectangle"      // 类型
	return graphics
}

// getPolygon 获取Polygon类型图形数据
func getPolygon(data []any, graphics map[string]any) map[string]any {
	graphics["visible"] = data[0]       // 是否显示
	graphics["offset"] = data[1]        // 偏移量
	graphics["rotation"] = data[2]      // 旋转角度
	graphics["lineColor"] = data[3]     // 线的颜色
	graphics["fillColor"] = data[4]     // 填充颜色
	graphics["linePattern"] = data[5]   // 线的样式
	graphics["fillPattern"] = data[6]   // 填充样式
	graphics["lineThickness"] = data[7] // 线的粗细
	graphics["extentsPoints"] = data[8] // 范围坐标
	graphics["smooth"] = data[9]        // 平滑度
	graphics["type"] = "Polygon"        // 类型
	return graphics
}

// getLine 获取Line类型图形数据
func getLine(data []any, graphics map[string]any) map[string]any {
	graphics["visible"] = data[0]       // 是否显示
	graphics["offset"] = data[1]        // 偏移量
	graphics["rotation"] = data[2]      // 旋转角度
	graphics["extentsPoints"] = data[3] // 范围坐标
	graphics["lineColor"] = data[4]     // 线的颜色
	graphics["linePattern"] = data[5]   // 线的样式
	graphics["lineThickness"] = data[6] // 线的粗细
	graphics["arrow"] = data[7]         // 箭头样式， 开始样式与结束样式
	graphics["arrowSize"] = data[8]     // 箭头大小
	graphics["smooth"] = data[9]        // 平滑度
	graphics["type"] = "Line"           // 类型
	return graphics
}

// getEllipse 获取Ellipse类型图形数据
func getEllipse(data []any, graphics map[string]any) map[string]any {
	graphics["visible"] = data[0]       // 是否显示
	graphics["offset"] = data[1]        // 偏移量
	graphics["rotation"] = data[2]      // 旋转角度
	graphics["lineColor"] = data[3]     // 线的颜色
	graphics["fillColor"] = data[4]     // 填充颜色
	graphics["linePattern"] = data[5]   // 线的样式
	graphics["fillPattern"] = data[6]   // 填充样式
	graphics["lineThickness"] = data[7] // 线的粗细
	graphics["extentsPoints"] = data[8] // 范围坐标
	graphics["startingAngle"] = data[9] // 起始角度
	graphics["endingAngle"] = data[10]  // 结束角度
	graphics["closure"] = data[11]      // 闭合
	graphics["type"] = "Ellipse"        // 类型
	return graphics
}

// getText 获取Text类型图形数据
func getText(data []any, graphics map[string]any, modelElements *elements) map[string]any {
	graphics["visible"] = data[0]                                       // 是否显示
	graphics["offset"] = data[1]                                        // 偏移量
	graphics["rotation"] = data[2]                                      // 旋转角度
	graphics["lineColor"] = data[3]                                     // 线的颜色
	graphics["fillColor"] = data[4]                                     // 填充颜色
	graphics["linePattern"] = data[5]                                   // 线的样式
	graphics["fillPattern"] = data[6]                                   // 填充样式
	graphics["lineThickness"] = data[7]                                 // 线的粗细
	graphics["extentsPoints"] = data[8]                                 // 范围坐标
	graphics["text"] = getTextString(data[9], modelElements)            // 文本文字
	graphics["fontSize"] = data[10]                                     // 字体大小
	graphics["textColor"] = data[11]                                    // 文本颜色
	graphics["fontName"] = data[12]                                     // 文字字体
	graphics["glyph"] = data[13]                                        // 字形， 粗体、斜体、下划线之类的
	graphics["horizontalAlignment"] = data[14].(map[string]any)["name"] // 水平对齐
	graphics["type"] = "Text"                                           // 类型
	return graphics
}

// getTextString 获取Text类型图形数据中的文字字符串内容， 有可能包含有组件参数需要获取对应的值
func getTextString(textData any, modelElements *elements) string {
	if _, ok := textData.(string); !ok {
		// s, _ := sonic.Marshal(textData)
		// fmt.Printf("%s,", s)
		return ""
	}
	text := textData.(string)
	originalTextString := text
	if modelElements == nil {
		return originalTextString
	}
	textList := stringOperation.PluralSplit(originalTextString, []string{"/", ",", "\t", "\n", "\r", " "})
	for _, t := range textList {
		pSignIndex := strings.Index(t, "%")
		if pSignIndex != -1 {
			varName := t[pSignIndex+1:]
			varValue := ""
			if varName != "name" {
				varName = strings.TrimSuffix(varName, "%")
				if varName != "" {
					varValue = getModelGraphicsParameters(varName, modelElements)
					if varValue == "" {
						varValue = varName
					}
				}
				oldVarName := "%" + varName
				originalTextString = strings.Replace(originalTextString, oldVarName, varValue, 1)
			}
		}
	}
	return originalTextString
}

// getModelGraphicsParameters 获取Text类型图形数据中组件参数的值
func getModelGraphicsParameters(varName string, modelElements *elements) string {
	modifiers, ok := modelElements.Modifiers.(map[string]any)
	if ok {
		varValue, ok := modifiers[varName]
		if ok {
			if vMap, ok := varValue.(map[string]any); ok {
				// s, _ := sonic.Marshal(varValue)
				// fmt.Printf("%s,", s)
				return vMap["value"].(string)
			}
			return varValue.(string)
		}
	}
	if modelElements.Type.BasicType {
		return varName
	}
	typeInstance := modelElements.Type
	valueStr := ""
	for i := 0; i < len(typeInstance.Elements); i++ {
		e := typeInstance.Elements[i]
		if e.BaseClass != nil && e.BaseClass.BasicType && e.Kind == "extends" {
			typeInstance.Elements = append(typeInstance.Elements, e.BaseClass.Elements...)
			continue
		}
		if e.Name == varName {
			value, _ := getParameterValue(e.Value)
			switch value.(type) {
			case string:
				return value.(string)
			case float64:
				valueStr = strconv.FormatFloat(value.(float64), 'f', -1, 64)
			case bool:
				valueStr = strconv.FormatBool(value.(bool))
			}
		}
	}
	return valueStr
}

// getParameterValue 获取Text类型图形数据中组件参数的值的核心逻辑，返回值内容和值类型
func getParameterValue(value any) (any, string) {
	switch value.(type) {
	case map[string]any:
		if v, ok := value.(map[string]any)["value"]; ok {
			return v, "Normal"
		}
		if v, ok := value.(map[string]any)["binding"]; ok {
			switch v.(type) {
			case map[string]any:
				vMap := v.(map[string]any)
				if vMap["kind"] == "enum" {
					return vMap["name"], "Enumeration"
				}
			case bool:
				return v, "CheckBox"
			}

			return v, "Normal"
		}
	}
	return "", "Normal"
}

// getBitmap 获取Bitmap类型图形数据
func getBitmap(data []any, graphics map[string]any) map[string]any {
	imageBase64 := data[5].(string)
	graphics["imageBase64"] = imageBase64
	if imageBase64 == "" {
		graphics["imageBase64"] = getImage(data[4].(string)) // 图片数据
	}
	graphics["visible"] = data[0]       // 是否显示
	graphics["offset"] = data[1]        // 偏移量
	graphics["rotation"] = data[2]      // 旋转角度
	graphics["extentsPoints"] = data[3] // 范围坐标
	graphics["type"] = "Bitmap"         // 类型
	return graphics
}

// getImage 获取Image类型图形数据
func getImage(dataImagePath string) string {
	imagePath := omc.OMC.UriToFilename(dataImagePath)
	imageBytes, err := os.ReadFile(imagePath)
	if err != nil {
		log.Println("dataImagePath 错误：", dataImagePath)
		log.Println("err：", err)
	}
	return base64.StdEncoding.EncodeToString(imageBytes)
}
