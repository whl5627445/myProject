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

// GitPackageVersion 实现接口
func (s *MyServer) GitPackageVersion(ctx context.Context, in *GitPackageVersionRequest) (*GitPackageVersionReply, error) {
	libraryVersion := omc.OMC.GetClassInformation(in.PackageName)[14].(string)
	return &GitPackageVersionReply{Version: libraryVersion}, nil
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
