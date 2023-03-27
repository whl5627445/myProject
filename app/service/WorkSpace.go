package service

import (
	"yssim-go/library/fileOperation"
)

func CreatWorkSpace(userName, SpaceName string) (string, bool) {
	path := "public/UserFiles/UploadFile/" + userName + "/WorkSpace/" + SpaceName
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
