package API

import (
	"log"
	"net/http"
	"time"
	"yssim-go/app/DataBaseModel"
	"yssim-go/app/service"
	"yssim-go/config"
	"yssim-go/library/timeConvert"

	"gorm.io/gorm"

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

func GetUserSpaceNewView(c *gin.Context) {
	/*
		# 获取用户所有的用户空间条目
	*/

	var res responseData
	userName := c.GetHeader("username")
	keyWords := c.Query("keywords")
	collect := c.Query("collect")
	var recentSpaceList []DataBaseModel.YssimUserSpace
	var allSpaceList []DataBaseModel.YssimUserSpace
	db := DB.Where("username = ?", userName)
	if keyWords != "" {
		db.Where("space_name LIKE ?", "%"+keyWords+"%")
	}
	if collect == "1" {
		db.Where("collect = ?", true)
	}
	db = db.Session(&gorm.Session{})
	db.Order("create_time desc").Find(&allSpaceList)
	db.Where("last_login_time <> ?", 0).Order("last_login_time desc").Limit(5).Find(&recentSpaceList)
	allAppSpace := make([]map[string]interface{}, 0)
	recentAppSpace := make([]map[string]interface{}, 0)
	for _, space := range allSpaceList {
		updateTime := space.UpdatedAt.Local().Unix()
		editTime := timeConvert.UseTimeFormatNew(int(updateTime), int(time.Now().Local().Unix()), 1)
		d := map[string]interface{}{
			"id":          space.ID,
			"name":        space.SpaceName,
			"description": space.Description,
			"background":  space.Background,
			"icon":        space.Icon,
			"icon_color":  space.IconColor,
			"collect":     space.Collect,
			"edit_time":   "编辑于" + editTime + "前",
		}
		allAppSpace = append(allAppSpace, d)
	}
	for _, space := range recentSpaceList {
		updateTime := space.UpdatedAt.Local().Unix()
		editTime := timeConvert.UseTimeFormatNew(int(updateTime), int(time.Now().Local().Unix()), 1)
		d := map[string]interface{}{
			"id":          space.ID,
			"name":        space.SpaceName,
			"description": space.Description,
			"background":  space.Background,
			"icon":        space.Icon,
			"icon_color":  space.IconColor,
			"collect":     space.Collect,
			"edit_time":   "编辑于" + editTime + "前",
		}
		recentAppSpace = append(recentAppSpace, d)
	}
	data := map[string][]map[string]interface{}{
		"all_space":    allAppSpace,
		"recent_space": recentAppSpace,
	}
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func CreateUserSpaceView(c *gin.Context) {
	/*
		# 创建用户空间
	*/
	userName := c.GetHeader("username")
	var res responseData
	var item CreateUserSpaceModel
	err := c.BindJSON(&item)
	if err != nil {
		return
	}
	var space DataBaseModel.YssimUserSpace
	DB.Where("username = ? AND space_name = ?", userName, item.SpaceName).First(&space)
	if space.ID != "" {
		res.Err = "名称已存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	space.ID = uuid.New().String()
	space.SpaceName = item.SpaceName
	space.UserName = userName
	space.Description = item.Description
	space.Background = item.Background
	space.Icon = item.Icon
	space.IconColor = item.IconColor
	space.Collect = false
	err = DB.Create(&space).Error
	res.Data = map[string]string{
		"id": space.ID,
	}
	FilePath, ok := service.CreatWorkSpace(userName, space.SpaceName)
	if ok {
		defaultWorkSpacePackage := DataBaseModel.YssimModels{
			ID:          uuid.New().String(),
			PackageName: "Workspace",
			SysUser:     userName,
			FilePath:    FilePath,
			UserSpaceId: space.ID,
			Default:     true,
		}
		err = DB.Create(&defaultWorkSpacePackage).Error
	}
	if err != nil || !ok {
		DB.Delete(&space)
		res.Err = "创建失败，请稍后再试"
		res.Status = 1
		c.JSON(http.StatusOK, res)
		return
	}
	res.Msg = "创建成功"
	c.JSON(http.StatusOK, res)
}

func EditUserSpaceView(c *gin.Context) {
	/*
		# 编辑用户空间
	*/
	userName := c.GetHeader("username")
	var res responseData
	var item EditUserSpaceModel
	err := c.BindJSON(&item)
	if err != nil {

		return
	}
	var space DataBaseModel.YssimUserSpace
	err = DB.Model(&space).Where("username = ? AND id = ?", userName, item.SpaceId).Updates(item).Error
	if err != nil {
		log.Println("建模空间更新数据库失败，错误： ", err)
		res.Err = "编辑失败，请重试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	res.Msg = "编辑成功"
	c.JSON(http.StatusOK, res)
}

func CollectUserSpaceView(c *gin.Context) {
	/*
		# 收藏用户空间
	*/
	userName := c.GetHeader("username")
	var res responseData
	var item CollectUserSpaceData
	err := c.BindJSON(&item)
	if err != nil {
		return
	}
	var space DataBaseModel.YssimUserSpace
	err = DB.Model(&space).Where("id IN ? AND username = ? ", item.SpaceId, userName).Update("collect", item.Collect).Error
	if err != nil {
		log.Println("建模空间更新数据库失败，错误： ", err)
		res.Err = "编辑失败，请重试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	res.Msg = "收藏成功"
	if !item.Collect {
		res.Msg = "取消收藏成功"
	}
	c.JSON(http.StatusOK, res)
}

func DeleteUserSpaceView(c *gin.Context) {
	/*
		# 删除用户空间
	*/
	userName := c.GetHeader("username")
	var res responseData
	var item DeleteUserSpaceModel
	err := c.BindJSON(&item)
	if err != nil {
		return
	}
	var space DataBaseModel.YssimUserSpace
	for _, id := range item.SpaceId {
		result := service.GetWorkSpaceId(&id)
		if result {
			service.Clear()
		}
	}
	DB.Model(&space).Where("id IN ? AND username = ?", item.SpaceId, userName).Delete(&space)
	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res)

}

func LoginUserSpaceView(c *gin.Context) {
	/*
		# 进入用户空间
		## space_id: 用户空间id
	*/
	var res responseData
	var item LoginUserSpaceModel
	err := c.BindJSON(&item)
	if err != nil {
		return
	}
	userName := c.GetHeader("username")
	result := service.SetWorkSpaceId(&item.SpaceId)
	if result {
		res.Msg = "初始化完成"
		c.JSON(http.StatusOK, res)
		return
	}

	var space DataBaseModel.YssimUserSpace
	DB.Model(space).Where("id = ? AND username = ?", item.SpaceId, userName).UpdateColumn("last_login_time", time.Now().Local().Unix())

	var packageModelAll []DataBaseModel.YssimModels
	DB.Where("sys_or_user IN ?  AND default_version = ? AND userspace_id IN ?", []string{"sys", userName}, true, []string{"0", space.ID}).Find(&packageModelAll)
	service.ModelLibraryInitialization(packageModelAll)

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
	//res.Data = config.EXAMPLES
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
