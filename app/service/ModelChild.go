package service

import (
	"strings"
	"yssim-go/library/omc"
)

type modelChild struct {
	Name      string `json:"name,omitempty"`
	ModelName string `json:"model_name,omitempty"`
	HasChild  bool   `json:"haschild,omitempty"`
	Type      string `json:"type,omitempty"`
}

func GetModelChild(modelName string) []modelChild {
	childAllList := omc.OMC.GetClassNames(modelName, false)
	var dataList []modelChild
	if childAllList != nil {
		for i := 0; i < len(childAllList); i++ {
			modelChildName := modelName + "." + childAllList[i]
			classInformation := GetClassInformation(modelChildName)
			modelType := strings.TrimSpace(classInformation[0].(string))
			data := modelChild{
				Name:      childAllList[i],
				ModelName: modelChildName,
				HasChild:  false,
				Type:      modelType,
			}
			childList := omc.OMC.GetClassNames(modelChildName, false)
			if len(childList) > 0 {
				data.HasChild = true
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

func GetCalibrationModelChild(modelName string) []modelChild {
	childAllList := omc.OMC.GetClassNames(modelName, false)
	var dataList []modelChild
	if childAllList != nil {
		for i := 0; i < len(childAllList); i++ {
			modelChildName := modelName + "." + childAllList[i]
			classInformation := GetClassInformation(modelChildName)
			modelType := strings.TrimSpace(classInformation[0].(string))
			data := modelChild{
				Name:      childAllList[i],
				ModelName: modelChildName,
				HasChild:  false,
				Type:      modelType,
			}
			childList := omc.OMC.GetClassNames(modelChildName, false)
			for _, c := range childList {
				t := GetModelType(c)
				if t == "model" {
					data.HasChild = true
					break
				}
			}
			dataList = append(dataList, data)
		}
	}

	return dataList
}
