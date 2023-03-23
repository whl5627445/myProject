package API

import (
	"log"
	"net/http"
	"time"
	"yssim-go/app/DataBaseModel"
	"yssim-go/app/service"
	"yssim-go/config"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var DB = config.DB

func GetUserSpaceView(c *gin.Context) {
	/*
		# 获取用户所有的用户空间条目
	*/
	userName := c.GetHeader("username")
	var res responseData
	var modelData []map[string]string
	var userSpace []DataBaseModel.YssimUserSpace
	_ = DB.Where("username = ?", userName).Find(&userSpace)
	for _, space := range userSpace {
		modelData = append(modelData, map[string]string{"id": space.ID, "name": space.SpaceName})
	}
	res.Data = modelData
	c.JSON(http.StatusOK, res)

}

func CreateUserSpaceView(c *gin.Context) {
	/*
		# 创建用户空间
		## space_name: 用户空间名称
	*/
	userName := c.GetHeader("username")
	var res responseData
	var item userSpaceModel
	err := c.BindJSON(&item)
	if err != nil {
		return
	}
	var oneSpace DataBaseModel.YssimUserSpace
	var allSpace []DataBaseModel.YssimUserSpace
	_ = DB.Where("username = ? AND space_name = ?", userName, item.SpaceName).First(&oneSpace).Error
	_ = DB.Where("username = ?", userName).Find(&allSpace).Error
	if oneSpace.ID != "" || len(allSpace) >= 5 {
		res.Err = "空间名称已存在或数量超过5个"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	space := DataBaseModel.YssimUserSpace{
		ID:        uuid.New().String(),
		SpaceName: item.SpaceName,
		UserName:  userName,
	}
	err = DB.Create(&space).Error
	if err != nil {
		res.Err = "创建失败，请稍后再试"
		res.Status = 1
		c.JSON(http.StatusOK, res)
		return
	}
	res.Data = map[string]string{
		"id":   space.ID,
		"name": space.SpaceName,
	}
	res.Msg = "创建成功"
	c.JSON(http.StatusOK, res)
}

func DeleteUserSpaceView(c *gin.Context) {
	/*
		# 删除用户空间
		## space_id: 用户空间id
	*/
	userName := c.GetHeader("username")
	var res responseData
	var item userSpaceModel
	err := c.BindJSON(&item)
	if err != nil {
		return
	}
	var space DataBaseModel.YssimUserSpace
	DB.Where("id = ? AND username = ?", item.SpaceId, userName).Delete(&space)
	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res)

}

func LoginUserSpaceView(c *gin.Context) {
	/*
		# 进入用户空间
		## space_id: 用户空间id
	*/
	var res responseData
	var item userSpaceModel
	err := c.BindJSON(&item)
	if err != nil {
		return
	}
	userName := c.GetHeader("username")
	var spaceLast DataBaseModel.YssimUserSpace
	DB.Where("username = ?", userName).Order("last_login_time desc").First(&spaceLast)
	if item.SpaceId == spaceLast.ID {
		res.Msg = "初始化完成"
		c.JSON(http.StatusOK, res)
		return
	}

	var packageModelList []DataBaseModel.YssimModels
	var space DataBaseModel.YssimUserSpace
	DB.Where("sys_or_user = ? ", userName).Find(&packageModelList)
	DB.Where("id = ? AND username = ?", item.SpaceId, userName).First(&space)
	service.SaveModelToFileALL(packageModelList)
	var sysPackageModelAll []DataBaseModel.YssimModels
	var userPackageModelAll []DataBaseModel.YssimModels
	config.DB.Where("sys_or_user = ?  AND default_version = ?", "sys", true).Find(&sysPackageModelAll)
	//config.DB.Where("sys_or_user = ? AND userspace_id = ?", userName, item.SpaceId).Find(&userPackageModelAll)
	packageModelAll := append(sysPackageModelAll, userPackageModelAll...)
	service.ModelLibraryInitialization(packageModelAll)
	space.LastLoginTime = time.Now().Local().Unix()
	DB.Save(&space)
	res.Msg = "初始化完成"
	c.JSON(http.StatusOK, res)
}

func GetUserRecentlyOpenedView(c *gin.Context) {
	/*
		#获取用户空间的最近打开
	*/
	userName := c.GetHeader("username")
	var res responseData
	var modelData []map[string]string
	var userSpace []DataBaseModel.YssimUserSpace
	DB.Where("username = ?", userName).Order("last_login_time desc").Find(&userSpace)
	for _, space := range userSpace {
		modelData = append(modelData, map[string]string{"id": space.ID, "name": space.SpaceName})
	}
	res.Data = modelData
	c.JSON(http.StatusOK, res)
}

func ExamplesView(c *gin.Context) {
	/*
		# 获取示例
	*/
	var res responseData
	res.Data = config.EXAMPLES
	c.JSON(http.StatusOK, res)
}

func GetUserSettingsView(c *gin.Context) {
	/*
		# 获取用户配置
	*/
	var res responseData
	var setting DataBaseModel.YssimUserSettings
	username := c.GetHeader("username")
	DB.Where("username =? ", username).First(&setting)
	oneData := map[string]interface{}{
		"grid_display": setting.GridDisplay,
	}
	res.Data = oneData
	c.JSON(http.StatusOK, res)
}

func SetUserSettingsView(c *gin.Context) {
	/*
		# 设置用户配置
	*/
	var res responseData
	username := c.GetHeader("username")
	var setting userSettingsModel
	var settingRecord DataBaseModel.YssimUserSettings
	err := c.BindJSON(&setting)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	res.Data = true
	DB.Where("username =? ", username).First(&settingRecord)
	if settingRecord.UserName != "" { //存在则修改
		res.Msg = "修改成功。"
		settingRecord.GridDisplay = setting.GridDisplay
		err := DB.Where("username =? ", username).Save(&settingRecord).Error
		if err != nil {
			log.Println("err:", err)
			res.Data = false
			res.Msg = "修改失败。"
		}

	} else { //不存在则创建
		settingNew := DataBaseModel.YssimUserSettings{
			UserName:    username,
			GridDisplay: setting.GridDisplay,
		}
		res.Msg = "创建成功。"
		err := DB.Create(&settingNew).Error
		if err != nil {
			log.Println("errr", err)
			res.Data = false
			res.Msg = "创建失败。"
		}
	}

	c.JSON(http.StatusOK, res)
}
