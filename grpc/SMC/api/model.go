package api

import (
	"context"

	smc "yssim-go/grpc/SMC"
)

// GetModelInstance 获取模型实例接口
func GetModelInstance(className string) (*smc.ClassDefinition, error) {
	req := &smc.ClassNameRequest{
		ClassName: className,
	}
	result, err := smc.SMC.GetModelInstance(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return result.GetClass(), nil
}

// GetClassNames 获取指定模型名称下的子模型名称，可指定是否递归查询，返回子模型名称列表
func GetClassNames(parentName string, isAll bool) ([]string, error) {
	req := &smc.ClassNamesRequest{
		ParentName: parentName,
		IsAll:      isAll,
	}
	result, err := smc.SMC.GetClassNames(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return result.GetNameList(), nil
}

// RenameComponentInClass 重命名指定模型下的组件
func RenameComponentInClass(className, oldComponentName, newComponentName string) (bool, error) {
	req := &smc.RenameComponentInClassRequest{
		ClassName: className,
		OldName:   oldComponentName,
		NewName:   newComponentName,
	}
	result, err := smc.SMC.RenameComponentInClass(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}

// GetSubTypeOf 获取指定名称的子类型，可指定筛选继承的父模型
func GetSubTypeOf(className, parentName string) ([]string, error) {
	req := &smc.GetSubTypeOfRequest{
		ClassName:  className,
		ParentName: parentName,
	}
	result, err := smc.SMC.GetSubTypeOf(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return result.GetSubTypeOfList(), nil
}

// ExistClass 判断模型是否存在
func ExistClass(className string) (bool, error) {
	req := &smc.ExistClassRequest{
		ClassName: className,
	}
	result, err := smc.SMC.ExistClass(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}

// CopyClass 复制模型
func CopyClass(name, className, parentName string) (bool, error) {
	req := &smc.CopyClassRequest{
		Name:       name,
		ClassName:  className,
		ParentName: parentName,
	}
	result, err := smc.SMC.CopyClass(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}

// GetInheritedClasses 获取指定模型父类的模型名称列表
func GetInheritedClasses(className string) ([]string, error) {
	req := &smc.GetInheritedClassesRequest{
		ClassName: className,
	}
	result, err := smc.SMC.GetInheritedClasses(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return result.GetNameList(), nil
}

// IsPackage 判断模型是否为包
func IsPackage(className string) (bool, error) {
	req := &smc.IsPackageRequest{
		ClassName: className,
	}
	result, err := smc.SMC.IsPackage(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}

// DeleteClass 删除指定模型
func DeleteClass(className string) (bool, error) {
	req := &smc.DeleteClassRequest{
		ClassName: className,
	}
	result, err := smc.SMC.DeleteClass(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}

// GetClassRestriction 获取指定模型名称的模型类型
func GetClassRestriction(className string) (string, error) {
	req := &smc.GetClassRestrictionRequest{
		ClassName: className,
	}
	result, err := smc.SMC.GetClassRestriction(context.Background(), req)
	if err != nil {
		return "", err
	}
	return result.GetRestriction(), nil
}

// GetSourceFile 获取指定模型名称的源文件路径
func GetSourceFile(className string) (string, error) {
	req := &smc.GetSourceFileRequest{
		ClassName: className,
	}
	result, err := smc.SMC.GetSourceFile(context.Background(), req)
	if err != nil {
		return "", err
	}
	return result.GetPath(), nil
}

// SetSourceFile 设置指定模型名称的源文件路径
func SetSourceFile(className, path string) (bool, error) {
	req := &smc.SetSourceFileRequest{
		ClassName: className,
		Path:      path,
	}
	result, err := smc.SMC.SetSourceFile(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}
