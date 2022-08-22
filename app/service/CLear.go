package service

import (
	"fmt"
	"yssim-go/app/DataBaseModel"
	"yssim-go/config"
	"yssim-go/library/omc"
)

func Clear(spaceId string) bool {
	username := config.USERNAME
	omc.OMC.Clear()
	var packageModel []DataBaseModel.YssimModels
	if spaceId == "" {
		var userSpace DataBaseModel.YssimUserSpace
		config.DB.Where("username = ?", username).Order("last_login_time desc").First(&userSpace)
		spaceId = userSpace.ID
	}
	config.DB.Where("sys_or_user IN ? AND userspace_id IN ?", []string{"sys", username}, []string{"0", spaceId}).Find(&packageModel)
	fmt.Println("初始化标准库...")
	for _, models := range packageModel {
		res, ok := omc.OMC.SendExpressionNoParsed(fmt.Sprintf("loadModel(%s, {\"\"},true,\"\",false)", models.PackageName))
		if ok {
			fmt.Printf("初始化模型库： %s  %s", models.PackageName, res)
		} else {
			fmt.Println("模型库：" + models.PackageName + "  初始化失败")
			return false
		}
	}
	fmt.Println("标准库初始化完成")
	return true
}
