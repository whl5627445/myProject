package resource

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"yssim-go/app/DataBaseModel"
	"yssim-go/app/DataType"
	serviceV2 "yssim-go/app/v2/service"
	"yssim-go/config"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var DB = config.DB

func GetRootListView(c *gin.Context) {
	/*
		# 获取资源库顶层文件夹列表
	*/
	var res DataType.ResponseData

	// 预创建的顶层文档
	rootPreFoldersList := []string{"工况", "动力成品件", "电气成品件", "液压成品件", "环控成品件", "燃油成品件"}
	var rootPreFolders []DataBaseModel.YssimResourceLib
	DB.Unscoped().Where("parent_id = ? AND name IN ?", "", rootPreFoldersList).Find(&rootPreFolders)
	if len(rootPreFolders) == 0 {
		for _, name := range rootPreFoldersList {
			preRootFolder := DataBaseModel.YssimResourceLib{
				ParentId:    "",
				ID:          uuid.New().String(),
				UserName:    "",
				FolderFile:  "folder",
				FullPath:    "",
				Name:        name,
				Description: "",
				FilePath:    "",
			}

			DB.Create(&preRootFolder)
		}
	}

	// 获取预创建的顶层文件夹
	var preRootFolders []DataBaseModel.YssimResourceLib
	if err := DB.Where("parent_id = ? AND name IN ?", "", rootPreFoldersList).Order("create_time desc").Find(&preRootFolders).Error; err != nil {
		res.Err = "查询失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 获取其他顶层文件夹
	var rootFolders []DataBaseModel.YssimResourceLib
	if err := DB.Where("parent_id = ? AND name NOT IN ?", "", rootPreFoldersList).Order("create_time desc").Find(&rootFolders).Error; err != nil {
		res.Err = "查询失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 返回数据
	rootFolderListData := make([]map[string]any, 0)

	for _, r := range preRootFolders {
		data := map[string]any{
			"id":          r.ID,
			"parent_id":   r.ParentId,
			"username":    r.UserName,
			"name":        r.Name,
			"type":        r.FolderFile,
			"type_str":    "文件夹",
			"full_path":   r.FullPath,
			"description": r.Description,
			"create_time": r.CreatedAt,
			"update_time": r.UpdatedAt,
		}

		rootFolderListData = append(rootFolderListData, data)
	}

	for _, r := range rootFolders {
		data := map[string]any{
			"id":          r.ID,
			"parent_id":   r.ParentId,
			"username":    r.UserName,
			"name":        r.Name,
			"type":        r.FolderFile,
			"type_str":    "文件夹",
			"full_path":   r.FullPath,
			"description": r.Description,
			"create_time": r.CreatedAt,
			"update_time": r.UpdatedAt,
		}

		rootFolderListData = append(rootFolderListData, data)
	}

	res.Data = rootFolderListData
	c.JSON(http.StatusOK, res)
}

func GetResourceListView(c *gin.Context) {
	/*
		# 获取子文件夹文件列表
	*/
	var res DataType.ResponseData
	parentId := c.Query("parent_id")
	pageNumStr := c.Query("page_num") //页码
	pageNum, _ := strconv.Atoi(pageNumStr)
	pageSizeStr := c.Query("page_size") //每页条数
	pageSize, _ := strconv.Atoi(pageSizeStr)
	if pageSize <= 0 {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	var total int64 //总条数
	DB.Where("parent_id = ?", parentId).Find(&DataBaseModel.YssimResourceLib{}).Count(&total)
	pageCount := math.Ceil(float64(total) / float64(pageSize)) //总页数

	// 获取子文件夹文件列表
	var subResources []DataBaseModel.YssimResourceLib
	if err := DB.Limit(pageSize).Offset((pageNum-1)*pageSize).Where("parent_id = ?", parentId).Order("folder_file desc, create_time desc").Find(&subResources).Error; err != nil {
		res.Err = "查询失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	subFolders := make([]map[string]any, 0)
	subFiles := make([]map[string]any, 0)

	for _, r := range subResources {
		switch r.FolderFile {
		case "folder":
			data := map[string]any{
				"id":          r.ID,
				"parent_id":   r.ParentId,
				"username":    r.UserName,
				"name":        r.Name,
				"type":        r.FolderFile,
				"type_str":    "文件夹",
				"full_path":   r.FullPath,
				"description": r.Description,
				"create_time": r.CreatedAt,
				"update_time": r.UpdatedAt,
			}
			if r.ParentId == "" {
				data["full_path"] = "/"
			}
			subFolders = append(subFolders, data)
		case "file":
			data := map[string]any{
				"id":          r.ID,
				"parent_id":   r.ParentId,
				"username":    r.UserName,
				"name":        r.Name,
				"type":        r.FolderFile,
				"type_str":    "文件",
				"full_path":   r.FullPath,
				"description": r.Description,
				"create_time": r.CreatedAt,
				"update_time": r.UpdatedAt,
			}
			subFiles = append(subFiles, data)
		}
	}

	resData := map[string]any{
		"page_count": pageCount,
		"total":      total,
		"folder":     subFolders,
		"file":       subFiles,
	}

	res.Data = resData
	c.JSON(http.StatusOK, res)
}

func SearchSubListView(c *gin.Context) {
	/*
		# 搜索文件夹和文件
	*/
	var res DataType.ResponseData
	keyWords := c.Query("keywords")

	var total int64 //总条数
	DB.Where("name LIKE ?", "%"+keyWords+"%").Find(&DataBaseModel.YssimResourceLib{}).Count(&total)

	var subResourceFolders []DataBaseModel.YssimResourceLib
	if err := DB.Where("folder_file = ? AND name LIKE ?", "folder", "%"+keyWords+"%").Order("create_time desc").Find(&subResourceFolders).Error; err != nil {
		res.Err = "查询失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	var subResourceFiles []DataBaseModel.YssimResourceLib
	if err := DB.Where("folder_file = ? AND name LIKE ?", "file", "%"+keyWords+"%").Order("create_time desc").Find(&subResourceFiles).Error; err != nil {
		res.Err = "查询失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	datalist := make([]map[string]any, 0)

	for _, r := range subResourceFolders {
		data := map[string]any{
			"id":          r.ID,
			"parent_id":   r.ParentId,
			"username":    r.UserName,
			"name":        r.Name,
			"type":        r.FolderFile,
			"type_str":    "文件夹",
			"full_path":   r.FullPath,
			"description": r.Description,
			"create_time": r.CreatedAt,
			"update_time": r.UpdatedAt,
		}
		if r.ParentId == "" {
			data["full_path"] = "/"
		}

		// 获取上级每一层文件夹的id
		path_list := []map[string]string{}
		strs := strings.Split(data["full_path"].(string), "/")
		if len(strs) != 0 {
			for _, subPath := range strs {
				if subPath == "" {
					continue
				}
				var subResourceFile DataBaseModel.YssimResourceLib
				if err := DB.Where("name = ?", subPath).First(&subResourceFile).Error; err != nil {
					res.Err = "查询失败"
					res.Status = 2
					c.JSON(http.StatusOK, res)
					return
				}

				currentPath := map[string]string{"name": subResourceFile.Name, "id": subResourceFile.ID}
				path_list = append(path_list, currentPath)
			}
		}

		itself := map[string]string{"name": r.Name, "id": r.ID}
		path_list = append(path_list, itself)
		data["path_list"] = path_list
		datalist = append(datalist, data)
	}

	for _, r := range subResourceFiles {
		data := map[string]any{
			"id":          r.ID,
			"parent_id":   r.ParentId,
			"username":    r.UserName,
			"name":        r.Name,
			"type":        r.FolderFile,
			"type_str":    "文件",
			"full_path":   r.FullPath,
			"description": r.Description,
			"create_time": r.CreatedAt,
			"update_time": r.UpdatedAt,
		}

		// 获取上级每一层文件夹的id
		path_list := []map[string]string{}
		strs := strings.Split(data["full_path"].(string), "/")
		if len(strs) != 0 {
			for _, subPath := range strs {
				if subPath == "" {
					continue
				}
				var subResourceFile DataBaseModel.YssimResourceLib
				if err := DB.Where("name = ?", subPath).First(&subResourceFile).Error; err != nil {
					res.Err = "查询失败"
					res.Status = 2
					c.JSON(http.StatusOK, res)
					return
				}

				currentPath := map[string]string{"name": subResourceFile.Name, "id": subResourceFile.ID}
				path_list = append(path_list, currentPath)
			}
		}

		itself := map[string]string{"name": r.Name, "id": r.ID}
		path_list = append(path_list, itself)
		data["path_list"] = path_list
		datalist = append(datalist, data)
	}

	res.Data = map[string]any{"total": total, "data": datalist}
	c.JSON(http.StatusOK, res)
}

func EditResourceInfoView(c *gin.Context) {
	/*
		# 编辑文件夹文件基本信息
		开发人： 周强
	*/
	var res DataType.ResponseData
	var item DataType.EditResourceInfoData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	// 验证应用名称命名规则
	matchSpaceName, _ := regexp.MatchString("^[_0-9a-zA-Z\u4e00-\u9fa5]+$", item.Name) // 由中文、字母、数字、下划线验证
	if !matchSpaceName {
		res.Err = "名称只能由中文、字母、数字、下划线组成"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	var resourceLib DataBaseModel.YssimResourceLib
	if DB.Where("id = ?", item.ID).First(&resourceLib); resourceLib.ID == "" {
		res.Err = "当前资源不存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 验证名称是否已存在
	var resourceLibName DataBaseModel.YssimResourceLib
	DB.Where("parent_id = ? AND name = ?", item.ParentId, item.Name).First(&resourceLibName)
	if resourceLibName.ID != "" && resourceLibName.ID != item.ID {
		res.Err = "名称已存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	resourceLib.Name = item.Name
	resourceLib.Description = item.Description
	if err = DB.Save(&resourceLib).Error; err != nil {
		log.Println("编辑资源库文件夹文件基本信息时数据库出现错误：", err)
		res.Err = "编辑失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	var updatingPathResources []*DataBaseModel.YssimResourceLib
	DB.Where("parent_id = ?", item.ID).Find(&updatingPathResources)

	// 广度优先更新所有子节点路径
	level := 0
	for len(updatingPathResources) > 0 {
		var updatingPathResourcesTemp []*DataBaseModel.YssimResourceLib
		for _, updatingResource := range updatingPathResources {
			//更新路径
			curFullPathsOriginal := strings.Split(updatingResource.FullPath, "/")
			curFullPaths := []string{}
			for _, path := range curFullPathsOriginal {
				if path != "" {
					curFullPaths = append(curFullPaths, path)
				}
			}
			curFullPaths[len(curFullPaths)-1-level] = item.Name

			newPath := ""
			for _, path := range curFullPaths {
				newPath = newPath + fmt.Sprintf("%s%s", "/", path)
			}
			updatingResource.FullPath = newPath

			// 添加所有子节点
			if updatingResource.FolderFile == "folder" {
				var subResources []*DataBaseModel.YssimResourceLib
				DB.Where("parent_id = ?", updatingResource.ID).Find(&subResources)
				updatingPathResourcesTemp = append(updatingPathResourcesTemp, subResources...)
			}
		}

		if err = DB.Save(&updatingPathResources).Error; err != nil {
			log.Println("编辑资源库文件夹文件基本信息时数据库出现错误：", err)
			res.Err = "编辑失败"
			res.Status = 2
			c.JSON(http.StatusOK, res)
			return
		}
		updatingPathResources = updatingPathResourcesTemp
		level += 1
	}

	res.Msg = "编辑成功"
	c.JSON(http.StatusOK, res)
}

func DeleteResourceView(c *gin.Context) {
	/*
		# 删除资源文件夹或文件
		开发人： 周强
	*/
	var res DataType.ResponseData
	var item DataType.DeleteResourceData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	// 删除记录
	var resourceLib DataBaseModel.YssimResourceLib
	if err = DB.Where("id = ? ", item.ID).Delete(&resourceLib).Error; err != nil {
		log.Println("删除资源文件夹或文件时数据库出现错误：", err)
		res.Err = "删除失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	var deletingResources []DataBaseModel.YssimResourceLib
	DB.Where("parent_id = ?", item.ID).Find(&deletingResources)

	// 广度优先删除所有子节点
	for len(deletingResources) > 0 {
		var deletingResourcesTemp []DataBaseModel.YssimResourceLib
		for _, deletingResource := range deletingResources {
			DB.Where("id = ?", deletingResource.ID).Delete(&DataBaseModel.YssimResourceLib{})
			if deletingResource.FolderFile == "folder" {
				var deletingResourceTemps []DataBaseModel.YssimResourceLib
				DB.Where("parent_id = ?", deletingResource.ID).Find(&deletingResourceTemps)
				deletingResourcesTemp = append(deletingResourcesTemp, deletingResourceTemps...)
			}
		}
		deletingResources = deletingResourcesTemp
	}

	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res)
}

func CreateResourceFolderView(c *gin.Context) {
	/*
		# 创建资源文件夹
		开发人： 周强
	*/
	var res DataType.ResponseData
	username := c.GetHeader("username")
	var item DataType.CreateResourceFolderData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}

	if utf8.RuneCountInString(item.Name) > 128 {
		res.Err = "名称太长"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 验证名称命名规则
	matchSpaceName, _ := regexp.MatchString("^[_0-9a-zA-Z\u4e00-\u9fa5]+$", item.Name) // 由中文、字母、数字、下划线验证
	if !matchSpaceName {
		res.Err = "名称只能由中文、字母、数字、下划线组成"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	var sameNameSubResource DataBaseModel.YssimResourceLib
	if DB.Where("parent_id = ? AND name = ?", item.ParentId, item.Name).First(&sameNameSubResource); sameNameSubResource.ID != "" {
		res.Err = "名称已存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	var newResourceFolder = DataBaseModel.YssimResourceLib{
		ParentId:    item.ParentId,
		ID:          uuid.New().String(),
		UserName:    username,
		FolderFile:  "folder",
		FullPath:    "",
		Name:        item.Name,
		Description: "",
		FilePath:    "",
	}

	var parentResource DataBaseModel.YssimResourceLib
	if item.ParentId != "" {
		DB.Where("id = ?", item.ParentId).First(&parentResource)
		newResourceFolder.FullPath = parentResource.FullPath + "/" + parentResource.Name
	}

	if err = DB.Create(&newResourceFolder).Error; err != nil {
		log.Println("资源文件夹数据库记录创建失败：", err)
		res.Err = "创建失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	res.Data = newResourceFolder.ID
	res.Msg = "创建成功"
	c.JSON(http.StatusOK, res)
}

func UploadResourceFileView(c *gin.Context) {
	/*
		# 上传资源文件
	*/
	var res DataType.ResponseData
	userName := c.GetHeader("username")
	var fileinfo DataType.UploadResourceFileData
	if err := c.Bind(&fileinfo); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}

	varFile, err := c.FormFile("file")
	if err != nil {
		log.Println("上传映射配置表文件时出现错误：", err)
		c.JSON(http.StatusBadRequest, "")
		return
	}

	if utf8.RuneCountInString(fileinfo.Filename) > 128 {
		res.Err = "名称太长"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 验证名称命名规则
	matchSpaceName, _ := regexp.MatchString("^[_0-9a-zA-Z\u4e00-\u9fa5]+$", fileinfo.Filename) // 由中文、字母、数字、下划线验证
	if !matchSpaceName {
		res.Err = "名称只能由中文、字母、数字、下划线组成"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 验证名称是否已存在
	var resourceLibName DataBaseModel.YssimResourceLib
	if DB.Where("name = ? AND parent_id = ?", fileinfo.Filename, fileinfo.ParentId).First(&resourceLibName); resourceLibName.ID != "" {
		res.Err = "名称已存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 限制文件格式
	if !strings.HasSuffix(varFile.Filename, ".txt") {
		res.Err = "暂时只支持 *.txt 格式文件上传"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 验证文件本身名称命名规则
	matchSpaceName, _ = regexp.MatchString("^[_0-9a-zA-Z]+$", strings.TrimSuffix(varFile.Filename, ".txt")) // 由字母、数字、下划线验证
	if !matchSpaceName {
		res.Err = "文件名称只能由字母、数字、下划线组成"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	var resourceLib = DataBaseModel.YssimResourceLib{
		ParentId:    fileinfo.ParentId,
		ID:          uuid.New().String(),
		UserName:    userName,
		FolderFile:  "file",
		FullPath:    "",
		Name:        fileinfo.Filename,
		Description: fileinfo.Description,
		FilePath:    "",
	}

	var parentResource DataBaseModel.YssimResourceLib
	DB.Where("id = ?", fileinfo.ParentId).First(&parentResource)
	resourceLib.FullPath = parentResource.FullPath + "/" + parentResource.Name

	filePath, ok := serviceV2.SaveResourceFile(varFile, userName, resourceLib.ID)
	if !ok {
		res.Err = "上传失败，请重新上传"
		res.Status = 2
		c.JSON(http.StatusOK, res)
	}

	resourceLib.FilePath = filePath
	if err = DB.Create(&resourceLib).Error; err != nil {
		log.Println("上传资源文件时数据库出现错误：", err)
		res.Err = "上传失败，请重新上传"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	res.Msg = "上传成功"
	res.Data = resourceLib.ID
	c.JSON(http.StatusOK, res)
}

func GetResourceFileContentView(c *gin.Context) {
	/*
		# 获取资源文件的内容
		开发人： 周强
	*/
	var res DataType.ResponseData
	resourceFileId := c.Query("id")

	// 获取资源文件的基本信息
	var resourceFile DataBaseModel.YssimResourceLib
	if err := DB.Where("id = ?", resourceFileId).First(&resourceFile).Error; err != nil {
		log.Println("获取资源文件信息时数据库出现错误：", err)
		res.Err = "资源文件不存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 获取资源文件的内容
	res.Data = serviceV2.GetResourceFileContent(resourceFile.FilePath)
	c.JSON(http.StatusOK, res)
}

func ParseResourceFileContentView(c *gin.Context) {
	/*
		# 解析资源文件的内容
		开发人： 周强
	*/
	var res DataType.ResponseData
	resourceFileId := c.Query("id")

	// 获取资源文件的基本信息
	var resourceFile DataBaseModel.YssimResourceLib
	if err := DB.Where("id = ?", resourceFileId).First(&resourceFile).Error; err != nil {
		log.Println("获取资源文件信息时数据库出现错误：", err)
		res.Err = "资源文件不存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 获取资源文件的内容
	res.Data = serviceV2.ParseResourceFileContent(resourceFile.FilePath)
	c.JSON(http.StatusOK, res)
}

func CopyLibFileToResourcesView(c *gin.Context) {
	/*
		# 将资源库文件拷贝到Resources
		开发人： 周强
	*/
	var res DataType.ResponseData
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")
	var item DataType.CopyLibFileToResourcesData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}

	// 获取资源文件信息
	var resourceFile DataBaseModel.YssimResourceLib
	if err := DB.Where("id = ?", item.ID).First(&resourceFile).Error; err != nil {
		res.Err = "没有找到资源"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	if resourceFile.FolderFile == "folder" {
		res.Err = "不能复制文件夹"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 查询要插入的package
	var packageModel DataBaseModel.YssimModels
	DB.Where("id = ? AND sys_or_user = ? AND userspace_id = ?", item.PackageId, username, userSpaceId).First(&packageModel)
	if packageModel.ID == "" {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "not found")
		return
	}

	names := strings.Split(resourceFile.FilePath, "/")
	fileName := names[len(names)-1]
	result, newfileName := serviceV2.CopyLibFileToResources(packageModel.PackageName, item.Parent, resourceFile.FilePath, fileName)
	if result {
		res.Msg = "复制文件成功"
		pathList := []string{}
		if item.Parent != "" {
			pathList = append(pathList, item.Parent)
		}

		pathList = append(pathList, newfileName)
		data := map[string]string{
			"path": "modelica://" + packageModel.PackageName + "/Resources/" + strings.Join(pathList, "/"),
			"name": newfileName,
		}
		res.Data = data
		c.JSON(http.StatusOK, res)
		return
	}

	res.Msg = "复制失败"
	c.JSON(http.StatusOK, res)
}
