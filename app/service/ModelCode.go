package service

import (
	"strings"
	"yssim-go/library/fileOperation"
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

func SaveModelCode(modelName, path string) bool {
	codeData := GetModelCode(modelName)
	ok := fileOperation.WriteFile(path, codeData)
	if !ok {
		return false
	}
	return true
}
