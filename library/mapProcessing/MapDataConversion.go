package mapProcessing

import (
	"reflect"
	"strconv"
)

func MapDataConversion(m map[string]interface{}, mode string) map[string]string {
	resMap := make(map[string]string)
	for k, v := range m {
		typeArray := reflect.TypeOf(v).String()
		if typeArray == "[]interface {}" {
			if mode == "om" {
				resMap[k] = v.([]interface{})[0].(string)
			} else {
				indexNumStr := v.([]interface{})[1].(string)
				indexNum, _ := strconv.Atoi(indexNumStr)
				resMap[k] = strconv.Itoa(indexNum + 1)
			}
		} else {
			resMap[k] = v.(string)
		}
	}
	return resMap
}
