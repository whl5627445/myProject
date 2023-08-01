package grpcInterface

import (
	"context"
	"log"
	"yssim-go/grpc/GoOmcGrpc/omc"
)

// MyServer 定义MyServer，用来实现proto文件，里面实现的Greeter服务里面的接口
type MyServer struct {
	UnimplementedOmcGreeterServer
}

// GetPackageVersion 实现接口
func (s *MyServer) GetPackageVersion(ctx context.Context, in *GetPackageVersionRequest) (*GetPackageVersionReply, error) {
	libraryVersion := omc.OMC.GetClassInformation(in.PackageName)[14].(string)
	return &GetPackageVersionReply{Version: libraryVersion}, nil
}

// LoadFile 实现接口
func (s *MyServer) LoadFile(ctx context.Context, in *LoadFileRequest) (*LoadFileReply, error) {
	log.Println(in.PackagePath)
	res := omc.OMC.LoadFile(in.PackagePath)
	return &LoadFileReply{LoadFileRes: res}, nil
}

// ParseFile 实现接口
func (s *MyServer) ParseFile(ctx context.Context, in *ParseFileRequest) (*ParseFileReply, error) {
	log.Println(in.FilePath)
	packageName, res := omc.OMC.ParseFile(in.FilePath)
	return &ParseFileReply{PackageName: packageName, ParseRes: res}, nil
}

// DeleteClass 实现接口
func (s *MyServer) DeleteClass(ctx context.Context, in *DeleteClassRequest) (*DeleteClassReply, error) {
	log.Println(in.PackageName)
	res := omc.OMC.DeleteClass(in.PackageName)
	return &DeleteClassReply{DeleteRes: res}, nil
}
