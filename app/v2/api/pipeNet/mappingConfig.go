package pipeNet

import (
	"log"
	"net/http"
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
