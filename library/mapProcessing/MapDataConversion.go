package mapProcessing

import (
	"fmt"
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
				indexNum := v.([]interface{})[1].(float64)
				resMap[k] = strconv.Itoa(int(indexNum + 1))
			}
		} else if typeArray == "bool" {
			resMap[k] = fmt.Sprintf("%t", v)
		} else {
			resMap[k] = v.(string)
		}
	}
	return resMap
}
