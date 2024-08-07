package service

import (
	"context"
	"log"
	"strings"
	"yssim-go/config"
	"yssim-go/library/omc"

	"github.com/bytedance/sonic"
)

func GetMessagesStringInternal() []map[string]string {
	data := omc.OMC.GetMessagesStringInternal()
	if len(data) < 1 {
		return []map[string]string{}
	}
	dataList := strings.Split(data, ";,")
	var messageList []map[string]string
	for i := 0; i < len(dataList); i++ {
		dl := strings.Split(strings.TrimSpace(dataList[i]), ",\n")
		messageMap := make(map[string]string)
		for j := 0; j < len(dl); j++ {
			jl := strings.TrimSpace(dl[j])
			switch true {
			case strings.Contains(jl, "MODELICAPATH") || strings.Contains(jl, "installPackage") || strings.Contains(jl, "Downloaded"):
				continue
			case strings.Contains(jl, "Automatically ") || strings.Contains(jl, "Lexer "):
				continue
			case strings.Contains(jl, "openModelica") || strings.Contains(jl, "OMC"):
				continue
			case strings.Contains(jl, "fully compatible without conversion script needed"):
				continue
			case strings.HasPrefix(jl, "message"):
				mes := strings.ReplaceAll(jl, "message = ", "")
				messageMap["message"] = mes[1 : len(mes)-1]
			case strings.HasPrefix(jl, "level"):
				level := strings.Split(jl, ".")
				messageMap["type"] = level[len(level)-1]
			}
		}
		if len(messageMap["message"]) > 0 {
			messageList = append(messageList, messageMap)
		}
	}
	return messageList
}

func MessageNotice(mes any) bool {
	mesJson, _ := sonic.Marshal(mes)
	userName := config.USERNAME
	reply, err := config.R.LPush(context.Background(), userName+"_"+"notification", mesJson).Result()
	if err != nil {
		log.Println("消息发送失败 reply： ", reply)
		log.Println("消息发送失败 err ： ", err)
		return false
	}
	return true
}
