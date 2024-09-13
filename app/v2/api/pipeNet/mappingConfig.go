package pipeNet

import (
	"fmt"
	"log"
	"net/http"
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

	// 获取映射配置表记录信息
	var mappingConfigList []DataBaseModel.YssimMappingConfig
	if err := DB.Where("id IN ? AND username = ?", item.MappingConfigIdList, userName).Find(&mappingConfigList).Error; err != nil {
		log.Println("获取映射配置表的详细参数信息时数据库出现错误：", err)
		res.Err = "映射配置表不存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 开始下载
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment")
	c.Header("Content-Transfer-Encoding", "binary")

	for _, mappingConfig := range mappingConfigList {
		c.File(mappingConfig.Path)
	}
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
