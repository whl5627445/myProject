package service

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"
	"yssim-go/library/fileOperation"
	"yssim-go/library/omc"
)

func GetResourcesList(packageName, parent string) []map[string]string {
	resourcesPath := resourcesDir(packageName, parent)
	data, err := fileOperation.GetDirChild(resourcesPath)
	if err != nil {
		log.Println("获取Resources子节点失败，错误： ", err)
		return nil
	}
	return data
}

func GetSourceFile(packageName string) string {
	return omc.OMC.GetSourceFile(packageName)
}

func resourcesDir(packageName, parent string) string {
	path := GetSourceFile(packageName)
	pathList := strings.Split(path, "/")
	packagePath := pathList[:len(pathList)-1]
	packagePath = append(packagePath, "Resources")
	if parent != "" {
		packagePath = append(packagePath, parent)
	}
	resourcesPath := strings.Join(packagePath, "/")
	return resourcesPath
}

func UploadResourcesFile(packageName, parent string, File *multipart.FileHeader) bool {
	pType := omc.OMC.IsPackage(packageName)
	if !pType {
		return false
	}
	file, _ := File.Open()
	fileData, _ := io.ReadAll(file)
	fileSavePath := resourcesDir(packageName, parent)
	name := fileSavePath + "/" + File.Filename
	return fileOperation.WriteFileByte(name, fileData)
}

func CreateResourcesDir(packageName, dirName, parent string) (bool, error) {
	path := resourcesDir(packageName, parent)
	dirPath := path + "/" + dirName
	ok, err := fileOperation.CreateFilePath(dirPath)
	return ok, err
}

func DeleteResourcesDirAndFile(packageName, name string) bool {
	path := resourcesDir(packageName, "")
	dirPath := path + "/" + name
	ok := fileOperation.DeletePathAndFile(dirPath)
	return ok
}

func GetResourcesFile(packageName, parent string) []byte {
	path := resourcesDir(packageName, parent)
	file, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
	}
	return file
}
