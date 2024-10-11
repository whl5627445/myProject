package taskManagement

//func StartTaskGrpc() (TaskAssignmentsClient, context.Context) {
//	TaskDispatcherIp, TaskDispatcherPort := config.GetHealthInstance("TaskDispatcher")
//	address := TaskDispatcherIp + ":" + TaskDispatcherPort
//	//fmt.Println(address)
//	//address = "192.168.121.12:30099"
//	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials())) // 建立链接
//	if err != nil {
//		fmt.Println("did not connect.", err)
//		return nil, nil
//	}
//
//	client := NewTaskAssignmentsClient(conn) // 初始化客户端
//	ctx := context.Background()              // 初始化元数据
//	return client, ctx
//}
//
//var TaskClient, TaskCtx = StartTaskGrpc()
