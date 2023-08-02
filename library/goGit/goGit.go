package goGit

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"

	"log"
)

func GitPlainClone(address, FilePath, branchName string) (bool, error) {
	err := error(nil)
	// 克隆远程仓库到本地
	if branchName == "" {
		_, err = git.PlainClone(FilePath, false, &git.CloneOptions{
			URL: address,
		})
	} else {
		_, err = git.PlainClone(FilePath, false, &git.CloneOptions{
			URL:           address,
			ReferenceName: plumbing.NewBranchReferenceName(branchName), // 指定要克隆的分支
			SingleBranch:  true,                                        // 只拉取指定分支
		})
	}
	if err != nil {
		log.Println("克隆存储库失败:", err)
		return false, err
	} else {
		return true, nil
	}
}

func GetLastTag(path string) (string, error) {
	// 打开本地仓库
	repo, err := git.PlainOpen(path)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// 获取所有标签引用
	tagRefs, err := repo.Tags()
	if err != nil {
		log.Println(err)
		return "", err
	}
	tag := ""
	// 遍历标签引用并打印标签名
	tagRefs.ForEach(func(tagRef *plumbing.Reference) error {
		tag = tagRef.Name().Short()
		return nil
	})
	return tag, nil
}
