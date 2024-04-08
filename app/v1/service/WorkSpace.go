package service

import (
	"time"
	"yssim-go/config"
	"yssim-go/library/fileOperation"
)

func CreatWorkSpace(userName, SpaceName string) (string, bool) {
	path := "static/UserFiles/UploadFile/" + userName + "/Workspace/" + time.Now().Local().Format("20060102150405") + "/Workspace/" + SpaceName
	fileOperation.CreateFilePath(path + "/Resources")
	FilePath := path + "/Workspace.mo"
	fileOperation.CreateFile(FilePath)
	modelStr := CreateWorkSpace("Workspace", "", "", "package", "", false, false)
	ok := fileOperation.WriteFile(FilePath, modelStr)
	//ok = SaveModelSource("Workspace", FilePath)
	return FilePath, ok

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
