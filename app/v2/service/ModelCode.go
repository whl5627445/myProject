package serviceV2

import (
	"yssim-go/config"
)

// ModelSave 用omc提供的API将模型源码保存的到对应文件， 并发安全
func ModelSave(modelName string) {
	//ok := omc.OMC.Save(modelName)
	config.ModelCodeChan <- modelName
}
