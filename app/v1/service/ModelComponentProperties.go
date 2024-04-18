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
		if data[i].([]any)[3].(string) == newComponentName && data[i].([]any)[3].(string) != oldComponentName && oldComponentName != newComponentName {
			return false, "名称重复，请更换另一个名称"
		}
	}
	for i := 0; i < len(data); i++ {
		if data[i].([]any)[3].(string) == oldComponentName {
			_ = omc.OMC.RenameComponentInClass(className, oldComponentName, newComponentName)
			return true, ""
		}
	}
	return false, "设置失败，请检查是否修改了继承模型的组件"
}

func SetComponentProperties(className, newComponentName, oldComponentName, final, protected, replaceable, variability, inner, outer, causality, comment, dimensions string) (bool, []string) {
	var msgList []string
	renameResult, iMsg := renameComponentInClass(className, oldComponentName, newComponentName)
	scpResult, pMsg := omc.OMC.SetComponentProperties(className, newComponentName, final, protected, replaceable, variability, inner, outer, causality)
	cms, cMsg := omc.OMC.SetComponentComment(className, newComponentName, comment)
	dms, dMsg := omc.OMC.SetComponentDimensions(className, newComponentName, dimensions)
	if scpResult && renameResult && cms && dms {
		return true, msgList
	} else {
		msgList = removeEmptyString(iMsg, pMsg, cMsg, dMsg)
	}
	return false, msgList
}

func removeEmptyString(strings ...string) []string {
	var list []string
	for _, str := range strings {
		if str != "" {
			list = append(list, str)
		}
	}
	return list
}
