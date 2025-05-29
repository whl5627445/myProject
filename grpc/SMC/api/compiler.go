package api

import (
	"context"

	smc "yssim-go/grpc/SMC"
)

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
