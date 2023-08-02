package service

import (
	"errors"
	"log"
	"path/filepath"
	"yssim-go/grpc/GoOmcGrpc/grpcOmc"
	"yssim-go/library/fileOperation"
	"yssim-go/library/goGit"
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

func RepositoryClone(address, FilePath, branchName string) bool {

	res, _ := goGit.GitPlainClone(address, FilePath, branchName)
	if res {
		return true
	} else {
		return false
	}
	// omc解析

}

func GetTag(path string) string {
	tag, err := goGit.GetLastTag(path)
	if err != nil {
		log.Println("获取标签失败：", err)
	}
	return tag
}

func GitPackageFileParse(repositoryName, repositoryPath string) (string, string, string, string, bool) {
	packagePath := ""
	if fileOperation.Exists(repositoryPath + "/" + repositoryName + ".mo") {
		packagePath = repositoryPath + repositoryName + ".mo"
		log.Println(repositoryName+".mo文件存在", packagePath)
	} else {
		packageFilePath, err := fileOperation.FindFile("package.mo", repositoryPath)
		if err != nil {
			log.Println("FindFile err", err)
			return "", "", "", "未找到package", false
		}
		packagePath = packageFilePath + "package.mo"
		log.Println("package.mo文件存在", packagePath)
	}

	packageName, ok := grpcOmc.ParseFile(packagePath)
	msg := ""
	packageVersion := ""
	if ok {
		ok = grpcOmc.LoadFile(packagePath)
		packageVersion = grpcOmc.GetPackageVersion(packageName)
		DeleteClassReturn := grpcOmc.DeleteClass(packageName)
		log.Println("DeleteClassReturn结果:", DeleteClassReturn)
	}
	if !ok {
		msg = "语法错误，请重新检查后上传"
	}
	log.Println("处理结果:", packageName, packagePath, packageVersion, msg, ok)

	return packageName, packagePath, packageVersion, msg, ok

}
