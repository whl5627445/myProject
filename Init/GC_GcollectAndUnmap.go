package Init

import (
	"time"
	"yssim-go/library/omc"
)

func GcCollectAndUnmap() {
	time.Sleep(60 * time.Second)
	for {
		// 每200秒强制执行GC一次
		if omc.OMCInstance.Start {
			omc.OMC.SendExpressionNoParsed("GC_gcollect_and_unmap()")
		}
		time.Sleep(time.Second * 200)
	}
}
