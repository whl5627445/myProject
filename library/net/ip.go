package net

import (
	"fmt"
	"net"

	"google.golang.org/grpc/grpclog"
)

func GetLocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		grpclog.Error("获取本机ip出错， Error: ", err)
		return "nacos"
	}
	for _, addr := range addrs {
		// 检查地址类型并跳过环回地址
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				fmt.Println("IP address:", ipNet.IP.String())
				return ipNet.IP.String()
			}
		}
	}
	return "nacos"
}
