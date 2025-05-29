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

// UnLoadLibrary 卸载指定名称的库
func UnLoadLibrary(packageName string) (bool, error) {
	req := &smc.UnLoadLibraryRequest{
		PackageName: packageName,
	}
	result, err := smc.SMC.UnLoadLibrary(context.Background(), req)
	if err != nil {
		return false, err
	}
	if !result.Result {
		return false, err
	}
	return result.GetResult(), nil
}
