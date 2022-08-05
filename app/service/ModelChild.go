package service

import (
	"strings"
	"yssim-go/library/omc"
)

func GetModelChild(modelName string) []map[string]interface{} {

	childAllList := omc.OMC.GetClassNames(modelName, true)
	var dataList []map[string]interface{}
	if childAllList != nil {
		modelLen := len(strings.Split(modelName, "."))
		childMap := map[string]map[string]any{}
		var childList []string
		for i := 0; i < len(childAllList); i++ {
			PrefixOK := strings.HasPrefix(childAllList[i], modelName)
			modelChildLen := len(strings.Split(childAllList[i], "."))
			if PrefixOK == true && modelChildLen != modelLen {
				modelNameList := strings.Split(childAllList[i], ".")
				modelNameAll := strings.Join(modelNameList[:modelLen+1], ".")
				modelName := modelNameList[len(modelNameList)-1]
				data := map[string]interface{}{
					"model_name":     modelName,
					"model_name_all": modelNameAll,
					"haschild":       false,
				}
				_, ok := childMap[modelNameAll]
				if ok && modelChildLen > modelLen+1 {
					childMap[modelNameAll]["haschild"] = true
				}
				if modelChildLen == modelLen+1 && !ok {
					childMap[modelNameAll] = data
					childList = append(childList, modelNameAll)
				}
			}
		}
		for _, c := range childList {
			dataList = append(dataList, childMap[c])
		}
	}
	return dataList
}

func GetModelHasChild(modelName string) bool {
	childList := omc.OMC.GetClassNames(modelName, false)
	if len(childList) > 0 {
		return true
	}
	return false
}
