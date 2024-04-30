package service

import (
	"strings"

	"yssim-go/config"
	"yssim-go/library/omc"
)

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
		if e.([]any)[3] == newComponentName {
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

func GetICList(name string) []string {
	nameList := []string{name}
	dataList := []string{name}
	for {
		var data []string
		for _, n := range nameList {
			data = append(data, omc.OMC.GetInheritedClasses(n)...)

		}
		if len(data) == 0 {
			break
		}
		dataList = append(data, dataList...)
		nameList = data
	}
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
	return dataList
}

// ModelSave 用omc提供的API将模型源码保存的到对应文件， 并发安全
func ModelSave(modelName string) {
	config.ModelCodeChan <- modelName
}

// SetPackageUses 在目标模型添加annotation，设置Uses
func SetPackageUses(fromModelName, toModelName string) {

	var fromPackageUses []string
	var toPackageUses [][]string
	// 判断是否是自身组件添加到自身库,是就不写入了
	if !strings.HasPrefix(toModelName, getFirstOrderName(fromModelName)) {
		// 获取组件所属库的版本号
		fromPackageInformation := GetClassInformation(getFirstOrderName(fromModelName))
		// 版本号为空，就不添加了
		if fromPackageInformation[14].(string) != "" {
			fromPackageUses = []string{getFirstOrderName(fromModelName), fromPackageInformation[14].(string)}
		}
		// 获取目标模型的PackageUse
		toPackageUses = GetPackageUses(getFirstOrderName(toModelName))
		// 判断toPackageUses中是否含有fromPackageUses，有就不添加了
		if len(fromPackageUses) != 0 {
			if !isExistFromUses(fromPackageUses, toPackageUses) {
				toPackageUses = append(toPackageUses, fromPackageUses)
			}
		}
		// 转成字符串
		uses := getPackageUsesString(toPackageUses)
		if uses != "" {
			omc.OMC.SetUses(getFirstOrderName(toModelName), uses)
		}
	}
}

func getFirstOrderName(modelName string) string {
	if strings.Contains(modelName, ".") {
		modelName = modelName[:strings.Index(modelName, ".")]
	}
	return modelName
}
func isExistFromUses(fromPackageUses []string, toPackageUses [][]string) bool {
	var flag = false
	for _, use := range toPackageUses {
		if fromPackageUses[0] == use[0] {
			flag = true
		}
	}
	return flag
}

func GetPackageUses(packageName string) [][]string {
	// 获取包用到的包
	return omc.OMC.GetUses(packageName)
}

func getPackageUsesString(toPackageUses [][]string) string {
	var result string
	for index, packageUse := range toPackageUses {
		result += packageUse[0] + "(version=\"" + packageUse[1] + "\")"
		if index < len(toPackageUses)-1 {
			result += ","
		}
	}
	return result
}

// GetClassInformation 获取模型信息
func GetClassInformation(modelName string) []any {
	classInformation := omc.OMC.GetClassInformation(modelName)
	return classInformation
}
