package service

import (
	"yssim-go/library/omc"
)

func GetIcon(modelName string) string {
	iconData := omc.OMC.GetIconAnnotation(modelName)
	if len(iconData) > 7 {
		imageData := iconData[8]
		if imageData.([]interface{})[0] == "Bitmap" {
			image := imageData.([][]string)[1][5]
			return image
		}
	}
	return ""
}
