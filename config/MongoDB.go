package config

//
//func openMangoDB() *mongo.Client {
//	// 设置客户端连接配置
//	clientOptions := options.Client().ApplyURI("mongodb://root:simtek_cloud_sim@124.70.211.127:27017/")
//
//	// 连接到 MongoDB
//	ctx := context.Background()
//	client, err := mongo.Connect(ctx, clientOptions)
//	if err != nil {
//		log.Println("连接 MongoDB 失败： %s", err)
//	}
//
//	// 检查连接是否成功
//	err = client.Ping(ctx, nil)
//	if err != nil {
//		log.Println("连接 MongoDB 失败：%s", err)
//	}
//
//	fmt.Println("成功连接到 MongoDB")
//
//	return client
//}

//var MB = openMangoDB()
