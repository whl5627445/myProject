package fileOperation

import (
	"container/list"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Exists 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) // os.Stat获取文件信息
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

func DeleteProjectPath(filePath string) bool {
	filePathList := strings.Split(filePath, "/")
	path := strings.Join(filePathList[:6], "/")
	exists := Exists(path)
	if exists {
		err := os.RemoveAll(path)
		if err != nil {
			log.Println("删除文件夹出错， err: ", err)
			return false
		}
	}
	return false

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

func CreateFilePath(filePath string) bool {
	exists := Exists(filePath)
	isDir := IsDir(filePath)
	if exists && isDir {
		return true
	}
	err := os.MkdirAll(filePath, 0777)
	if err != nil {
		log.Println("仿真文件夹路径创建失败：", err)
		return false
	}
	return true
}

func CreateFile(filePath string) (io.ReadWriteCloser, bool) {
	filePathList := strings.Split(filePath, "/")
	path := strings.Join(filePathList[:len(filePathList)-1], "/")
	exists := Exists(path)
	if !exists {
		ok := CreateFilePath(path)
		if !ok {
			return nil, false
		}
	}
	nfs, err := os.Create(filePath)
	if err != nil {
		log.Println("创建文件失败: ", err)
		return nil, false
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

func FindFile(fileName, rootPath string) (string, error) {
	_, err := os.Lstat(rootPath)
	// 既不是文件，也不是文件夹
	if err != nil {
		os.RemoveAll(rootPath)
		return "", errors.New("文件格式不正确")
	}
	// 文件夹添加到队列里
	queue := list.New()
	queue.PushBack(rootPath)
	for queue.Len() > 0 {
		dir := queue.Front().Value.(string)
		files, _ := os.ReadDir(dir)
		for _, file := range files {
			if file.IsDir() {
				queue.PushBack(filepath.Join(dir, file.Name()))
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
	// 既不是文件，也不是文件夹
	if err != nil {
		log.Println("err", err)
		return nil, err
	}
	// 文件夹添加到队列里
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

// CopyDir 复制文件夹
func CopyDir(srcDir string, destDir string) error {
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		destPath := filepath.Join(destDir, path[len(srcDir):])

		if info.IsDir() {
			err := os.MkdirAll(destPath, info.Mode())
			if err != nil {
				return err
			}
		} else {
			err := copyFile(path, destPath, info.Mode())
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

// 复制文件
func copyFile(srcPath string, destPath string, mode os.FileMode) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	err = destFile.Chmod(mode)
	if err != nil {
		return err
	}

	return nil
}

// FindFileBySuffixName 查找文件夹下后缀名为suffixName的文件
func FindFileBySuffixName(suffixName, rootPath string) []string {
	// 用于保存符合条件的文件路径
	var eligibleFiles []string
	var files []string

	// 获取目录中的所有文件
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		fmt.Println("无法读取目录:", err)
		return eligibleFiles
	}

	// 遍历文件列表
	for _, file := range files {
		// 判断文件后缀名
		if strings.HasSuffix(file, suffixName) {
			eligibleFiles = append(eligibleFiles, file)
		}
	}
	return eligibleFiles
}

// 设置文件和文件夹权限为 777
func SetPermissions(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 设置权限为 777
		if err = os.Chmod(path, 0777); err != nil {
			return err
		}

		return nil
	})
}
