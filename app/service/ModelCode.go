package service

import (
	"strings"
	"yssim-go/library/omc"
)

func GetModelCode(modelName string) string {
	codeData := omc.OMC.List(modelName)
	codeData = strings.ReplaceAll(codeData, "\\\"", "\"")
	codeData = strings.ReplaceAll(codeData, "\\\"", "\"")
	if len(codeData) > 2 {
		codeData = codeData[1 : len(codeData)-2]
	}
	return codeData
}
