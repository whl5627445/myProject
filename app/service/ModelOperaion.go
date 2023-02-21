package service

import (
	"strings"
	"yssim-go/library/omc"
)

func inheritanceModelNameFixes(copiedClassName, className string) {
	classInformation := omc.OMC.GetClassRestriction(copiedClassName)
	classStrOld := omc.OMC.List(className)
	if classInformation != "model" {
		return
	}
	classNamList := strings.Split(copiedClassName, ".")
	packageName := classNamList[0]
	components := omc.OMC.GetElements(className)
	inheritedClasses := omc.OMC.GetInheritedClasses(className)
	var inheritedNameFixesList []string
	componentNameFixesMap := map[string][]interface{}{}
	for i := 0; i < len(inheritedClasses); i++ {
		if !strings.HasPrefix(inheritedClasses[i], packageName) {
			inheritedNameFixesList = append(inheritedNameFixesList, inheritedClasses[i])
		}
	}
	for c := 0; c < len(components); c++ {
		component := components[c].([]interface{})
		cname := component[2].(string)
		if !strings.HasPrefix(cname, packageName) {
			componentNameFixesMap[cname] = component
		}
	}
	classStrNew := classStrOld
	inheritedFixesMap := make(map[string]map[string]string, 0)
	for n := len(classNamList) - 1; n > 0; n-- {
		parentName := strings.Join(classNamList[:n], ".")
		modelNameAll := omc.OMC.GetClassNames(parentName, true)
		for _, name := range modelNameAll {
			for c := 0; c < len(inheritedClasses); c++ {
				if strings.HasSuffix(name, inheritedClasses[c]) && name != inheritedClasses[c] {
					inheritedFixesMap[inheritedClasses[c]] = map[string]string{"name": name}
					inheritedFixesMap[inheritedClasses[c]]["type"] = ""
					inheritedClasses = append(inheritedClasses[:c], inheritedClasses[c+1:]...)
				}
			}
		}
		if len(inheritedFixesMap) == len(inheritedClasses) {
			break
		}
	}
	for k, v := range inheritedFixesMap {
		classStrNew = strings.ReplaceAll(classStrOld, "extends "+k+";", "extends "+v["name"]+";")
	}

	componentFixesMap := make(map[string]map[string]string, 0)
	for n := len(classNamList) - 1; n > 0; n-- {
		parentName := strings.Join(classNamList[:n], ".")
		modelNameAll := omc.OMC.GetClassNames(parentName, true)
		for _, name := range modelNameAll {
			for k, fixesModelData := range componentNameFixesMap {
				fixesName := fixesModelData[2].(string)
				if strings.HasSuffix(name, fixesName) && name != k {
					componentFixesMap[fixesName] = map[string]string{"name": name}
					componentFixesMap[fixesName]["type"] = fixesModelData[10].(string)
					delete(componentNameFixesMap, k)
				}
			}
		}
		if len(componentNameFixesMap) == 0 {
			break
		}
	}
	for k, v := range componentFixesMap {
		if v["type"] != "parameter" {
			classStrNew = strings.ReplaceAll(classStrNew, ";\n  "+k+" ", ";\n  "+v["name"]+" ")

		} else {
			classStrNew = strings.ReplaceAll(classStrNew, "parameter "+k+" ", "parameter "+v["name"]+" ")

		}
	}
	result := omc.OMC.CopyLoadString(classStrNew, className)
	if !result {
		omc.OMC.CopyLoadString(classStrOld, className)
	}
	return
}

func copyModel(copiedClassName, className, parentName string) (bool, string) {
	classNameAll := className

	if parentName == "" {
		parentName = "TopLevel"
	} else {
		parentInformation := omc.OMC.GetClassRestriction(parentName)
		if parentInformation != "package" {
			return false, "模型父节点不是包类型或根节点"
		}
		classNameAll = parentName + "." + className
	}
	existClass := ExistClass(classNameAll)
	if existClass {
		return false, "模型名称已存在"
	}
	copyResult := omc.OMC.CopyClass(copiedClassName, className, parentName)
	if copyResult {
		inheritanceModelNameFixes(copiedClassName, className)
		return true, "模型复制成功"
	} else {
		return false, "模型复制失败"
	}
}

func deleteModel(className string) (bool, string) {
	existClass := ExistClass(className)
	if existClass {
		result := omc.OMC.DeleteClass(className)
		if result {
			return true, "模型删除成功"
		} else {
			return false, "模型删除失败"
		}
	} else {
		return true, "模型删除成功"
	}
}

func SaveModel(className, copiedClassName, parentName, packageName, copeOrDelete, fileName string) (bool, string) {
	result := false
	msg := ""
	switch {
	case copeOrDelete == "copy":
		result, msg = copyModel(copiedClassName, className, parentName)
	case copeOrDelete == "delete":
		result, msg = deleteModel(className)
	default:
		result = false
		msg = "未知操作"
	}
	if result {
		switch {
		case parentName != "":
			go SaveModelToFile(packageName, fileName)
		default:
			go SaveModelCode(className, fileName)
		}
	}
	return result, msg
}

func ExistClass(className string) bool {
	return omc.OMC.ExistClass(className)
}

func GetModelType(modelName string) string {
	modelType := omc.OMC.GetClassRestriction(modelName)
	modelType = strings.TrimSpace(modelType)
	return modelType
}

func GetClassInformation(modelName string) []interface{} {
	classInformation := omc.OMC.GetClassInformation(modelName)
	return classInformation
}
