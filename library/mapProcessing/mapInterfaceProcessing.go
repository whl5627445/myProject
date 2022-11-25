package mapProcessing

import (
	"reflect"
	"strconv"
)

func MapInterfaceProcessing(m map[string]interface{}, mode string) map[string]string {
	resMap := make(map[string]string)
	for k, v := range m {
		typeArray := reflect.TypeOf(m[k]).String()
		if typeArray == "[2]string" {
			if mode == "om" {
				resMap[k] = m[k].([2]string)[0]
			} else {
				indexNumStr := m[k].([2]string)[1]
				indexNum, _ := strconv.Atoi(indexNumStr)
				resMap[k] = strconv.Itoa(indexNum + 1)
			}
		} else {
			resMap[k] = v.(string)
		}
	}
	return resMap
}
