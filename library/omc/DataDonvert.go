package omc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func dataToGo(data []byte) ([]interface{}, error) {
	var resData []interface{}
	data = replaceDynamicSelectData(data)
	data = dialogErrorReplace(data)
	data = iconErrorReplace(data)
	resStr := parseString(data)
	if resStr == "" {
		return []interface{}{}, nil
	}
	err := json.Unmarshal([]byte(resStr), &resData)
	if err != nil {
		log.Println("数据转换失败: ", err)
		log.Println("data:  ", string(data))
		log.Println("Str:  ", resStr)
		return nil, err
	}
	return resData, err
}

//func ParseString(data []byte) string {
//	Str := strings.Builder{}
//	data = bytes.ReplaceAll(data, []byte("\n"), []byte("\\n"))
//	data = bytes.ReplaceAll(data, []byte("\r"), []byte(""))
//	data = bytes.ReplaceAll(data, []byte("\\\\\\\\"), []byte("\\"))
//	data = bytes.ReplaceAll(data, []byte("\\\\\\"), []byte("\\"))
//	data = bytes.ReplaceAll(data, []byte(", "), []byte(","))
//	data = bytes.TrimSuffix(data, []byte(","))
//	data = bytes.TrimSuffix(data, []byte("\\n"))
//	if len(data) <= 5 {
//		d := string(data)
//		if d == "\"\"" || d == "Error" || d == "{}" || d == "[]" || d == "" {
//			return ""
//		}
//	}
//	mark := false
//	lData := len(data)
//	for i := 0; i < lData; i++ {
//		if data[i] == '"' {
//			strIndex := bytes.Index(data[i+1:], []byte("\""))
//			dataEndIndex := i + strIndex + 1
//			str := data[i : dataEndIndex+1]
//			if mark == true {
//				str = bytes.ReplaceAll(str, []byte("\""), []byte("\\\""))
//			}
//			Str.Write(str)
//			i += strIndex + 1
//			continue
//		}
//		//if data[i] == '"' {
//		//	for {
//		//		strIndex := bytes.Index(data[i+1:], []byte("\""))
//		//		dataEndIndex := i + strIndex + 1
//		//		str := data[i:dataEndIndex]
//		//		Str.Write(str)
//		//		i += strIndex + 1
//		//		if data[dataEndIndex-1] != byte('\\') && (data[dataEndIndex+1] == byte(',') || (data[dataEndIndex+1] == byte('}'))) || strIndex == -1 {
//		//			//if mark == true {
//		//			//	str = bytes.ReplaceAll(str, []byte("\""), []byte("\\\""))
//		//			//}
//		//			Str.Write([]byte("\""))
//		//			break
//		//		}
//		//	}
//		//	continue
//		//}
//		switch {
//		case data[i] == '{':
//			switch {
//			case (i == 0 && data[i+1] == '"') || data[i+1] == '"' || data[i+1] == '{':
//				mark = false
//				Str.WriteString("[")
//			default:
//				mark = true
//				Str.WriteString("[\"")
//			}
//		case data[i] == '}':
//			switch {
//			case i == lData-1 && data[i-1] == '"' || (data[i-1] == '"' && mark == false) || data[i-1] == ')' || data[i-1] == '}':
//				mark = true
//				Str.WriteString("]")
//			default:
//				mark = false
//				Str.WriteString("\"]")
//			}
//
//		case data[i] == '(':
//			switch {
//			case i == 0:
//				if data[i+1] != '(' && data[i+1] != '{' && data[i+1] != '"' {
//					Str.WriteString("[\"")
//				} else {
//					Str.WriteString("[")
//				}
//			case (data[i+1] == '(' || data[i+1] == '{' || data[i+1] == '"') && data[i-1] == ',':
//				Str.WriteString("[")
//			case (data[i+1] == '"' && data[i-1] != '"') || (data[i+1] == '{'):
//				Str.WriteString("\",[")
//			default:
//				Str.WriteString("\",[\"")
//			}
//			mark = false
//		case data[i] == ')':
//			switch {
//			case i == lData-1 && (data[i-1] == '}' || data[i-1] == '"'):
//				mark = true
//				Str.WriteString("]")
//			case data[i-1] != '}' && data[i-1] != '"' && data[i-1] != ')':
//				mark = false
//				Str.WriteString("\"]")
//			case data[i+1] == '}' || data[i+1] == ')':
//				mark = true
//				Str.WriteString("]")
//			default:
//				mark = true
//				Str.WriteString("]")
//			}
//		case data[i] == ',':
//			switch {
//			case data[i-1] != '}' && data[i-1] != ')' && string(data[i+1]) != "\\" && data[i+1] != '{' && data[i+1] != '(' && data[i-1] != '"' && data[i+1] != '"':
//				Str.WriteString("\",\"")
//			case (data[i-1] == '}' || data[i-1] == ')' || data[i-1] == '"') && data[i+1] != '(' && data[i+1] != '{' && data[i+1] != '"' && string(data[i+1]) != "\\":
//				Str.WriteString(",\"")
//			case (data[i+1] == '(' || data[i+1] == '{' || data[i+1] == '"') && data[i-1] != ')' && data[i-1] != '}' && data[i-1] != '"' && string(data[i+1]) != "\\":
//				Str.WriteString("\",")
//			default:
//				Str.WriteString(",")
//			}
//			mark = false
//		default:
//			mark = true
//			Str.WriteString(string(data[i]))
//		}
//	}
//	resStr := Str.String()
//	resStr = strings.ReplaceAll(resStr, "[\"\"]", "[]")
//	return resStr
//}

func parseString(data []byte) string {
	mark := false
	Str := strings.Builder{}
	data = bytes.ReplaceAll(data, []byte("\n"), []byte("\\n"))
	data = bytes.ReplaceAll(data, []byte("\r"), []byte(""))
	data = bytes.ReplaceAll(data, []byte("\\\\\\\\"), []byte("\\"))
	data = bytes.ReplaceAll(data, []byte("\\\\\\"), []byte("\\"))
	data = bytes.ReplaceAll(data, []byte(", "), []byte(","))
	//data = bytes.ReplaceAll(data, []byte("(\""), []byte("("))
	//data = bytes.ReplaceAll(data, []byte("\")"), []byte(")"))
	data = bytes.TrimSuffix(data, []byte(","))
	data = bytes.TrimSuffix(data, []byte("\\n"))
	if len(data) <= 5 {
		d := string(data)
		if d == "\"\"" || d == "Error" || d == "{}" || d == "[]" || d == "" {
			return ""
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
			}
			Str.Write(str)
			i += strIndex + 1
			continue
		}
		closure := false
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
			case data[i-1] == ' ':
				fmt.Println(string(data[:i+1]))
				Str.WriteString("(")
				closure = true
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
			case closure == true:
				Str.WriteString(")")
				closure = false
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
	return resStr
}

func replaceDynamicSelectData(data []byte) []byte {
	index := bytes.Index(data, []byte("DynamicSelect"))
	wordIndex := index + 13
	startIndex := index + 14
	var defaultData []byte
	var allData []byte
	if index != -1 {
		num := 0
		switch true {
		case data[startIndex] == '{':
			for i := startIndex; i < len(data); i++ {
				if data[i] == '{' {
					num += 1
				}
				if data[i] == '}' {
					num -= 1
				}
				if num == 0 {
					defaultData = data[startIndex : i+1]
					break
				}
			}
		case data[startIndex] == '(':
			for i := startIndex; i < len(data); i++ {
				if data[i] == '(' {
					num += 1
				}
				if data[i] == ')' {
					num -= 1
				}
				if num == 0 {
					defaultData = data[startIndex : i+1]
					break
				}
			}
		case data[startIndex] == '"':
			i := bytes.Index(data[startIndex+1:], []byte("\""))
			defaultData = data[startIndex : startIndex+i+2]
		default:
			i := bytes.Index(data[startIndex+1:], []byte(","))
			defaultData = data[startIndex : startIndex+i+1]
		}
		for i := wordIndex; i < len(data); i++ {
			if data[i] == '(' {
				num += 1
			}
			if data[i] == ')' {
				num -= 1
			}
			if num == 0 {
				allData = data[index : i+1]
				break
			}
		}
		data = bytes.ReplaceAll(data, allData, defaultData)
		return replaceDynamicSelectData(data)
	}
	return data
}

func dialogErrorReplace(data []byte) []byte {
	index := bytes.Index(data, []byte("{Dialog(\"error evaluating: annotation"))
	if index != -1 {
		endIndex := bytes.Index(data[index+1:], []byte("}"))
		replaceStr := data[index : index+endIndex+2]
		data = bytes.Replace(data, replaceStr, []byte("{}"), 1)
		return dialogErrorReplace(data)
	}
	return data
}

func iconErrorReplace(data []byte) []byte {
	index := bytes.Index(data, []byte("Icon(\"error evaluating: annotation("))
	if index != -1 {

		return []byte("{}")
	}
	return data
}
