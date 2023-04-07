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

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func DeletePathAndFile(path string) bool {
	exists := Exists(path)
	if !exists {
		return true
	}
	isDir := IsDir(path)
	if isDir {
		err := os.RemoveAll(path)
		if err != nil {
			log.Println("删除文件夹出错， err: ", err)
			return false
		}
	} else {
		err := os.Remove(path)
		if err != nil {
			log.Println("删除文件出错， err: ", err)
			return false
		}
	}
	return true
}

func CreateFilePath(filePath string) (bool, error) {
	exists := Exists(filePath)
	isDir := IsDir(filePath)
	if exists && isDir {
		return false, errors.New("文件夹已存在")
	}
	err := os.MkdirAll(filePath, 0777)
	if err != nil {
		log.Println(err)
	}
	return true, nil
}

func CreateFile(filePath string) (io.ReadWriteCloser, bool) {
	filePathList := strings.Split(filePath, "/")
	path := strings.Join(filePathList[:len(filePathList)-1], "/")
	exists := Exists(path)
	if !exists {
		ok, err := CreateFilePath(path)
		if !ok {
			log.Println("err: ", err)
			return nil, false
		}
	}
	nfs, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	err = os.Chmod(filePath, 0777)
	if err != nil {
		log.Println("err: ", err)
		return nil, false
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
		err := os.WriteFile(fileName, data, 0777)
		if err != nil {
			return false
		}
	}
	err := nfs.Close()
	if err != nil {

		log.Println("fileName: ", fileName)
		log.Println("err: ", err)
		return false
	}
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

func Zip(filePath string, toPath string) error {
	filePathList := []string{}
	file, _ := os.ReadDir(filePath)
	for _, name := range file {
		filePathList = append(filePathList, filePath+"/"+name.Name())
	}
	err := archiver.Archive(filePathList, toPath)
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
		files, _ := os.ReadDir(queue.Front().Value.(string))
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

func GetDirChild(rootPath string) ([]map[string]string, error) {
	_, err := os.Lstat(rootPath)
	//既不是文件，也不是文件夹
	if err != nil {
		log.Println("err", err)
		return nil, err
	}
	//文件夹添加到队列里
	dataList := make([]map[string]string, 0, 1)
	queue := list.New()
	queue.PushBack(rootPath)
	for queue.Len() > 0 {

		name := queue.Front().Value.(string)
		files, _ := ioutil.ReadDir(name)
		for _, file := range files {
			if file.IsDir() {
				dataList = append(dataList, map[string]string{"name": file.Name(), "type": "dir"})
			} else {
				dataList = append(dataList, map[string]string{"name": file.Name(), "type": "file"})
			}
		}
		queue.Remove(queue.Front())
	}
	return dataList, nil
}
