package grpcOmc

import (
	"fmt"
	"yssim-go/grpc/GoOmcGrpc/grpcInterface"
)

func LoadFile(path string) bool {
	requestData := &grpcInterface.LoadFileRequest{
		PackagePath: path,
	}
	replyData, err := Client.LoadFile(Ctx, requestData)
	if err != nil {
		fmt.Println("调用grpc服务(LoadFile)出错：", err)
		return false
	} else {
		return replyData.LoadFileRes
	}
}

func GetPackageVersion(PackageName string) string {
	requestData := &grpcInterface.GetPackageVersionRequest{
		PackageName: PackageName,
	}
	replyData, err := Client.GetPackageVersion(Ctx, requestData)
	if err != nil {
		fmt.Println("调用grpc服务(GitPackageVersion)出错：", err)
		return ""
	} else {
		return replyData.Version
	}
}

func ParseFile(filePath string) (string, bool) {
	requestData := &grpcInterface.ParseFileRequest{
		FilePath: filePath,
	}
	replyData, err := Client.ParseFile(Ctx, requestData)
	if err != nil {
		fmt.Println("调用grpc服务(ParseFile)出错：", err)
		return "", false
	} else {
		return replyData.PackageName, replyData.ParseRes
	}
}

func DeleteClass(packageName string) bool {
	requestData := &grpcInterface.DeleteClassRequest{
		PackageName: packageName,
	}
	replyData, err := Client.DeleteClass(Ctx, requestData)
	if err != nil {
		fmt.Println("调用grpc服务(DeleteClass)出错：", err)
		return false
	} else {
		return replyData.DeleteRes
	}
}
