package service

import (
	"time"
	"yssim-go/config"
	"yssim-go/library/fileOperation"
)

func CreatWorkSpace(userName, SpaceName string) (string, bool) {
	path := "static/UserFiles/UploadFile/" + userName + "/Workspace/" + time.Now().Local().Format("20060102150405") + "/" + SpaceName
	fileOperation.CreateFilePath(path + "/Resources")
	FilePath := path + "/Workspace.mo"
	fileOperation.CreateFile(FilePath)
	ok := CreateModelAndPackage("Workspace", "", "", "package", "", "", false, false, false)
	if ok {
		ok = SaveModelSource("Workspace", FilePath)
		return FilePath, ok
	}
	return "", ok
}

func SetWorkSpaceId(spaceId *string) bool {
	result := GetWorkSpaceId(spaceId)
	if !result {
		config.UserSpaceId = *spaceId
	}
	return result
}

func GetWorkSpaceId(spaceId *string) bool {
	userSpaceId := config.UserSpaceId
	if *spaceId == userSpaceId {
		return true
	}
	return false
}
