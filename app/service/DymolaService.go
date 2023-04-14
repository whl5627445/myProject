package service

import (
	"log"
	"strconv"
	"strings"
	"time"
	"yssim-go/config"

	"github.com/google/uuid"

	"github.com/wangluozhe/requests"

	"github.com/wangluozhe/requests/url"
)

func DymolaFmuExport(fmuPar map[string]interface{}, token, userName, fmuName, packageName, modelName, fileName, filePath string) (resultFmuFileData []byte, res bool) {
	data := map[string]interface{}{
		"username":    userName,
		"fmuPar":      fmuPar,
		"modelName":   fmuName,
		"fileName":    "",
		"modelToOpen": modelName,
		"token":       token,
		"taskId":      uuid.New().String(),
	}
	urlStr := packageName + "/" + strings.ReplaceAll(modelName, ".", "-") + "/" + strconv.Itoa(int(time.Now().UnixNano()))
	req := url.NewRequest()
	params := url.NewParams()
	params.Set("url", userName+"/"+urlStr)
	req.Params = params
	Headers := url.NewHeaders()
	Headers.Set("Authorization", token)
	req.Headers = Headers
	req.Timeout = 600 * time.Second
	if filePath != "" {
		files := url.NewFiles()
		files.SetFile("file", fileName+".mo", filePath, "")
		req.Files = files
		uploadFileRes, err := requests.Post(config.DymolaSimutalionConnect+"/file/upload", req)
		if err != nil {
			log.Println(err)
			return resultFmuFileData, false
		}
		uploadFileResJson, _ := uploadFileRes.Json()
		log.Println("dymola服务上传文件结果：", uploadFileResJson)
		uploadResultCode, ok := uploadFileResJson["code"]
		if ok && uploadResultCode.(float64) == 200 {
			data["fileName"] = urlStr + "/" + fileName + ".mo"
		} else {
			return resultFmuFileData, false

		}
	}
	req = url.NewRequest()
	req.Json = data
	req.Headers = Headers
	req.Timeout = time.Minute * 10
	exportFmuRes, err := requests.Post(config.DymolaSimutalionConnect+"/dymola/translateModelFMU", req)
	if err != nil {
		log.Println(err)
		return resultFmuFileData, false
	}
	exportResult, _ := exportFmuRes.Json()
	log.Println("dymola服务编译FMU结果：", exportResult)
	ResultCode, ok := exportResult["code"]
	if err != nil || len(exportResult) == 0 || (ok && ResultCode.(float64) != 200) {
		return resultFmuFileData, false
	}
	req = url.NewRequest()
	req.Headers = Headers
	req.Timeout = time.Minute * 10
	fmuFileUrl := config.DymolaSimutalionConnect + "/file/download?fileName=" + exportResult["msg"].(string)
	fmuFileRes, err := requests.Get(fmuFileUrl, req)
	if err != nil {
		return resultFmuFileData, false
	}
	resultFmuFileData = fmuFileRes.Content
	return resultFmuFileData, true
}
