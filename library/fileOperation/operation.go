package fileOperation

import (
	"bytes"
	"container/list"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/klauspost/compress/zip"
	"github.com/mholt/archiver/v3"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
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
		log.Println(err)
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
	err := errors.New("")
	if strings.HasSuffix(filePath, ".zip") {
		err = unZip(filePath, toPath)
	} else {
		err = archiver.Unarchive(filePath, toPath)
	}
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func unZip(zipFile string, destDir string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()
	var decodeName string
	for _, f := range zipReader.File {
		if f.Flags == 0 || f.Flags == 8 {
			//如果标致位是0  则是默认的本地编码   默认为gbk
			i := bytes.NewReader([]byte(f.Name))
			decoder := transform.NewReader(i, simplifiedchinese.GBK.NewDecoder())
			content, _ := ioutil.ReadAll(decoder)
			decodeName = string(content)
		} else {
			//如果标志为是 1 << 11也就是 2048  则是utf-8编码
			decodeName = f.Name
		}

		fpath := filepath.Join(destDir, decodeName)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
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

// 复制文件夹
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
