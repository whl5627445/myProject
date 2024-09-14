package pipeNet

import (
	"archive/zip"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"yssim-go/app/DataBaseModel"
	"yssim-go/app/DataType"
	serviceV2 "yssim-go/app/v2/service"
)

func UploadInfoFileView(c *gin.Context) {
	/*
		# 上传管网信息文件
	*/
	var res DataType.ResponseData
	userName := c.GetHeader("username")
	var fileInfo DataType.UploadPipeNetInfoFileData
	if err := c.Bind(&fileInfo); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}

	varFile, err := c.FormFile("file")
	if err != nil {
		log.Println("上传管网信息文件时出现错误：", err)
		c.JSON(http.StatusBadRequest, "")
		return
	}

	// 验证名称是否已存在
	var pipeNetInfoFileRecord DataBaseModel.YssimPipeNetCad
	if DB.Where("name = ? AND username = ?", fileInfo.Filename, userName).First(&pipeNetInfoFileRecord); pipeNetInfoFileRecord.ID != "" {
		res.Err = "管网信息文件名称已存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 限制文件格式
	if !strings.HasSuffix(varFile.Filename, ".xml") {
		res.Err = "暂时只支持 *.xml 格式文件上传"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 校验文件内容 待完成
	//if !serviceV2.CheckMappingConfigContent(varFile) {
	//	res.Err = "内容格式有误，请检查后再上传"
	//	res.Status = 2
	//	c.JSON(http.StatusOK, res)
	//	return
	//}

	var newPipeNetInfoFileRecord = DataBaseModel.YssimPipeNetCad{
		ID:          uuid.New().String(),
		UserName:    userName,
		Name:        fileInfo.Filename,
		Description: fileInfo.Description,
		Path:        "",
	}

	filePath, ok := serviceV2.SavePipeNetInfoFile(varFile, userName, newPipeNetInfoFileRecord.ID)
	if !ok {
		res.Err = "上传失败，请重新上传"
		res.Status = 2
		c.JSON(http.StatusOK, res)
	}

	newPipeNetInfoFileRecord.Path = filePath
	if err = DB.Create(&newPipeNetInfoFileRecord).Error; err != nil {
		log.Println("上传管网信息文件时数据库出现错误：", err)
		res.Err = "上传失败，请重新上传"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	res.Msg = "上传成功"
	res.Data = newPipeNetInfoFileRecord.ID
	c.JSON(http.StatusOK, res)
}

func DownloadInfoFileView(c *gin.Context) {
	/*
		# 下载映射配置表,支持单个/多个，下载zip文件
		开发人： xqd
	*/
	var res DataType.ResponseData
	userName := c.GetHeader("username")
	var item DataType.DownloadPipeNetInfoFileData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	// 创建临时的ZIP文件
	zipFile, err := os.CreateTemp("", "xml_files_*.zip")
	if err != nil {
		res.Err = "下载失败，请稍后再试"
		c.JSON(http.StatusOK, res)
		return
	}
	defer os.Remove(zipFile.Name())

	// 创建ZIP writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 获取映射配置表记录信息
	var pipeNetInfoFileRecordList []DataBaseModel.YssimPipeNetCad
	if err = DB.Where("id IN ? AND username = ?", item.PipeNetInfoFileIdList, userName).Find(&pipeNetInfoFileRecordList).Error; err != nil {
		log.Println("获管网信息文件的详细参数信息时数据库出现错误：", err)
		res.Err = "管网信息文件不存在"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 遍历每个目录，添加XML文件到ZIP包
	for i, dir := range pipeNetInfoFileRecordList {
		file, err := os.Open(dir.Path)
		if err != nil {
			res.Err = "下载失败，请稍后再试"
			c.JSON(http.StatusOK, res)
			return
		}
		defer file.Close()

		// 获取文件名
		_, fileName := filepath.Split(dir.Path)
		// 获取文件的扩展名
		ext := filepath.Ext(fileName) // .xml

		// 获取文件名（不含扩展名）
		baseName := fileName[:len(fileName)-len(ext)] // test

		// 设置新文件名，加上序号
		newFileName := fmt.Sprintf("%s_%d%s", baseName, i, ext) // test_1.xml

		// 创建ZIP文件条目
		zipFileWriter, err := zipWriter.Create(newFileName)
		if err != nil {
			res.Err = "下载失败，请稍后再试"
			c.JSON(http.StatusOK, res)
			return
		}

		// 将文件内容写入到ZIP条目
		_, err = io.Copy(zipFileWriter, file)
		if err != nil {
			res.Err = "下载失败，请稍后再试"
			c.JSON(http.StatusOK, res)
			return
		}
	}

	// 关闭ZIP writer，确保所有内容都写入文件
	zipWriter.Close()

	// 设置下载响应头
	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", "attachment; filename=xml_files.zip")
	c.Header("Content-Length", "0")
	c.File(zipFile.Name())

}

func DeleteInfoFileView(c *gin.Context) {
	/*
		# 下载映射配置表
		开发人： xqd
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
	c.Header("Content-Transfer-Encoding", "binary")

	for _, mappingConfig := range mappingConfigList {
		c.Header("Content-Disposition", "attachment; filename="+path.Base(mappingConfig.Path))
		c.File(mappingConfig.Path)
	}
}
