package mapProcessing

import (
	"fmt"
	"reflect"
	"strings"
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

func ComponentParamsToMap(componentParams []map[string]any) map[string]string {
	// 从数据库中保存的参数信息中提取出设置参数要用的key和value
	resMap := make(map[string]string)
	for i := 0; i < len(componentParams); i++ {
		componentName := componentParams[i]["name"].(string)
		componentParamsList := componentParams[i]["parameters"].([]any)
		for j := 0; j < len(componentParamsList); j++ {
			v := componentParamsList[j].(map[string]any)["value"]
			k := componentName + "." + componentParamsList[j].(map[string]any)["name"].(string)
			if v == nil {
				resMap[k] = ""
				continue
			}
			typeArray := reflect.TypeOf(v).String()
			switch typeArray {
			case "map[string]interface {}":
				parts := strings.Split(k, ".")
				parts[len(parts)-1] = "fixed"
				k_ := strings.Join(parts, ".")
				if v.(map[string]any)["value"] == nil {
					resMap[k_] = ""
					continue
				}
				resMap[k] = v.(map[string]any)["value"].(string)

				if v.(map[string]any)["isFixed"] == "" {
					resMap[k_] = ""
				}
				if v.(map[string]any)["isFixed"] == false {
					resMap[k_] = "false"
				}
				if v.(map[string]any)["isFixed"] == true {
					resMap[k_] = "true"
				}

			case "bool":
				resMap[k] = fmt.Sprintf("%t", v)
			default:
				resMap[k] = v.(string)
			}

		}
	}

	return resMap
}

// 对多个实验参数取并集，componentParamsMap中的实验参数要按照实验创建时间从小到大排序
func GetUnionComponentParams(componentParamsMap [][]map[string]any) map[string]string {

	// 并集
	unionSet := map[string]string{}

	// 取实验参数并集
	for _, componentParams := range componentParamsMap {
		resMap := ComponentParamsToMap(componentParams)
		for k, v := range resMap {
			unionSet[k] = v
		}
	}

	return unionSet
}
