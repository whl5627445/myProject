package Init

import (
	"log"
	"yssim-go/app/DataBaseModel"
	"yssim-go/app/service"
	"yssim-go/config"
)

func modelLibraryInit() {
	username := config.USERNAME

	var userSpace DataBaseModel.YssimUserSpace
	var packageModelAll []DataBaseModel.YssimModels
	config.DB.Where("username = ?", username).Order("last_login_time desc").First(&userSpace)

	config.DB.Where("sys_or_user IN ? AND userspace_id IN ?", []string{"sys", username}, []string{"0", userSpace.ID}).Find(&packageModelAll)

	log.Println("初始化模型库...")
	service.ModelLibraryInitialization(packageModelAll)
	log.Println("模型库初始化完成")
}
