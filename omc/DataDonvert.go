package omc

import (
	"encoding/json"
	"fmt"
	"strings"
)

func DataToGo(data string) ([]interface{}, error) {
	var resData []interface{}
	resStr := ""
	data = strings.ReplaceAll(data, "\n", "\\n")
	data = strings.ReplaceAll(data, "\r", "")
	data = strings.ReplaceAll(data, "\\\\\\\\", "\\")
	data = strings.ReplaceAll(data, "\\\\\\", "\\")
	data = strings.ReplaceAll(data, ", ", ",")
	data = strings.TrimSuffix(data, ",")
	data = strings.TrimSuffix(data, "\\n")

	if data == "\"\"" || data == "Error" || data == "{}" {
		return resData, nil
	}
	mark := false
	lData := len(data)
	for i := 0; i < lData; i++ {

		if data[i] == '"' {
			strIndex := strings.Index(data[i+1:], "\"")
			dataEndIndex := i + strIndex + 1
			str := data[i : dataEndIndex+1]
			if mark == true {
				str = strings.ReplaceAll(str, "\"", "\\\"")
			}
			resStr += str
			i += strIndex + 1
			continue
		}
		switch {
		case data[i] == '{':
			switch {
			case (i == 0 && data[i+1] == '"') || data[i+1] == '"' || data[i+1] == '{':
				mark = false
				resStr += "["
			default:
				mark = true
				resStr += "[\""
			}
		case data[i] == '}':
			switch {
			case i == lData-1 && data[i-1] == '"' || (data[i-1] == '"' && mark == false) || data[i-1] == ')' || data[i-1] == '}':
				mark = true
				resStr += "]"
			default:
				mark = false
				resStr += "\"]"
			}

		case data[i] == '(':
			switch {
			case i == 0 || data[i+1] == '(' || data[i+1] == '{' || data[i-1] == ',':
				resStr += "["
			case data[i+1] == '"' && data[i-1] != '"':
				resStr += "\",["
			default:
				resStr += "\",[\""
			}
			mark = false
		case data[i] == ')':
			switch {
			case i == lData-1 && (data[i-1] == '}' || data[i-1] == '"'):
				mark = true
				resStr += "]"
			case data[i-1] != '}' && data[i-1] != '"':
				mark = false
				resStr += "\"]"
			case data[i+1] == '}' || data[i+1] == ')':
				mark = true
				resStr += "]"
			default:
				mark = true
				resStr += "]"
			}

		case data[i] == ',':
			switch {
			case data[i-1] != '}' && data[i-1] != ')' && string(data[i+1]) != "\\" && data[i+1] != '{' && data[i+1] != '(' && data[i-1] != '"' && data[i+1] != '"':
				resStr += "\",\""
			case (data[i-1] == '}' || data[i-1] == ')' || data[i-1] == '"') && data[i+1] != '(' && data[i+1] != '{' && data[i+1] != '"' && string(data[i+1]) != "\\":
				resStr += ",\""
			case (data[i+1] == '(' || data[i+1] == '{' || data[i+1] == '"') && data[i-1] != ')' && data[i-1] != '}' && data[i-1] != '"' && string(data[i+1]) != "\\":
				resStr += "\","
			default:
				resStr += ","
			}
			mark = false
		default:
			if string(data[i]) == "," {
				mark = false
			} else {
				mark = true
			}
			resStr += string(data[i])
		}
	}
	resStr = strings.ReplaceAll(resStr, "\"true\"", "true")
	resStr = strings.ReplaceAll(resStr, "\"false\"", "false")
	if strings.HasPrefix(resStr, "[") == false {
		resStr = "[" + resStr + "]"
	}
	// fmt.Println("resStr:   " + string(resStr))
	if resStr == "[]" || resStr == "[\"\"]" {
		return nil, nil
	}
	b := []byte(resStr)
	err := json.Unmarshal(b, &resData)
	if err != nil {
		fmt.Println("数据转换失败: ", err)
		fmt.Println("data:  ", data)
		fmt.Println("resStr:  ", resStr)
		return resData, err
	}
	return resData, err

}
