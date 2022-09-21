package API

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
	"yssim-go/app/DataBaseModel"
	"yssim-go/app/service"
	"yssim-go/config"
)

var DB = config.DB

func GetUserSpaceView(c *gin.Context) {
	/*
		# 获取用户所有的用户空间条目
	*/
	username := c.GetHeader("username")
	var res ResponseData
	var modelData []map[string]string
	var userSpace []DataBaseModel.YssimUserSpace
	_ = DB.Where("UserName = ?", username).Find(&userSpace)
	for _, space := range userSpace {
		modelData = append(modelData, map[string]string{"id": space.ID, "name": space.SpaceName})
	}
	res.Data = modelData
	c.JSON(http.StatusOK, res)

}

func CreateUserSpaceView(c *gin.Context) {
	/*
		# 创建用户空间
	*/
	username := c.GetHeader("username")
	var res ResponseData
	var item UserSpaceModel
	err := c.BindJSON(&item)
	if err != nil {
		return
	}
	var oneSpace DataBaseModel.YssimUserSpace
	var allSpace []DataBaseModel.YssimUserSpace
	_ = DB.Where("username = ? AND space_name = ?", username, item.SpaceName).First(&oneSpace).Error
	_ = DB.Where("username = ?", username).Find(&allSpace).Error
	if oneSpace.ID != "" || len(allSpace) >= 5 {
		res.Err = "空间名称已存在或数量超过5个"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	space := DataBaseModel.YssimUserSpace{
		ID:        uuid.New().String(),
		SpaceName: item.SpaceName,
		UserName:  username,
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
	*/
	username := c.GetHeader("username")
	var res ResponseData
	var item UserSpaceModel
	err := c.BindJSON(&item)
	if err != nil {
		return
	}
	var space DataBaseModel.YssimUserSpace
	DB.Where("id = ? AND username = ?", item.SpaceId, username).Delete(&space)
	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res)

}

func LoginUserSpaceView(c *gin.Context) {
	/*
		# 进入用户空间
	*/
	var res ResponseData
	var item UserSpaceModel
	err := c.BindJSON(&item)
	if err != nil {
		return
	}
	username := c.GetHeader("username")
	var packageModel []DataBaseModel.YssimModels
	var spaceModel DataBaseModel.YssimUserSpace
	DB.Where("sys_or_user IN ? AND userspace_id IN ?", []string{"sys", username}, []string{"0", item.SpaceId}).Find(&packageModel)
	DB.Where("id = ? AND username = ?", item.SpaceId, username).First(&spaceModel)
	service.ModelLibraryInitialization(packageModel)
	spaceModel.LastLoginTime = time.Now().Local().Unix()
	DB.Save(&spaceModel)
	res.Msg = "初始化完成"
	c.JSON(http.StatusOK, res)
}

func ExamplesView(c *gin.Context) {
	/*
		# 获取示例
	*/
	var res ResponseData
	res.Data = config.EXAMPLES
	c.JSON(http.StatusOK, res)
}

func GetUserRecentlyOpenedView(c *gin.Context) {
	/*
		#获取用户空间的最近一次打开
	*/
	username := c.GetHeader("username")
	var res ResponseData
	var modelData []map[string]string
	var userSpace []DataBaseModel.YssimUserSpace
	DB.Where("username = ?", username).Order("last_login_time desc").Find(&userSpace)
	for _, space := range userSpace {
		modelData = append(modelData, map[string]string{"id": space.ID, "name": space.SpaceName})
	}
	res.Data = modelData
	c.JSON(http.StatusOK, res)
}
