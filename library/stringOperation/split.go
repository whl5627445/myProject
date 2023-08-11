package stringOperation

import (
	"sort"
	"strings"
)

func PluralSplit(s string, sep []string) []string {
	strData := s
	for i := 0; i < len(sep); i++ {
		strData = strings.ReplaceAll(strData, sep[i], "0x86")
	}
	return strings.Split(strData, "0x86")
}

// GetLastModelName 获取Modelica.Blocks.Examples.PID_Controller中PID_Controller
func GetLastModelName(className string) string {
	return className[strings.LastIndex(className, ".")+1:]
}

func GetSecondLastModelName(className string) string {
	name := strings.Split(className, ".")
	return strings.Join(name[len(name)-2:], ".")
}

func GetThirdLastModelName(className string) string {
	name := strings.Split(className, ".")
	return strings.Join(name[len(name)-3:], ".")
}

func GetComponentType(modelType string) string {
	modelTypeList := strings.Split(modelType, " ")
	return modelTypeList[len(modelTypeList)-1]
}

func in(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}

func Distinct(target string, strArray, secondStrArray *[]string) string {
	ok := in(target, *strArray)
	if ok {
		return GetLastModelName(target)
	} else {
		for _, s := range *strArray {
			if strings.HasSuffix(s, GetLastModelName(target)) {
				ok = in(target, *secondStrArray)
				if ok {
					return GetSecondLastModelName(target)
				} else {
					for _, s2 := range *secondStrArray {
						if strings.HasSuffix(s2, GetSecondLastModelName(target)) {
							return GetThirdLastModelName(target)
						}
					}
				}
				*secondStrArray = append(*secondStrArray, target)
				return GetSecondLastModelName(target)
			}
		}
		*strArray = append(*strArray, target)
		return GetLastModelName(target)
	}
}

func ContainsString(target string) bool {
	collection := []string{"function", "package"}
	for _, str := range collection {
		if str == target {
			return true
		}
	}
	return false
}
