package instance

import (
	"encoding/base64"
	"log"
	"os"
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
	graphics["thickness"] = data[7]     // 线的粗细
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
	graphics["thickness"] = data[7]     // 线的粗细
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
	graphics["thickness"] = data[6]     // 线的粗细
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
	graphics["thickness"] = data[7]     // 线的粗细
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
	graphics["thickness"] = data[7]                                     // 线的粗细
	graphics["extentsPoints"] = data[8]                                 // 范围坐标
	graphics["textString"] = modelElements.getTextString(data[9])       // 文本文字
	graphics["fontSize"] = data[10]                                     // 字体大小
	graphics["textColor"] = data[11]                                    // 文本颜色
	graphics["fontName"] = data[12]                                     // 文字字体
	graphics["glyph"] = data[13]                                        // 字形， 粗体、斜体、下划线之类的
	graphics["horizontalAlignment"] = data[14].(map[string]any)["name"] // 水平对齐
	graphics["type"] = "Text"                                           // 类型
	return graphics
}

// getTextString 获取Text类型图形数据中的文字字符串内容， 有可能包含有组件参数需要获取对应的值
func (e *elements) getTextString(textData any) string {
	if _, ok := textData.(string); !ok {
		return ""
	}
	text := textData.(string)
	originalTextString := text
	if e == nil {
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
					varValue = e.getModelGraphicsParameters(varName)
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
func (e *elements) getModelGraphicsParameters(varName string) string {
	value := varName
	unitStr := ""
	if p, ok := e.ElementsParameter[varName]; ok {
		if p.Value != nil && p.Value != "" {
			pValue, ok := p.Value.(string)
			if ok {
				value = pValue
			}
		}
		if p.DefaultValue != "" && p.DefaultValue != nil && value == varName {
			pDefaultValue, ok := p.DefaultValue.(string)
			if ok {
				value = pDefaultValue
			}
		}
		if unit, ok := p.ParameterUnit["unit"]; ok {
			uMap, ok := unit.(map[string]any)
			if ok {
				unitStr = uMap["value"].(string)
			}
		}
		if strings.Contains(value, ".") && value != varName {
			strList := strings.Split(value, ".")
			value = strList[len(strList)-1]
		}
		if unitStr != "" {
			value += " " + unitStr
		}
	}
	return value
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
