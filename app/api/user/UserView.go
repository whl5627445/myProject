package API

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
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
	userSpaceId := c.GetHeader("space_id")
	var res ResponseData
	result := service.Clear(userSpaceId)
	if !result {
		res.Err = "模型库初始化失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	res.Msg = "初始化完成"
	c.JSON(http.StatusOK, res)
}
