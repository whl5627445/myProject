package service

import (
	"yssim-go/config"
	"yssim-go/library/omc"
)

func renameComponentInClass(className, oldComponentName, newComponentName string) bool {
	data := omc.OMC.GetElements(className)
	renameResult := false
	_, ok := config.ModelicaKeywords[newComponentName]
	if ok {
		return false
	}
	for i := 0; i < len(data); i++ {
		if data[i].([]interface{})[3] == oldComponentName {
			renameResult = omc.OMC.RenameComponentInClass(className, oldComponentName, newComponentName)
		}
	}
	return renameResult
}

func SetComponentProperties(className, newComponentName, oldComponentName, final, protected, replaceable, variability, inner, outer, causality string) bool {
	renameResult := true
	ScpResult := true
	renameResult = renameComponentInClass(className, oldComponentName, newComponentName)
	ScpResult = omc.OMC.SetComponentProperties(className, newComponentName, final, protected, replaceable, variability, inner, outer, causality)
	if ScpResult && renameResult {

		return true
	}
	return false
}
