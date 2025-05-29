package api

import (
	"context"

	smc "yssim-go/grpc/SMC"
)

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
