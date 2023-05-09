package API

import (
	"log"
	"net/http"
	"strings"
	"time"
	"yssim-go/app/DataBaseModel"
	"yssim-go/config"
	"yssim-go/library/timeConvert"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var DB = config.DB

func AppModelMarkView(c *gin.Context) {
	/*
		# 标记模型为应用可用数据源, 并编译为可执行文件，若编译失败则不作为数据源显示
	*/
	var res responseData
	userName := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var item AppModelMarkData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var record DataBaseModel.AppDataSource
	DB.Where("package_id = ? AND username = ? AND user_space_id = ?", item.PackageId, userName, userSpaceId).First(&record)
	if record.ID != "" && !item.MandatorySave {
		res.Msg = "该模型已可以创建应用,继续会更新之前的记录,已发布的页面不受影响, 是否继续？"
		c.JSON(http.StatusOK, res)
		return
	}
	if record.CompileStatus == 1 && !item.MandatorySave {
		res.Msg = "该模型已导入数据源并正在进行编译,继续保存会放弃编译阶段作业, 是否继续？"
		c.JSON(http.StatusOK, res)
		return
	}
	CompilePath := "static/modelDataSource/" + userName + "/" + strings.ReplaceAll(item.ModelName, ".", "-") + "/" + time.Now().Local().Format("20060102150405")
	if record.ID == "" {
		dataSource := DataBaseModel.AppDataSource{
			ID:            uuid.New().String(),
			UserName:      userName,
			UserSpaceId:   userSpaceId,
			PackageId:     item.PackageId,
			ModelName:     item.ModelName,
			CompilerType:  item.CompilerType,
			CompilePath:   CompilePath,
			CompileStatus: 0,
		}
		err = DB.Create(&dataSource).Error
	} else {
		record.CompilePath = CompilePath
		err = DB.Save(&record).Error
	}

	if err != nil {
		log.Println("标记数据源时创建数据库记录失败： ", err)
		res.Err = "创建失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	res.Msg = "创建成功"
	c.JSON(http.StatusOK, res)
}

func GetAppSpaceView(c *gin.Context) {
	/*
		# 获取用户所有的应用空间条目
	*/
	var res responseData
	userName := c.GetHeader("username")
	//userSpaceId := c.GetHeader("space_id")
	keyWords := c.Query("keywords")
	release := c.Query("release")
	collect := c.Query("collect")
	var recentAppSpaceList []DataBaseModel.AppSpace
	var allAppSpaceList []DataBaseModel.AppSpace
	db := DB.Where("username = ?", userName)
	if keyWords != "" {
		db.Where("space_name LIKE ?", "%"+keyWords+"%")
	}
	if collect == "1" {
		db.Where("collect = ?", true)
	}
	switch {
	case release == "1":
		db.Where("is_release = ?", true)
	case release == "0":
		db.Where("is_release = ?", false)
	}
	db.Order("update_time desc").Find(&allAppSpaceList)
	db.Where("last_login_time <> ?", 0).Order("last_login_time desc").Limit(5).Find(&recentAppSpaceList)
	allAppSpace := make([]map[string]interface{}, 0)
	recentAppSpace := make([]map[string]interface{}, 0)
	for _, space := range allAppSpaceList {
		updateTime := space.UpdatedAt.Local().Unix()
		editTime := timeConvert.UseTimeFormat(int(updateTime), int(time.Now().Local().Unix()))
		d := map[string]interface{}{
			"id":          space.ID,
			"name":        space.SpaceName,
			"description": space.Description,
			"background":  space.Background,
			"icon":        space.Icon,
			"icon_color":  space.IconColor,
			"collect":     space.Collect,
			"edit_time":   editTime + "前",
		}
		allAppSpace = append(allAppSpace, d)
	}
	for _, space := range recentAppSpaceList {
		updateTime := space.UpdatedAt.Local().Unix()
		editTime := timeConvert.UseTimeFormat(int(updateTime), int(time.Now().Local().Unix()))
		d := map[string]interface{}{
			"id":          space.ID,
			"name":        space.SpaceName,
			"description": space.Description,
			"background":  space.Background,
			"icon":        space.Icon,
			"icon_color":  space.IconColor,
			"collect":     space.Collect,
			"edit_time":   editTime + "前",
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

func CreateAppSpaceView(c *gin.Context) {
	/*
		# 创建应用空间条目
	*/
	var res responseData
	userName := c.GetHeader("username")
	var item CreateAppSpaceData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var space DataBaseModel.AppSpace
	DB.Where("space_name = ? AND username = ?", item.SpaceName, userName).First(&space)
	if space.ID != "" {
		res.Err = "空间名称已存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	space.ID = uuid.New().String()
	space.SpaceName = item.SpaceName
	space.UserName = userName
	space.Description = item.SpaceDescription
	space.Background = item.Background
	space.Icon = item.Icon
	space.IconColor = item.IconColor
	space.Collect = false
	err = DB.Create(&space).Error
	if err != nil {
		log.Println("应用空间数据库记录创建失败：", err)
		res.Err = "创建失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	res.Data = map[string]string{
		"id": space.ID,
	}
	res.Msg = "创建成功"
	c.JSON(http.StatusOK, res)
}

func EditAppSpaceView(c *gin.Context) {
	/*
		# 修改app应用空间条目
	*/
	var res responseData
	userName := c.GetHeader("username")
	var item EditAppSpaceData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var space DataBaseModel.AppSpace
	err = DB.Model(space).Where("id = ? AND username = ?", item.SpaceId, userName).Updates(item).Error
	if err != nil {
		log.Println("更新app空间时保存数据库出现错误：", err)
		res.Err = "更新失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	res.Msg = "更新成功"
	c.JSON(http.StatusOK, res)

}

func DeleteAppSpaceView(c *gin.Context) {
	/*
		# 删除app应用空间条目
	*/
	var res responseData
	userName := c.GetHeader("username")
	var item DeleteAppSpaceData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var space DataBaseModel.AppSpace
	DB.Where("id = ? AND username = ?", item.SpaceId, userName).First(&space)
	err = DB.Delete(&space).Error
	if err != nil {
		log.Println("删除app空间时保存数据库出现错误：", err)
		res.Err = "删除失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res)
}

func CreateAppPageView(c *gin.Context) {
	/*
		# 创建app应用空间中的页面
	*/
	var res responseData
	userName := c.GetHeader("username")
	spaceId := c.GetHeader("space_id")
	var item CreateAppPageData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var page DataBaseModel.AppPage
	page.ID = uuid.New().String()
	page.AppSpaceId = spaceId
	page.PageName = item.PageName
	page.UserName = userName
	page.PagePath = item.Tag
	err = DB.Create(&page).Error
	if err != nil {
		log.Println("创建app页面时保存数据库出现错误：", err)
		res.Err = "创建失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	res.Data = map[string]string{
		"id": page.ID,
	}
	res.Msg = "创建成功"
	c.JSON(http.StatusOK, res)
}

func GetAppPageView(c *gin.Context) {
	/*
		# 查询app应用空间中的页面
	*/

	var res responseData
	userName := c.GetHeader("username")
	spaceId := c.GetHeader("space_id")
	release := c.Query("release")
	keyWords := c.Query("keywords")
	var releaseCount, noReleaseCount int64
	var pageList []DataBaseModel.AppPage
	db := DB.Where("app_space_id = ? AND username = ?", spaceId, userName).Order("create_time desc")
	rc := DB.Model(DataBaseModel.AppPage{}).Where("app_space_id = ? AND username = ?", spaceId, userName).Where("is_release = ?", true)
	nrc := DB.Model(DataBaseModel.AppPage{}).Where("app_space_id = ? AND username = ?", spaceId, userName).Where("is_release = ?", false)
	if keyWords != "" {
		db.Where("page_name LIKE ?", "%"+keyWords+"%")
		rc.Where("page_name LIKE ?", "%"+keyWords+"%")
		nrc.Where("page_name LIKE ?", "%"+keyWords+"%")
	}
	rc.Count(&releaseCount)
	nrc.Count(&noReleaseCount)
	switch {
	case release == "":
		db.Find(&pageList)
	case release == "true":
		db.Where("is_release = ?", true).Find(&pageList)
	case release == "false":
		db.Where("is_release = ?", false).Find(&pageList)
	}
	var pageDataList []map[string]interface{}
	for _, page := range pageList {
		p := map[string]interface{}{
			"id":          page.ID,
			"name":        page.PageName,
			"create_time": page.CreatedAt.Local().Format("2006年01月02日"),
		}
		pageDataList = append(pageDataList, p)
	}
	res.Data = map[string]interface{}{
		"data":             pageDataList,
		"all_count":        releaseCount + noReleaseCount,
		"release_count":    releaseCount,
		"no_release_count": noReleaseCount,
	}
	c.JSON(http.StatusOK, res)
}

func EditAppPageView(c *gin.Context) {
	/*
		# 修改app应用空间中的页面、web设计器页面
	*/
	var res responseData
	userName := c.GetHeader("username")
	var item EditAppPageData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println("设置app空间页面时出现数据错误：", err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var page DataBaseModel.AppPage
	DB.Where("id = ? AND app_space_id = ? AND username = ?", item.PageId, item.SpaceId, userName).First(&page)
	if page.ID == "" {
		log.Println("设置app空间页面时未查询到数，相关数据是：", item)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	page.PageName = item.PageName
	page.PagePath = item.Tag
	err = DB.Save(&page).Error
	if err != nil {
		log.Println("设置app空间页面时保存数据库出现错误：", err)
		res.Err = "设置失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	res.Msg = "设置成功"
	c.JSON(http.StatusOK, res)
}

func DeleteAppPageView(c *gin.Context) {
	/*
		# 删除app应用空间中的页面
	*/
	var res responseData
	userName := c.GetHeader("username")
	var item DeleteAppPageData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var page DataBaseModel.AppPage
	DB.Where("id = ? AND app_space_id = ? AND username = ?", item.PageId, item.SpaceId, userName).First(&page)
	err = DB.Delete(&page).Error
	if err != nil {
		log.Println("删除app空间页面时保存数据库出现错误：", err)
		res.Err = "删除失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res)
}
