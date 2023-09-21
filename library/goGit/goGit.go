package goGit

import (
	"context"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"time"

	"log"
)

func GitInitVersionControl(path, username, password string) (bool, error) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		log.Println(err)
		return false, err
	}

	// 获取工作目录
	worktree, err := repo.Worktree()
	if err != nil {
		log.Println(err)
		return false, err
	}
	// 添加所有文件到仓库
	_, err = worktree.Add(".")
	if err != nil {
		log.Println(err)
		return false, err
	}
	// 创建提交
	_, err = worktree.Commit("Initial commit", &git.CommitOptions{
		Author: &object.Signature{
			Name: username,
			//Email: "your.email@example.com",
		},
	})
	if err != nil {
		log.Println(err)
		return false, err
	}

	// 推送到远程仓库
	err = repo.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth: &http.BasicAuth{
			Username: username,
			Password: password,
		},
	})
	if err != nil {
		log.Println(err)
		return false, err
	}
	return true, nil
}

func GitPlainClone(address, FilePath, branchName string) (bool, error) {
	ctx := context.TODO()
	ctx, cancel := context.WithTimeout(ctx, 180*time.Second)
	defer cancel()

	err := error(nil)
	// 克隆远程仓库到本地
	if branchName == "" {
		_, err = git.PlainCloneContext(ctx, FilePath, false, &git.CloneOptions{
			URL: address,
		})
	} else {
		_, err = git.PlainCloneContext(ctx, FilePath, false, &git.CloneOptions{
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
