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

func in(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}

func Distinct(target string, strArray *[]string) string {
	ok := in(target, *strArray)
	if ok {
		return GetLastModelName(target)
	} else {
		for _, s := range *strArray {
			if strings.HasSuffix(s, GetLastModelName(target)) {
				return GetSecondLastModelName(target)
			}
		}
		*strArray = append(*strArray, target)
		return GetLastModelName(target)
	}
}
