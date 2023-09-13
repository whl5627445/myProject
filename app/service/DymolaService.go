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

	"github.com/google/uuid"

	"github.com/wangluozhe/requests"

	"github.com/wangluozhe/requests/url"
)

//func DymolaFmuExport(fmuPar map[string]any, token, userName, fmuName, packageName, modelName, fileName, filePath string) (string, bool) {
//	data := map[string]any{
//		"username":    userName,
//		"fmuPar":      fmuPar,
//		"modelName":   fmuName,
//		"fileName":    "",
//		"modelToOpen": modelName,
//		"token":       token,
//		"taskId":      uuid.New().String(),
//	}
//	urlStr := packageName + "/" + strings.ReplaceAll(modelName, ".", "-") + "/" + strconv.Itoa(int(time.Now().UnixNano()))
//	req := url.NewRequest()
//	params := url.NewParams()
//	params.Set("url", userName+"/"+urlStr)
//	req.Params = params
//	Headers := url.NewHeaders()
//	Headers.Set("Authorization", token)
//	req.Headers = Headers
//	req.Timeout = 600 * time.Second
//	if filePath != "" {
//		files := url.NewFiles()
//		files.SetFile("file", fileName+".mo", filePath, "")
//		req.Files = files
//		uploadFileRes, err := requests.Post(config.DymolaSimutalionConnect+"/file/upload", req)
//		if err != nil {
//			log.Println(err)
//			return "", false
//		}
//		uploadFileResJson, _ := uploadFileRes.Json()
//		log.Println("dymola服务上传文件结果：", uploadFileResJson)
//		uploadResultCode, ok := uploadFileResJson["code"]
//		if ok && uploadResultCode.(float64) == 200 {
//			data["fileName"] = urlStr + "/" + fileName + ".mo"
//		} else {
//			return "", false
//
//		}
//	}
//	req = url.NewRequest()
//	req.Json = data
//	req.Headers = Headers
//	req.Timeout = time.Minute * 10
//	exportFmuRes, err := requests.Post(config.DymolaSimutalionConnect+"/dymola/translateModelFMU", req)
//	if err != nil {
//		log.Println(err)
//		return "", false
//	}
//	exportResult, _ := exportFmuRes.Json()
//	log.Println("dymola服务编译FMU结果：", exportResult)
//	ResultCode, ok := exportResult["code"]
//	if err != nil || len(exportResult) == 0 || (ok && ResultCode.(float64) != 200) {
//		return "", false
//	}
//	req = url.NewRequest()
//	req.Headers = Headers
//	req.Timeout = time.Minute * 10
//	fmuFileUrl := config.DymolaSimutalionConnect + "/file/download?fileName=" + exportResult["msg"].(string)
//	fmuFileRes, err := requests.Get(fmuFileUrl, req)
//	if err != nil {
//		return "", false
//	}
//	resultFmuFileData := fmuFileRes.Content
//	newFilePath := "static/tmp/" + userName + "/" + strings.ReplaceAll(modelName, ".", "-") + "/" + time.Now().Local().Format("20060102150405") + "/" + fmuName + ".fmu"
//	fileOperation.WriteFileByte(newFilePath, resultFmuFileData)
//	return newFilePath, true
//}

func DymolaFmuExportWithLibrary(fmuPar map[string]any, envLibrary map[string]string, token, userName, fmuName, packageName, modelName, fileName, filePath string) (string, bool, string) {
	data := map[string]any{
		"username":        userName,
		"fmuPar":          fmuPar,
		"modelName":       fmuName,
		"fileName":        "",
		"modelToOpen":     modelName,
		"token":           token,
		"taskId":          uuid.New().String(),
		"dymolaLibraries": []map[string]string{},
	}
	messageMap := make(map[string]interface{})
	var folders []string // 所有要加载的用户模型的包文件地址
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

	// 将Modelica放在列表的第一位
	for i, d := range dymolaLibraries {
		if d["libraryName"] == "Modelica" {
			dymolaLibraries = append(dymolaLibraries[:i], dymolaLibraries[i+1:]...)
			dymolaLibraries = append(dymolaLibraries[:0], append([]map[string]string{d}, dymolaLibraries[0:]...)...)
			break
		}
	}

	req := url.NewRequest()
	params := url.NewParams()
	params.Set("url", paramsUrl)
	req.Params = params
	Headers := url.NewHeaders()
	Headers.Set("Authorization", token)
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
		files.SetFile("file", fileName+".zip", uploadFileName, "")
		req.Files = files
		uploadFileRes, err := requests.Post(config.DymolaSimutalionConnect+"/file/upload", req)
		if err != nil {
			log.Println(err)
			return "", false, errTips
		}
		uploadFileResJson, _ := uploadFileRes.Json()
		log.Println("dymola服务上传文件结果：", uploadFileResJson)
		uploadResultCode, ok := uploadFileResJson["code"]
		if ok && uploadResultCode.(float64) == 200 {
			data["fileName"] = paramsUrl + "/" + strings.Join(strings.Split(filePath, "/")[5:], "/")
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
	data["dymolaLibraries"] = dymolaLibraries
	req = url.NewRequest()
	req.Json = data
	req.Headers = Headers
	req.Timeout = time.Minute * 10
	exportFmuRes, err := requests.Post(config.DymolaSimutalionConnect+"/dymola/translateModelFMU", req)
	if err != nil {
		errTips = "下载失败，请查看日志"
		messageMap["message"] = "Description Failed to send the http request"
		messageMap["status"] = true
		messageMap["type"] = "error"
		MessageNotice(messageMap)
		return "", false, errTips
	}
	exportResult, _ := exportFmuRes.Json()
	log.Println("dymola服务编译FMU结果：", exportResult)
	ResultCode, ok := exportResult["code"]
	if err != nil || len(exportResult) == 0 || (ok && ResultCode.(float64) != 200) {
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
	fmuFileUrl := config.DymolaSimutalionConnect + "/file/download?fileName=" + exportResult["msg"].(string)
	fmuFileRes, err := requests.Get(fmuFileUrl, req)
	if err != nil {
		return "", false, errTips
	}
	resultFmuFileData := fmuFileRes.Content
	newFilePath := "static/tmp/" + userName + "/" + strings.ReplaceAll(modelName, ".", "-") + "/" + time.Now().Local().Format("20060102150405") + "/" + fmuName + ".fmu"
	fileOperation.WriteFileByte(newFilePath, resultFmuFileData)
	return newFilePath, true, ""
}
