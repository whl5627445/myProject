package Init

import (
	"time"
	"yssim-go/library/omc"
)

func checkOMC() {
	for {
		if omc.OMCInstance.Start && (time.Now().Local().Unix()-omc.OMCInstance.UseTime.Unix()) > 14400 {
			omc.StopOMC()
		}
		time.Sleep(time.Second * 60)
	}
}
