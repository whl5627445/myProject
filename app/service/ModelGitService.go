package service

import (
	"errors"
	"log"
	"path/filepath"
	"time"
	"yssim-go/library/fileOperation"
	"yssim-go/library/goGit"
	"yssim-go/library/omc"
)

func GetRepositoryName(cloneURL string) (string, error) {
	base := filepath.Base(cloneURL)
	ext := filepath.Ext(base)

	if ext != ".git" {
		return "", errors.New("Invalid Git URL")
	}

	repoName := base[0 : len(base)-len(ext)]
	return repoName, nil
}

func RepositoryClone(address, branchName, userName string) (string, string, bool) {
	// 获取这个储存库的名称
	repositoryName, err := GetRepositoryName(address)
	if err != nil {
		log.Println("储存库的名称解析错误:", err)
		return "", "", false
	}
	// 创建本地存储库路径
	repositoryPath := "static/UserFiles/UploadFile/" + userName + "/" + time.Now().Local().Format("20060102150405") + "/" + repositoryName + "/"
	fileOperation.CreateFilePath(repositoryPath)
	// 克隆到本地
	res, _ := goGit.GitPlainClone(address, repositoryPath, branchName)
	if res {
		return repositoryPath, repositoryName, true
	} else {
		return repositoryPath, repositoryName, false
	}

}

func GetTag(path string) string {
	tag, err := goGit.GetLastTag(path)
	if err != nil {
		log.Println("获取标签失败：", err)
	}
	return tag
}

func GitPackageFileParse(repositoryName, repositoryPath string) (string, string, string, bool) {
	packagePath := ""
	if fileOperation.Exists(repositoryPath + "/" + repositoryName + ".mo") {
		packagePath = repositoryPath + repositoryName + ".mo"
	} else {
		packageFilePath, err := fileOperation.FindFile("package.mo", repositoryPath)
		if err != nil {
			log.Println("FindFile err", err)
			return "", "", "未找到package", false
		}
		packagePath = packageFilePath + "package.mo"
	}

	packageName, ok := omc.OMC.ParseFile(packagePath)
	msg := ""
	//packageVersion := ""
	//if ok {
	//	ok = grpcOmc.LoadFile(packagePath)
	//	packageVersion = grpcOmc.GetPackageVersion(packageName)
	//	grpcOmc.DeleteClass(packageName)
	//}
	if !ok {
		msg = "语法错误，请重新检查后上传"
	}

	return packageName, packagePath, msg, ok

}
