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

func GetResourcesImages(path string) []byte {
	realPath := omc.OMC.UriToFilename(path)
	file, err := os.ReadFile(realPath)
	if err != nil {
		log.Println(err)
	}
	return file
}

// GetResourcesImagesPath  递归获取所有文件和路径
func GetResourcesImagesPath(packageName, keyWord string) []map[string]interface{} {
	dirList := []string{""}
	dataList := []map[string]interface{}{}
	for i := 0; i < len(dirList); i++ {
		parent := dirList[i]
		resourcesPath := resourcesDir(packageName, parent)
		data, err := fileOperation.GetDirChild(resourcesPath)
		if err != nil {
			log.Println("获取Resources子节点失败，错误： ", err)
			return nil
		}
		d := []string{}
		for _, p := range data {
			path := ""
			if i > 0 {
				path = parent + "/" + p["name"]
			} else {
				path = p["name"]
			}
			if p["type"] == "dir" {
				dirList = append(dirList, path)
				continue
			}
			if strings.Contains(path, keyWord) && strings.HasSuffix(p["name"], ".png") {
				d = append(d, "modelica://"+packageName+"/Resources/"+path)
			}
		}
		if len(d) > 0 {
			if parent == "" {
				parent = "Resources"
			}
			dataList = append(dataList, map[string]interface{}{"images_path": d, "path_name": parent})
		}
	}
	return dataList
}

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

func GetResourcesFile(packageName, path string) []byte {
	pathAll := resourcesDir(packageName, path)
	file, err := os.ReadFile(pathAll)
	if err != nil {
		log.Println(err)
	}
	return file
}
