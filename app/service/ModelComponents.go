package service

import (
	"strconv"
	"strings"

	"yssim-go/config"
	"yssim-go/library/omc"
)

func GetElements(className, componentName string) []interface{} {
	classNameList := GetICList(className)
	var componentsData []interface{}
	for i := 0; i < len(classNameList); i++ {
		classnameData := omc.OMC.GetElements(classNameList[i])
		componentsData = append(componentsData, classnameData...)
	}

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

	// componentData = append(componentData, icnamelist)
	return componentData
}

func GetComponentName(modelName, className string) string {
	modelNameList := GetICList(modelName)
	componentsData := omc.OMC.GetElementsList(modelNameList)
	nameList := strings.Split(className, ".")
	name := strings.ToLower(nameList[len(nameList)-1])
	nameNum := 0
	nameMap := map[string]bool{}
	if _, ok := config.ModelicaKeywords[name]; ok {
		nameNum += 1
	}
	for _, c := range componentsData {
		cList := c
		n := cList[3].(string)
		if len(cList) >= 2 && strings.HasPrefix(n, name) {
			nameMap[n] = true
			nameNum += 1
		}
	}
	for i := 1; i < nameNum+1; i++ {
		n := name + strconv.Itoa(i)
		_, ok := config.ModelicaKeywords[n]
		_, mapOk := nameMap[n]
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

func AddComponent(componentName, componentClassName, modelNameAll, origin, rotation string, extent []string) (bool, string) {
	v, msg := addComponentVerification(componentClassName, componentName, modelNameAll)
	if !v {
		return false, msg
	}
	result := false
	if omc.OMC.GetClassRestriction(componentClassName) != "connector" {
		result = omc.OMC.AddComponent(componentName, componentClassName, modelNameAll, origin, rotation, extent)
	} else {
		result = omc.OMC.AddInterfacesComponent(componentName, componentClassName, modelNameAll, origin, rotation, extent)
	}
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
	result := false
	if omc.OMC.GetClassRestriction(componentClassName) != "connector" {
		result = omc.OMC.UpdateComponent(componentName, componentClassName, modelNameAll, origin, rotation, extent)
	} else {
		result = omc.OMC.UpdateInterfacesComponent(componentName, componentClassName, modelNameAll, origin, rotation, extent)
	}
	return result
}

func GetICList(name string) []string {
	nameList := []string{name}
	dataList := omc.OMC.GetInheritedClassesListAll(nameList)
	// dataList去重
	var datalistLen = len(dataList)
	for i := 0; i < datalistLen; i++ {
		for j := i + 1; j < datalistLen; j++ {
			if dataList[i] == dataList[j] {
				dataList = append(dataList[:i], dataList[i+1:]...)
				datalistLen--
				i--
				break
			}
		}
	}
	var dataListNew []string
	for i := len(dataList) - 1; i >= 0; i-- {
		dataListNew = append(dataListNew, dataList[i])
	}
	return dataListNew
}
