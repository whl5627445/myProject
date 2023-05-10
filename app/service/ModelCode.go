package service

import (
	"errors"
	"io"
	"log"
	"os"
	"strings"
	"time"
	"yssim-go/config"
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

// SaveModelCode  保存模型到指定文件
func SaveModelCode(modelName, path string) bool {
	pathList := strings.Split(path, "/")
	numPath := strings.Join(pathList[:len(pathList)-1], "/")
	filesList, _ := os.ReadDir(numPath)
	ok := false
	if len(filesList) == 0 {
		_, ok = fileOperation.CreateFile(path)
	}
	ok = SaveModelSource(modelName, path)
	return ok
}

// SaveModelSource 用omc提供的API将模型源码保存的到对应文件， 并发安全
func SaveModelSource(modelName, path string) bool {
	ok := omc.OMC.SetSourceFile(modelName, path)
	omc.OMC.Save(modelName)
	//ok := go omc.OMC.SaveModel(path, modelName)
	return ok
}

func SaveModelToFile(modelName, path string) bool {
	pathList := strings.Split(path, "/")
	numPath := strings.Join(pathList[:len(pathList)-1], "/")
	filesList, _ := os.ReadDir(numPath)
	ok := false
	if len(filesList) == 0 {
		_, ok = fileOperation.CreateFile(path)
	}
	code := GetModelCode(modelName)
	ok = fileOperation.WriteFile(path, code)
	return ok
}

// ModelSave 用omc提供的API将模型源码保存的到对应文件， 并发安全
func ModelSave(modelName string) {
	//ok := omc.OMC.Save(modelName)
	config.ModelCodeChan <- modelName
}

func DeletePackageFile(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}

func PackageFileParse(fileName, saveFilePathBase string, file io.Reader) (string, string, string, bool) {
	dirName := strings.Split(fileName, ".")[0]
	saveFilePath := saveFilePathBase + "/" + dirName
	zipPackagePath := saveFilePath + "/" + fileName
	fileOperation.CreateFilePath(saveFilePath)
	fileData, _ := io.ReadAll(file)
	fileOperation.WriteFile(zipPackagePath, string(fileData))

	packagePath := ""
	packageFilePath := ""
	if strings.HasSuffix(fileName, ".mo") {
		pathList := strings.Split(zipPackagePath, "/")
		packagePath = zipPackagePath
		packageFilePath = strings.Join(pathList[:len(pathList)-1], "/")
	} else {
		err := fileOperation.UnZip(zipPackagePath, saveFilePath)
		if err != nil {
			log.Println("UnZip err", err)
			return "", "", "", false
		}
		os.Remove(zipPackagePath)
		packageFilePath, err = fileOperation.FindFile("package.mo", saveFilePath)
		if err != nil {
			log.Println("FindFile err", err)
			return "", "", "未找到package", false
		}
		packagePath = packageFilePath + "/package.mo"
	}
	packageName, ok := omc.OMC.ParseFile(packagePath)
	msg := ""
	if ok {
		pathList := strings.Split(packagePath, "/")
		packagePathNew := packagePath
		if pathList[len(pathList)-2] != packageName {
			pathList[len(pathList)-2] = packageName
			packageFilePathNew := strings.Join(pathList[:len(pathList)-1], "/")
			packagePathNew = strings.Join(pathList, "/")
			os.Rename(packageFilePath, packageFilePathNew)
			packagePath = packagePathNew
		}
		ok = omc.OMC.LoadFile(packagePathNew)
	}
	if !ok {
		msg = "语法错误，请重新检查后上传"
	}
	return packageName, packagePath, msg, ok
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
		//omc.OMC.AddClassAnnotation(createPackageNameAll, "annotate=__Dymola_state(true)")
		omc.OMC.AddClassAnnotation(createPackageNameAll, "singleInstance(true)")
	}
	return res
}

func ZipPackage(packageName, path string) (string, error) {
	tmpPath := "static/tmp/" + time.Now().Local().Format("20060102150405") + "/" + packageName + ".zip"
	packagePathList := strings.Split(path, "/")
	packagePath := strings.Join(packagePathList[:len(packagePathList)-2], "/")
	err := fileOperation.Zip(packagePath, tmpPath)
	if err != nil {
		return "", errors.New("模型包压缩失败，错误为：" + err.Error())
	}
	return tmpPath, nil
}

func ZipPackageStream(packageName, path string) (string, error) {
	tmpPath, err := ZipPackage(packageName, path)
	if err != nil {
		return "", errors.New("压缩文件包失败，错误为：" + err.Error())
	}
	//data, err := os.ReadFile(tmpPath)
	if err != nil {
		return "", errors.New("读取文件失败，错误为：" + err.Error())
	}
	//os.Remove(tmpPath)
	return tmpPath, nil
}
