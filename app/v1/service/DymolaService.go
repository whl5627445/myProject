package service

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
	"yssim-go/config"
	"yssim-go/library/fileOperation"

	"github.com/mholt/archiver/v3"

	"github.com/wangluozhe/requests"

	"github.com/wangluozhe/requests/url"
)

func DymolaFmuExportWithLibrary(fmuPar map[string]any, envLibrary map[string]string, userName, saveName, modelName string) (string, bool, string) {
	workPath := "static/tmp/" + userName + "/" + strings.ReplaceAll(modelName, ".", "-") + "/" + time.Now().Local().Format("20060102150405") + "/"
	data := map[string]any{
		"fmuPar":      fmuPar,
		"modelToOpen": modelName,
		"saveName":    saveName,
		"envLibrary":  envLibrary,
		"workPath":    workPath,
	}
	messageMap := make(map[string]interface{})
	errTips := "下载失败，请稍后再试"
	// 创建目录
	if ok := fileOperation.CreateFilePath(workPath); !ok {
		log.Println("创建路径失败")
		return "", false, errTips
	}
	// 目录设置777权限
	err := fileOperation.SetPermissions(workPath)
	if err != nil {
		return "", false, errTips
	}

	Headers := url.NewHeaders()
	req := url.NewRequest()
	req.Json = data
	req.Headers = Headers
	req.Timeout = time.Minute * 10
	exportFmuRes, err := requests.Post(config.FmuExportConnect+"/dymola/buildModelFMU", req)
	if err != nil {
		errTips = "下载失败，请查看日志"
		messageMap["message"] = "导出失败"
		messageMap["status"] = true
		messageMap["type"] = "error"
		MessageNotice(messageMap)
		return "", false, errTips
	}
	exportResult, _ := exportFmuRes.Json()
	log.Println("dymola服务编译FMU结果：", exportResult)
	ResultCode, ok := exportResult["code"]
	if len(exportResult) == 0 || (ok && ResultCode.(float64) != 200) {
		errTips = "下载失败，请查看日志"
		messageMap["message"] = exportResult["log"].(string)
		messageMap["status"] = true
		messageMap["type"] = "error"
		MessageNotice(messageMap)
		messageMap["message"] = "导出失败"
		messageMap["status"] = true
		messageMap["type"] = "error"
		MessageNotice(messageMap)
		return "", false, errTips
	}

	return exportResult["path"].(string), true, ""
}

func OmcFmuExportWithLibrary(fmuPar map[string]any, envLibrary map[string]string, userName, fmuName, packageName, modelName, filePath string) (string, bool, string) {
	data := map[string]any{
		"fmuPar":       fmuPar,
		"fmuName":      fmuName,
		"workPath":     "",
		"modelToOpen":  modelName,
		"omcLibraries": []map[string]string{},
	}
	messageMap := make(map[string]interface{})
	// 所有要加载的用户模型的包文件地址
	var folders []string
	var dymolaLibraries []map[string]string
	// 编译问题提示 "下载失败，请查看日志"
	// 系统问题提示 "下载失败，请稍后再试"
	errTips := "下载失败，请稍后再试"
	now := time.Now()
	timestamp := now.Format("20060102150405")
	// 新建zip文件地址
	uploadFileName := filepath.Join("static/tmp", timestamp, packageName+".zip")
	delUploadFileName := filepath.Join("static/tmp", timestamp)
	// dymola服务器上新建的路径
	paramsUrl := filepath.Join(userName, packageName, timestamp)
	data["workPath"] = paramsUrl

	for key, val := range envLibrary {
		// 初始化加载用户模型，key是名称，val是mo文件地址
		if strings.Contains(val, "/") {
			// 将路径分割  "static/UserFiles/UploadFile/songyi/20230427165917/Applications1/Applications1/package.mo"
			frontPath := strings.Join(strings.Split(val, "/")[:6], "/")  // static/UserFiles/UploadFile/songyi/20230427165917/Applications1
			behindPath := strings.Join(strings.Split(val, "/")[5:], "/") // Applications1/Applications1/package.mo
			// 插入到zip文件列表
			folders = append(folders, filepath.Join(frontPath))
			dymolaLibraries = append(dymolaLibraries, map[string]string{
				"libraryName":    "",
				"libraryVersion": "",
				"userFile":       filepath.Join(paramsUrl, behindPath),
			})
		} else {
			// 初始化加载系统模型，key是名称，val是版本号
			dymolaLibraries = append(dymolaLibraries, map[string]string{
				"libraryName":    key,
				"libraryVersion": val,
				"userFile":       "",
			})
		}
	}

	req := url.NewRequest()
	params := url.NewParams()
	params.Set("url", paramsUrl)
	req.Params = params
	Headers := url.NewHeaders()
	req.Headers = Headers
	req.Timeout = 600 * time.Second
	if filePath != "" {
		// 压缩文件
		err := archiver.Archive(folders, uploadFileName)
		if err != nil {
			fmt.Println(err)
			return "", false, errTips
		}

		files := url.NewFiles()
		files.SetFile("file", packageName+".zip", uploadFileName, "")
		req.Files = files
		uploadFileRes, err := requests.Post(config.FmuExportConnect+"/file/upload", req)
		if err != nil {
			log.Println(err)
			return "", false, errTips
		}
		uploadFileResJson, _ := uploadFileRes.Json()
		uploadResultCode, ok := uploadFileResJson["code"]
		if ok && uploadResultCode.(float64) == 200 {
		} else {
			return "", false, errTips
		}

		err = os.RemoveAll(delUploadFileName)
		if err != nil {
			fmt.Println(err)
		}

	} else {
		var filteredDymolaLibraries []map[string]string
		for _, element := range dymolaLibraries {
			if element["userFile"] == "" {
				filteredDymolaLibraries = append(filteredDymolaLibraries, element)
			}
		}
		dymolaLibraries = filteredDymolaLibraries
	}
	data["omcLibraries"] = dymolaLibraries
	req = url.NewRequest()
	req.Json = data
	req.Headers = Headers
	req.Timeout = time.Minute * 10
	exportFmuRes, err := requests.Post(config.FmuExportConnect+"/omc/buildModelFMU", req)
	if err != nil {
		errTips = "下载失败，请查看日志"
		messageMap["message"] = "Description Failed to send the http request"
		messageMap["status"] = true
		messageMap["type"] = "error"
		MessageNotice(messageMap)
		return "", false, errTips
	}
	exportResult, _ := exportFmuRes.Json()
	ResultCode, _ := exportResult["code"]
	if ResultCode.(float64) != 200 {
		errTips = "下载失败，请查看日志"
		messageMap["message"] = exportResult["log"].(string)
		messageMap["status"] = true
		messageMap["type"] = "error"
		MessageNotice(messageMap)
		return "", false, errTips
	}
	req = url.NewRequest()
	req.Headers = Headers
	req.Timeout = time.Minute * 10
	fmuFileUrl := config.FmuExportConnect + "/file/download?fileName=" + exportResult["msg"].(string)
	fmuFileRes, err := requests.Get(fmuFileUrl, req)
	if err != nil {
		return "", false, errTips
	}
	resultFmuFileData := fmuFileRes.Content
	newFilePath := "static/tmp/" + userName + "/" + strings.ReplaceAll(modelName, ".", "-") + "/" + time.Now().Local().Format("20060102150405") + "/" + fmuName + ".fmu"
	fileOperation.WriteFileByte(newFilePath, resultFmuFileData)
	return newFilePath, true, ""
}
