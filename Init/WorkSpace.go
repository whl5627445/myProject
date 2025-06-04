package Init

import (
	"time"

	"yssim-go/grpc/WorkSpace"
)

func initWorkSpace() {
	WorkSpace.WS.Inquire()
}
func checkWorkSpace() {
	for {
		// if omc.OMCInstance.UseTime != nil || omc.OMCInstance.Start || omc.OMCInstance.Cmd != nil {
		//	log.Println("username: ", config.USERNAME)
		//	log.Printf("omc.OMCInstance.Cmd %#v", omc.OMCInstance.Cmd)
		//	log.Printf("omc已运行 %d 秒", time.Now().Local().Unix()-omc.OMCInstance.UseTime.Unix())
		// }

		for _, space := range WorkSpace.WS.ContainerMap {
			if time.Since(space.LastTime) > 3600*time.Second {
				WorkSpace.WS.Remove(space.Id)
			}
		}
		time.Sleep(time.Second * 60)
	}
}
