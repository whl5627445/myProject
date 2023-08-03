package mapProcessing

import (
	"fmt"
	"reflect"
	"yssim-go/app/DataType"
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

func IsExistKey(resultData []map[string]*DataType.GetUMLData, className string, rootExtendsModelData DataType.ExtendsModelData) bool {

	for _, m := range resultData {
		if value, ok := m[className]; ok {
			extendsModelList := value.ExtendsModelData
			flag := false
			index := 0
			for i := 0; i < len(extendsModelList); i++ {
				if extendsModelList[i].ClassName == rootExtendsModelData.ClassName {
					flag = true
					index = i
				}
			}
			if flag {
				extendsModelList[index].Count = extendsModelList[index].Count + 1
			} else {
				value.ExtendsModelData = append(value.ExtendsModelData, rootExtendsModelData)
			}
			return ok
		}
	}
	return false
}
