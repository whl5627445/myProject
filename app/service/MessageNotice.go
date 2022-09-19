package service

import (
	"context"
	"encoding/json"
	"log"
	"yssim-go/config"
)

func MessageNotice(mes interface{}) bool {
	mesJson, _ := json.Marshal(mes)
	username := config.USERNAME
	reply, err := config.R.LPush(context.Background(), username+"_"+"notification", mesJson).Result()
	if err != nil {
		log.Println("消息发送失败 reply： ", reply)
		log.Println("消息发送失败 err ： ", err)
		return false
	}
	return true
}
