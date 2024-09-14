package serviceV2

import (
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"os"
	"sort"
	"strings"
	"yssim-go/library/fileOperation"
)

type MappingConfigData struct {
	Version            any                  `json:"version"`
	MappingDescription string               `json:"mappingDescription"`
	MappingDefinitions []*MappingDefinition `json:"mappingDefinitions"`
}

type MappingDefinition struct {
	Kind   string     `json:"kind"`
	Type   string     `json:"type"`
	Usages *UsagePipe `json:"usages"`
}

type UsagePipe struct {
	Default         *SystemModel `json:"default,omitempty"`
	PipeModel       *PipeModel   `json:"PipeModel,omitempty"`
	BranchConnector *PipeModel   `json:"BranchConnector,omitempty"`
}

type SystemModel struct {
	ModelicaClass string `json:"modelicaClass"`
	PhysicalID    string `json:"physicalID"`
	Activate      bool   `json:"activate"`
}

type PipeModel struct {
	ModelicaClass string         `json:"modelicaClass"`
	PhysicalID    string         `json:"physicalID"`
	Ports         []*MappingPair `json:"ports"`
	Parameters    []*MappingPair `json:"parameters"`
	Activate      bool           `json:"activate"`
}

type MappingPair struct {
	SourceName string `json:"sourceName"`
	TargetName string `json:"targetName"`
}

// 验证映射配置表内容
func CheckMappingConfigContent(fileHeader *multipart.FileHeader) bool {
	file, _ := fileHeader.Open()
	rawData, _ := io.ReadAll(file)
	m := MappingConfigData{}

	if err := json.Unmarshal(rawData, &m); err != nil {
		log.Println("验证映射配置表内容时出现错误：", err)
		return false
	}

	if m.MappingDefinitions == nil {
		log.Println("验证映射配置表内容时出现错误：找不到mappingDefinitions字段")
		return false
	}

	for _, mappingDefinition := range m.MappingDefinitions {
		if mappingDefinition.Kind == "" || mappingDefinition.Type == "" || mappingDefinition.Usages == nil {
			log.Println("验证映射配置表内容时出现错误：找不到Kind, Type, 或Usages字段")
			return false
		}
	}

	return true
}

// 保存映射配置表
func SaveMappingConfig(fileHeader *multipart.FileHeader, userName, mappingConfigId string) (filepath string, ok bool) {
	file, _ := fileHeader.Open()
	data, _ := io.ReadAll(file)

	filePath := "static" + "/mappingConfig/" + userName + "/" + mappingConfigId + "/" + fileHeader.Filename
	if ok := fileOperation.WriteFileByte(filePath, data); !ok {
		log.Println("保存映射配置表文件时出现错误")
		return "", false
	}
	return filePath, true
}

// 复制映射配置表
func CopyMappingConfig(mappingConfigPath, userName, newMappingConfigId string) (dstPath string, ok bool) {
	path := "static" + "/mappingConfig/" + userName + "/" + newMappingConfigId + "/"
	strs := strings.Split(mappingConfigPath, "/")
	dstPath = path + "/" + strs[len(strs)-1]

	if ok := fileOperation.CreateFilePath(path); !ok {
		log.Println("复制映射配置表时出现错误：创建文件父路径失败")
		return "", false
	}

	if err := fileOperation.CopyDir(mappingConfigPath, dstPath); err != nil {
		log.Println("复制映射配置表时出现错误：", err)
		return "", false
	}

	return dstPath, true
}

type MappingConfigParseData struct {
	System string  `json:"system"`
	Medium string  `json:"medium"`
	Parts  []*Part `json:"parts"`
}

type Part struct {
	Kind          string  `json:"kind"`
	Name          string  `json:"name"`
	ModelicaClass string  `json:"modelica_class"`
	ParameterList []*Pair `json:"parameter_list"`
	PortList      []*Pair `json:"port_list"`
}

type Pair struct {
	SourceName string `json:"source_name"`
	TargetName string `json:"target_name"`
}

// 获取映射配置表管道信息详情
func GetMappingConfigDetails(path string) (res *MappingConfigParseData, err error) {
	// 读取文件内容
	contentByte, err := os.ReadFile(path)
	if err != nil {
		log.Println("获取映射配置表管道信息详情出错: ", err)
		return nil, err
	}

	// 将文件内容映射到MappingConfigData结构体中
	m := MappingConfigData{}
	if err := json.Unmarshal(contentByte, &m); err != nil {
		log.Println("获取映射配置表管道信息详情出错: ", err)
		return nil, err
	}

	// 获取管道信息
	res = &MappingConfigParseData{
		System: "",
		Medium: "",
		Parts:  []*Part{},
	}

	for _, item := range m.MappingDefinitions {
		itemKind := item.Kind
		switch itemKind {
		case "System":
			res.System = item.Usages.Default.ModelicaClass
		case "Medium":
			res.Medium = item.Usages.Default.ModelicaClass
		case "Pipe", "Part":
			// 获取管道信息
			usagePipe := item.Usages
			onePipe := &Part{
				Kind:          item.Kind,
				Name:          item.Type,
				ModelicaClass: usagePipe.PipeModel.ModelicaClass,
				ParameterList: []*Pair{},
				PortList:      []*Pair{},
			}

			for _, each := range usagePipe.PipeModel.Parameters {
				pair := &Pair{
					SourceName: each.SourceName,
					TargetName: each.TargetName,
				}
				onePipe.ParameterList = append(onePipe.ParameterList, pair)
			}

			for _, each := range usagePipe.PipeModel.Ports {
				pair := &Pair{
					SourceName: each.SourceName,
					TargetName: each.TargetName,
				}
				onePipe.PortList = append(onePipe.PortList, pair)
			}

			res.Parts = append(res.Parts, onePipe)
		}
	}

	return res, nil
}

func FindFirstCopyNum(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	switch n {
	case 0:
		return 1
	case 1:
		if nums[0] == 0 || nums[0] == 1 {
			return nums[0] + 1
		}
		return 1
	default:
		if nums[0] > 1 {
			return 1
		}
		i, j := 0, 1
		for j < n {
			if nums[i]+1 == nums[j] {
				i, j = i+1, j+1
			} else {
				return nums[i] + 1
			}
		}
		return nums[i] + 1
	}
}
