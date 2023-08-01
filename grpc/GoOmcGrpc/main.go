package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"yssim-go/grpc/GoOmcGrpc/grpcInterface"
	"yssim-go/grpc/GoOmcGrpc/omc"
)

func StartOMC() bool {
	if omc.OMCInstance.Start {
		return true
	}
	result := make(chan bool)
	go omc.StartOMC(result)
	return <-result
}

func main() {
	// 监听127.0.0.1:50051地址
	lis, err := net.Listen("tcp", "127.0.0.1:50052")
	if err != nil {
		log.Println("failed to listen: %v", err)
	} else {
		log.Println("监听端口50052")
	}
	result := StartOMC()
	if result {
		log.Println("服务启动成功")
	} else {
		log.Println("服务启动失败,请联系管理员")
	}

	// 实例化grpc服务端
	g := grpc.NewServer()
	s := grpcInterface.MyServer{}
	// 注册Greeter服务
	grpcInterface.RegisterOmcGreeterServer(g, &s)

	// 往grpc服务端注册反射服务
	reflection.Register(g)

	// 启动grpc服务
	log.Println("启动grpc服务")
	err = g.Serve(lis)
	if err != nil {
		log.Println("服务开启失败: %v", err)
	} else {
		log.Println("服务开启成功")
	}
}
