package Init

import (
	"log"
	"strconv"
	"time"

	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/util"
	"yssim-go/grpc/taskManagement"

	"yssim-go/config"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func Register() {
	// if config.DEBUG != "" {
	// 	return
	// }
	for {
		Client, err := clients.NewNamingClient(
			vo.NacosClientParam{
				ClientConfig:  &config.ClientConfig,
				ServerConfigs: config.ServerConfigs,
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

// ListenTaskDispatcher 监听调度中心的实例变化，如果实例列表发生变化，则运行对应的函数逻辑
func ListenTaskDispatcher() {
	subscribeParam := &vo.SubscribeParam{
		ServiceName: "TaskDispatcher",
		GroupName:   "DEFAULT_GROUP",
		SubscribeCallback: func(services []model.SubscribeService, err error) {
			log.Println("监听到变化:%s \n", util.ToJsonString(services))
			// taskManagement.GetHealthInstance("TaskDispatcher")
			// 如果变化，
			TaskManagement.ConnectTaskDispatcherClientList(services)
		},
	}
	err := config.NamingClient.Subscribe(subscribeParam)

	if err != nil {
		return
	} else {
		log.Println(err)
	}
}
