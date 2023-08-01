package grpcOmc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"yssim-go/grpc/GoOmcGrpc/grpcInterface"
)

const (
	address = "127.0.0.1:50052"
)

func StartGrpc() (grpcInterface.OmcGreeterClient, context.Context) {
	conn, err := grpc.Dial(
		address, grpc.WithTransportCredentials(insecure.NewCredentials())) // 建立链接
	if err != nil {
		log.Println("did not connect.", err)
		return nil, nil
	}

	client := grpcInterface.NewOmcGreeterClient(conn) // 初始化客户端
	ctx := context.Background()                       // 初始化元数据
	return client, ctx
}

var Client, Ctx = StartGrpc()
