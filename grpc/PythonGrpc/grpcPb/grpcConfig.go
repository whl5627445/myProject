package grpcPb

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"yssim-go/config"
)

func StartGrpc() (GreeterClient, context.Context) {
	address := config.GrpcServerName + ":" + config.GrpcPort
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials())) // 建立链接
	if err != nil {
		log.Println("did not connect.", err)
		return nil, nil
	}

	client := NewGreeterClient(conn) // 初始化客户端
	ctx := context.Background()      // 初始化元数据
	return client, ctx
}

var Client, Ctx = StartGrpc()
