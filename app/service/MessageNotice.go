package service

import (
	"context"
	"encoding/json"
	"fmt"
	"yssim-go/config"
)

func MessageNotice(mes interface{}) bool {
	mesJson, _ := json.Marshal(mes)
	username := config.USERNAME
	reply, err := config.R.LPush(context.Background(), username+"_"+"notification", mesJson).Result()
	if err != nil {
		fmt.Println("消息发送失败 reply： ", reply)
		fmt.Println("消息发送失败 err ： ", err)
		return false
	}
	return true
}
