package config

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"strconv"
)

var clientConfig = constant.ClientConfig{
	BeatInterval:        1000,
	NamespaceId:         "public", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
	TimeoutMs:           5000,
	NotLoadCacheAtStart: true,
}
var serverConfigs = getServerConfigs()

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

func GetHealthInstance(serviceName string) (string, string) {
	Client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		fmt.Println("服务注册出现错误： ", err)
		return "", ""
	}
	instance, err := Client.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: serviceName,
	})
	if err != nil {
		fmt.Println("获取服务实例出现错误： ", err)
		return "", ""
	}
	return instance.Ip, strconv.FormatUint(instance.Port, 10)
}
