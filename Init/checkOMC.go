package Init

import (
	"time"
	"yssim-go/library/omc"
)

func checkOMC() {
	for {
		//if omc.OMCInstance.UseTime != nil || omc.OMCInstance.Start || omc.OMCInstance.Cmd != nil {
		//	log.Println("username: ", config.USERNAME)
		//	log.Printf("omc.OMCInstance.Cmd %#v", omc.OMCInstance.Cmd)
		//	log.Printf("omc已运行 %d 秒", time.Now().Local().Unix()-omc.OMCInstance.UseTime.Unix())
		//}

		if omc.OMCInstance.Start && (time.Now().Local().Unix()-omc.OMCInstance.UseTime.Unix()) > 14400 {
			omc.StopOMC()
		}
		time.Sleep(time.Second * 60)
	}
}
