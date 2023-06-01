package Init

import (
	"log"
	"time"
	"yssim-go/app/service"
	"yssim-go/library/omc"
)

func OMCMessagesInit() {
	log.Println("初始化消息通知服务")
	for {
		time.Sleep(time.Second * 5)
		if omc.OMCInstance.Start {
			messageList := service.GetMessagesStringInternal()
			for _, mes := range messageList {
				_ = service.MessageNotice(mes)
			}
		}
	}
}
