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

func GetComponentName(modelName, className string) string {
	componentData := omc.OMC.GetElements(modelName)
	nameList := strings.Split(className, ".")
	name := strings.ToLower(nameList[len(nameList)-1])
	nameNum := 0
	nameMap := map[string]bool{}
	if _, ok := config.ModelicaKeywords[className]; ok {
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
		n := name + strconv.Itoa(i)
		_, ok := config.ModelicaKeywords[n]
		mapOk := nameMap[n]
		if !ok && !mapOk {
			return n
		}
	}
	return name
}

func addComponentVerification(oldComponentName, newComponentName, modelName string) (bool, string) {
	classType := omc.OMC.GetClassRestriction(oldComponentName)
	modelType := omc.OMC.GetClassRestriction(modelName)
	if modelType == "package" {
		return false, " package类型不能新增组件"
	}
	noType := config.ClassTypeAll[classType]
	if !noType {
		return false, "不能插入：" + oldComponentName + ", 这是一个 \"" + classType + " \"类型。组件视图层只允许有model、record、expandable connector、class、connector、function或者block。"
	}
	elementsData := omc.OMC.GetElements(modelName)
	for _, e := range elementsData {
		if e.([]interface{})[3] == newComponentName {
			return false, "新增组件失败，名称 \"" + newComponentName + "\" 已经存在或是 Modelica 关键字。 请选择其他名称。"
		}
	}
	return true, ""
}

func AddComponent(oldComponentName, newComponentName, modelName, origin, rotation string, extent []string) (bool, string) {
	v, msg := addComponentVerification(oldComponentName, newComponentName, modelName)
	if !v {
		return false, msg
	}
	result := omc.OMC.AddComponent(modelName, newComponentName, oldComponentName, origin, rotation, extent)
	if !result {
		msg = "新增组件失败"
	}
	return result, msg
}

func DeleteComponent(componentName, modelNameAll string) bool {
	result := omc.OMC.DeleteComponent(componentName, modelNameAll)
	return result
}

func UpdateComponent(componentName, componentClassName, modelNameAll, origin, rotation string, extent []string) bool {
	result := omc.OMC.UpdateComponent(componentName, componentClassName, modelNameAll, origin, rotation, extent)
	return result
}
