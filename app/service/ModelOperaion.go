package service

import (
	"yssim-go/library/omc"
)

func copyModel(copiedClassName, className, parentName string) (bool, string) {

	if parentName == "" {
		parentName = "TopLevel"
	} else {
		parentInformation := omc.OMC.GetClassInformation(parentName)
		if len(parentInformation) > 0 && parentInformation[0].(string) != "package" {
			return false, "模型父节点不是包类型或根节点"
		}
	}
	existClass := omc.OMC.ExistClass(parentName + "." + className)
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
	existClass := omc.OMC.ExistClass(className)
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
		ok := true
		if parentName != "" {
			ok = SaveModelCode(packageName, fileName)
		} else {
			ok = SaveModelCode(className, fileName)
		}
		if !ok {
			result = false
			msg = "未知操作"
		}
	}
	return result, msg
}