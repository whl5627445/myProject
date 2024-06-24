package serviceV2

import (
	"log"

	serviceV1 "yssim-go/app/v1/service"
	"yssim-go/library/omc"
)

// RestartOMC 重启OMC服务
func RestartOMC() {
	serviceV1.StopOMC()
	serviceV1.StartOMC()
}

// LibraryInitializationSingle 加载单个模型
func LibraryInitializationSingle(packageName, packageVersion, filePath string) bool {
	// 根据模型文件加载单个模型
	if filePath != "" {
		ok := omc.OMC.LoadFileNoPwd(filePath)
		log.Printf("初始化模型库：%s %s  %t \n", packageName, packageVersion, ok)
		return ok
	}

	// 根据模型名称和版本加载单个模型
	ok := omc.OMC.LoadModel(packageName, packageVersion)
	log.Printf("初始化模型库：%s %s  %t \n", packageName, packageVersion, ok)
	return ok
}
