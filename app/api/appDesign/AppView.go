package API

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"yssim-go/app/DataBaseModel"
	"yssim-go/app/service"
	"yssim-go/config"
	"yssim-go/library/timeConvert"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var DB = config.DB

func AppModelMarkView(c *gin.Context) {
	/*
		# 标记模型为应用可用数据源, 并编译为可执行文件，若编译失败则不作为数据源显示
	*/
	// TODO： 徐庆达
	var res responseData
	userName := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var item AppModelMarkData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	// 检测数据源名称是否重复
	var record DataBaseModel.AppDataSource
	DB.Where("package_id = ? AND username = ? AND ground_name = ? AND data_source_name = ?", item.PackageId, userName, item.GroundName, item.DataSourceName).First(&record)
	if record.ID != "" {
		c.JSON(http.StatusBadRequest, "名称重复")
		return
	}
	CompilePath := "static/UserFiles/modelDataSource/" + userName + "/" + strings.ReplaceAll(item.ModelName, ".", "-") + "/" + time.Now().Local().Format("20060102150405") + "/"
	// 数据源表中创建一条记录
	dataSource := DataBaseModel.AppDataSource{
		ID:             uuid.New().String(),
		UserName:       userName,
		UserSpaceId:    userSpaceId,
		PackageId:      item.PackageId,
		ModelName:      item.ModelName,
		CompilePath:    CompilePath,
		ExperimentId:   item.ExperimentId,
		GroundName:     item.GroundName,
		DataSourceName: item.DataSourceName,
		CompileStatus:  0,
	}
	err = DB.Create(&dataSource).Error
	record = dataSource
	if err != nil {
		log.Println("标记数据源时创建数据库记录失败： ", err)
		res.Err = "创建失败"
		c.JSON(http.StatusOK, res)
		return
	}
	_, err = service.GrpcTranslate(record)
	if err != nil {
		log.Println("提交任务失败： ", err)
		res.Err = "创建失败"
		c.JSON(http.StatusOK, res)
		return
	}
	res.Msg = "提交任务成功，请等待编译完成！"
	c.JSON(http.StatusOK, res)

}

func MultipleSimulateView(c *gin.Context) {
	/*
		# 多轮仿真接口
	*/
	// TODO： 徐庆达
	var res responseData
	//userName := c.GetHeader("username")
	//userSpaceId := c.GetHeader("space_id")
	var item AppMultipleSimulateData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	err = service.GrpcRunResult(item.AppPageId, item.SingleSimulationInputData)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "多轮仿真失败！")
		return
	}
	res.Msg = "任务提交成功，等待多轮仿真完成。"
	c.JSON(http.StatusOK, res)
	return

}

func GetAppSimulateResultView(c *gin.Context) {
	/*
		# 读取AppPage仿真结果
	*/
	// TODO： 徐庆达
	var res responseData
	var item GetSimResData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	data, err := service.AppSimulateResult(item.AppPageId, item.Variable)
	if err != nil {
		log.Println(err)
		res.Msg = "读取失败。"
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res.Data = data
	res.Msg = "读取成功。"
	c.JSON(http.StatusOK, res)
}

func GetAppReleaseResultView(c *gin.Context) {
	/*
		# 读取多轮仿真结果csv数据接口
	*/
	// TODO： 徐庆达
	var res responseData
	var item GetReleaseResData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	data, err := service.AppReleaseResult(item.AppPageId)
	if err != nil {
		log.Println(err)
		res.Msg = "读取失败。"
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res.Data = data
	res.Msg = "读取成功。"
	c.JSON(http.StatusOK, res)
}
func GetModelStateView(c *gin.Context) {
	/*
	   ## 获取仿真状态 返回2表示在仿真中或者发布中  返回4表示仿真结束或者发布结束
	*/
	// TODO： 徐庆达
	appPageId := c.Query("app_page_id")
	var appPageRecord DataBaseModel.AppPage
	DB.Where("id = ?", appPageId).First(&appPageRecord)
	var res responseData
	resData := map[string]int{
		"ReleaseState":  2,
		"SimulateState": 2,
	}
	// 如果状态是3（失败）或者4（仿真结束），返回4，否则返回2
	if appPageRecord.ReleaseState == 4 || appPageRecord.ReleaseState == 3 {
		resData["ReleaseState"] = 4
	}
	if appPageRecord.SimulateState == 4 || appPageRecord.ReleaseState == 3 {
		resData["SimulateState"] = 4
	}
	res.Data = resData
	c.JSON(http.StatusOK, res)

}
func GetDataSourceGroupView(c *gin.Context) {
	/*
		# 获取用户数据源分组
	*/
	// TODO： 徐庆达
	var res responseData
	userName := c.GetHeader("username")
	var group []DataBaseModel.AppDataSource
	DB.Select("ground_name").Where("username = ? AND compile_status = ?", userName, 4).Group("ground_name").Find(&group)
	groupData := make([]string, 0)

	for _, g := range group {
		groupData = append(groupData, g.GroundName)
	}
	res.Data = groupData
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
	db = db.Session(&gorm.Session{})
	db.Order("update_time desc").Find(&allAppSpaceList)
	db.Where("last_login_time <> ?", 0).Order("last_login_time desc").Limit(5).Find(&recentAppSpaceList)
	allAppSpace := make([]map[string]interface{}, 0)
	recentAppSpace := make([]map[string]interface{}, 0)
	for _, space := range allAppSpaceList {
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
	for _, space := range recentAppSpaceList {
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
	DB.Where("space_name = ? AND username = ? AND id <> ?", item.SpaceName, userName, item.SpaceId).First(&space)
	if space.ID != "" {
		res.Err = "应用名称已存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	DB.Where("id = ? AND username = ?", item.SpaceId, userName).First(&space)
	if space.ID != "" {
		space.SpaceName = item.SpaceName
		space.Description = item.Description
		space.Background = item.Background
		space.Icon = item.Icon
		space.IconColor = item.IconColor
		space.Collect = item.Collect
		err = DB.Save(&space).Error
		if err != nil {
			log.Println("更新app空间时保存数据库出现错误：", err)
			res.Err = "编辑失败"
			res.Status = 2
			c.JSON(http.StatusOK, res)
			return
		}
	}
	res.Msg = "编辑成功"
	c.JSON(http.StatusOK, res)
}

func AppSpaceCollectView(c *gin.Context) {
	/*
		# 收藏app应用空间条目
	*/
	var res responseData
	userName := c.GetHeader("username")
	var item AppSpaceCollectData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var space DataBaseModel.AppSpace
	err = DB.Model(&space).Where("id IN ? AND username = ?", item.SpaceId, userName).Updates(map[string]interface{}{"collect": item.Collect}).Error
	if err != nil {
		log.Println("更新app空间时保存数据库出现错误：", err)
		res.Err = "收藏失败"
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
	err = DB.Where("id IN ? AND username = ?", item.SpaceId, userName).Delete(&space).Error
	//err = DB.Delete(&space)
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
	var item CreateAppPageData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println("创建页面时验证数据出错：", err)
		c.JSON(http.StatusBadRequest, "验证失败")
		return
	}
	var page DataBaseModel.AppPage
	DB.Where("app_space_id = ? AND page_name = ? AND username = ?", item.SpaceId, item.PageName, userName).Or("app_space_id = ? AND page_path = ? AND username = ?", item.SpaceId, item.Tag, userName).First(&page)
	if page.ID != "" {
		switch {
		case page.PageName == item.PageName:
			res.Err = "应用页面名称已存在"
		case page.PagePath == item.Tag:
			res.Err = "应用页面标识已存在"
		}
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	var pageNew DataBaseModel.AppPage
	pageNew.ID = uuid.New().String()
	pageNew.AppSpaceId = item.SpaceId
	pageNew.PageName = item.PageName
	pageNew.UserName = userName
	pageNew.PagePath = item.Tag
	err = DB.Create(&pageNew).Error
	if err != nil {
		log.Println("创建app页面时保存数据库出现错误：", err)
		res.Err = "创建失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	res.Data = map[string]string{
		"id": pageNew.ID,
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
	release := c.Query("release")
	spaceId := c.Query("space_id")
	keyWords := c.Query("keywords")
	var releaseCount, noReleaseCount int64
	var pageList []DataBaseModel.AppPage
	DB.Model(DataBaseModel.AppSpace{}).Where("id = ? AND username = ?", spaceId, userName).UpdateColumn("last_login_time", time.Now().Local().Unix())

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
	case release == "1":
		db.Where("is_release = ?", true).Find(&pageList)
	case release == "0":
		db.Where("is_release = ?", false).Find(&pageList)
	}
	var pageDataList []map[string]interface{}
	for _, page := range pageList {
		p := map[string]interface{}{
			"id":          page.ID,
			"name":        page.PageName,
			"create_time": page.CreatedAt.Local().Format("2006年01月02日"),
			"update_time": page.UpdatedAt.Local().Format("2006年01月02日"),
			"tag":         page.PagePath,
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
		log.Println("编辑app空间页面时出现数据错误：", err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var page DataBaseModel.AppPage
	var pageTagAndName DataBaseModel.AppPage
	DB.Where("app_space_id = ? AND page_path = ? AND username = ? AND id <> ?", item.SpaceId, item.Tag, userName, item.PageId).Or("app_space_id = ? AND page_name = ? AND username = ? AND id <> ?", item.SpaceId, item.PageName, userName, item.PageId).First(&pageTagAndName)
	if pageTagAndName.PagePath == item.Tag || pageTagAndName.PageName == item.PageName {
		switch {
		case pageTagAndName.PageName == item.PageName:
			res.Err = "应用页面名称已存在"
		case pageTagAndName.PagePath == item.Tag:
			res.Err = "应用页面标识已存在"
		}
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	DB.Where("id = ? AND app_space_id = ? AND username = ?", item.PageId, item.SpaceId, userName).First(&page)
	if page.ID == "" {
		log.Println("编辑app空间页面时未查询到数，相关数据是：", item)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	page.PageName = item.PageName
	page.PagePath = item.Tag
	err = DB.Save(&page).Error
	if err != nil {
		log.Println("设置app空间页面时保存数据库出现错误：", err)
		res.Err = "编辑失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	res.Msg = "编辑成功"
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
		log.Println("DeleteAppPageView err", err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var page DataBaseModel.AppPage
	err = DB.Model(&page).Where("id = ?  AND username = ?", item.PageId, userName).Delete(&page).Error
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

func EditAppPageDesignView(c *gin.Context) {
	/*
		# 修改app页面设计
	*/
	var res responseData
	userName := c.GetHeader("username")
	var item EditAppPageDesignData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	var page DataBaseModel.AppPage
	DB.Where("id = ? AND username = ?", item.Id, userName).First(&page)
	if page.ID != "" {
		page.PageWidth = item.PageWidth
		page.PageHeight = item.PageHeight
		page.Background = item.Background
		page.BackgroundColor = item.BackgroundColor
		err = DB.Save(&page).Error
		if err != nil {
			log.Println("修改app页面设计出现错误：", err)
			res.Err = "编辑失败"
			res.Status = 2
			c.JSON(http.StatusOK, res)
			return
		}
	}
	res.Msg = "编辑成功"
	c.JSON(http.StatusOK, res)

}

func CreatePageComponentView(c *gin.Context) {
	/*
		# app应用页面新增组件
	*/
	var res responseData
	userName := c.GetHeader("username")
	//spaceId := c.GetHeader("space_id")
	var item CreatePageComponentData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println("数据错误：", err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var page DataBaseModel.AppPage
	//DB.Where("id = ? AND username = ?", item.PageId, userName).First(&page)
	DB.Where("id = ? AND app_space_id = ? AND username = ?", item.PageId, item.SpaceId, userName).First(&page)
	if page.ID == "" {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var pageComponent DataBaseModel.AppPageComponent
	pageComponent.ID = uuid.New().String()
	pageComponent.PageId = item.PageId
	pageComponent.Height = item.Height
	pageComponent.Width = item.Width
	pageComponent.Type = item.Type
	pageComponent.PositionX = item.PositionX
	pageComponent.PositionY = item.PositionY
	pageComponent.Angle = item.Angle
	pageComponent.HorizontalFlip = item.HorizontalFlip
	pageComponent.VerticalFlip = item.VerticalFlip
	pageComponent.Opacity = item.Opacity
	pageComponent.OtherConfiguration = item.OtherConfiguration
	pageComponent.ZIndex = item.ZIndex
	pageComponent.Styles = item.Styles
	pageComponent.Events = item.Events
	pageComponent.ChartConfig = item.ChartConfig
	pageComponent.Option = item.Option
	pageComponent.ComponentPath = item.ComponentPath
	pageComponent.Hide = item.Hide
	pageComponent.Lock = item.Lock
	pageComponent.IsGroup = item.IsGroup
	err = DB.Create(&pageComponent).Error
	if err != nil {
		log.Println("创建app页面时保存数据库出现错误：", err)
		res.Err = "创建失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	res.Data = map[string]string{
		"id": pageComponent.ID,
	}
	res.Msg = "创建成功"
	c.JSON(http.StatusOK, res)
}

func GetPageComponentView(c *gin.Context) {
	/*
		# app应用页面组件查询组件
	*/

	var res responseData
	userName := c.GetHeader("username")
	spaceId := c.Query("space_id")
	pageId := c.Query("page_id")
	var page DataBaseModel.AppPage
	var componentList []DataBaseModel.AppPageComponent
	DB.Where("id = ? AND app_space_id = ? AND username = ?", pageId, spaceId, userName).First(&page)
	DB.Where("page_id = ?", page.ID).Find(&componentList)

	var componentDataList []map[string]interface{}

	for _, c := range componentList {
		p := map[string]interface{}{
			"id":                  c.ID,
			"type":                c.Type,
			"width":               c.Width,
			"height":              c.Height,
			"position_x":          c.PositionX,
			"position_y":          c.PositionY,
			"angle":               c.Angle,
			"horizontal_flip":     c.HorizontalFlip,
			"vertical_flip":       c.VerticalFlip,
			"opacity":             c.Opacity,
			"other_configuration": c.OtherConfiguration,
			"z_index":             c.ZIndex,
			"styles":              c.Styles,
			"events":              c.Events,
			"chart_config":        c.ChartConfig,
			"option":              c.Option,
			"component_path":      c.ComponentPath,
			"hide":                c.Hide,
			"lock":                c.Lock,
			"is_group":            c.IsGroup,
		}
		componentDataList = append(componentDataList, p)
	}
	res.Data = map[string]interface{}{
		"components": componentDataList,
		"page": map[string]interface{}{
			"background":       page.Background,
			"background_color": page.BackgroundColor,
			"height":           page.PageHeight,
			"width":            page.PageWidth,
			"release":          page.Release,
		},
	}
	c.JSON(http.StatusOK, res)
}

func EditPageComponentView(c *gin.Context) {
	/*
		# app应用页面编辑组件
	*/
	var res responseData
	userName := c.GetHeader("username")
	var item EditPageComponentData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println("", err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var page DataBaseModel.AppPage
	DB.Where("id = ? AND app_space_id = ? AND username = ?", item.PageId, item.SpaceId, userName).First(&page)
	if page.ID == "" {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	//err = DB.Model(DataBaseModel.AppPageComponent{}).Select("*").Where("id = ? AND page_id = ?", item.Id, item.PageId).Updates(&item).Error
	err = DB.Model(DataBaseModel.AppPageComponent{}).Where("id = ? AND page_id = ?", item.Id, item.PageId).Updates(map[string]interface{}{
		"type": item.Type,
		//"input_output": item.InputOutput,
		"width":               item.Width,
		"height":              item.Height,
		"position_x":          item.PositionX,
		"position_y":          item.PositionY,
		"angle":               item.Angle,
		"horizontal_flip":     item.HorizontalFlip,
		"vertical_flip":       item.VerticalFlip,
		"opacity":             item.Opacity,
		"other_configuration": item.OtherConfiguration,
		"z_index":             item.ZIndex,
		"styles":              item.Styles,
		"events":              item.Events,
		"chart_config":        item.ChartConfig,
		"option":              item.Option,
		"component_path":      item.ComponentPath,
		"hide":                item.Hide,
		"lock":                item.Lock,
		"is_group":            item.IsGroup,
	}).Error

	if err != nil {
		log.Println("编辑app页面中组件时保存数据库出现错误：", err)
		res.Err = "编辑失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	res.Msg = "更新成功"
	c.JSON(http.StatusOK, res)
}

func DeletePageComponentView(c *gin.Context) {
	/*
		# app应用页面删除组件
	*/
	var res responseData
	userName := c.GetHeader("username")
	var item DeletePageComponentData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	//var space DataBaseModel.AppSpace
	var page DataBaseModel.AppPage
	var components DataBaseModel.AppPageComponent
	//DB.Where("id = ? AND username = ? ",item.)
	DB.Where("id = ? AND username = ? ", item.PageId, userName).First(&page)
	err = DB.Model(DataBaseModel.AppPageComponent{}).Where("id IN ? AND page_id = ?", item.ComponentsList, page.ID).Delete(&components).Error
	//err = DB.Model(DataBaseModel.AppPageComponent{}).Where("id IN ? AND page_id = ?", item.ComponentsList, item.PageId).Delete(&components).Error
	if err != nil {
		log.Println("删除app空间页面组件时保存数据库出现错误：", err)
		res.Err = "删除失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res)
}

func GetDatasourceView(c *gin.Context) {
	/*
		# 获取数据源相关信息
	*/
	// TODO： 徐庆达
	var res responseData
	userName := c.GetHeader("username")
	groundName := c.Query("ground_name")
	searchName := c.Query("search_name")
	var dataList []map[string]interface{}
	var appDataSourceRecord []DataBaseModel.AppDataSource
	DB.Where("username = ? AND ground_name = ? AND compile_status = ? AND data_source_name LIKE ?", userName, groundName, 4, "%"+searchName+"%").Order("create_time desc").Find(&appDataSourceRecord)
	for i := 0; i < len(appDataSourceRecord); i++ {
		data := map[string]interface{}{
			"id":                 appDataSourceRecord[i].ID,
			"source_name":        appDataSourceRecord[i].DataSourceName,
			"compile_model_name": appDataSourceRecord[i].ModelName,
		}
		dataList = append(dataList, data)
	}
	res.Data = dataList
	c.JSON(http.StatusOK, res)

}

func DatasourceDeleteView(c *gin.Context) {
	/*
		# 删除数据源
	*/
	// TODO： 徐庆达
	var res responseData
	var item DeleteDatasourceData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	// 删除数据库记录
	var page DataBaseModel.AppDataSource
	err = DB.Model(DataBaseModel.AppDataSource{}).Where("id = ?", item.DataSourceID).Delete(&page).Error
	if err != nil {
		log.Println("删除app数据源出现错误：", err)
		res.Err = "删除失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	// 删除文件
	err = os.RemoveAll(page.CompilePath)
	if err != nil {
		log.Println(err)
	}
	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res)

}

func DataSourceRenameView(c *gin.Context) {
	/*
		# 重命名数据源
	*/
	// TODO： 徐庆达
	var item DataSourceRenameData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	var res responseData
	err = DB.Model(&DataBaseModel.AppDataSource{}).Where("id = ?", item.DataSourceID).Update("data_source_name", item.NewName).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "修改失败")
		return
	}
	res.Msg = "修改成功"
	c.JSON(http.StatusOK, res)

}

func GetDatasourceInputOutputView(c *gin.Context) {
	/*
		# 获取数据源输入与输出接口
	*/
	// TODO： 徐庆达

	//username := c.GetHeader("username")
	//userSpaceId := c.GetHeader("space_id")
	recordId := c.Query("record_id")
	parentNode := c.Query("parent_node")
	keyWords := c.Query("key_words")
	var record DataBaseModel.AppDataSource
	DB.Where("id = ? AND compile_status = ?", recordId, "4").First(&record)
	if record.ID == "" {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}

	var res responseData
	if record.CompilePath != "" {
		if record.CompilePath == "DM" {
			//DM生成的fmu解压后的xml文件
			result := service.DymolaSimulationResultTree(record.CompilePath+"result_init.xml", parentNode, keyWords)
			res.Data = result
		} else {
			//OMC仿真完输出的xml文件
			result := service.SimulationResultTree(record.CompilePath+"result_init.xml", parentNode, keyWords)
			res.Data = result
		}

	} else {
		res.Err = "查询失败"
		res.Status = 2
	}
	c.JSON(http.StatusOK, res)
}

func GetPageInputOutputView(c *gin.Context) {
	/*
		# 获取app应用页面的数据源输入与输出接口
	*/
	var res responseData
	userName := c.GetHeader("username")
	pageId := c.Query("id")
	var page DataBaseModel.AppPage
	DB.Where("id = ? AND username = ?", pageId, userName).First(&page)
	data := map[string]interface{}{
		"input":          page.Input,
		"output":         page.Output,
		"data_source_id": page.DataSourceId,
	}
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func SetPageInputOutputView(c *gin.Context) {
	/*
		# 设置app应用页面的数据源输入与输出接口
	*/
	var res responseData
	userName := c.GetHeader("username")
	var item SetPageInputOutputData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	var page DataBaseModel.AppPage
	err = DB.Where("id = ? AND app_space_id= ? AND username = ?", item.PageId, item.SpaceId, userName).First(&page).Error
	if err != nil {
		log.Println("设置app应用页面的数据源时保存数据库出现错误：", err)
		res.Err = "设置失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	page.Input = item.Input
	page.Output = item.Output
	page.DataSourceId = item.DataSourceId
	DB.Save(&page)
	res.Msg = "设置成功"
	c.JSON(http.StatusOK, res)
}

func GetPageComponentInputOutputView(c *gin.Context) {
	/*
		# app应用页面组件的数据源输入与输出查询接口
	*/
	var res responseData
	id := c.Query("id")
	pageId := c.Query("page_id")
	spaceId := c.Query("space_id")
	var component DataBaseModel.AppPageComponent
	var page DataBaseModel.AppPage
	DB.Where("app_space_id = ? AND id = ? ", spaceId, pageId).First(&page)
	if page.ID == "" {
		log.Println("未查询到页面数据")
		c.JSON(http.StatusBadRequest, "")
		return
	}
	DB.Where("id = ? AND page_id= ? ", id, pageId).First(&component)
	if component.ID == "" {
		log.Println("未查询到数据")
		c.JSON(http.StatusBadRequest, "")
		return
	}

	input := map[string]interface{}{
		"inputName": component.InputName,
		"max":       component.Max,
		"min":       component.Min,
		"interval":  component.Interval,
	}
	output := map[string]interface{}{
		"output": component.Output,
	}
	data := map[string]interface{}{
		"input":  input,
		"output": output,
	}
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func SetPageComponentInputOutputView(c *gin.Context) {
	/*
		# app应用页面组件的数据源输入与输出绑定接口
	*/
	var res responseData
	var component DataBaseModel.AppPageComponent

	var item SetPageComponentsInputOutputData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		res.Err = "设置参数错误!"
		c.JSON(http.StatusBadRequest, res)
		return
	}
	err = DB.Where("id = ? AND page_id = ?", item.Id, item.PageId).First(&component).Error
	if err != nil {
		log.Println("app应用页面组件的数据源输入与输出绑定时保存数据库出现错误：", err)
		res.Err = "设置失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	component.InputName = item.InputName
	component.Output = item.Output
	component.Max = item.Min
	component.Min = item.Min
	component.Interval = item.Interval
	DB.Save(&component)
	res.Msg = "设置成功"
	c.JSON(http.StatusOK, res)
}

func AppPagePreviewView(c *gin.Context) {
	/*
		# 预览页面组件相关接口
	*/
	var res responseData
	userName := c.GetHeader("username")
	spaceId := c.Query("space_id")
	pageId := c.Query("page_id")
	var space DataBaseModel.AppSpace
	var page DataBaseModel.AppPage
	var components []DataBaseModel.AppPageComponent
	DB.Where("id = ? AND username = ?", spaceId, userName).First(&space)
	DB.Where("id = ? AND app_space_id = ?", pageId, space.ID).First(&page)
	DB.Where("page_id = ?", page.ID).Find(&components)
	pageData := make(map[string]interface{}, 0)
	pageData["width"] = page.PageWidth
	pageData["height"] = page.PageHeight
	pageData["background"] = page.Background
	pageData["color"] = page.Color
	componentsData := make([]map[string]interface{}, 0)
	for _, component := range components {
		d := make(map[string]interface{}, 0)
		d["id"] = component.ID
		d["type"] = component.Type
		d["input_name"] = component.InputName
		d["output_name"] = component.Option
		d["max"] = component.Max
		d["min"] = component.Min
		d["interval"] = component.Interval
		d["width"] = component.Width
		d["height"] = component.Height
		d["position_x"] = component.PositionX
		d["position_y"] = component.PositionY
		d["angle"] = component.Angle
		d["horizontal_flip"] = component.HorizontalFlip
		d["vertical_flip"] = component.VerticalFlip
		d["opacity"] = component.Opacity
		d["z_index"] = component.ZIndex
		d["styles"] = component.Styles
		d["events"] = component.Events
		d["chart_config"] = component.ChartConfig
		d["option"] = component.Option
		d["component_path"] = component.ComponentPath
		d["hide"] = component.Hide
		d["lock"] = component.Lock
		d["other_configuration"] = component.OtherConfiguration
		componentsData = append(componentsData, d)
	}
	res.Data = map[string]interface{}{
		"page":       pageData,
		"components": componentsData,
	}
	c.JSON(http.StatusOK, res)
}

func SetComponentBasicInformationView(c *gin.Context) {
	var res responseData
	//userName := c.GetHeader("username")
	var item CreateComponentBasesData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println("", err)
		c.JSON(http.StatusBadRequest, "")
		return
	}

	var component DataBaseModel.AppComponentBases
	component.ID = uuid.New().String()
	component.TopLevelName = item.TopLevelName
	component.SecondLevelName = item.SecondLevelName
	component.Type = item.Type
	component.Width = item.Width
	component.Height = item.Height
	component.Margin = item.Margin
	component.PositionX = item.PositionX
	component.PositionY = item.PositionY
	component.Angle = item.Angle
	component.HorizontalFlip = item.HorizontalFlip
	component.VerticalFlip = item.VerticalFlip
	component.Opacity = item.Opacity
	component.OtherConfiguration = item.OtherConfiguration
	//component.CreatedAt = time.Now().Local().Format("20060102150405")
	err = DB.Save(&component).Error
	if err != nil {
		log.Println("创建web组件设计页面数据库失败！", err)
		res.Err = "创建失败,其稍后再试!"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	data := map[string]interface{}{
		"componentId": component.ID,
	}
	res.Data = data
	res.Msg = "创建成功!"
	c.JSON(http.StatusOK, res)
}

func GetComponentBasicInformationView(c *gin.Context) {
	var res responseData
	//id := c.Query("id")

	var component []DataBaseModel.AppComponentBases
	DB.Find(&component)

	//if err != nil {
	//	c.JSON(http.StatusBadRequest, "not found")
	//	return
	//}
	res.Data = component
	c.JSON(http.StatusOK, res)
}
