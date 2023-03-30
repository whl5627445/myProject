package service

import (
	"time"
	"yssim-go/config"
	"yssim-go/library/fileOperation"
)

func CreatWorkSpace(userName, SpaceName string) (string, bool) {
	path := "public/UserFiles/UploadFile/" + userName + "/WorkSpace/" + time.Now().Local().Format(time.RFC3339) + "/" + SpaceName
	fileOperation.CreateFilePath(path + "/Resources")
	FilePath := path + "/WorkSpace.mo"
	fileOperation.CreateFile(FilePath)
	ok := CreateModelAndPackage("WorkSpace", "", "", "package", "", "", false, false, false)
	if ok {
		ok = SaveModelToFile("WorkSpace", FilePath)
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
