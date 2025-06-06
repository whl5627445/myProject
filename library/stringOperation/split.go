package stringOperation

import (
	"sort"
	"strconv"
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

func SliceIndexString(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}

// 获取小数字符串的精度
func GetDecimalStrPrecision(decimalStr string) int {
	parts := strings.Split(decimalStr, ".")
	if len(parts) < 2 {
		return 0
	}
	return len(parts[1])
}

// 判断字符串是否为float
func IsFloat(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}
