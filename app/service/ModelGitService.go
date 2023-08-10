package service

import (
	"errors"
	"log"
	"os"
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

func RepositoryClone(address, branchName, userName string) (string, string, string, bool) {
	msg := ""
	// 获取这个储存库的名称
	repositoryName, err := GetRepositoryName(address)
	if err != nil {
		msg = "Git存储库地址错误！"
		return "", "", msg, false
	}
	// 创建本地存储库路径
	repositoryPath := "static/UserFiles/UploadFile/" + userName + "/" + time.Now().Local().Format("20060102150405") + "/" + repositoryName + "/"
	fileOperation.CreateFilePath(repositoryPath)
	// 克隆到本地
	res, err := goGit.GitPlainClone(address, repositoryPath, branchName)
	if res {
		return repositoryPath, repositoryName, "", true
	} else {
		//克隆失败清除垃圾文件
		err_ := os.RemoveAll(repositoryPath)
		if err_ != nil {
			log.Println("删除本地存储库路径出错:", err_)
		}

		msg = err.Error()
		return "", "", msg, false
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
			return "", "", "未找到合理的.mo文件！", false
		}
		packagePath = packageFilePath + "package.mo"
	}

	packageName, ok := omc.OMC.ParseFile(packagePath)
	msg := ""
	if !ok {
		//克隆失败清除垃圾文件
		err_ := os.RemoveAll(repositoryPath)
		if err_ != nil {
			log.Println("删除本地存储库路径出错:", err_)
		}
		msg = "语法错误，请检查模型！"
	}

	return packageName, packagePath, msg, ok

}
