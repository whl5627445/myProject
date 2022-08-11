package service

import (
	"strconv"
	"strings"
	"yssim-go/config"
	"yssim-go/library/omc"
)

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

func GetComponentName(className string, componentName string) string {
	componentData := omc.OMC.GetComponents(className)
	nameList := strings.Split(componentName, ".")
	name := strings.ToLower(nameList[len(nameList)-1])
	nameNum := 0
	nameMap := map[string]bool{}
	if _, ok := config.ModelicaKeywords[componentName]; ok {
		nameNum += 1
	}
	for _, c := range componentData {
		cList := c.([]interface{})
		n := cList[1].(string)
		if len(cList) >= 2 && strings.HasPrefix(n, name) {
			nameMap[n] = true
			nameNum += 1
		}
	}
	for i := 1; i < nameNum+1; i++ {
		n := componentName + strconv.Itoa(i)
		if _, ok := config.ModelicaKeywords[n]; !ok {
			return n
		}
	}
	return name
}

func AddComponentVerification(oldComponentName string, newComponentName string, modelNameAll string) (bool, string) {
	return true, ""
}

func AddComponent(oldComponentName string, newComponentName string, modelNameAll string) (bool, string) {
	return true, ""
}

func DeleteComponent(oldComponentName string, newComponentName string, modelNameAll string) bool {
	return true
}

func UpdateComponent(oldComponentName string, newComponentName string, modelNameAll string) bool {
	return true
}
