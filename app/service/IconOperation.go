package service

import (
	"encoding/base64"
	"os"
	"yssim-go/library/omc"
)

func GetIcon(modelName, packageName, version string) string {
	fileName := "static/ModelicaIcons/" + packageName + "/" + version + "/" + modelName + ".svg"
	imageBytes, _ := os.ReadFile(fileName)
	if len(imageBytes) == 0 {
		iconData := omc.OMC.GetIconAnnotation(modelName)
		if len(iconData) > 1 {
			for i := 0; i < len(iconData); i += 2 {
				imageData := iconData[i]
				if imageData == "Bitmap" {
					image := iconData[i+1].([]interface{})[5].(string)
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
