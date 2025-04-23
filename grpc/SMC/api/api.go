package api

import (
	"context"
	"errors"

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

func GetModelInstance(modelName string) (*smc.ClassDefinition, error) {
	req := &smc.ModelInstanceRequest{
		ModelName: modelName,
	}
	result, err := smc.SMC.GetModelInstance(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return result.Model, nil
}
