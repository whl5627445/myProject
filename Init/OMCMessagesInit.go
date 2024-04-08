package Init

import (
	"log"
	"time"

	"yssim-go/app/v1/service"
	"yssim-go/library/omc"
)

func OMCMessagesInit() {
	log.Println("初始化消息通知服务")
	for {
		time.Sleep(time.Second * 5)
		if omc.OMCInstance.Start && omc.OMCInstance.Cmd != nil {
			omc.OMCInstance.Mu.Lock()
			messageList := service.GetMessagesStringInternal()
			for _, mes := range messageList {
				_ = service.MessageNotice(mes)
			}
			omc.OMCInstance.Mu.Unlock()
		}
	}
}
