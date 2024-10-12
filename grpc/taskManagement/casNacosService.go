package taskManagement

import (
	"github.com/nacos-group/nacos-sdk-go/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strconv"
)

type SubscribeServiceAndClient struct {
	Client     TaskAssignmentsClient
	Service    model.SubscribeService
	UpdateTime int64
}

var CasService *SubscribeServiceAndClient

// connectCasServiceClient 初始化CAS服务的实例
func connectCasServiceClient(subscribeService model.SubscribeService) *SubscribeServiceAndClient {
	conn, err := grpc.Dial(subscribeService.Ip+":"+strconv.FormatUint(subscribeService.Port, 10), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("did not connect: %v", err)
	}
	c := NewTaskAssignmentsClient(conn)
	i := &SubscribeServiceAndClient{
		Client:  c,
		Service: subscribeService,
	}
	return i
}

// 初始化计算节点的实例列表
func ConnectTaskDispatcherClientList(instanceList []model.SubscribeService) {
	for _, instance := range instanceList {
		if !instance.Healthy {
			continue
		}
		ins := connectCasServiceClient(instance)
		CasService = ins
	}
}
