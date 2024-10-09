package pipeNet

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
	"yssim-go/app/DataBaseModel"
	"yssim-go/app/DataType"
	"yssim-go/app/v1/service"
	serviceV2 "yssim-go/app/v2/service"
	"yssim-go/config"
	"yssim-go/library/fileOperation"

	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var DB = config.DB

func UploadMappingConfigView(c *gin.Context) {
	/*
		# 上传映射配置表文件
	*/
	var res DataType.ResponseData
	userName := c.GetHeader("username")
	var fileinfo DataType.UploadMappingConfigData
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

	// 验证应用名称命名规则
	matchSpaceName, _ := regexp.MatchString("^[_0-9a-zA-Z\u4e00-\u9fa5]+$", fileinfo.Filename) // 由中文、字母、数字、下划线验证
	if !matchSpaceName {
		res.Err = "名称只能由中文、字母、数字、下划线组成"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 验证名称是否已存在
	var mappingConfigName DataBaseModel.YssimMappingConfig
	if DB.Where("name = ? AND username = ?", fileinfo.Filename, userName).First(&mappingConfigName); mappingConfigName.ID != "" {
		res.Err = "映射配置表名称已存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 限制文件格式
	if !strings.HasSuffix(varFile.Filename, ".json") {
		res.Err = "暂时只支持 *.json 格式文件上传"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 校验文件内容
	if !serviceV2.CheckMappingConfigContent(varFile) {
		res.Err = "内容格式有误，请检查后再上传"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	var mappingConfig DataBaseModel.YssimMappingConfig = DataBaseModel.YssimMappingConfig{
		ID:          uuid.New().String(),
		UserName:    userName,
		Name:        fileinfo.Filename,
		Description: fileinfo.Description,
		Path:        "",
	}

	filePath, ok := serviceV2.SaveMappingConfig(varFile, userName, mappingConfig.ID)
	if !ok {
		res.Err = "上传失败，请重新上传"
		res.Status = 2
		c.JSON(http.StatusOK, res)
	}

	mappingConfig.Path = filePath
	if err = DB.Create(&mappingConfig).Error; err != nil {
		log.Println("上传映射配置表时数据库出现错误：", err)
		res.Err = "上传失败，请重新上传"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	res.Msg = "上传成功"
	res.Data = mappingConfig.ID
	c.JSON(http.StatusOK, res)
}

func DownloadMappingConfigView(c *gin.Context) {
	/*
		# 下载映射配置表
		开发人： 周强
	*/
	var res DataType.ResponseData
	userName := c.GetHeader("username")
	var item DataType.DownloadMappingConfigData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	// 创建ZIP文件
	zipPath := "static/mappingConfig/tmp/" + time.Now().Local().Format("20060102150405") + "/" + "mapping_configs.zip"
	zipFile, ok := fileOperation.CreateFile(zipPath)
	if !ok {
		log.Println("创建临时的ZIP文件时出现错误：", err)
		res.Err = "下载失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	defer zipFile.Close()

	// 创建ZIP writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 获取映射配置表记录信息
	var mappingConfigList []DataBaseModel.YssimMappingConfig
	if err := DB.Where("id IN ? AND username = ?", item.MappingConfigIdList, userName).Find(&mappingConfigList).Error; err != nil {
		log.Println("获取映射配置表记录信息时数据库出现错误：", err)
		res.Err = "映射配置表不存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 遍历每个目录，添加json文件到ZIP包
	for _, mappingConfig := range mappingConfigList {
		file, err := os.Open(mappingConfig.Path)
		if err != nil {
			log.Println("打开映射配置表文件时出现错误：", err)
			res.Err = "下载失败，请稍后再试"
			res.Status = 2
			c.JSON(http.StatusOK, res)
			return
		}
		defer file.Close()

		// 设置文件名
		fileName := mappingConfig.Name + "_" + path.Base(mappingConfig.Path)

		// 创建ZIP文件条目
		zipFileWriter, err := zipWriter.Create(fileName)
		if err != nil {
			log.Println("创建ZIP文件条目时出现错误：", err)
			res.Err = "下载失败，请稍后再试"
			res.Status = 2
			c.JSON(http.StatusOK, res)
			return
		}

		// 将文件内容写入到ZIP条目
		_, err = io.Copy(zipFileWriter, file)
		if err != nil {
			log.Println("将文件内容写入到ZIP条目时出现错误：", err)
			res.Err = "下载失败，请稍后再试"
			res.Status = 2
			c.JSON(http.StatusOK, res)
			return
		}
	}

	// 关闭ZIP writer，确保所有内容都写入文件
	zipWriter.Close()
	res.Data = map[string]string{"url": zipPath}
	c.JSON(http.StatusOK, res)
}

func DeleteMappingConfigView(c *gin.Context) {
	/*
		# 删除映射配置表
		开发人： 周强
	*/
	var res DataType.ResponseData
	userName := c.GetHeader("username")
	var item DataType.DeleteMappingConfigData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	// 删除记录
	var mappingConfig DataBaseModel.YssimMappingConfig
	if err = DB.Where("id IN ? AND username = ?", item.MappingConfigIdList, userName).Delete(&mappingConfig).Error; err != nil {
		log.Println("删除映射配置表时数据库出现错误：", err)
		res.Err = "删除失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res)
}

func CopyMappingConfigView(c *gin.Context) {
	/*
		# 复制映射配置表
		开发人： 周强
	*/
	var res DataType.ResponseData
	userName := c.GetHeader("username")
	var item DataType.CopyMappingConfigData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	var mappingConfigList []DataBaseModel.YssimMappingConfig
	DB.Where("id IN ? AND username = ?", item.MappingConfigIdList, userName).Find(&mappingConfigList)

	var newMappingConfigList []DataBaseModel.YssimMappingConfig
	for _, mappingConfig := range mappingConfigList {
		// 生成复制出来的副本的名称
		var newName string
		var mappingConfigName DataBaseModel.YssimMappingConfig
		if DB.Where("username = ? AND name = ?", userName, mappingConfig.Name+"_副本").First(&mappingConfigName); mappingConfigName.ID == "" {
			newName = mappingConfig.Name + "_副本"
		} else {
			var mappingConfigNameList []DataBaseModel.YssimMappingConfig
			DB.Where("username = ? AND name REGEXP ?", userName, mappingConfig.Name+"_副本"+"[0-9]+").Find(&mappingConfigNameList)
			nums := []int{}
			for _, mappingConfigName := range mappingConfigNameList {
				strs := strings.Split(mappingConfigName.Name, "副本")
				num, _ := strconv.Atoi(strs[len(strs)-1])
				nums = append(nums, num)
			}

			// 获取待创建的副本的编号
			num := serviceV2.FindFirstCopyNum(nums)
			newName = fmt.Sprintf("%s%d", mappingConfig.Name+"_副本", num)
		}

		var newMappingConfig DataBaseModel.YssimMappingConfig = DataBaseModel.YssimMappingConfig{
			ID:          uuid.New().String(),
			UserName:    userName,
			Name:        newName,
			Description: mappingConfig.Description,
			Path:        "",
		}

		newPath, ok := serviceV2.CopyMappingConfig(mappingConfig.Path, userName, newMappingConfig.ID)
		if !ok {
			res.Err = "复制失败"
			res.Status = 2
			c.JSON(http.StatusOK, res)
			return
		}

		newMappingConfig.Path = newPath
		newMappingConfigList = append(newMappingConfigList, newMappingConfig)
	}

	if err = DB.Create(&newMappingConfigList).Error; err != nil {
		log.Println("复制映射配置表时数据库出现错误：", err)
		res.Err = "复制失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	res.Msg = "复制成功"
	c.JSON(http.StatusOK, res)
}

func GetMappingConfigListView(c *gin.Context) {
	/*
		# 获取映射配置表列表
		开发人： 周强
	*/
	var res DataType.ResponseData
	userName := c.GetHeader("username")
	keyWords := c.Query("keywords")
	pageNumStr := c.Query("page_num") //页码
	pageNum, _ := strconv.Atoi(pageNumStr)
	pageSizeStr := c.Query("page_size") //每页条数
	pageSize, _ := strconv.Atoi(pageSizeStr)
	if pageSize <= 0 {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	var total int64 //总条数s
	DB.Where("username = ? AND name LIKE ?", userName, "%"+keyWords+"%").Find(&DataBaseModel.YssimMappingConfig{}).Count(&total)
	pageCount := math.Ceil(float64(total) / float64(pageSize)) //总页数

	var mappingConfigList []DataBaseModel.YssimMappingConfig
	if err := DB.Limit(pageSize).Offset((pageNum-1)*pageSize).Where("username = ? AND name LIKE ?", userName, "%"+keyWords+"%").Order("create_time desc").Find(&mappingConfigList).Error; err != nil {
		log.Println("获取映射配置表列表时数据库出现错误：", err)
		res.Err = "获取映射配置表列表失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	mappingConfigListData := make([]map[string]any, 0)
	for _, m := range mappingConfigList {
		data := map[string]any{
			"id":          m.ID,
			"username":    m.UserName,
			"name":        m.Name,
			"description": m.Description,
			"create_time": m.CreatedAt,
			"update_time": m.UpdatedAt,
		}

		mappingConfigListData = append(mappingConfigListData, data)
	}

	data := make(map[string]any)
	data["mapping_config_list"] = mappingConfigListData
	data["page_count"] = pageCount
	data["total"] = total

	res.Data = data
	c.JSON(http.StatusOK, res)
}

func EditMappingConfigView(c *gin.Context) {
	/*
		# 编辑映射配置表基本信息
		开发人： 周强
	*/
	var res DataType.ResponseData
	userName := c.GetHeader("username")
	var item DataType.EditMappingConfigData
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

	var mappingConfig DataBaseModel.YssimMappingConfig
	if DB.Where("id = ? AND username = ?", item.ID, userName).First(&mappingConfig); mappingConfig.ID == "" {
		res.Err = "映射配置表不存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 验证应用名称是否已存在
	var mappingConfigName DataBaseModel.YssimMappingConfig
	DB.Where("name = ? AND username = ?", item.Name, userName).First(&mappingConfigName)
	if mappingConfigName.ID != "" && mappingConfigName.ID != item.ID {
		res.Err = "映射配置表名称已存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	mappingConfig.Name = item.Name
	mappingConfig.Description = item.Description
	if err = DB.Save(&mappingConfig).Error; err != nil {
		log.Println("编辑映射配置表基本信息时数据库出现错误：", err)
		res.Err = "编辑失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	res.Msg = "编辑成功"
	c.JSON(http.StatusOK, res)
}

func GetMappingConfigDetailsView(c *gin.Context) {
	/*
		# 获取映射配置表的详细参数信息
		开发人： 周强
	*/
	var res DataType.ResponseData
	userName := c.GetHeader("username")
	mappingConfigId := c.Query("mapping_config_id")
	// 获取映射配置表的基本信息
	var mappingConfig DataBaseModel.YssimMappingConfig
	if err := DB.Where("id = ? AND username = ?", mappingConfigId, userName).First(&mappingConfig).Error; err != nil {
		log.Println("获取映射配置表详细参数信息时数据库出现错误：", err)
		res.Err = "映射配置表不存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 获取映射配置表的详细参数信息
	data, err := serviceV2.GetMappingConfigDetails(mappingConfig.ID, mappingConfig.Name, mappingConfig.Description, mappingConfig.Path)
	if err != nil {
		res.Err = "获取映射配置表详细参数信息失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	res.Data = data
	c.JSON(http.StatusOK, res)
}

func EditMappingConfigDetailsView(c *gin.Context) {
	/*
		# 编辑映射配置表的详细参数信息
		开发人： 周强
	*/
	var res DataType.ResponseData
	userName := c.GetHeader("username")
	var item DataType.EditMappingConfigDetailsData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	// 获取映射配置表的基本信息
	var mappingConfig DataBaseModel.YssimMappingConfig
	if err := DB.Where("id = ? AND username = ?", item.ID, userName).First(&mappingConfig).Error; err != nil {
		log.Println("获取映射配置表详细参数信息时数据库出现错误：", err)
		res.Err = "映射配置表不存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	if ok, err := serviceV2.EditMappingConfigDetails(mappingConfig.Path, &item, item.Op); !ok {
		if err != nil {
			res.Err = err.Error()
		} else {
			switch item.Op {
			case "add":
				res.Err = "添加失败"
			case "replace":
				res.Err = "编辑失败"
			case "remove":
				res.Err = "删除失败"
			default:
				res.Err = "编辑失败"
			}
		}
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	switch item.Op {
	case "add":
		res.Msg = "添加成功"
	case "replace":
		res.Msg = "编辑成功"
	case "remove":
		res.Msg = "删除成功"
	default:
		res.Msg = "编辑成功"
	}
	c.JSON(http.StatusOK, res)
}

func GetInstanceMappingView(c *gin.Context) {
	/*
		# 获取映射配置表的详细参数信息
		开发人： 周强
	*/
	var res DataType.ResponseData
	userName := c.GetHeader("username")
	pipeNetInfoId := c.Query("pipe_net_info_id")
	mappingConfigId := c.Query("mapping_config_id")

	// 获取管网信息文件基本信息
	var pipeNetInfoFileRecord DataBaseModel.YssimPipeNetCad
	DB.Where("id = ? AND username = ?", pipeNetInfoId, userName).First(&pipeNetInfoFileRecord)
	if pipeNetInfoFileRecord.ID == "" {
		res.Err = "not found"
		c.JSON(http.StatusOK, res)
		return
	}

	// 获取映射配置表的基本信息
	var mappingConfig DataBaseModel.YssimMappingConfig
	if err := DB.Where("id = ? AND username = ?", mappingConfigId, userName).First(&mappingConfig).Error; err != nil {
		log.Println("获取映射配置表详细参数信息时数据库出现错误：", err)
		res.Err = "映射配置表不存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 生成实例映射表
	data, err := serviceV2.GetInstanceMapping(pipeNetInfoFileRecord.ID, mappingConfig.ID, pipeNetInfoFileRecord.Path, mappingConfig.Path)
	if err != nil {
		res.Err = "获取映射配置表详细参数信息失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	res.Data = data
	c.JSON(http.StatusOK, res)
}

func GetInstanceMappingLogView(c *gin.Context) {
	/*
		# 获取映射配置表的详细参数信息
		开发人： 周强
	*/
	var res DataType.ResponseData
	//userName := c.GetHeader("username")
	pipeNetInfoId := c.Query("pipe_net_info_id")
	mappingConfigId := c.Query("mapping_config_id")

	// 生成实例映射表
	data := serviceV2.GetInstanceMappingLog(mappingConfigId, pipeNetInfoId)
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func CreatePipeNetModelView(c *gin.Context) {
	/*
		# 创建管网模型
		开发人： 周强
	*/
	var res DataType.ResponseData
	var item DataType.CreatePipeNetModelData
	err := c.BindJSON(&item)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	username := c.GetHeader("username")
	userSpaceId := c.GetHeader("space_id")

	// 获取管网信息文件基本信息
	var pipeNetInfoFileRecord DataBaseModel.YssimPipeNetCad
	DB.Where("id = ? AND username = ?", item.PipeNetInfoId, username).First(&pipeNetInfoFileRecord)
	if pipeNetInfoFileRecord.ID == "" {
		res.Err = "not found"
		c.JSON(http.StatusOK, res)
		return
	}

	// 获取映射配置表的基本信息
	var mappingConfig DataBaseModel.YssimMappingConfig
	if err := DB.Where("id = ? AND username = ?", item.MappingConfigId, username).First(&mappingConfig).Error; err != nil {
		log.Println("获取映射配置表详细参数信息时数据库出现错误：", err)
		res.Err = "映射配置表不存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	// 获取映射配置表的详细参数信息
	mappingConfigData, err := serviceV2.GetMappingConfigDetails(mappingConfig.ID, mappingConfig.Name, mappingConfig.Description, mappingConfig.Path)
	if err != nil {
		res.Err = "获取映射配置表详细参数信息失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 生成实例映射表
	data, err := serviceV2.GetInstanceMapping(pipeNetInfoFileRecord.ID, mappingConfig.ID, pipeNetInfoFileRecord.Path, mappingConfig.Path)
	if err != nil {
		res.Err = "获取映射配置表详细参数信息失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 创建空模型
	matchSpaceName1, _ := regexp.MatchString("^[_a-zA-Z0-9]+$", item.Name) // 字母、数字、下划线验证
	matchSpaceName2, _ := regexp.MatchString("^[a-zA-Z_]", item.Name)      // 字母、下划线验证
	if !matchSpaceName1 {
		res.Err = "模型名称只能由字母数字下划线组成"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	if !matchSpaceName2 {
		res.Err = "模型名称只能由字母下划线开头"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	createPackageName := item.Name
	createPackageNameALL := item.Name

	var packageRecord DataBaseModel.YssimModels
	var newPackage = DataBaseModel.YssimModels{
		ID:          uuid.New().String(),
		PackageName: createPackageName,
		SysUser:     username,
		FilePath:    "static/UserFiles/UploadFile/" + username + "/" + time.Now().Local().Format("20060102150405") + "/" + createPackageName + "/" + createPackageName + ".mo",
		UserSpaceId: userSpaceId,
	}
	DB.Where("package_name = ? AND sys_or_user IN ? AND userspace_id IN ?", item.Name, []string{"sys", username}, []string{"0", userSpaceId}).First(&packageRecord)
	if packageRecord.PackageName != "" && item.Vars.InsertTo == "" {
		res.Err = "模型名称已存在，请修改后再试。"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	if item.Vars.InsertTo != "" {
		var insertPackageRecord DataBaseModel.YssimModels
		DB.Where("id = ? AND sys_or_user = ? AND userspace_id = ?", item.PackageId, username, userSpaceId).First(&insertPackageRecord)
		createPackageNameALL = item.Vars.InsertTo + "." + item.Name
		modelChildList := service.GetModelChild(item.Vars.InsertTo)
		for i := 0; i < len(modelChildList); i++ {
			if modelChildList[i].Name == item.Name {
				res.Err = "名称已存在，请修改后再试。"
				res.Status = 2
				c.JSON(http.StatusOK, res)
				return
			}
		}
		newPackage = insertPackageRecord
	} else {
		DB.Create(&newPackage)
	}
	result := service.CreateModelAndPackage(createPackageName, item.Vars.InsertTo, item.Vars.Expand, item.StrType, createPackageNameALL, item.Comment, item.Vars.Partial, item.Vars.Encapsulated, item.Vars.State)
	if result {
		saveResult := service.SaveModelCode(createPackageNameALL, newPackage.FilePath)
		if saveResult {
			res.Msg = "创建成功"
			if item.Vars.InsertTo == "" {
				res.Data = map[string]string{
					"model_name": newPackage.PackageName,
					"id":         newPackage.ID,
				}
			} else {
				res.Data = map[string]string{
					"model_name": item.Vars.InsertTo + "." + item.Name,
					"id":         newPackage.ID,
				}
			}
			packageInformation := service.GetPackageInformation()
			packageInformationJson, _ := sonic.Marshal(packageInformation)
			DB.Model(DataBaseModel.YssimUserSpace{}).Where("id = ? AND username = ?", userSpaceId, username).Update("package_information", packageInformationJson)
		} else {
			DB.Delete(&newPackage)
			res.Err = "创建模型失败，请稍后再试"
			res.Status = 2
			c.JSON(http.StatusOK, res)
			return
		}
	} else {
		res.Err = "创建模型失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 向模型中写入Modelica代码
	//serviceV2.WritePipeNetModeCode(item.Name, createPackageNameALL, mappingConfigData.Medium, newPackage.PackageName, newPackage.FilePath, data["mapping_tree"])
	if ok := serviceV2.WritePipeNetModeCodeNew(item.Name, createPackageNameALL, mappingConfigData.System, mappingConfigData.Medium, &newPackage, newPackage.FilePath, data["mapping_tree"]); !ok {
		res.Err = "创建模型失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 创建更新下载记录
	var newPipeNetCadDownload = DataBaseModel.YssimPipeNetCadDownload{
		ID:          uuid.New().String(),
		UserName:    username,
		Name:        pipeNetInfoFileRecord.Name,
		Description: pipeNetInfoFileRecord.Description,
		PackageId:   newPackage.ID,
		ModelName:   createPackageNameALL,
	}
	// 复制当前管网信息文件
	newPath, ok := serviceV2.CopyPipeNetInfoFile(pipeNetInfoFileRecord.Path, username, newPipeNetCadDownload.ID)
	if !ok {
		res.Err = "创建失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	newPipeNetCadDownload.PipeNetPath = newPath
	// 复制映射表
	newMappingPath, ok := serviceV2.CopyMappingConfig(mappingConfig.Path, username, newPipeNetCadDownload.ID)
	if !ok {
		res.Err = "创建失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	newPipeNetCadDownload.PipeNetPath = newPath
	newPipeNetCadDownload.MappingPath = newMappingPath
	DB.Create(&newPipeNetCadDownload)

	res.Data = map[string]any{
		"encryption": false,
		"model":      serviceV2.GetModelInstanceData(createPackageNameALL),
		"package_id": newPackage.ID,
	}
	c.JSON(http.StatusOK, res)
}
