package omc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func DataToGo(data []byte) ([]interface{}, error) {
	var resData []interface{}
	Str := strings.Builder{}
	data = bytes.ReplaceAll(data, []byte("\n"), []byte("\\n"))
	data = bytes.ReplaceAll(data, []byte("\r"), []byte(""))
	data = bytes.ReplaceAll(data, []byte("\\\\\\\\"), []byte("\\"))
	data = bytes.ReplaceAll(data, []byte("\\\\\\"), []byte("\\"))
	data = bytes.ReplaceAll(data, []byte(", "), []byte(","))
	data = bytes.TrimSuffix(data, []byte(","))
	data = bytes.TrimSuffix(data, []byte("\\n"))
	if len(data) <= 5 {
		d := string(data)
		if d == "\"\"" || d == "Error" || d == "{}" || d == "[]" || d == "" {
			return resData, nil
		}
	}
	mark := false
	lData := len(data)
	for i := 0; i < lData; i++ {
		if data[i] == '"' {
			strIndex := bytes.Index(data[i+1:], []byte("\""))
			dataEndIndex := i + strIndex + 1
			str := data[i : dataEndIndex+1]
			if mark == true {
				str = bytes.ReplaceAll(str, []byte("\""), []byte("\\\""))
			}
			Str.Write(str)
			i += strIndex + 1
			continue
		}
		//if data[i] == '"' {
		//	for {
		//		strIndex := bytes.Index(data[i+1:], []byte("\""))
		//		dataEndIndex := i + strIndex + 1
		//		str := data[i:dataEndIndex]
		//		Str.Write(str)
		//		i += strIndex + 1
		//		if data[dataEndIndex-1] != byte('\\') && (data[dataEndIndex+1] == byte(',') || (data[dataEndIndex+1] == byte('}'))) || strIndex == -1 {
		//			//if mark == true {
		//			//	str = bytes.ReplaceAll(str, []byte("\""), []byte("\\\""))
		//			//}
		//			Str.Write([]byte("\""))
		//			break
		//		}
		//	}
		//	continue
		//}
		switch {
		case data[i] == '{':
			switch {
			case (i == 0 && data[i+1] == '"') || data[i+1] == '"' || data[i+1] == '{':
				mark = false
				Str.WriteString("[")
			default:
				mark = true
				Str.WriteString("[\"")
			}
		case data[i] == '}':
			switch {
			case i == lData-1 && data[i-1] == '"' || (data[i-1] == '"' && mark == false) || data[i-1] == ')' || data[i-1] == '}':
				mark = true
				Str.WriteString("]")
			default:
				mark = false
				Str.WriteString("\"]")
			}

		case data[i] == '(':
			switch {
			case i == 0 && data[i+1] != '(' && data[i+1] != '{':
				Str.WriteString("[\"")
			case data[i+1] == '(' || data[i+1] == '{' || data[i-1] == ',':
				Str.WriteString("[")
			case data[i+1] == '"' && data[i-1] != '"':
				Str.WriteString("\",[")
			default:
				Str.WriteString("\",[\"")
			}
			mark = false
		case data[i] == ')':
			switch {
			case i == lData-1 && (data[i-1] == '}' || data[i-1] == '"'):
				mark = true
				Str.WriteString("]")
			case data[i-1] != '}' && data[i-1] != '"':
				mark = false
				Str.WriteString("\"]")
			case data[i+1] == '}' || data[i+1] == ')':
				mark = true
				Str.WriteString("]")
			default:
				mark = true
				Str.WriteString("]")
			}
		case data[i] == ',':
			switch {
			case data[i-1] != '}' && data[i-1] != ')' && string(data[i+1]) != "\\" && data[i+1] != '{' && data[i+1] != '(' && data[i-1] != '"' && data[i+1] != '"':
				Str.WriteString("\",\"")
			case (data[i-1] == '}' || data[i-1] == ')' || data[i-1] == '"') && data[i+1] != '(' && data[i+1] != '{' && data[i+1] != '"' && string(data[i+1]) != "\\":
				Str.WriteString(",\"")
			case (data[i+1] == '(' || data[i+1] == '{' || data[i+1] == '"') && data[i-1] != ')' && data[i-1] != '}' && data[i-1] != '"' && string(data[i+1]) != "\\":
				Str.WriteString("\",")
			default:
				Str.WriteString(",")
			}
			mark = false
		default:
			mark = true
			Str.WriteString(string(data[i]))
		}
	}
	resStr := Str.String()
	resStr = strings.ReplaceAll(resStr, "[\"\"]", "[]")
	err := json.Unmarshal([]byte(resStr), &resData)
	if err != nil {
		fmt.Println("数据转换失败: ", err)
		fmt.Println("data:  ", string(data))
		fmt.Println("Str:  ", Str.String())
		return nil, err
	}
	return resData, err
}
