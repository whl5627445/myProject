package user

import (
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"yssim-go/app/DataBaseModel"
	"yssim-go/app/DataType"
	"yssim-go/app/v1/service"
	"yssim-go/config"
	"yssim-go/library/timeConvert"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var dbUser = config.DB

func GetUserSpaceView(c *gin.Context) {
	/*
		# 获取用户所有的用户空间条目
	*/
	userName := config.USERNAME

	var res DataType.ResponseData
	var modelData []map[string]string
	var userSpace []DataBaseModel.YssimUserSpace
	_ = dbUser.Where("username = ?", userName).Find(&userSpace)
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

	var res DataType.ResponseData
	userName := config.USERNAME
	keyWords := c.Query("keywords")
	collect := c.Query("collect")
	var recentSpaceList []DataBaseModel.YssimUserSpace
	var allSpaceList []DataBaseModel.YssimUserSpace
	db := dbUser.Where("username = ?", userName)
	if keyWords != "" {
		db.Where("space_name LIKE ?", "%"+keyWords+"%")
	}
	if collect == "1" {
		db.Where("collect = ?", true)
	}
	db = db.Session(&gorm.Session{})
	db.Order("create_time desc").Find(&allSpaceList)
	db.Where("last_login_time <> ?", 0).Order("last_login_time desc").Limit(5).Find(&recentSpaceList)
	allAppSpace := make([]map[string]any, 0)
	recentAppSpace := make([]map[string]any, 0)
	for _, space := range allSpaceList {
		updateTime := space.UpdatedAt.Local().Unix()
		editTime := timeConvert.UseTimeFormatNew(int(updateTime), int(time.Now().Local().Unix()), 1)
		d := map[string]any{
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
		d := map[string]any{
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
	data := map[string][]map[string]any{
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
	userName := config.USERNAME
	var res DataType.ResponseData
	var item DataType.CreateUserSpaceModel
	err := c.BindJSON(&item)
	if err != nil {
		return
	}
	matchSpaceName, _ := regexp.MatchString("^[_0-9a-zA-Z\u4e00-\u9fa5]+$", item.SpaceName) // 由中文、字母、数字、下划线验证
	if !matchSpaceName {
		res.Err = "工作空间名称只能由中文、字母、数字、下划线组成"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	var space DataBaseModel.YssimUserSpace
	dbUser.Where("username = ? AND space_name = ?", userName, item.SpaceName).First(&space)
	if space.ID != "" {
		res.Err = "工作空间名称已存在"
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
	err = dbUser.Create(&space).Error
	res.Data = map[string]string{
		"id":   space.ID,
		"name": space.SpaceName,
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
		err = dbUser.Create(&defaultWorkSpacePackage).Error
	}
	if err != nil || !ok {
		dbUser.Delete(&space)
		res.Err = "创建失败，请稍后再试"
		res.Status = 2
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
	userName := config.USERNAME
	var res DataType.ResponseData
	var item DataType.EditUserSpaceModel
	err := c.BindJSON(&item)
	if err != nil {

		return
	}
	matchSpaceName, _ := regexp.MatchString("^[_0-9a-zA-Z\u4e00-\u9fa5]+$", item.SpaceName) // 由中文、字母、数字、下划线验证
	if !matchSpaceName {
		res.Err = "工作空间名称只能由中文、字母、数字、下划线组成"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	var space DataBaseModel.YssimUserSpace
	dbUser.Where("space_name = ? AND username = ? AND id <> ?", item.SpaceName, userName, item.SpaceId).First(&space)
	if space.ID != "" {
		res.Err = "工作空间名称已存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	err = dbUser.Model(&space).Where("username = ? AND id = ?", userName, item.SpaceId).Updates(item).Error
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
	userName := config.USERNAME
	var res DataType.ResponseData
	var item DataType.CollectUserSpaceData
	err := c.BindJSON(&item)
	if err != nil {
		return
	}
	var space DataBaseModel.YssimUserSpace
	err = dbUser.Model(&space).Where("id IN ? AND username = ? ", item.SpaceId, userName).Update("collect", item.Collect).Error
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
	userName := config.USERNAME
	var res DataType.ResponseData
	var item DataType.DeleteUserSpaceModel
	err := c.BindJSON(&item)
	if err != nil {
		return
	}
	var space DataBaseModel.YssimUserSpace
	// for _, id := range item.SpaceId {
	//	result := service.GetWorkSpaceId(&id)
	//	if result {
	//		service.Clear()
	//	}
	// }
	dbUser.Model(&space).Where("id IN ? AND username = ?", item.SpaceId, userName).Delete(&space)
	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res)

}

func GetUserRecentlyOpenedView(c *gin.Context) {
	/*
		#获取用户空间的最近打开
	*/
	userName := config.USERNAME
	var res DataType.ResponseData
	var modelData []map[string]string
	var userSpace []DataBaseModel.YssimUserSpace
	dbUser.Where("username = ?", userName).Order("last_login_time desc").Find(&userSpace)
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
	var res DataType.ResponseData
	// res.Data = config.EXAMPLES
	c.JSON(http.StatusOK, res)
}

func GetUserSettingsView(c *gin.Context) {
	/*
		# 获取用户配置
	*/
	var res DataType.ResponseData
	var setting DataBaseModel.YssimUserSettings
	userName := config.USERNAME
	dbUser.Where("username =? ", userName).First(&setting)
	oneData := map[string]any{
		"grid_display": setting.GridDisplay,
	}
	res.Data = oneData
	c.JSON(http.StatusOK, res)
}

func SetUserSettingsView(c *gin.Context) {
	/*
		# 设置用户配置
	*/
	var res DataType.ResponseData
	userName := config.USERNAME
	var setting DataType.UserSettingsModel
	var settingRecord DataBaseModel.YssimUserSettings
	err := c.BindJSON(&setting)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	res.Data = true
	dbUser.Where("username =? ", userName).First(&settingRecord)
	if settingRecord.ID != "" { // 存在则修改
		res.Msg = "修改成功。"
		settingRecord.GridDisplay = setting.GridDisplay
		err := dbUser.Where("username =? ", userName).Save(&settingRecord).Error
		if err != nil {
			log.Println("err:", err)
			res.Data = false
			res.Msg = "修改失败。"
			res.Status = 2
		}

	} else { // 不存在则创建
		settingNew := DataBaseModel.YssimUserSettings{
			ID:          uuid.New().String(),
			UserName:    userName,
			GridDisplay: setting.GridDisplay,
		}
		res.Msg = "创建成功。"
		err := dbUser.Create(&settingNew).Error
		if err != nil {
			log.Println("err", err)
			res.Data = false
			res.Msg = "创建失败。"
			res.Status = 2
		}
	}

	c.JSON(http.StatusOK, res)
}

func BackgroundUploadView(c *gin.Context) {
	/*
		# 上传背景图接口
		## path: 文件相对路径
	*/
	var res DataType.ResponseData
	userName := config.USERNAME
	varFile, err := c.FormFile("file")
	if !strings.HasSuffix(varFile.Filename, ".jpg") && !strings.HasSuffix(varFile.Filename, ".jpeg") && !strings.HasSuffix(varFile.Filename, ".png") &&
		!strings.HasSuffix(varFile.Filename, ".jfif") && !strings.HasSuffix(varFile.Filename, ".pjp") && !strings.HasSuffix(varFile.Filename, ".pjpeg") {
		res.Err = "暂时只支持*.jpg、*.jfif、*.pjpeg、*.pjp、*.jpeg、*.png格式文件上传"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	if varFile.Size > 3500000 {
		res.Err = "上传文件过大，请上传小于3.5M的文件"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	filePath := "static/UserFiles/Images/" + userName + time.Now().Local().Format("20060102150405") + varFile.Filename
	ok := service.SaveBackground(filePath, varFile)
	if ok {
		res.Msg = "上传成功"
		res.Data = map[string]string{
			"path": filePath,
		}
		c.JSON(http.StatusOK, res)
		return
	}
	res.Err = "上传失败，请重新上传"
	res.Status = 2
	c.JSON(http.StatusOK, res)

}

func StartOMCView(c *gin.Context) {
	/*
		# 启动用户的omc实例并连接
	*/
	var res DataType.ResponseData
	result := service.StartOMC()
	if result {
		res.Msg = "服务启动成功"
	} else {
		res.Err = "服务启动失败,请联系管理员"
		res.Status = 2
	}
	c.JSON(http.StatusOK, res)
}

func StopOMCView(c *gin.Context) {
	/*
		# 暂停用户的omc实例并连接
	*/
	var res DataType.ResponseData
	service.StopOMC()
	res.Msg = "服务暂停成功"
	c.JSON(http.StatusOK, res)
}

func RestartOMCView(c *gin.Context) {
	/*
		# 重启用户的omc实例并连接
	*/
	var res DataType.ResponseData

	service.StopOMC()
	result := service.StartOMC()
	if result {
		res.Msg = "服务重启成功"
	} else {
		res.Err = "服务重启失败,请联系管理员"
		res.Status = 2
	}

	c.JSON(http.StatusOK, res)
}
