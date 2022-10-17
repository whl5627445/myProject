package service

import (
	"yssim-go/library/omc"
)

func GetIcon(modelName string) string {
	iconData := omc.OMC.GetIconAnnotation(modelName)
	if len(iconData) > 1 {
		for i := 0; i < len(iconData); i += 2 {
			imageData := iconData[i]
			if imageData == "Bitmap" {
				image := iconData[i+1].([]interface{})[5].(string)
				return image
			}
		}
	}
	return ""
}

func UploadIcon(modelName, iconData string) bool {
	annotateStr := "Icon(graphics = {Bitmap(origin = {0, 0}, extent = {{-100,100},{100,-100}}, imageSource = \"" + iconData + "\")})"
	result := omc.OMC.AddClassAnnotation(modelName, annotateStr)
	return result
}
