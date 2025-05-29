package api

import (
	"context"

	smc "yssim-go/grpc/SMC"
)

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

// SetElementComment 设置模型元素的描述
func SetElementComment(className, elementName, comment string) (bool, error) {
	req := &smc.SetElementCommentRequest{
		ClassName:   className,
		ElementName: elementName,
		Comment:     comment,
	}
	result, err := smc.SMC.SetElementComment(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}

// SetElementVariability 设置模型元素的 variability
func SetElementVariability(className, elementName, variability string) (bool, error) {
	req := &smc.SetElementVariabilityRequest{
		ClassName:   className,
		ElementName: elementName,
		Variability: variability,
	}
	result, err := smc.SMC.SetElementVariability(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}

// SetElementCausality 设置模型元素的 causality
func SetElementCausality(className, elementName, causality string) (bool, error) {
	req := &smc.SetElementCausalityRequest{
		ClassName:   className,
		ElementName: elementName,
		Causality:   causality,
	}
	result, err := smc.SMC.SetElementCausality(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}

// SetElementProperties 设置模型元素的 properties
func SetElementProperties(className, elementName string, replaceable, final, protected bool) (bool, error) {
	req := &smc.SetElementPropertiesRequest{
		ClassName:   className,
		ElementName: elementName,
		Replaceable: replaceable,
		Final:       final,
		Protected:   protected,
	}
	result, err := smc.SMC.SetElementProperties(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}

// SetElementInnerAndOuter 设置模型元素的 inner 和 outer
func SetElementInnerAndOuter(className, elementName string, inner, outer bool) (bool, error) {
	req := &smc.SetElementInnerAndOuterRequest{
		ClassName:   className,
		ElementName: elementName,
		Inner:       inner,
		Outer:       outer,
	}
	result, err := smc.SMC.SetElementInnerAndOuter(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}

// SetElementDimensions 设置模型元素的 dimensions
func SetElementDimensions(className, elementName, dimensions string) (bool, error) {
	req := &smc.SetElementDimensionsRequest{
		ClassName:   className,
		ElementName: elementName,
		Dimensions:  dimensions,
	}
	result, err := smc.SMC.SetElementDimensions(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}

// SetElementModifier 设置模型元素的 modifier
func SetElementModifier(className, elementName string, modifierList []*smc.Argument) (bool, error) {
	req := &smc.SetElementModifierRequest{
		ClassName:    className,
		ElementName:  elementName,
		ModifierList: modifierList,
	}
	result, err := smc.SMC.SetElementModifier(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}

// SetExtendsModifier 设置模型指定父类的 modifier
func SetExtendsModifier(className, typeSpecifier string, modifierList []*smc.ArgumentOrInheritanceModification) (bool, error) {
	req := &smc.SetExtendsModifierRequest{
		ClassName:     className,
		TypeSpecifier: typeSpecifier,
		ModifierList:  modifierList,
	}
	result, err := smc.SMC.SetExtendsModifier(context.Background(), req)
	if err != nil {
		return false, err
	}
	return result.GetResult(), nil
}
