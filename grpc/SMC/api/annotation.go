package api

import (
	"context"

	smc "yssim-go/grpc/SMC"
)

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

// GetLoadLibrariesAndVersion 获取已加载的库和版本
func GetLoadLibrariesAndVersion() (map[string]string, error) {
	req := &smc.LoadLibrariesAndVersionRequest{}
	result, err := smc.SMC.GetLoadLibrariesAndVersion(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return result.GetLibrariesAndVersion(), nil
}

// GetUses 获取类的依赖项
func GetUses(className string) ([]*smc.Argument, error) {
	req := &smc.GetUsesRequest{ClassName: className}
	result, err := smc.SMC.GetUses(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return result.GetUses(), nil
}

// SetUses 设置类的依赖项
func SetUses(className string, uses []*smc.Argument) (bool, error) {
	req := &smc.SetUsesRequest{ClassName: className, Uses: uses}
	result, err := smc.SMC.SetUses(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}

// GetDiagramAnnotation 获取类注解中的图表层数据
func GetDiagramAnnotation(className string) ([]*smc.Argument, error) {
	req := &smc.GetDiagramAnnotationRequest{ClassName: className}
	result, err := smc.SMC.GetDiagramAnnotation(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return result.GetDiagram(), nil
}

// GetIconAnnotation 获取类注解中的图标层数据
func GetIconAnnotation(className string) ([]*smc.Argument, error) {
	req := &smc.GetIconAnnotationRequest{ClassName: className}
	result, err := smc.SMC.GetIconAnnotation(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return result.GetIcon(), nil
}

// GetDefaultComponentName 获取类的默认组件名称
func GetDefaultComponentName(className string) (string, error) {
	req := &smc.GetDefaultComponentNameRequest{ClassName: className}
	result, err := smc.SMC.GetDefaultComponentName(context.Background(), req)
	if err != nil {
		return "", err
	}
	return result.GetName(), nil
}

// GetClassDocumentation 获取类的文档
func GetClassDocumentation(className string) (*smc.Argument, error) {
	req := &smc.GetClassDocumentationRequest{ClassName: className}
	result, err := smc.SMC.GetClassDocumentation(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return result.GetDocument(), nil
}
