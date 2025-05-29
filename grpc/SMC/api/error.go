package api

import (
	"context"

	smc "yssim-go/grpc/SMC"
)

// GetErrors 获取当前错误信息以及提示信息
func GetErrors() ([]*smc.ErrorMessage, error) {
	req := &smc.GetErrorsRequest{}
	result, err := smc.SMC.GetErrors(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return result.GetMessageList(), nil
}
