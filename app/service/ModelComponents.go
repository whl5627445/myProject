package service

import (
	"strconv"
	"strings"
	"yssim-go/config"
	"yssim-go/library/omc"
)

func GetElements(className, componentName string) []interface{} {
	componentsData := omc.OMC.GetElements(className)
	var componentData []interface{}
	for i := 0; i < len(componentsData); i++ {
		cData := componentsData[i].([]interface{})
		switch {
		case componentName != "" && cData[3] == componentName:
			componentData = cData
			break
		case !(cData[5] == "protected" || cData[6] == "true" || cData[10] == "parameter") && componentName == "":
			componentData = append(componentData, cData)
		}
	}
	return componentData
}

func GetComponentName(className, componentName string) string {
	componentData := omc.OMC.GetElements(className)
	nameList := strings.Split(componentName, ".")
	name := strings.ToLower(nameList[len(nameList)-1])
	nameNum := 0
	nameMap := map[string]bool{}
	if _, ok := config.ModelicaKeywords[componentName]; ok {
		nameNum += 1
	}
	for _, c := range componentData {
		cList := c.([]interface{})
		n := cList[3].(string)
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

func AddComponentVerification(oldComponentName, newComponentName, modelNameAll string) (bool, string) {
	classInformation := omc.OMC.GetClassInformation(oldComponentName)
	classTypeAll := map[string]bool{"model": true, "class": true, "connector": true, "block": true}
	if classInformation != nil {
		classType := classInformation[0].(string)
		noType := classTypeAll[classType]
		if noType {
			return false, "不能插入：" + oldComponentName + ", 这是一个 \"" + classType + " \"类型。组件视图层只允许有model、class、connector或者block。"
		}
	}
	elementsData := omc.OMC.GetElements(modelNameAll)
	for _, e := range elementsData {
		if e.([]interface{})[3] == newComponentName {
			return false, "新增组件失败，名称 \"" + newComponentName + "\" 已经存在或是 Modelica 关键字。 请选择其他名称。"
		}
	}
	return true, ""
}

func AddComponent(oldComponentName, newComponentName, modelNameAll, origin, rotation string, extent []string) (bool, string) {
	v, msg := AddComponentVerification(oldComponentName, newComponentName, modelNameAll)
	if !v {
		return false, msg
	}
	result := omc.OMC.AddComponent(modelNameAll, newComponentName, oldComponentName, origin, rotation, extent)
	if !result {
		msg = "新增组件失败"
	}
	return result, msg
}

func DeleteComponent(componentName, modelNameAll string) bool {
	result := omc.OMC.DeleteComponent(componentName, modelNameAll)
	return result
}

func UpdateComponent(componentName, ComponentClassName, modelNameAll, origin, rotation string, extent []string) bool {
	result := omc.OMC.UpdateComponent(componentName, ComponentClassName, modelNameAll, origin, rotation, extent)
	return result
}
