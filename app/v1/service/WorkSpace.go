package service

import (
	"time"

	"yssim-go/grpc/WorkSpace"
	"yssim-go/library/fileOperation"
)

func CreatWorkSpace(userName, SpaceName string) (string, bool) {
	path := "static/UserFiles/UploadFile/" + userName + "/Workspace/" + time.Now().Local().Format("20060102150405") + "/Workspace/" + SpaceName
	fileOperation.CreateFilePath(path + "/Resources")
	FilePath := path + "/Workspace.mo"
	fileOperation.CreateFile(FilePath)
	modelStr := CreateWorkSpace("Workspace", "", "", "package", "", false, false)
	ok := fileOperation.WriteFile(FilePath, modelStr)
	// ok = SaveModelSource("Workspace", FilePath)
	return FilePath, ok

}

func StartSMC(id string) (bool, error) {
	return WorkSpace.WS.Create(id)
}

func StopSMC(id string) bool {
	return WorkSpace.WS.Remove(id)
}
