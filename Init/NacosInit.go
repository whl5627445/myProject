package Init

import (
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc/grpclog"
	"yssim-go/config"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var clientConfig = constant.ClientConfig{
	BeatInterval:        1000,
	NamespaceId:         "public", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
	TimeoutMs:           5000,
	NotLoadCacheAtStart: true,
}
var serverConfigs = getServerConfigs()

func getServerConfigs() []constant.ServerConfig {
	if config.NacosIp != "" && config.NacosPort != "" {
		Port, err := strconv.ParseUint(config.NacosPort, 10, 64)
		if err != nil {
			grpclog.Error("nacos port 解析失败： ", err.Error())
		}
		return []constant.ServerConfig{{
			IpAddr: config.NacosIp,
			Port:   Port,
		}}
	}

	return []constant.ServerConfig{{
		IpAddr: "nacos",
		Port:   8848,
	}}
}

func Register() {
	// if config.DEBUG != "" {
	// 	return
	// }
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
			Ip:          config.ServiceIp,
			Port:        port,
			ServiceName: config.USERNAME,
			Weight:      10,
			Enable:      true,
			Healthy:     true,
			Ephemeral:   true,
			Metadata:    map[string]string{},
		})
		log.Println("注册结果： ", success)
		if success {
			log.Println("服务注册成功")
			break
		} else {
			log.Println("服务注册失败正在准备重新尝试: ", err)
			time.Sleep(time.Second * 10)
		}
	}
}
