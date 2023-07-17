package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func openMangoDB() *mongo.Client {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://root:simtek_cloud_sim@mongodb:27017/")

	// 连接到 MongoDB
	ctx := context.Background()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(fmt.Sprintf("连接 MongoDB 失败： %s", err))
	}

	// 检查连接是否成功
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(fmt.Sprintf("连接 MongoDB 失败：%s", err))
	}

	fmt.Println("成功连接到 MongoDB")

	return client
}

var MB = openMangoDB()
