package serviceV2

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"yssim-go/app/v1/service"
	"yssim-go/library/fileOperation"
	"yssim-go/library/stringOperation"
	"yssim-go/library/xmlOperation"

	"github.com/google/uuid"
)

// 定义XML结构
type Root struct {
	Components []Component `xml:"Components>Component"`
	Connectors []Connector `xml:"Connectors>Connector"`
}

type Component struct {
	PathName      string      `xml:"Pathname,attr"`
	InstanceName  string      `xml:"InstanceName,attr"`
	Name          string      `xml:"-"` //防止字段在 XML 中被序列化
	Id            string      `xml:"-"` //防止字段在 XML 中被序列化
	LegalName     string      `xml:"LegalName,attr"`
	PartNumberCAD string      `xml:"PartNumber.CAD,attr"`
	PartNumberCAE string      `xml:"PartNumber.CAE,attr"`
	TypeCAD       string      `xml:"Type.CAD,attr"`
	TypeCAE       string      `xml:"Type.CAE,attr"`
	Comments      string      `xml:"Comments,attr"`
	Properties    []Property  `xml:"Properties>Property"`
	Parameters    []Parameter `xml:"Parameters>Parameter"`
}

type Property struct {
	Type   string  `xml:"Type,attr"`
	Name   string  `xml:"Name,attr"`
	Id     string  `xml:"-"` //防止字段在 XML 中被序列化
	Points []Point `xml:"Point"`
	//Value    string  `xml:"Value,attr,omitempty"`
	//Comments string  `xml:"comments,attr,omitempty"`
	Radius string `xml:"Radius,attr,omitempty"`
	Unit   string `xml:"Unit,attr,omitempty"`
	Angle  string `xml:"Angle,attr,omitempty"`
}

type Point struct {
	Name  string `xml:"Name,attr"`
	Id    string `xml:"-"` //防止字段在 XML 中被序列化
	Value string `xml:"Value,attr"`
}

type Parameter struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:"Value,attr"`
	Unit  string `xml:"Unit,attr"`
	Id    string `xml:"-"` //防止字段在 XML 中被序列化
}

type Connector struct {
	Comments string `xml:"comments,attr"`
	From     Node   `xml:"From"`
	To       Node   `xml:"To"`
}

type Node struct {
	Component string `xml:"Component,attr"`
	LegalName string `xml:"LegalName,attr"`
	Point     string `xml:"Point,attr"`
}

func CheckInfoFileXml(fileHeader *multipart.FileHeader) bool {
	// 验证管网信息文件内容
	file, _ := fileHeader.Open()
	rawData, _ := io.ReadAll(file)
	m := Root{}
	err := xml.Unmarshal(rawData, &m)

	if err != nil {
		log.Println("验证管网信息文件内容时出现错误：", err)
		return false
	}

	if m.Components == nil {
		log.Println("验证管网信息文件内容时出现错误：找不到Components字段")
		return false
	}

	for _, component := range m.Components {
		if component.InstanceName == "" {
			log.Println("验证映射配置表内容时出现错误：InstanceName字段为空")
			return false
		}
	}

	return true
}

func ParseInfoFileXml(path string) (Root, error) {
	var root Root
	xmlFile, err := os.Open(path) // 打开XML文件
	if err != nil {
		return root, errors.New("解析失败")
	}
	defer xmlFile.Close()

	err = xmlOperation.ParseXML(path, &root)
	if err != nil {
		return root, errors.New("解析失败")
	}
	for i := 0; i < len(root.Components); i++ {
		root.Components[i].LegalName = stringOperation.SanitizeName(root.Components[i].PathName)
		root.Components[i].Name = root.Components[i].PathName
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
	for i := 0; i < len(root.Connectors); i++ {
		root.Connectors[i].From.LegalName = stringOperation.SanitizeName(root.Connectors[i].From.Component)
		root.Connectors[i].To.LegalName = stringOperation.SanitizeName(root.Connectors[i].To.Component)
	}
	return root, nil
}

// SaveInfoFileXml 将 Root 结构保存为 XML 文件
func SaveInfoFileXml(path string, root Root) error {
	// 打开文件以写入
	xmlFile, err := os.Create(path)
	if err != nil {
		log.Println(err)
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

// 复制管网信息文件
func CopyPipeNetInfoFile(pipeNetInfoFilePath, userName, pipeNetInfoFileId string) (dstPath string, ok bool) {
	path := "static" + "/pipeNetInfoFile/" + userName + "/" + pipeNetInfoFileId + "/"
	strs := strings.Split(pipeNetInfoFilePath, "/")
	dstPath = path + "/" + strs[len(strs)-1]

	if ok := fileOperation.CreateFilePath(path); !ok {
		log.Println("复制管网信息文件时出现错误：创建文件父路径失败")
		return "", false
	}

	if err := fileOperation.CopyDir(pipeNetInfoFilePath, dstPath); err != nil {
		log.Println("复制管网信息文件时出现错误：", err)
		return "", false
	}

	return dstPath, true
}

func UpdatePipeNetInfoFile(pipeNetInfoFilePath, mappingConfigPath, simResPath string) (dstPath string, logList []string, ok bool) {
	var infoList []string
	// 解析管网信息文件
	pipeNetXml1, err := ParseInfoFileXml(pipeNetInfoFilePath)
	if err != nil {
		return "", infoList, false
	}

	// 将管网信息文件第二棵树的Name替换为LegalName
	for i := 0; i < len(pipeNetXml1.Components); i++ {
		pipeNetXml1.Components[i].Name = pipeNetXml1.Components[i].LegalName
	}

	// 解析映射配置表
	mappingConfig, err := GetMappingConfigDetails("", "", "", mappingConfigPath)
	haha := map[string]*Part{}
	for _, part := range mappingConfig.Parts {
		haha[part.Name] = part
	}

	// 找到要获取的参数
	// 开始进行匹配
	for i := 0; i < len(pipeNetXml1.Components); i++ {
		var found bool
		if part, ok_ := haha[pipeNetXml1.Components[i].TypeCAD]; ok_ {

			for j := 0; j < len(pipeNetXml1.Components[i].Parameters); j++ {
				var foundPamameter bool
				for _, partInfo := range part.ParameterList {
					// 如果找到匹配对象则执行下面的操作
					if pipeNetXml1.Components[i].Parameters[j].Name == partInfo.SourceName {
						// 从结果中获取数据
						variableNmae := pipeNetXml1.Components[i].Name + "." + partInfo.TargetName
						log.Println("获取参数值", variableNmae)
						data, ok__ := service.ReadSimulationResult([]string{variableNmae}, simResPath+"result_res.mat")
						log.Println(data[1])
						if ok__ && len(data) > 0 {
							ordinate := data[1]
							if len(ordinate) < 1 {
								break
							}
							//abscissa := data[0]
							value := ordinate[len(ordinate)-1]
							pipeNetXml1.Components[i].Parameters[j].Value = strconv.FormatFloat(value, 'f', -1, 64)
							foundPamameter = true

						}
						break

					}
				}

				if !foundPamameter {
					info := fmt.Sprintf("映射配置表中没有找到参数信息：零件: %s CAD类型: %s CAD参数: %s CAE类型: %s CAE参数: 缺失",
						pipeNetXml1.Components[i].Name, pipeNetXml1.Components[i].TypeCAD, pipeNetXml1.Components[i].Parameters[j].Name, pipeNetXml1.Components[i].TypeCAE)
					infoList = append(infoList, info)
				}

			}
			found = true
		}
		if !found {
			info := fmt.Sprintf("映射配置表中没有找到零件类型信息: 零件: %s CAD类型: %s CAE类型: 缺失", pipeNetXml1.Components[i].InstanceName, pipeNetXml1.Components[i].TypeCAD)
			infoList = append(infoList, info)
		}
	}
	// 将pipeNetXml1保存为新的xml文件
	updatePipeNetInfoFilePath := "static" + "/tmp/" + uuid.New().String()
	if ok = fileOperation.CreateFilePath(updatePipeNetInfoFilePath); !ok {
		log.Println("更新管网信息文件时出现错误：创建文件父路径失败")
		return "", infoList, false
	}
	err = SaveInfoFileXml(updatePipeNetInfoFilePath+"/updatePipeNetInfoFile.xml", pipeNetXml1)
	if err != nil {
		fmt.Println(err)
		return "", infoList, false
	}

	return updatePipeNetInfoFilePath + "/updatePipeNetInfoFile.xml", infoList, true
}
