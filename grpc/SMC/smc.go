package smc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/grpclog"
	"yssim-go/config"
)

func newSMC() SMCClient {
	conn, err := grpc.NewClient(config.SmcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(
		grpc.MaxCallSendMsgSize(1024*1024*50),
		grpc.MaxCallRecvMsgSize(1024*1024*50),
		grpc.UseCompressor(gzip.Name),
	))
	if err != nil {
		grpclog.Fatalf("did not connect: %v", err)
	}
	conn.Connect()
	c := NewSMCClient(conn)
	return c

}

var SMC = newSMC()
