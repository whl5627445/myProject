package service

import (
	"yssim-go/library/omc"
)

func GetIcon(modelName string) string {
	iconData := omc.OMC.GetIconAnnotation(modelName)
	if len(iconData) > 8 {
		imageData := iconData[8]
		if imageData.([]interface{})[0] == "Bitmap" {
			image := imageData.([]interface{})[1].([]interface{})[5].(string)
			return image
		}
	}
	return ""
}

func UploadIcon(modelName, iconData string) bool {
	annotateStr := "Icon(graphics = {Bitmap(origin = {0, 0}, extent = {{-100,100},{100,-100}}, imageSource = \"" + iconData + "\")})"
	result := omc.OMC.AddClassAnnotation(modelName, annotateStr)
	return result
}
