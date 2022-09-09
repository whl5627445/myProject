package Init

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"time"
	"yssim-go/config"
)

var clientConfig = constant.ClientConfig{
	BeatInterval:        1000,
	NamespaceId:         "public", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
	TimeoutMs:           5000,
	NotLoadCacheAtStart: true,
}
var serverConfigs = []constant.ServerConfig{
	{
		IpAddr: "nacos",
		Port:   8848,
	},
}

func NacosRegister() {
	for {
		Client, err := clients.NewNamingClient(
			vo.NacosClientParam{
				ClientConfig:  &clientConfig,
				ServerConfigs: serverConfigs,
			},
		)
		if err != nil {
			fmt.Println("服务注册出现错误： ", err)
			time.Sleep(time.Second * 10)
			continue
		}
		success, err := Client.RegisterInstance(vo.RegisterInstanceParam{
			Ip:          config.USERNAME,
			Port:        uint64(config.PORT),
			ServiceName: config.USERNAME,
			Weight:      10,
			Enable:      true,
			Healthy:     true,
			Ephemeral:   true,
			Metadata:    map[string]string{},
		})
		if success {
			fmt.Println("服务注册成功")
			break
		} else {
			fmt.Println("服务注册失败正在准备重新尝试")
			time.Sleep(time.Second * 10)
			continue
		}
	}
}
