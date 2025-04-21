package api

import (
	"context"
	"errors"

	"github.com/bytedance/sonic"
	smc "yssim-go/grpc/SMC"
)

func LoadFile(path string) (bool, error) {
	req := &smc.LoadFileRequest{
		Path: path,
	}
	result, err := smc.SMC.LoadFile(context.Background(), req)
	if err != nil {
		return false, err
	}
	if !result.Result {
		return false, errors.New(result.Msg)
	}
	return true, nil
}

func GetModelInstance(modelName string) (*ClassDefinition, error) {
	req := &smc.ModelInstanceRequest{
		ModelName: modelName,
	}
	result, err := smc.SMC.GetModelInstance(context.Background(), req)
	if err != nil {
		return nil, err
	}
	m := &ClassDefinition{}
	err = sonic.Unmarshal(result.Model, m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
