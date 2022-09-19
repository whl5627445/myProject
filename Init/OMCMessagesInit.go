package Init

import (
	"log"
	"time"
	"yssim-go/app/service"
)

func OMCMessagesInit() {
	log.Println("初始化消息通知服务")
	for {
		time.Sleep(time.Second * 2)
		messageList := service.GetMessagesStringInternal()
		for _, mes := range messageList {
			_ = service.MessageNotice(mes)
		}
	}
}
