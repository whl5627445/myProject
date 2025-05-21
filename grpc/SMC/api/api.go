package api

import (
	"context"
	"errors"

	smc "yssim-go/grpc/SMC"
)

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

func GetModelInstance(modelName string) (*smc.ClassDefinition, error) {
	req := &smc.ModelNameRequest{
		ModelName: modelName,
	}
	result, err := smc.SMC.GetModelInstance(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return result.GetModel(), nil
}

func GetModelAST(modelName string) (*smc.ClassDefinition, error) {
	req := &smc.ModelNameRequest{
		ModelName: modelName,
	}
	result, err := smc.SMC.GetModelAST(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return result.GetModel(), nil
}

func GetModelCode(modelName string) (string, error) {
	req := &smc.ModelNameRequest{
		ModelName: modelName,
	}
	result, err := smc.SMC.GetModelCode(context.Background(), req)
	if err != nil {
		return "", err
	}

	return result.GetCode(), nil
}

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

func ParserModelCode(code, className string) (*smc.ClassDefinition, error) {
	req := &smc.ParserModelCodeRequest{
		ModelCode: code,
		ModelName: className,
	}
	result, err := smc.SMC.ParserModelCode(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return result.GetModel(), nil
}

func GetClassNames(parentName string, isAll bool) ([]string, error) {
	req := &smc.ModelNamesRequest{
		ParentName: parentName,
		IsAll:      isAll,
	}
	result, err := smc.SMC.GetClassNames(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return result.GetNameList(), nil
}

func Save(className string) (bool, error) {
	req := &smc.ModelNameRequest{
		ModelName: className,
	}
	result, err := smc.SMC.Save(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}

func GetSubTypeOf(className, parentName string) ([]string, error) {
	req := &smc.GetSubTypeOfRequest{
		ModelName:  className,
		ParentName: parentName,
	}
	result, err := smc.SMC.GetSubTypeOf(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return result.GetSubTypeOfList(), nil
}

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

func GetModelicaPath(path string) ([]string, error) {
	req := &smc.GetModelicaPathRequest{}
	result, err := smc.SMC.GetModelicaPath(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return result.GetPathList(), nil
}

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

func DeleteElement(modelName, elementName string) (bool, error) {
	req := &smc.DeleteElementRequest{
		ModelName:   modelName,
		ElementName: elementName,
	}
	result, err := smc.SMC.DeleteElement(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}

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

func DeleteConnection(modelName, connectStart, connectEnd string) (bool, error) {
	req := &smc.DeleteConnectionRequest{
		ModelName: modelName,
		Left:      connectStart,
		Right:     connectEnd,
	}
	result, err := smc.SMC.DeleteConnection(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}

func SetClassAnnotation(modelName string, annotation *smc.Argument) (bool, error) {
	req := &smc.SetClassAnnotationRequest{
		SetClassAnnotation: annotation,
		ModelName:          modelName,
	}
	result, err := smc.SMC.SetClassAnnotation(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}
