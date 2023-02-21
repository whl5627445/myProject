package service

import (
	"strings"
	"yssim-go/library/omc"
)

func GetModelChild(modelName string) []map[string]interface{} {
	childAllList := omc.OMC.GetClassNames(modelName, false)
	var dataList []map[string]interface{}
	if childAllList != nil {
		for i := 0; i < len(childAllList); i++ {
			modelChildName := modelName + "." + childAllList[i]
			classInformation := GetClassInformation(modelChildName)
			modelType := strings.TrimSpace(classInformation[0].(string))
			data := map[string]interface{}{
				"name":       childAllList[i],
				"model_name": modelChildName,
				"haschild":   false,
				"type":       modelType,
			}
			childList := omc.OMC.GetClassNames(modelChildName, false)
			if len(childList) > 0 {
				data["haschild"] = true
			}
			dataList = append(dataList, data)
		}
	}

	return dataList
}

func GetModelHasChild(modelName string) bool {
	childList := omc.OMC.GetClassNames(modelName, false)
	return len(childList) > 0
}
