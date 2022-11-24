package mapProcessing

import "reflect"

func MapInterfaceProcessing(m map[string]interface{}, mode string) map[string]string {
	resMap := make(map[string]string)
	for k, v := range m {
		typeArray := reflect.TypeOf(m[k]).String()
		if typeArray == "[2]string" {
			if mode == "om" {
				resMap[k] = m[k].([2]string)[0]
			} else {
				resMap[k] = m[k].([2]string)[1]
			}
		} else {
			resMap[k] = v.(string)
		}
	}
	return resMap
}
