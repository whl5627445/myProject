package serviceV2

import (
	"encoding/xml"
	"errors"
	"github.com/google/uuid"
	"io"
	"log"
	"mime/multipart"
	"os"
	"yssim-go/library/fileOperation"
	"yssim-go/library/stringOperation"
	"yssim-go/library/xmlOperation"
)

// 定义XML结构
type Root struct {
	Components []Component `xml:"Components>Component"`
	Connectors []Connector `xml:"Connectors>Connector"`
}

type Component struct {
	InstanceName string      `xml:"InstanceName,attr"`
	Name         string      `xml:"-"` //防止字段在 XML 中被序列化
	Id           string      `xml:"-"` //防止字段在 XML 中被序列化
	LegalName    string      `xml:"LegalName,attr"`
	PartNumber   string      `xml:"PartNumber.CAD,attr"`
	TypeCAD      string      `xml:"Type.CAD,attr"`
	TypeCAE      string      `xml:"Type.CAE,attr"`
	Comments     string      `xml:"comments,attr"`
	Properties   []Property  `xml:"Properties>Property"`
	Parameters   []Parameter `xml:"parameters>parameter"`
}

type Property struct {
	Type     string  `xml:"Type,attr"`
	Name     string  `xml:"Name,attr"`
	Id       string  `xml:"-"` //防止字段在 XML 中被序列化
	Points   []Point `xml:"Point"`
	Value    string  `xml:"Value,attr,omitempty"`
	Comments string  `xml:"comments,attr"`
}

type Point struct {
	Name  string `xml:"Name,attr"`
	Id    string `xml:"-"` //防止字段在 XML 中被序列化
	Value string `xml:"Value,attr"`
}

type Parameter struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:"Value,attr"`
	Id    string `xml:"-"` //防止字段在 XML 中被序列化
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
	for i := 0; i < len(root.Components); i++ {
		root.Components[i].LegalName = stringOperation.SanitizeName(root.Components[i].InstanceName)
		root.Components[i].Name = root.Components[i].InstanceName
		root.Components[i].Id = uuid.New().String()
		for j := 0; j < len(root.Components[i].Properties); j++ {
			root.Components[i].Properties[j].Id = uuid.New().String()
			for k := 0; k < len(root.Components[i].Properties[j].Points); k++ {
				root.Components[i].Properties[j].Points[k].Id = uuid.New().String()
			}
		}
		for j := 0; j < len(root.Components[i].Parameters); j++ {
			root.Components[i].Parameters[j].Id = uuid.New().String()
		}

	}
	return nil, root
}

// SaveInfoFileXml 将 Root 结构保存为 XML 文件
func SaveInfoFileXml(path string, root Root) error {
	// 打开文件以写入
	xmlFile, err := os.Create(path)
	if err != nil {
		return errors.New("无法创建文件")
	}
	defer xmlFile.Close()

	// 序列化结构为 XML 格式
	output, err := xml.MarshalIndent(root, "", "  ") // 以缩进的方式生成XML
	if err != nil {
		return errors.New("XML序列化失败")
	}

	// 写入 XML 头
	xmlFile.Write([]byte(xml.Header))
	// 将序列化后的 XML 数据写入文件
	_, err = xmlFile.Write(output)
	if err != nil {
		return errors.New("写入XML文件失败")
	}

	return nil
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
