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
	"yssim-go/app/DataBaseModel"
	"yssim-go/app/DataType"
	serviceV2 "yssim-go/app/v2/service"
	"yssim-go/config"

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

	// 创建临时的ZIP文件
	zipFile, err := os.CreateTemp("", "mapping_configs_*.zip")
	if err != nil {
		log.Println("创建临时的ZIP文件时出现错误：", err)
		res.Err = "下载失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}
	defer os.Remove(zipFile.Name())

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
		fileName := path.Base(mappingConfig.Path)

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

	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", "attachment; filename=mapping_configs.zip")
	c.Header("Content-Length", "0")
	c.File(zipFile.Name())
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
		if DB.Where("username = ? AND name = ?", userName, mappingConfig.Name+"-副本").First(&mappingConfigName); mappingConfigName.ID == "" {
			newName = mappingConfig.Name + "-副本"
		} else {
			var mappingConfigNameList []DataBaseModel.YssimMappingConfig
			DB.Where("username = ? AND name REGEXP ?", userName, mappingConfig.Name+"-副本"+"[0-9]+").Find(&mappingConfigNameList)
			nums := []int{}
			for _, mappingConfigName := range mappingConfigNameList {
				strs := strings.Split(mappingConfigName.Name, "副本")
				num, _ := strconv.Atoi(strs[len(strs)-1])
				nums = append(nums, num)
			}

			// 获取待创建的副本的编号
			num := serviceV2.FindFirstCopyNum(nums)
			newName = fmt.Sprintf("%s%d", mappingConfig.Name+"-副本", num)
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

	var total int64 //总条数s
	DB.Where("username = ? AND name LIKE ?", userName, "%"+keyWords+"%").Find(&DataBaseModel.YssimMappingConfig{}).Count(&total)
	pageCount := math.Ceil(float64(total) / 10) //总页数

	var mappingConfigList []DataBaseModel.YssimMappingConfig
	if err := DB.Limit(10).Offset((pageNum-1)*10).Where("username = ? AND name LIKE ?", userName, "%"+keyWords+"%").Order("create_time desc").Find(&mappingConfigList).Error; err != nil {
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
		res.Err = "应用名称只能由中文、字母、数字、下划线组成"
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
	data, err := serviceV2.GetMappingConfigDetails(mappingConfig.Path)
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

	if ok := serviceV2.EditMappingConfigDetails(mappingConfig.Path, &item, item.Op); !ok {
		res.Err = "编辑失败"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	res.Msg = "编辑成功"
	c.JSON(http.StatusOK, res)
}
