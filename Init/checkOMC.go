package Init

import (
	"time"
	"yssim-go/library/omc"
)

func checkOMC() {
	time.Sleep(60 * time.Second)
	for {
		check := omc.OMC.IsPackage("Modelica")
		if !check {
			omc.OMC = omc.OmcInit()
			ModelLibraryInit()
		}
		time.Sleep(time.Second * 2)
	}
}
