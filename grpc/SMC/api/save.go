package api

import (
	"context"

	smc "yssim-go/grpc/SMC"
)

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
