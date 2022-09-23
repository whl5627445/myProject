package service

import (
	"io"
	"io/ioutil"
	"os"
	"strconv"
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

// SaveModelCode 手动写入文件的形式将模型源码保存的到对应文件，并发不安全
func SaveModelCode(modelName, path string) bool {
	filesList, _ := ioutil.ReadDir("./")
	num := strconv.Itoa(len(filesList))
	os.Rename(path, path+".old"+num)
	codeData := GetModelCode(modelName)
	ok := fileOperation.WriteFile(path, codeData)
	if !ok {
		os.Rename(path+".old"+num, path)
		return false
	}
	return true
}

// SaveModelToFile 用omc提供的API将模型源码保存的到对应文件， 并发安全
func SaveModelToFile(modelName, path string) bool {
	ok := omc.OMC.SaveModel(path, modelName)
	if !ok {
		return false
	}
	return true
}

func PackageFileParse(fileName, saveFilePath, zipPackagePath string, file io.Reader) (string, bool) {
	fileOperation.CreateFilePath(saveFilePath)
	fileData, _ := ioutil.ReadAll(file)
	fileOperation.WriteFile(zipPackagePath, string(fileData))

	packagePath := ""
	if strings.HasSuffix(fileName, ".mo") {
		packagePath = zipPackagePath
	} else {
		fileOperation.UnZip(zipPackagePath, saveFilePath)
		saveFilePath, _ = fileOperation.FindFile("package.mo", saveFilePath)
		packagePath = saveFilePath + "/package.mo"
	}
	packageName, ok := omc.OMC.ParseFile(packagePath)
	if ok {
		omc.OMC.LoadFile(packagePath)
	}
	return packageName, ok
}

func ParseCodeString(code, path string) (string, bool) {
	return omc.OMC.ParseString(code, path)
}

func LoadCodeString(code, path string) bool {
	return omc.OMC.LoadString(code, path)
}

func CreateModelAndPackage(createPackageName, insertTo, expand, strType, createPackageNameAll, comment string, partial, encapsulated, state bool) bool {
	if expand != "" {
		expand = " extends " + expand + ";"
	}
	if comment != "" {
		comment = " \\\"" + comment + "\\\""
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
