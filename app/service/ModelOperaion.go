package service

import (
	"strings"
	"yssim-go/library/omc"
)

func inheritanceModelNameFixes(copiedClassName, className string) {
	classInformation := omc.OMC.GetClassRestriction(copiedClassName)
	classStrOld := omc.OMC.ListFile(className)
	classStrNew := classStrOld
	if classInformation != "model" {
		return
	}
	classNameList := strings.Split(copiedClassName, ".")
	packageName := classNameList[0]
	components := omc.OMC.GetElements(className)
	inheritedClasses := omc.OMC.GetInheritedClasses(className) // 找到父类的第一层， 只需修复这一层即可
	var inheritedNameFixesList []string
	for i := 0; i < len(inheritedClasses); i++ {
		if !strings.HasPrefix(inheritedClasses[i], packageName) && !strings.HasPrefix(inheritedClasses[i], "modelica") { // 查看父类的名称是否以包名开始， 不是，则安排修复
			inheritedNameFixesList = append(inheritedNameFixesList, inheritedClasses[i])
		}
	}
	componentNameFixesMap := map[string][]interface{}{} // 组件名称修复Map，防止重复
	for c := 0; c < len(components); c++ {
		component := components[c].([]interface{})
		cname := component[2].(string)
		if !strings.HasPrefix(cname, packageName) && !strings.HasPrefix(cname, "modelica") {
			componentNameFixesMap[cname] = component // 查看组件名称是否以包名开始， 不是，则安排修复
		}
	}

	for n := len(classNameList) - 1; n > 0; n-- {
		parentName := strings.Join(classNameList[:n], ".") // 以切割出来的模型名称为层级，逐级向上扩大查找
		modelNameAll := omc.OMC.GetClassNames(parentName, true)
		for _, name := range modelNameAll {
			for c := 0; c < len(inheritedNameFixesList); c++ {
				if strings.HasSuffix(name, "."+inheritedNameFixesList[c]) && name != inheritedNameFixesList[c] { // 如果name是以被修复的名称结尾，且不等于被修复名称自身，则被视为找到带前缀的名称
					classStrNew = strings.ReplaceAll(classStrOld, "extends "+"."+inheritedNameFixesList[c]+";", "extends "+name+";") // 查找成功，进行替换
					inheritedNameFixesList = append(inheritedNameFixesList[:c], inheritedNameFixesList[c+1:]...)                     // 替换完成，将被替换的名字移除
				}
			}
		}
		if len(inheritedNameFixesList) == 0 {
			break
		}
	}
	for n := len(classNameList) - 1; n > 0; n-- {
		parentName := strings.Join(classNameList[:n], ".")
		modelNameAll := omc.OMC.GetClassNames(parentName, true)
		for _, name := range modelNameAll {
			for k, fixesModelData := range componentNameFixesMap {
				if strings.HasSuffix(name, "."+k) && name != k {
					switch {
					case fixesModelData[10].(string) == "parameter": // 替换参数组件
						classStrNew = strings.ReplaceAll(classStrNew, "parameter "+k+" ", "parameter "+name+" ")
					case fixesModelData[9].(bool) == true: // 替换replaceable组件
						classStrNew = strings.ReplaceAll(classStrNew, "replaceable "+k+" ", "replaceable "+name+" ")
					default: // 替换普通组件
						classStrNew = strings.ReplaceAll(classStrNew, ";\n  "+k+" ", ";\n  "+name+" ")
					}
					delete(componentNameFixesMap, k)
				}
			}
		}
		if len(componentNameFixesMap) == 0 {
			break
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
		inheritanceModelNameFixes(copiedClassName, classNameAll)
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

func SaveModel(className, copiedClassName, parentName, copeOrDelete, fileName string) (bool, string) {
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
			path := omc.OMC.GetSourceFile(parentName)
			omc.OMC.SetSourceFile(parentName+"."+className, path)
			omc.OMC.Save(parentName)
		default:
			SaveModelCode(className, fileName)
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
