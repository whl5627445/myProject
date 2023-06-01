package Init

import (
	"log"
	"time"
	"yssim-go/library/omc"
)

func checkOMC() {
	time.Sleep(60 * time.Second)
	for {
		useTime := time.Now().Local().Unix() - omc.OMCInstance.UseTime.Unix()
		if omc.OMCInstance.Start && (useTime) > 14400 {
			log.Printf("omc运行已经超过 %d 秒，启动omc停止程序", useTime)
			omc.StopOMC()
		}
		time.Sleep(time.Second * 60)
	}
}
