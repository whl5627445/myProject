package Init

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	"strconv"
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
	if config.DEBUG != "" {
		return
	}
	for {
		Client, err := clients.NewNamingClient(
			vo.NacosClientParam{
				ClientConfig:  &clientConfig,
				ServerConfigs: serverConfigs,
			},
		)
		if err != nil {
			log.Println("服务注册出现错误： ", err)
			time.Sleep(time.Second * 10)
			continue
		}
		port, _ := strconv.ParseUint(config.PORT, 10, 64)
		success, err := Client.RegisterInstance(vo.RegisterInstanceParam{
			Ip:          config.USERNAME,
			Port:        port,
			ServiceName: config.USERNAME,
			Weight:      10,
			Enable:      true,
			Healthy:     true,
			Ephemeral:   true,
			Metadata:    map[string]string{},
		})
		if success {
			log.Println("服务注册成功")
			break
		} else {
			log.Println("服务注册失败正在准备重新尝试")
			time.Sleep(time.Second * 10)
			continue
		}
	}
}