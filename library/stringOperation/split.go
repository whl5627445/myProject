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

func GetComponentType(modelType string) string {
	modelTypeList := strings.Split(modelType, " ")
	return modelTypeList[len(modelTypeList)-1]
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

func SliceIndexString(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}
