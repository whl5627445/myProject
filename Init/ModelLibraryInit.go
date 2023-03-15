package Init

import (
	"log"
	"yssim-go/app/DataBaseModel"
	"yssim-go/app/service"
	"yssim-go/config"
)

func ModelLibraryInit() {
	username := config.USERNAME

	var userSpace DataBaseModel.YssimUserSpace
	var sysPackageModelAll []DataBaseModel.YssimModels
	var userPackageModelAll []DataBaseModel.YssimModels

	config.DB.Where("username = ?", username).Order("last_login_time desc").First(&userSpace)
	config.DB.Where("sys_or_user = ? AND userspace_id = ? AND default_version = ?", "sys", "0", true).Find(&sysPackageModelAll)
	//config.DB.Where("sys_or_user = ? AND userspace_id = ?", username, userSpace.ID).Find(&userPackageModelAll)
	log.Println("初始化模型库...")
	packageModelAll := append(sysPackageModelAll, userPackageModelAll...)
	service.ModelLibraryInitializationNew(packageModelAll)
	log.Println("模型库初始化完成")
}
