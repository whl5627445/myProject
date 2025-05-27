package api

import (
	"context"
	"errors"

	smc "yssim-go/grpc/SMC"
)

// LoadFile 加载指定路径的模型库或者模型
func LoadFile(path string) (string, error) {
	req := &smc.LoadFileRequest{
		Path: path,
	}
	result, err := smc.SMC.LoadFile(context.Background(), req)
	if err != nil {
		return "", err
	}
	if result.PackageName == "" {
		return "", errors.New(result.Msg)
	}
	return result.GetPackageName(), nil
}

// LoadLibrary 加载modelicaPath下的库，可以指定版本
func LoadLibrary(packageName, version string) (bool, error) {
	req := &smc.LoadLibraryRequest{
		PackageName: packageName,
		Version:     version,
	}
	result, err := smc.SMC.LoadLibrary(context.Background(), req)
	if err != nil {
		return false, err
	}
	if !result.Result {
		return false, errors.New(result.Msg)
	}
	return result.GetResult(), nil
}

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

// GetModelAST 获取模型的AST对象
func GetModelAST(className string) (*smc.ClassDefinition, error) {
	req := &smc.ClassNameRequest{
		ClassName: className,
	}
	result, err := smc.SMC.GetModelAST(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return result.GetModel(), nil
}

// GetModelCode 获取指定模型源码
func GetModelCode(className string) (string, error) {
	req := &smc.ClassNameRequest{
		ClassName: className,
	}
	result, err := smc.SMC.GetModelCode(context.Background(), req)
	if err != nil {
		return "", err
	}

	return result.GetCode(), nil
}

// ParserCode 解析给定源码字符，并指定代码的源路径文件名
func ParserCode(code, path string) (*smc.ClassDefinition, error) {
	req := &smc.ParserCodeRequest{
		Code:       code,
		SourcePath: path,
	}
	result, err := smc.SMC.ParserCode(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return result.GetModel(), nil
}

// ParserModelCode 解析模型代码，用于模型代码的更新
func ParserModelCode(code, className string) (*smc.ClassDefinition, error) {
	req := &smc.ParserModelCodeRequest{
		ClassCode: code,
		ClassName: className,
	}
	result, err := smc.SMC.ParserModelCode(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return result.GetModel(), nil
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

// Save 保存指定模型到SourcePath，更新文件内的代码
func Save(className string) (bool, error) {
	req := &smc.ClassNameRequest{
		ClassName: className,
	}
	result, err := smc.SMC.Save(context.Background(), req)
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

// SetModelicaPath 设置ModelicaPath，用于自动搜索模型库
func SetModelicaPath(path string) (bool, error) {
	req := &smc.SetModelicaPathRequest{
		Path: path,
	}
	result, err := smc.SMC.SetModelicaPath(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}

// GetModelicaPath 获取已设置的ModelicaPath
func GetModelicaPath() ([]string, error) {
	req := &smc.GetModelicaPathRequest{}
	result, err := smc.SMC.GetModelicaPath(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return result.GetPathList(), nil
}

// AddElement 向指定模型添加元素
func AddElement(e *smc.AddElement) (bool, error) {
	req := &smc.AddElementRequest{
		AddElement: e,
	}
	result, err := smc.SMC.AddElement(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}

// SetElementAnnotation 设置模型元素的注解
func SetElementAnnotation(annotation *smc.SetElementAnnotation) (bool, error) {
	req := &smc.SetElementAnnotationRequest{
		SetElementAnnotation: annotation,
	}
	result, err := smc.SMC.SetElementAnnotation(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}

// DeleteElement 删除模型元素
func DeleteElement(modelName, elementName string) (bool, error) {
	req := &smc.DeleteElementRequest{
		ClassName:   modelName,
		ElementName: elementName,
	}
	result, err := smc.SMC.DeleteElement(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}

// AddConnection 向指定模型添加连接方程
func AddConnection(c *smc.AddConnection) (bool, error) {
	req := &smc.AddConnectionRequest{
		AddConnection: c,
	}
	result, err := smc.SMC.AddConnection(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}

// SetConnectionAnnotation 设置模型方程的注解
func SetConnectionAnnotation(annotation *smc.SetConnectionAnnotation) (bool, error) {
	req := &smc.SetConnectionAnnotationRequest{
		SetConnectionAnnotation: annotation,
	}
	result, err := smc.SMC.SetConnectionAnnotation(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}

// DeleteConnection 删除模型方程
func DeleteConnection(className, connectStart, connectEnd string) (bool, error) {
	req := &smc.DeleteConnectionRequest{
		ClassName: className,
		Left:      connectStart,
		Right:     connectEnd,
	}
	result, err := smc.SMC.DeleteConnection(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}

// SetClassAnnotation 设置类注释
func SetClassAnnotation(className string, annotation *smc.Argument) (bool, error) {
	req := &smc.SetClassAnnotationRequest{
		SetClassAnnotation: annotation,
		ClassName:          className,
	}
	result, err := smc.SMC.SetClassAnnotation(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}
