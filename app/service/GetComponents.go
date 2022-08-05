package service

import "yssim-go/library/omc"

func GetComponents(className string, componentName string) []interface{} {
	componentsData := omc.OMC.GetComponents(className)
	var componentData []interface{}
	for i := 0; i < len(componentsData); i++ {
		cData := componentsData[i].([]interface{})
		switch {
		case componentName != "" && cData[1] == componentName:
			componentData = cData
			break
		case (cData[3] != "protected" || cData[4] != "True" || cData[8] != "parameter") && componentName == "":
			componentData = append(componentData, cData)
		}
	}
	return componentData
}
