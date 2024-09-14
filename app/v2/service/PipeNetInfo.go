package serviceV2

import (
	"errors"
	"io"
	"log"
	"mime/multipart"
	"os"
	"yssim-go/library/fileOperation"
	"yssim-go/library/xmlOperation"
)

// 定义XML结构
type Root struct {
	Components []Component `xml:"Components>Component"`
	Connectors []Connector `xml:"Connectors>Connector"`
}

type Component struct {
	InstanceName string      `xml:"InstanceName,attr"`
	PartNumber   string      `xml:"PartNumber.CAD,attr"`
	TypeCAD      string      `xml:"Type.CAD,attr"`
	TypeCAE      string      `xml:"Type.CAE,attr"`
	Properties   []Property  `xml:"Properties>Property"`
	Parameters   []Parameter `xml:"parameters>parameter"`
}

type Property struct {
	Type   string  `xml:"Type,attr"`
	Name   string  `xml:"Name,attr"`
	Points []Point `xml:"Point"`
	Value  string  `xml:"Value,attr,omitempty"`
}

type Point struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:"Value,attr"`
}

type Parameter struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:"Value,attr"`
}

type Connector struct {
	Comments string `xml:"comments,attr"`
	From     Node   `xml:"From"`
	To       Node   `xml:"To"`
}

type Node struct {
	Component string `xml:"Component,attr"`
	Point     string `xml:"Point,attr"`
}

func ParseInfoFileXml(path string) (error, Root) {
	var root Root
	xmlFile, err := os.Open(path) // 打开XML文件
	if err != nil {
		return errors.New("解析失败"), root
	}
	defer xmlFile.Close()

	err = xmlOperation.ParseXML(path, &root)
	if err != nil {
		return errors.New("解析失败"), root
	}
	return nil, root
}

// 保存管网信息文件
func SavePipeNetInfoFile(fileHeader *multipart.FileHeader, userName, pipeNetInfoFileId string) (filepath string, ok bool) {
	file, _ := fileHeader.Open()
	data, _ := io.ReadAll(file)

	filePath := "static" + "/pipeNetInfoFile/" + userName + "/" + pipeNetInfoFileId + "/" + fileHeader.Filename
	if ok = fileOperation.WriteFileByte(filePath, data); !ok {
		log.Println("保存管网信息文件时出现错误")
		return "", false
	}
	return filePath, true
}
