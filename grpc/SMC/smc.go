package smc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

func NewSMC(addr string) SMCClient {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(
		grpc.MaxCallSendMsgSize(1024*1024*50),
		grpc.MaxCallRecvMsgSize(1024*1024*50),
	))
	if err != nil {
		grpclog.Fatalf("did not connect: %v", err)
	}
	conn.Connect()
	c := NewSMCClient(conn)
	return c

}
