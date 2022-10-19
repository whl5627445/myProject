package service

import (
	"yssim-go/config"
	"yssim-go/library/omc"
)

func renameComponentInClass(className, oldComponentName, newComponentName string) (bool, string) {
	data := omc.OMC.GetElements(className)
	_, ok := config.ModelicaKeywords[newComponentName]
	if ok {
		return false, "名称为关键字，请更换另一个名称"
	}
	for i := 0; i < len(data); i++ {
		if data[i].([]interface{})[3].(string) == oldComponentName {
			_ = omc.OMC.RenameComponentInClass(className, oldComponentName, newComponentName)

			return true, ""
		}
	}
	return false, "设置失败，请检查是否修改了继承模型的组件"
}

func SetComponentProperties(className, newComponentName, oldComponentName, final, protected, replaceable, variability, inner, outer, causality string) (bool, string) {
	renameResult, msg := renameComponentInClass(className, oldComponentName, newComponentName)
	ScpResult := omc.OMC.SetComponentProperties(className, newComponentName, final, protected, replaceable, variability, inner, outer, causality)
	if ScpResult && renameResult {
		return true, msg
	}
	return false, msg
}
