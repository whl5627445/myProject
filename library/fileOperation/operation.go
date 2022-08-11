package fileOperation

import (
	"io/ioutil"
	"os"
	"strings"
)

func CreateFile(filePath string) bool {
	filePathList := strings.Split(filePath, "/")
	path := strings.Join(filePathList[:len(filePathList)-1], "/")
	//fileName := filePathList[len(filePathList)-1]
	err := os.MkdirAll(path, 0777)
	if err != nil {
		panic(err)
	}
	return true
}

func WriteFile(fileName string, data string) bool {
	ok := CreateFile(fileName)
	if ok {
		err := ioutil.WriteFile(fileName, []byte(data), 0644)
		if err != nil {
			panic(err)
		}
	}
	return true
}
