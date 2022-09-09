package xmlOperation

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"
)

func ParseXML(path string, obj interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return errors.New("文件打开错误: " + err.Error())
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return errors.New("读取消息错误错误：" + err.Error())
	}
	err = xml.Unmarshal(data, obj)
	return nil
}
