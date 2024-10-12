package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func openMangoDB() *mongo.Client {
	// 设置MongoDB连接信息
	address := "mongo:27017"
	if DebugMongo != "" {
		address = DebugMongo
	}
	clientOptions := options.Client().ApplyURI("mongodb://" + address + "/?connect=direct")

	// 连接MongoDB数据库
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)

	}

	// 检查连接
	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println("Failed to ping MongoDB:", err)

	}

	// fmt.Println("成功连接到 MongoDB")

	return client
}

var MB = openMangoDB()
