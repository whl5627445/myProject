package Init

import (
	"time"
	"yssim-go/config"
	"yssim-go/library/omc"
)

func ModelCodeAutoSave() {
	for {
		modelNameNew := <-config.ModelCodeChan
		time.Sleep(time.Second * 1)
		omc.OMC.Save(modelNameNew)
	}
}
