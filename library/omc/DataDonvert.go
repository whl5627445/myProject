package omc

import (
	"bytes"
	"log"
	"strings"

	"github.com/bytedance/sonic"
)

func dataToGo(oldData []byte) ([]any, error) {
	var resData []any
	data := replaceDynamicSelectData(oldData)
	data = dialogErrorReplace(data)
	resStr := parseString(data)
	if resStr == "" {
		return []any{}, nil
	}
	err := sonic.Unmarshal([]byte(resStr), &resData)
	if err != nil {
		log.Println("数据转换失败: ", err)
		log.Println("data:  ", string(oldData))
		log.Println("Str:  ", resStr)
		return nil, err
	}
	return resData, err
}

func parseString(data []byte) string {
	mark := false
	Str := strings.Builder{}
	data = bytes.ReplaceAll(data, []byte("\n"), []byte("\\n"))
	data = bytes.ReplaceAll(data, []byte("\r"), []byte(""))
	data = bytes.ReplaceAll(data, []byte("\\\\\\\\"), []byte("\\"))
	data = bytes.ReplaceAll(data, []byte("\\\\\\"), []byte("\\"))
	data = bytes.ReplaceAll(data, []byte(", "), []byte(","))
	data = bytes.ReplaceAll(data, []byte("{extent={{"), []byte("{extent{{"))
	data = bytes.TrimSuffix(data, []byte(","))
	data = bytes.TrimSuffix(data, []byte(" "))
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
				if i != 0 && data[i+1] == '{' && data[i-1] != '"' && data[i-1] != ',' {
					Str.WriteString("\",[")
					mark = false
				} else {
					Str.WriteString("[")
				}
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
				num := 1
				for p := i + 1; p < len(data); p++ {
					if data[p] == '(' {
						num += 1
					}
					if data[p] == ')' {
						num -= 1
					}
					if data[p] != '"' {
						Str.WriteString(string(data[p]))
					} else {
						Str.WriteString("\\\"")
					}
					if num == 0 {
						if p < len(data)-1 && data[p+1] != '"' {
							Str.WriteString("\"")
							mark = false
						}
						i = p
						break
					}
				}
			case (data[i+1] == '(' || data[i+1] == '{' || data[i+1] == '"') && data[i-1] == ',':
				Str.WriteString("[")
			case (data[i+1] == '"' && data[i-1] != '"') || (data[i+1] == '{'):
				Str.WriteString("\",[")
				mark = false
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
	index := bytes.Index(data, []byte("error evaluating: annotation"))
	wordIndex := index + 28
	var allData []byte
	if index != -1 {
		num := 0
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
		data = bytes.Replace(data, allData, []byte(""), 1)
		return dialogErrorReplace(data)
	}
	return data
}
