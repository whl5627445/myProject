package Init

import (
	"fmt"
	"yssim-go/app/DataBaseModel"
	"yssim-go/app/service"
	"yssim-go/config"
)

func ModelLibraryInit() {
	username := config.USERNAME

	var packageModelAll []DataBaseModel.YssimModels
	var userSpace DataBaseModel.YssimUserSpace
	config.DB.Where("username = ?", username).Order("last_login_time desc").First(&userSpace)

	spaceId := userSpace.ID
	config.DB.Where("sys_or_user IN ? AND userspace_id IN ?", []string{"sys", username}, []string{"0", spaceId}).Find(&packageModelAll)
	//config.DB.Find(&packageModelAll)
	fmt.Println("初始化标准库...")
	service.ModelLibraryInitialization(packageModelAll)
}
