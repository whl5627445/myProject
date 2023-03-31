package service

import (
	"encoding/base64"
	"log"
	"os"
	"strings"
	"yssim-go/library/omc"
)

func GetIcon(modelName, packageName, version string) string {
	fileName := "static/ModelicaIcons/" + packageName + "/" + version + "/" + modelName + ".svg"
	imageBytes, _ := os.ReadFile(fileName)
	if len(imageBytes) == 0 {
		iconData := omc.OMC.GetIconAnnotation(modelName)
		if len(iconData) >= 8 {
			bitmapData := iconData[8].([]interface{})
			for i := 0; i < len(bitmapData); i += 2 {
				imageData := bitmapData[i]
				if imageData == "Bitmap" {
					image := bitmapData[i+1].([]interface{})[5].(string)
					imageUri := bitmapData[i+1].([]interface{})[4].(string)
					if strings.HasPrefix(imageUri, "modelica://") {
						imageFile := omc.OMC.UriToFilename(imageUri)
						file, err := os.ReadFile(imageFile)
						if err != nil {
							log.Println("获取模型图表文件信息失败: ", err)
							return ""
						}
						fileBase64Str := base64.StdEncoding.EncodeToString(file)
						return "data:image/png;base64," + fileBase64Str
					}
					return "data:image/png;base64," + image
				}
			}
		}
		return ""
	}
	return "data:image/svg+xml;base64," + base64.StdEncoding.EncodeToString(imageBytes)
}

func UploadIcon(modelName, iconData string) bool {
	annotateStr := "Icon(graphics = {Bitmap(origin = {0, 0}, extent = {{-100,100},{100,-100}}, imageSource = \"" + iconData + "\")})"
	result := omc.OMC.AddClassAnnotation(modelName, annotateStr)
	return result
}
