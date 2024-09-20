package pipeNet

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"yssim-go/app/DataBaseModel"
	"yssim-go/app/DataType"
	serviceV2 "yssim-go/app/v2/service"
	"yssim-go/library/fileOperation"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	// 限制文件格式
	if !strings.HasSuffix(varFile.Filename, ".xml") {
		res.Err = "暂时只支持 *.xml 格式文件上传"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	// 校验文件内容
	if !serviceV2.CheckInfoFileXml(varFile) {
		res.Err = "内容格式有误，请检查后再上传"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	var newPipeNetInfoFileRecord = DataBaseModel.YssimPipeNetCad{
		ID:          uuid.New().String(),
		UserName:    userName,
		Name:        varFile.Filename,
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
		# 下载管网信息文件,支持单个/多个，下载zip文件
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
	zipPath := "static/pipeNetInfoFile/tmp/" + time.Now().Local().Format("20060102150405") + "/" + "xml_files.zip"
	zipFile, ok := fileOperation.CreateFile(zipPath)
	if !ok {
		res.Err = "下载失败，请稍后再试"
		c.JSON(http.StatusOK, res)
		return
	}
	defer zipFile.Close()

	// 创建ZIP writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 获取管网信息文件表记录信息
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
	res.Data = map[string]string{"url": zipPath}
	c.JSON(http.StatusOK, res)
}

func DeleteInfoFileView(c *gin.Context) {
	/*
		# 删除管网信息文件,支持单个/多个
		开发人： xqd
	*/
	var res DataType.ResponseData
	userName := c.GetHeader("username")
	var item DataType.DeletePipeNetInfoFileData
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	// 获取管网信息文件表记录信息
	var pipeNetInfoFileRecordList []DataBaseModel.YssimPipeNetCad
	if err = DB.Where("id IN ? AND username = ?", item.PipeNetInfoFileIdList, userName).Delete(&pipeNetInfoFileRecordList).Error; err != nil {
		log.Println("删除网信息文件的详细参数信息时数据库出现错误：", err)
		res.Err = "删除失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res)
}

func GetInfoFileListView(c *gin.Context) {
	/*
		# 获取管网信息文件列表,支持分页/关键词搜索
		开发人： xqd
	*/
	var res DataType.ResponseData
	userName := c.GetHeader("username")
	keyWords := c.Query("keywords")
	pageNumStr := c.Query("page_num") //页码
	pageNum, _ := strconv.Atoi(pageNumStr)

	var total int64 //总条数s
	DB.Where("username = ? AND name LIKE ?", userName, "%"+keyWords+"%").Find(&DataBaseModel.YssimPipeNetCad{}).Count(&total)
	pageCount := math.Ceil(float64(total) / 10) //总页数

	var pipeNetInfoFileRecordList []DataBaseModel.YssimPipeNetCad
	if err := DB.Limit(10).Offset((pageNum-1)*10).Where("username = ? AND name LIKE ?", userName, "%"+keyWords+"%").Order("create_time desc").Find(&pipeNetInfoFileRecordList).Error; err != nil {
		log.Println("获取管网信息文件列表时数据库出现错误：", err)
		res.Err = "获取管网信息文件列表失败，请稍后再试"
		res.Status = 2
		c.JSON(http.StatusOK, res)
		return
	}

	pipeNetInfoFileRecordListData := make([]map[string]any, 0)
	for _, m := range pipeNetInfoFileRecordList {
		data := map[string]any{
			"id":          m.ID,
			"username":    m.UserName,
			"name":        m.Name,
			"description": m.Description,
			"create_time": m.CreatedAt,
			"update_time": m.UpdatedAt,
		}

		pipeNetInfoFileRecordListData = append(pipeNetInfoFileRecordListData, data)
	}

	data := make(map[string]any)
	data["pipe_net_info_list"] = pipeNetInfoFileRecordListData
	data["page_count"] = pageCount
	data["total"] = total

	res.Data = data
	c.JSON(http.StatusOK, res)
}

func GetInfoView(c *gin.Context) {
	/*
		# 获取管网信息文件列表,支持分页/关键词搜索
		开发人： xqd
	*/
	var res DataType.ResponseData
	userName := c.GetHeader("username")
	pipeNetInfoId := c.Query("id")

	var pipeNetInfoFileRecord DataBaseModel.YssimPipeNetCad
	DB.Where("id = ? AND username = ?", pipeNetInfoId, userName).First(&pipeNetInfoFileRecord)
	if pipeNetInfoFileRecord.ID == "" {
		res.Err = "not found"
		c.JSON(http.StatusOK, res)
		return
	}

	// 解析管网信息文件xml
	data, err := serviceV2.ParseInfoFileXml(pipeNetInfoFileRecord.Path)
	if err != nil {
		res.Err = "解析xml失败"
		c.JSON(http.StatusOK, res)
		return
	}
	//重新保存为xml
	fmt.Println(pipeNetInfoFileRecord.Path)
	err = serviceV2.SaveInfoFileXml(pipeNetInfoFileRecord.Path, data)
	if err != nil {
		fmt.Println(err)
	}

	// 返回数据
	res.Data = data
	c.JSON(http.StatusOK, res)

}
