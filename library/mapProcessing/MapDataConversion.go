package mapProcessing

import (
	"fmt"
	"reflect"
)

//func MapDataConversion(m map[string]any, mode string) map[string]string {
//	resMap := make(map[string]string)
//	for k, v := range m {
//		typeArray := reflect.TypeOf(v).String()
//		if typeArray == "[]interface {}" {
//			if mode == "OM" {
//				resMap[k] = v.([]any)[0].(string)
//			} else {
//				indexNum := v.([]any)[1].(float64)
//				resMap[k] = strconv.Itoa(int(indexNum + 1))
//			}
//		} else if typeArray == "bool" {
//			resMap[k] = fmt.Sprintf("%t", v)
//		} else {
//			resMap[k] = v.(string)
//		}
//	}
//	return resMap
//}

func MapDataConversion(m map[string]any) map[string]string {
	resMap := make(map[string]string)
	for k, v := range m {
		typeArray := reflect.TypeOf(v).String()
		switch typeArray {
		case "[]interface {}":
			resMap[k] = v.([]any)[0].(string)
		case "bool":
			resMap[k] = fmt.Sprintf("%t", v)
		default:
			resMap[k] = v.(string)
		}
	}
	return resMap
}
