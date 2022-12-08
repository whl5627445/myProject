package service

import (
	"strings"
	"yssim-go/library/omc"
)

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
	result, msg := true, ""
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
