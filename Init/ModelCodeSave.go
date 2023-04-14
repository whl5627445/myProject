package Init

import (
	"time"
	"yssim-go/config"
	"yssim-go/library/omc"
)

func ModelCodeAutoSave() {
	modelName := ""
	timeTime := time.Now().Unix()
	for {
		now := time.Now().Unix()
		modelNameNew := <-config.ModelCodeChan
		if modelName == modelNameNew && (now-timeTime) < 1 {
			continue
		}
		time.Sleep(time.Second * 1)
		omc.OMC.Save(modelName)
		modelName = modelNameNew
	}
}
