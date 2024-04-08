package Init

import (
	"log"

	"yssim-go/app/DataBaseModel"
	"yssim-go/app/v1/service"
	"yssim-go/config"
)

func ModelLibraryInit() {
	userName := config.USERNAME

	var userSpace DataBaseModel.YssimUserSpace
	var packageModelAll []DataBaseModel.YssimModels

	config.DB.Where("username = ?", userName).Order("last_login_time desc").First(&userSpace)
	config.DB.Where("sys_or_user IN ? AND userspace_id IN ? AND default_version = ?", []string{"sys", userName}, []string{"0", userSpace.ID}, true).Find(&packageModelAll)

	log.Println("初始化模型库...")
	service.DefaultLibraryInitialization(packageModelAll)
	service.SetWorkSpaceId(&userSpace.ID)
	log.Println("模型库初始化完成")
}
