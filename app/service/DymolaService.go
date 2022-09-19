package service

import (
	"github.com/wangluozhe/requests"
	"log"
	"strings"
	"time"
	"yssim-go/config"

	"github.com/wangluozhe/requests/url"
)

func DymolaFmuExport(fmuPar map[string]interface{}, token, username, fmuName, packageName, modelName, fileName, filePath string) ([]byte, bool) {
	data := map[string]interface{}{
		"username":    username,
		"fmuPar":      fmuPar,
		"modelName":   fmuName,
		"fileName":    "",
		"modelToOpen": modelName,
		"token":       token,
	}
	res := true
	urlStr := packageName + "/" + strings.ReplaceAll(modelName, ".", "-") + "/" + time.Now().Local().Format("20060102150405")
	req := url.NewRequest()
	params := url.NewParams()
	params.Set("url", username+"/"+urlStr)
	req.Params = params
	req.Timeout = 600 * time.Second
	var resultFmuFileData []byte
	if filePath != "" {
		files := url.NewFiles()
		files.SetFile("file", fileName+".mo", filePath, "")
		req.Files = files
		uploadFileRes, err := requests.Post(config.DymolaConnect+"/file/upload", req)
		if err != nil {
			log.Println(err)
			return resultFmuFileData, false
		}
		uploadFileResJson, _ := uploadFileRes.Json()
		uploadResultCode, ok := uploadFileResJson["code"]
		if ok && uploadResultCode.(float64) == 200 {
			data["fileName"] = urlStr + "/" + fileName + ".mo"
		} else {
			res = false

		}
	}
	req = url.NewRequest()
	req.Json = data
	exportFmuRes, err := requests.Post(config.DymolaConnect+"/dymola/translateModelFMU", req)
	exportResult, _ := exportFmuRes.Json()
	ResultCode, ok := exportResult["code"]
	if err != nil || len(exportResult) == 0 || (ok && ResultCode.(float64) != 200) {
		log.Println(err)
		return resultFmuFileData, false
	}
	if res {
		req = url.NewRequest()
		fmuFileUrl := config.DymolaConnect + "/file/download?fileName=" + exportResult["msg"].(string)
		fmuFileRes, err := requests.Get(fmuFileUrl, req)
		if err != nil {
			return resultFmuFileData, false
		}
		resultFmuFileData = fmuFileRes.Content
	}
	return resultFmuFileData, true
}
