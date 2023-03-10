package fileOperation

import (
	"container/list"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mholt/archiver/v3"
)

func CreateFilePath(filePath string) bool {
	err := os.MkdirAll(filePath, 0755)
	if err != nil {
		panic(err)
	}
	return true
}

func CreateFile(filePath string) (io.ReadWriteCloser, bool) {
	filePathList := strings.Split(filePath, "/")
	path := strings.Join(filePathList[:len(filePathList)-1], "/")
	err := os.MkdirAll(path, 0755)
	if err != nil {
		panic(err)
		return nil, false
	}

	nfs, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	return nfs, true
}

func WriteFile(fileName string, data string) bool {
	nfs, ok := CreateFile(fileName)
	if ok {
		err := os.WriteFile(fileName, []byte(data), 0755)
		if err != nil {
			return false
		}
	}
	nfs.Close()
	return true
}

func WriteFileByte(fileName string, data []byte) bool {
	nfs, ok := CreateFile(fileName)
	if ok {
		err := os.WriteFile(fileName, data, 0755)
		if err != nil {
			return false
		}
	}
	nfs.Close()
	return true
}

func UnZip(filePath string, toPath string) error {
	err := archiver.Unarchive(filePath, toPath)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func FindFile(fileName, rootPath string) (string, error) {
	_, err := os.Lstat(rootPath)
	//既不是文件，也不是文件夹
	if err != nil {
		os.RemoveAll(rootPath)
		return "", errors.New("文件格式不正确")
	}
	//文件夹添加到队列里
	queue := list.New()
	queue.PushBack(rootPath)
	for queue.Len() > 0 {
		files, _ := ioutil.ReadDir(queue.Front().Value.(string))
		for _, file := range files {
			if file.IsDir() {
				queue.PushBack(filepath.Join(rootPath, file.Name()))
			} else {
				if file.Name() == fileName {
					return queue.Front().Value.(string), nil
				}
			}
		}
		queue.Remove(queue.Front())
	}
	return "", errors.New("文件不存在")
}
