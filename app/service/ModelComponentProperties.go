package service

import (
	"yssim-go/config"
	"yssim-go/library/omc"
)

func RenameComponentInClass(className string, oldComponentName string, newComponentName string) bool {
	data := omc.OMC.GetElements(className)
	for i := 0; i < len(data); i++ {
		_, ok := config.ModelicaKeywords[newComponentName]
		if data[i].([]interface{})[3] == newComponentName || ok {
			return false
		}
	}
	renameResult := omc.OMC.RenameComponentInClass(className, oldComponentName, newComponentName)
	return renameResult
}

func SetComponentProperties(className string, newComponentName string, oldComponentName string, final string, protected string, replaceable string, variability string, inner string, outer string, causality string) bool {
	renameResult := true
	ScpResult := true
	if oldComponentName != newComponentName {
		renameResult = RenameComponentInClass(className, oldComponentName, newComponentName)

	}
	ScpResult = omc.OMC.SetComponentProperties(className, newComponentName, final, protected, replaceable, variability, inner, outer, causality)
	if ScpResult && renameResult {

		return true
	}
	return false
}
