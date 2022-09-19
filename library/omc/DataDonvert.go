package omc

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"
)

func DataToGo(data []byte) ([]interface{}, error) {
	var resData []interface{}
	resStr, mark := ParseStringTest(data)
	if resStr == "" {
		return []interface{}{}, nil
	}
	err := json.Unmarshal([]byte(resStr), &resData)
	if err != nil {
		log.Println("数据转换失败: ", err)
		log.Println("data:  ", string(data))
		log.Println("Str:  ", resStr)
		log.Println("mark:  ", mark)
		return nil, err
	}
	return resData, err
}

func ParseString(data []byte) string {
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
			return ""
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
			case i == 0:
				if data[i+1] != '(' && data[i+1] != '{' && data[i+1] != '"' {
					Str.WriteString("[\"")
				} else {
					Str.WriteString("[")
				}
			case (data[i+1] == '(' || data[i+1] == '{' || data[i+1] == '"') && data[i-1] == ',':
				Str.WriteString("[")
			case (data[i+1] == '"' && data[i-1] != '"') || (data[i+1] == '{'):
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
			case data[i-1] != '}' && data[i-1] != '"' && data[i-1] != ')':
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
	return resStr
}

func ParseStringTest(data []byte) (string, bool) {
	mark := false
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
			return "", false
		}
	}

	lData := len(data)
	for i := 0; i < lData; i++ {
		if data[i] == '"' {
			strIndex := bytes.Index(data[i+1:], []byte("\""))
			dataEndIndex := i + strIndex + 1
			str := data[i : dataEndIndex+1]
			if mark == true {
				if data[i-1] != byte('\\') {
					str = bytes.ReplaceAll(str, []byte("\""), []byte("\\\""))
				}
				if string(data[dataEndIndex+1]) == "," {
					str = append(str, '"')
				}
				//mark = false
			}
			Str.Write(str)
			i += strIndex + 1
			continue
		}
		switch {
		case data[i] == '{':
			switch {
			case i > 0 && data[i-1] == ' ':
				if data[i+1] != '"' && data[i+1] != '{' && data[i+1] != '(' {
					Str.WriteString("\",[\"")
					mark = true
				} else {
					Str.WriteString("\",[")
					mark = false
				}
			case i == 0 && data[i+1] == '"' || data[i+1] == '"' || data[i+1] == '{' || (i != 0 && data[i+1] == '{' && data[i-1] != ' '):
				Str.WriteString("[")
			default:
				mark = true
				Str.WriteString("[\"")
			}

		case data[i] == '}':
			switch {
			case i < lData-1 && data[i+1] == ' ':
				if data[i-1] != '"' && data[i-1] != '}' && data[i-1] != ')' {
					Str.WriteString("\"],\"")
				} else {
					Str.WriteString("],\"")
				}
				mark = true
			case i == lData-1 && data[i-1] == '"' || (data[i-1] == '"' && mark == false) || data[i-1] == ')' || data[i-1] == '}' || (i > 1 && data[i-1] == '"' && data[i-2] == '"' && data[i-3] == '\\'):
				Str.WriteString("]")
			default:
				mark = false
				Str.WriteString("\"]")
			}

		case data[i] == '(':
			switch {
			case i == 0:
				if data[i+1] != '(' && data[i+1] != '{' && data[i+1] != '"' {
					Str.WriteString("[\"")
					mark = true
				} else {
					Str.WriteString("[")
				}
			case (data[i+1] == '(' || data[i+1] == '{' || data[i+1] == '"') && data[i-1] == ',':
				Str.WriteString("[")
			case (data[i+1] == '"' && data[i-1] != '"') || (data[i+1] == '{'):
				Str.WriteString("\",[")
				//if data[i+1] != '"' {
				//	mark = false
				//}
				mark = false
			case data[i+1] != '(' && data[i+1] != '{' && data[i+1] != '"' && data[i-1] == ' ':
				Str.WriteString("[\"")
				mark = true
			//case data[i-1] == ' ':
			//	Str.WriteString("wang")
			default:
				Str.WriteString("\",[\"")
				mark = true
			}
		case data[i] == ')':
			switch {
			case i == lData-1 && (data[i-1] == '}' || data[i-1] == '"'):
				Str.WriteString("]")
			case data[i-1] != '}' && data[i-1] != '"' && data[i-1] != ')':
				if i < lData-1 && data[i+1] == ' ' {
					Str.WriteString("\"],\"")
					mark = true
				} else {
					Str.WriteString("\"]")
					mark = false
				}
			case data[i+1] == '}' || data[i+1] == ')':
				Str.WriteString("]")

			default:
				Str.WriteString("]")
			}
		case data[i] == ',':
			switch {
			case data[i-1] != '}' && data[i-1] != ')' && data[i+1] != '\\' && data[i+1] != '{' && data[i+1] != '(' && (data[i-1] != '"' && mark == true) && data[i+1] != '"' && data[i+1] != ' ':
				Str.WriteString("\",\"")
				mark = true
			case (data[i-1] == '}' || data[i-1] == ')' || data[i-1] == '"') && data[i+1] != '(' && data[i+1] != '{' && data[i+1] != '"' && string(data[i+1]) != "\\":
				Str.WriteString(",\"")
				mark = true
			case (data[i+1] == '(' || data[i+1] == '{' || data[i+1] == '"') && data[i-1] != ')' && data[i-1] != '}' && data[i-1] != '"' && string(data[i+1]) != "\\":
				Str.WriteString("\",")
				mark = false
			default:
				Str.WriteString(",")
			}
		default:
			Str.WriteString(string(data[i]))
		}
	}
	resStr := Str.String()
	resStr = strings.ReplaceAll(resStr, "[\"\"]", "[]")
	return resStr, mark
}
