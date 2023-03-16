package service

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"yssim-go/app/DataBaseModel"
	"yssim-go/library/fileOperation"
	"yssim-go/library/omc"
)

func GetModelFileCode(modelName string) string {
	codeData := omc.OMC.ListFile(modelName)
	codeData = strings.ReplaceAll(codeData, "\\\"", "\"")
	codeData = strings.ReplaceAll(codeData, "\\\"", "\"")
	if len(codeData) > 2 {
		codeData = codeData[1 : len(codeData)-2]
	}
	return codeData
}

func GetModelCode(modelName string) string {
	codeData := omc.OMC.List(modelName)
	codeData = strings.ReplaceAll(codeData, "\\\"", "\"")
	codeData = strings.ReplaceAll(codeData, "\\\"", "\"")
	if len(codeData) > 2 {
		codeData = codeData[1 : len(codeData)-2]
	}
	return codeData
}

// SaveModelCode  保存模型，备份旧文件
func SaveModelCode(modelName, path string) bool {
	pathList := strings.Split(path, "/")
	numPath := strings.Join(pathList[:len(pathList)-1], "/")
	filesList, _ := ioutil.ReadDir(numPath)
	num := strconv.Itoa(len(filesList))
	if len(filesList) != 0 {
		err := os.Rename(path, path+".old"+num)
		if err != nil {
			log.Println("重命名文件结果", err)
			return false
		}
	}
	_, ok := fileOperation.CreateFile(path)
	if ok {
		pwd, _ := os.Getwd()
		ok = SaveModelToFile(modelName, pwd+"/"+path)
	}
	if !ok {
		os.Remove(path)
		err := os.Rename(path+".old"+num, path)
		if err != nil {
			return false
		}
		return false
	}
	return true
}

// SaveModelToFile 用omc提供的API将模型源码保存的到对应文件， 并发安全
func SaveModelToFile(modelName, path string) bool {
	ok := omc.OMC.SaveModel(path, modelName)
	return ok
}

func PackageFileParse(fileName, saveFilePath, zipPackagePath string, file io.Reader) (string, bool) {
	fileOperation.CreateFilePath(saveFilePath)
	fileData, _ := ioutil.ReadAll(file)
	fileOperation.WriteFile(zipPackagePath, string(fileData))

	packagePath := ""
	if strings.HasSuffix(fileName, ".mo") {
		packagePath = zipPackagePath
	} else {
		err := fileOperation.UnZip(zipPackagePath, saveFilePath)
		if err != nil {
			return "", false
		}
		packageFilePath, err := fileOperation.FindFile("package.mo", saveFilePath)
		if err != nil {
			return "", false
		}
		packagePath = packageFilePath + "/package.mo"
	}
	packageName, ok := omc.OMC.ParseFile(packagePath)
	if ok {
		ok = omc.OMC.LoadFile(packagePath)
	}
	return packageName, ok
}

func ParseCodeString(code, path string) (string, bool) {
	return omc.OMC.ParseString(code, path)
}

func LoadCodeString(code, path string) bool {
	return omc.OMC.LoadString(code, path)
}

func LoadFile(packagePath string) bool {
	return omc.OMC.LoadFile(packagePath)
}

func CreateModelAndPackage(createPackageName, insertTo, expand, strType, createPackageNameAll, comment string, partial, encapsulated, state bool) bool {
	if expand != "" {
		expand = " extends " + expand + ";"
	}
	if comment != "" {
		comment = " \"" + comment + "\""
	}
	modelStrBase := strType + " " + createPackageName + comment + expand + " end " + createPackageName + ";"
	modelStr := ""
	if insertTo != "" {
		modelStr = "within " + insertTo + "; "
	}
	if encapsulated {
		modelStr = modelStr + "encapsulated "
	}
	if partial {
		modelStr = modelStr + "partial "
	}
	modelStr = modelStr + modelStrBase
	res := omc.OMC.LoadString(modelStr, createPackageName)
	if state {
		omc.OMC.AddClassAnnotation(createPackageNameAll, "Icon(graphics={Text(extent={{-100,100},{100,-100}},textString=\"%name\")})")
		omc.OMC.AddClassAnnotation(createPackageNameAll, "annotate=__Dymola_state(true)")
		omc.OMC.AddClassAnnotation(createPackageNameAll, "singleInstance(true)")
	}
	if res {
		return true
	}
	return false
}

func SaveModelToFileALL(packageModel []DataBaseModel.YssimModels) {
	libraryAndVersions := GetLibraryAndVersions()
	for i := 0; i < len(packageModel); i++ {
		p, ok := libraryAndVersions[packageModel[i].PackageName]
		if ok && p == packageModel[i].Version {
			SaveModelToFile(packageModel[i].PackageName, packageModel[i].FilePath)
		}
	}
}
