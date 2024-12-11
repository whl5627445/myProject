package config

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/common/logger"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	"strconv"
	"time"
)

var ClientConfig = constant.ClientConfig{
	NamespaceId:         "public", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
	TimeoutMs:           5000,
	NotLoadCacheAtStart: true,
	AppendToStdout:      true,
	LogSampling:         &logger.SamplingConfig{Initial: 10, Thereafter: 3600 * 12, Tick: time.Second},
	LogLevel:            "error",
}
var ServerConfigs = getServerConfigs()

func getServerConfigs() []constant.ServerConfig {
	if NacosIp != "" && NacosPort != "" {
		Port, err := strconv.ParseUint(NacosPort, 10, 64)
		if err != nil {
			fmt.Println("nacos port 解析失败： ", err.Error())
		}
		return []constant.ServerConfig{{
			IpAddr: NacosIp,
			Port:   Port,
		}}
	}

	return []constant.ServerConfig{{
		IpAddr: "nacos",
		Port:   8848,
	}}
}

func createNamingClient() naming_client.INamingClient {
	for {
		var client, err = clients.NewNamingClient(
			vo.NacosClientParam{
				ClientConfig:  &ClientConfig,
				ServerConfigs: ServerConfigs,
			},
		)
		if err != nil {
			log.Println("连接注册中心出现错误： %s", err)
			time.Sleep(time.Second * 10)
			continue
		}
		return client
	}
}

var NamingClient = createNamingClient()
