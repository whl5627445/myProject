package serviceV2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"yssim-go/app/DataBaseModel"
	"yssim-go/app/DataType"
	"yssim-go/app/v1/service"
	serviceV1 "yssim-go/app/v1/service"
	"yssim-go/config"
	"yssim-go/library/fileOperation"

	"github.com/bytedance/sonic"
	jsonpatch "github.com/evanphx/json-patch"
)

type MappingConfigData struct {
	MappingDefinitions []*MappingDefinition `json:"mappingDefinitions"`
}

type MappingDefinition struct {
	Kind   string     `json:"kind"`
	Type   string     `json:"type"`
	Usages *UsagePipe `json:"usages"`
}

type UsagePipe struct {
	Default   *SystemModel `json:"default,omitempty"`
	PipeModel *PipeModel   `json:"PipeModel,omitempty"`
}

type SystemModel struct {
	ModelicaClass string `json:"modelicaClass"`
}

type PipeModel struct {
	ModelicaClass string         `json:"modelicaClass"`
	Ports         []*MappingPair `json:"ports"`
	Parameters    []*MappingPair `json:"parameters"`
}

type MappingPair struct {
	SourceName string `json:"sourceName"`
	TargetName string `json:"targetName"`
}

// 验证映射配置表内容
func CheckMappingConfigContent(fileHeader *multipart.FileHeader) (bool, string) {
	file, _ := fileHeader.Open()
	rawData, _ := io.ReadAll(file)
	m := MappingConfigData{}
	errMessage := ""

	// 验证基本格式
	if err := json.Unmarshal(rawData, &m); err != nil {
		errMessage = fmt.Sprintf("验证映射配置表时出现错误：%s", err.Error())
		log.Println(errMessage)
		return false, errMessage
	}

	// 验证mappingDefinitions字段
	if m.MappingDefinitions == nil {
		errMessage = "验证映射配置表时出现错误: 找不到mappingDefinitions字段"
		log.Println(errMessage)
		return false, errMessage
	}

	// 验证system，medium和零件类型
	systemMap := map[string]struct{}{}
	mediumMap := map[string]struct{}{}
	partMap := map[string]struct{}{}

	for _, mappingDefinition := range m.MappingDefinitions {
		if mappingDefinition.Kind == "" || mappingDefinition.Type == "" || mappingDefinition.Usages == nil {
			errMessage = "验证映射配置表时出现错误：找不到Kind, Type, 或Usages字段"
			log.Println(errMessage)
			return false, errMessage
		}

		switch mappingDefinition.Kind {
		case "System":
			if _, ok := systemMap["System"]; ok {
				errMessage = "验证映射配置表时出现错误: 重复定义了System"
				log.Println(errMessage)
				return false, errMessage
			}
			systemMap["System"] = struct{}{}
		case "Medium":
			if _, ok := mediumMap["Medium"]; ok {
				errMessage = "验证映射配置表时出现错误: 重复定义了Medium"
				log.Println(errMessage)
				return false, errMessage
			}
			mediumMap["Medium"] = struct{}{}
		case "Pipe":
			if _, ok := partMap[mappingDefinition.Type]; ok {
				errMessage = fmt.Sprintf("验证映射配置表时出现错误: 重复定义了类型 %s", mappingDefinition.Type)
				log.Println(errMessage)
				return false, errMessage
			}
			partMap[mappingDefinition.Type] = struct{}{}

			sourceNameMap := map[string]struct{}{}
			targetNameMap := map[string]struct{}{}
			for _, parameterPair := range mappingDefinition.Usages.PipeModel.Parameters {
				if _, ok := sourceNameMap[parameterPair.SourceName]; ok {
					errMessage = fmt.Sprintf("验证映射配置表时出现错误: 类型'%s'重复定义了参数'%s'", mappingDefinition.Type, parameterPair.SourceName)
					log.Println(errMessage)
					return false, errMessage
				}
				sourceNameMap[parameterPair.SourceName] = struct{}{}

				if _, ok := targetNameMap[parameterPair.TargetName]; ok {
					errMessage = fmt.Sprintf("验证映射配置表时出现错误: 类型'%s'重复定义了映射'%s'", mappingDefinition.Type, parameterPair.TargetName)
					log.Println(errMessage)
					return false, errMessage
				}
				targetNameMap[parameterPair.TargetName] = struct{}{}
			}
		}
	}

	if len(systemMap) == 0 {
		errMessage = "验证映射配置表时出现错误: 没有定义System"
		log.Println(errMessage)
		return false, errMessage
	}
	if len(mediumMap) == 0 {
		errMessage = "验证映射配置表时出现错误: 没有定义Medium"
		log.Println(errMessage)
		return false, errMessage
	}
	if len(partMap) == 0 {
		errMessage = "验证映射配置表时出现错误: 没有定义任何零件类型"
		log.Println(errMessage)
		return false, errMessage
	}

	return true, errMessage
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
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	System      string  `json:"system"`
	Medium      string  `json:"medium"`
	Parts       []*Part `json:"parts"`
}

type Part struct {
	Kind          string  `json:"kind"`
	Name          string  `json:"name"`
	NewName       string  `json:"newName"`
	ModelicaClass string  `json:"modelica_class"`
	ParameterList []*Pair `json:"parameter_list"`
	PortList      []*Pair `json:"port_list"`
}

type Pair struct {
	SourceName string `json:"source_name"`
	TargetName string `json:"target_name"`
}

// 获取映射配置表管道信息详情
func GetMappingConfigDetails(id, name, description, path string) (res *MappingConfigParseData, err error) {
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
		ID:          id,
		Name:        name,
		Description: description,
		System:      "",
		Medium:      "",
		Parts:       []*Part{},
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
				NewName:       item.Type,
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

// 编辑映射配置表管道信息详情
func EditMappingConfigDetails(path string, requestData *DataType.EditMappingConfigDetailsData, op string) (bool, error) {
	if op != "add" && op != "replace" && op != "remove" {
		log.Println("传入的json-patch操作方法错误，必须是 add replace remove")
		return false, nil
	}

	// 请求数据类型转换为MappingConfigParseData
	item := ConvertMappingConfigStruct(requestData)

	// 读取文件内容
	originalContentByte, err := os.ReadFile(path)
	if err != nil {
		log.Println("读取映射配置表文件时出现错误", err)
		return false, nil
	}
	// 将文件内容映射到MappingConfigData结构体中
	m := MappingConfigData{}
	if err := json.Unmarshal(originalContentByte, &m); err != nil {
		log.Println("将文件内容映射到MappingConfigData结构体中时出现错误", err)
		return false, nil
	}

	// 生成json-patch数据
	var patches []map[string]any

	switch op {
	case "add":
		// 处理添加新参数 add
		hasRepeatedPart, repeatedPartName := false, ""
		if patches, hasRepeatedPart, repeatedPartName = CreateAddJsonPatch(item, &m); hasRepeatedPart {
			return false, errors.New(fmt.Sprintf("添加的类型已存在: %s", repeatedPartName))
		}
	case "replace":
		// 处理更新现有参数 replace
		hasRepeatedPart, repeatedPartName := false, ""
		if patches, hasRepeatedPart, repeatedPartName = CreateReplaceJsonPatch(item, &m); hasRepeatedPart {
			return false, errors.New(fmt.Sprintf("类型名称已存在: %s", repeatedPartName))
		}
	case "remove":
		// 处理删除现有参数 remove
		patches = CreateRemoveJsonPatch(item, &m)
	}

	// 创建补丁对象
	patchesByte, err := json.Marshal(patches)
	fmt.Println(string(patchesByte))
	if err != nil {
		log.Println("json-patch创建补丁对象时出现错误", err)
		return false, nil
	}

	patchObj, err := jsonpatch.DecodePatch(patchesByte)
	if err != nil {
		log.Println("json-patch解析补丁对象时出现错误", err)
		return false, nil
	}

	// 应用补丁
	fmt.Println(string(originalContentByte))
	patchedData, err := patchObj.Apply(originalContentByte)
	if err != nil {
		log.Println("json-patch应用补丁修改源数据时出现错误", err)
		return false, nil
	}

	// 写回映射配置文件
	if ok := fileOperation.WriteFileByte(path, patchedData); !ok {
		log.Println("向映射配置文件中写回数据时出现错误", err)
		return false, nil
	}

	return true, nil
}

func ConvertMappingConfigStruct(item *DataType.EditMappingConfigDetailsData) *MappingConfigParseData {
	mappingConfigParseData := MappingConfigParseData{
		System: item.System,
		Medium: item.Medium,
		Parts:  []*Part{},
	}

	for _, part := range item.Parts {
		newPart := Part{
			Kind:          part.Kind,
			NewName:       part.NewName,
			Name:          part.Name,
			ModelicaClass: part.ModelicaClass,
			ParameterList: []*Pair{},
			PortList:      []*Pair{},
		}

		for _, parameter := range part.ParameterList {
			newParameter := Pair{SourceName: parameter.SourceName, TargetName: parameter.TargetName}
			newPart.ParameterList = append(newPart.ParameterList, &newParameter)
		}

		for _, port := range part.PortList {
			newPort := Pair{SourceName: port.SourceName, TargetName: port.TargetName}
			newPart.PortList = append(newPart.PortList, &newPort)
		}

		mappingConfigParseData.Parts = append(mappingConfigParseData.Parts, &newPart)
	}

	return &mappingConfigParseData
}

func GenSystemInfo(systemClass string) MappingDefinition {
	systemInfo := MappingDefinition{
		Kind:   "System",
		Type:   "系统",
		Usages: &UsagePipe{Default: &SystemModel{ModelicaClass: systemClass}},
	}

	return systemInfo
}

func GenMediumInfo(mediumClass string) MappingDefinition {
	systemInfo := MappingDefinition{
		Kind:   "Medium",
		Type:   "介质",
		Usages: &UsagePipe{Default: &SystemModel{ModelicaClass: mediumClass}},
	}

	return systemInfo
}

func GenPartInfo(requestPartInfo *Part) MappingDefinition {
	partInfo := MappingDefinition{
		Kind:   requestPartInfo.Kind,
		Type:   requestPartInfo.NewName,
		Usages: &UsagePipe{PipeModel: &PipeModel{ModelicaClass: requestPartInfo.ModelicaClass}},
	}

	// 获取零件参数映射信息
	partInfo.Usages.PipeModel.Parameters = []*MappingPair{}
	for _, newParameter := range requestPartInfo.ParameterList {
		newPair := MappingPair{
			SourceName: newParameter.SourceName,
			TargetName: newParameter.TargetName,
		}
		partInfo.Usages.PipeModel.Parameters = append(partInfo.Usages.PipeModel.Parameters, &newPair)
	}

	// 获取零件端点映射信息
	partInfo.Usages.PipeModel.Ports = []*MappingPair{}
	for _, newParameter := range requestPartInfo.PortList {
		newPair := MappingPair{
			SourceName: newParameter.SourceName,
			TargetName: newParameter.TargetName,
		}
		partInfo.Usages.PipeModel.Ports = append(partInfo.Usages.PipeModel.Ports, &newPair)
	}

	return partInfo
}

func CreateAddJsonPatch(item *MappingConfigParseData, m *MappingConfigData) (patches []map[string]any, hasRepeatedPart bool, repeatedPartName string) {
	patches = []map[string]any{}

	var systemAlreadyExist bool
	var mediumAlreadyExist bool
	for _, mappingDefinition := range m.MappingDefinitions {
		if mappingDefinition.Kind == "System" {
			systemAlreadyExist = true
		}

		if mappingDefinition.Kind == "Medium" {
			mediumAlreadyExist = true
		}
	}

	if !systemAlreadyExist {
		systemInfo := GenSystemInfo(item.System)
		// 生成json-patch格式的数据
		patch := map[string]any{
			"op":    "add",
			"path":  fmt.Sprintf("%s%d", "/mappingDefinitions/", 0),
			"value": systemInfo,
		}
		patches = append(patches, patch)
	}

	if !mediumAlreadyExist {
		mediumInfo := GenMediumInfo(item.Medium)
		// 生成json-patch格式的数据
		patch := map[string]any{
			"op":    "add",
			"path":  fmt.Sprintf("%s%d", "/mappingDefinitions/", 0),
			"value": mediumInfo,
		}
		patches = append(patches, patch)
	}

	for _, part := range item.Parts {
		for _, mappingDefinition := range m.MappingDefinitions {
			if part.Kind == mappingDefinition.Kind && part.NewName == mappingDefinition.Type {
				return patches, true, part.NewName
			}
		}

		// 生成零件的数据
		partInfo := GenPartInfo(part)

		// 生成json-patch格式的数据
		patch := map[string]any{
			"op":    "add",
			"path":  fmt.Sprintf("%s%d", "/mappingDefinitions/", 0),
			"value": partInfo,
		}

		patches = append(patches, patch)
	}

	return patches, false, ""
}

func CreateReplaceJsonPatch(item *MappingConfigParseData, m *MappingConfigData) (patches []map[string]any, hasRepeatedPart bool, repeatedPartName string) {
	patches = []map[string]any{}

	// 判断零件类型的新名字是否和已有的零件类型重复
	for _, part := range item.Parts {
		if part.NewName != part.Name {
			for _, mappingDefinition := range m.MappingDefinitions {
				if part.Kind == mappingDefinition.Kind && part.NewName == mappingDefinition.Type {
					return patches, true, part.NewName
				}
			}
		}
	}

	for index, mappingDefinition := range m.MappingDefinitions {
		// 创建系统信息补丁
		if mappingDefinition.Kind == "System" && item.System != "" {
			// 生成系统的数据
			systemInfo := GenSystemInfo(item.System)
			patch := map[string]any{
				"op":    "replace",
				"path":  fmt.Sprintf("%s%d", "/mappingDefinitions/", index),
				"value": systemInfo,
			}

			patches = append(patches, patch)
		}

		// 创建介质信息补丁
		if mappingDefinition.Kind == "Medium" && item.Medium != "" {
			// 生成介质的数据
			mediumInfo := GenMediumInfo(item.Medium)
			patch := map[string]any{
				"op":    "replace",
				"path":  fmt.Sprintf("%s%d", "/mappingDefinitions/", index),
				"value": mediumInfo,
			}

			patches = append(patches, patch)
		}

		// 创建零件信息补丁
		for _, part := range item.Parts {
			if part.Kind == mappingDefinition.Kind && part.Name == mappingDefinition.Type {
				partInfo := GenPartInfo(part)
				// 生成json-patch格式的数据
				patch := map[string]any{
					"op":    "replace",
					"path":  fmt.Sprintf("%s%d", "/mappingDefinitions/", index),
					"value": partInfo,
				}

				patches = append(patches, patch)
			}
		}
	}

	return patches, false, ""
}

func CreateRemoveJsonPatch(item *MappingConfigParseData, m *MappingConfigData) (patches []map[string]any) {
	patches = []map[string]any{}
	for index, mappingDefinition := range m.MappingDefinitions {
		// 创建系统信息补丁
		if mappingDefinition.Kind == "System" && item.System != "" {
			patch := map[string]any{
				"op":   "remove",
				"path": fmt.Sprintf("%s%d", "/mappingDefinitions/", index),
			}

			patches = append(patches, patch)
		}

		// 创建介质信息补丁
		if mappingDefinition.Kind == "Medium" && item.Medium != "" {
			patch := map[string]any{
				"op":   "remove",
				"path": fmt.Sprintf("%s%d", "/mappingDefinitions/", index),
			}

			patches = append(patches, patch)
		}

		// 创建零件信息补丁
		for _, part := range item.Parts {
			if part.Kind == mappingDefinition.Kind && part.Name == mappingDefinition.Type {
				patch := map[string]any{
					"op":   "remove",
					"path": fmt.Sprintf("%s%d", "/mappingDefinitions/", index),
				}

				patches = append(patches, patch)
			}
		}
	}

	return patches
}

// 应用补丁
func ApplyJsonPatch(patchesByte, originalContentByte []byte) []byte {
	patchObj, err := jsonpatch.DecodePatch([]byte(patchesByte))
	if err != nil {
		log.Println("json-patch解析补丁对象时出现错误", err)
		return nil
	}

	// 应用补丁
	patchedData, err := patchObj.Apply(originalContentByte)
	if err != nil {
		log.Println("json-patch应用补丁修改源数据时出现错误", err)
		return nil
	}

	return patchedData
}

// 获取映射配置表管道信息详情
func GetInstanceMapping(pipeNetInfoFileId, mappingConfigId, pipeNetInfoFilePath, mappingConfigPath string) (res map[string]*Root, err error) {

	// 解析管网信息文件
	pipeNetXml1, err := ParseInfoFileXml(pipeNetInfoFilePath)
	var pipeNetXml2 Root
	data, _ := json.Marshal(&pipeNetXml1)
	json.Unmarshal(data, &pipeNetXml2)

	// 将管网信息文件第二棵树的Name替换为LegalName
	for i := 0; i < len(pipeNetXml2.Components); i++ {
		pipeNetXml2.Components[i].Name = pipeNetXml2.Components[i].LegalName
	}

	// 解析映射配置表
	mappingConfig, err := GetMappingConfigDetails("", "", "", mappingConfigPath)
	haha := map[string]*Part{}
	for _, part := range mappingConfig.Parts {
		haha[part.Name] = part
	}

	logRedisKey := "pipenet" + "_" + mappingConfigId + "_" + pipeNetInfoFileId
	// 删除之前的日志
	config.R.Del(context.Background(), logRedisKey)
	config.R.RPush(context.Background(), logRedisKey, time.Now().Format("2006-01-02 15:04:05"))

	// 开始进行匹配
	for i := 0; i < len(pipeNetXml2.Components); i++ {
		var found bool
		if part, ok := haha[pipeNetXml2.Components[i].TypeCAD]; ok {
			pipeNetXml2.Components[i].TypeCAE = part.ModelicaClass

			// 获取并生成angle向量参数
			if pipeNetXml2.Components[i].TypeCAD == "弯管" {
				angleList := []string{}
				for j := 0; j < len(pipeNetXml2.Components[i].Properties); j++ {
					if pipeNetXml2.Components[i].Properties[j].Angle != "" {
						angleList = append(angleList, pipeNetXml2.Components[i].Properties[j].Angle)
					}
				}

				angleListLen := len(angleList)
				if angleListLen > 0 {
					angleStr := "{"
					for j, eachAngle := range angleList {
						if j != angleListLen-1 {
							angleStr = angleStr + eachAngle + ", "
						} else {
							angleStr = angleStr + eachAngle
						}
					}
					angleStr = angleStr + "}"
					angleParameterCAD := Parameter{Name: "角度", Value: angleStr, Unit: "m", Id: "1"}
					angleParameterCAE := Parameter{Name: "angles", Value: angleStr, Unit: "m", Id: "1"}
					pipeNetXml1.Components[i].Parameters = append(pipeNetXml1.Components[i].Parameters, angleParameterCAD)
					pipeNetXml2.Components[i].Parameters = append(pipeNetXml2.Components[i].Parameters, angleParameterCAE)
				}
			}

			// 生成常规参数
			for j := 0; j < len(pipeNetXml2.Components[i].Parameters); j++ {
				var foundPamameter bool
				for _, partInfo := range part.ParameterList {
					if pipeNetXml2.Components[i].Parameters[j].Name == partInfo.SourceName {
						pipeNetXml2.Components[i].Parameters[j].Name = partInfo.TargetName
						// 生成length向量参数
						if pipeNetXml2.Components[i].Parameters[j].Name == "length" && pipeNetXml2.Components[i].TypeCAD == "弯管" {
							pipeNetXml2.Components[i].Parameters[j].Value = "{" + pipeNetXml2.Components[i].Parameters[j].Value + "}"
						}
						foundPamameter = true
						break
					}
				}

				if !foundPamameter && pipeNetXml2.Components[i].Parameters[j].Name != "angles" {
					info := fmt.Sprintf("映射配置表中没有找到参数信息：零件: %s CAD类型: %s CAD参数: %s CAE类型: %s CAE参数: 缺失",
						pipeNetXml2.Components[i].Name, pipeNetXml2.Components[i].TypeCAD, pipeNetXml2.Components[i].Parameters[j].Name, pipeNetXml2.Components[i].TypeCAE)
					config.R.RPush(context.Background(), logRedisKey, info)
					pipeNetXml2.Components[i].Parameters[j].Name = ""
				}

			}
			found = true
		}
		if !found {
			info := fmt.Sprintf("映射配置表中没有找到零件类型信息: 零件: %s CAD类型: %s CAE类型: 缺失", pipeNetXml2.Components[i].InstanceName, pipeNetXml2.Components[i].TypeCAD)
			config.R.RPush(context.Background(), logRedisKey, info)
			pipeNetXml2.Components[i].TypeCAE = ""
			for j := 0; j < len(pipeNetXml2.Components[i].Parameters); j++ {
				pipeNetXml2.Components[i].Parameters[j].Name = ""
			}
		}
	}

	res = map[string]*Root{"file_tree": &pipeNetXml1, "mapping_tree": &pipeNetXml2}
	return res, nil
}

// 获取映射配置表管道信息详情
func GetInstanceMappingLog(mappingConfigId, pipeNetInfoFileId string) []string {
	logRedisKey := "pipenet" + "_" + mappingConfigId + "_" + pipeNetInfoFileId
	logs, _ := config.R.LRange(context.Background(), logRedisKey, 0, -1).Result()
	return logs
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

// 生成管网模型源码
func WritePipeNetModeCode(modeName, modeNameAll, medium, packageName, packageFilePath string, instanceMapping *Root) {
	// 向模型中写入Mdedium代码
	oldCode := serviceV1.GetModelCode(packageName)
	modelStr := "model " + modeName + "\n" + "replaceable package Medium = " + medium + ";\n" + "end " + modeName + ";"
	newCodeStr := ""
	if modeName != modeNameAll {
		newCodeStr = strings.ReplaceAll(oldCode, "model "+modeName+"\n  end "+modeName+";", modelStr)
	} else {
		newCodeStr = strings.ReplaceAll(oldCode, "model "+modeName+"\nend "+modeName+";", modelStr)
	}
	parseResult, ok := serviceV1.ParseCodeString(newCodeStr, packageFilePath)
	if ok && len(parseResult) > 0 {
		loadResult := serviceV1.LoadCodeString(newCodeStr, packageFilePath)
		if loadResult {
			serviceV1.ModelSave(parseResult)
		}
	}

	// 向模型中写入组件代码
	for _, component := range instanceMapping.Components {
		rotation := strconv.Itoa(0)
		data := serviceV1.GetIconNew(component.TypeCAE, component.LegalName, false)
		graphics := data["graphics"].(map[string]any)
		if graphics == nil {
			continue
		}
		graphics["originDiagram"] = "0, 0"
		graphics["original_name"] = component.LegalName
		graphics["name"] = component.LegalName
		graphics["type"] = "Transformation"
		graphics["ID"] = "0"
		graphics["rotateAngle"] = graphics["rotation"]
		extentDiagram := serviceV1.GetModelExtentToString(graphics["coordinate_system"])
		data["graphics"] = graphics
		result, msg := serviceV1.AddComponent(component.LegalName, component.TypeCAE, modeNameAll, "0, 0", rotation, extentDiagram)
		if !result {
			fmt.Println(msg)
			break
		} else {
			serviceV1.SetPackageUses(component.TypeCAE, modeNameAll)
			serviceV1.ModelSave(modeNameAll)
		}

		// 向组件中写入参数
		for _, parameter := range component.Parameters {
			result = serviceV1.SetElementModifierValue(modeNameAll, component.LegalName+"."+parameter.Name, parameter.Value)
			if !result {
				fmt.Printf("向组件中写入参数失败: %s %s %s\n", modeNameAll, parameter.Name, parameter.Value)
			}
		}
		result = serviceV1.SetElementModifierValue(modeNameAll, component.LegalName+".Medium", "redeclare package Medium = Medium")
		if !result {
			fmt.Printf("向组件中写入Medium参数失败\n")
		}
		serviceV1.ModelSave(modeNameAll)
	}

	// 向模型中写入连线代码
	portNameMapping := map[string]string{
		"Point1": "port_a",
		"Point2": "port_b",
		"Point3": "port_c",
	}
	for _, connector := range instanceMapping.Connectors {
		result := serviceV1.AddConnection(
			modeNameAll,
			connector.From.LegalName+"."+portNameMapping[connector.From.Point],
			connector.To.LegalName+"."+portNameMapping[connector.To.Point],
			"0,0,127",
			[]string{},
		)
		if !result {
			fmt.Printf("添加组件连线失败: %s %s %s\n", modeNameAll, connector.From.LegalName, connector.To.LegalName)
		}
	}
	serviceV1.ModelSave(modeNameAll)
}

// 生成管网模型源码, 代码拼接版
func WritePipeNetModeCodeNew(modeName, modeNameAll, system, medium string, packageRecord *DataBaseModel.YssimModels, packageFilePath string, instanceMapping *Root) bool {
	// 全局Medium代码
	modelicaCode := "model " + modeName + "\n" + "replaceable package Medium = " + medium + ";\n"

	// system组件代码
	if system != "" {
		modelicaCode = modelicaCode + "inner " + system + " system " + "annotation(Placement(visible = true, transformation(origin = {0, 0}, extent = {{-10, -10}, {10, 10}}, rotation = 0)));\n"
	}

	// 组件代码
	componentTypeCADMap := map[string]string{}
	for _, component := range instanceMapping.Components {
		componentTypeCADMap[component.LegalName] = component.TypeCAD
		modelicaCode = modelicaCode + component.TypeCAE + " " + component.LegalName + "("
		for _, parameter := range component.Parameters {
			if parameter.Name != "" && parameter.Value != "" {
				modelicaCode = modelicaCode + parameter.Name + "=" + parameter.Value + ", "
			}
		}
		modelicaCode = modelicaCode + "redeclare package Medium = Medium)" + " \"" + component.InstanceName + "\" " + "annotation(Placement(visible = true, transformation(origin = {0, 0}, extent = {{-10, -10}, {10, 10}}, rotation = 0)));\n"
	}

	// equation代码
	modelicaCode = modelicaCode + "equation\n"

	// 连线代码
	portNameMapping := map[string]string{
		"Point1": "port_a",
		"Point2": "port_b",
		"Point3": "port_c",
	}

	// 加油控制活门端口
	/*closedVolumePortNameMapping := map[string]string{
		"Point1": "ports[1]",
		"Point2": "ports[2]",
		"Point3": "ports[3]",
		"Point4": "ports[4]",
	}*/

	for _, connector := range instanceMapping.Connectors {
		// 对普通管的连线端点进行处理
		fromPoint := connector.From.LegalName + "." + portNameMapping[connector.From.Point]
		toPoint := connector.To.LegalName + "." + portNameMapping[connector.To.Point]

		// 对加油控制活门的连线端点进行特殊处理
		/*if componentTypeCADMap[connector.From.LegalName] == "加油控制活门" {
			fromPoint = connector.From.LegalName + "." + closedVolumePortNameMapping[connector.From.Point]
		}

		if componentTypeCADMap[connector.To.LegalName] == "加油控制活门" {
			toPoint = connector.To.LegalName + "." + closedVolumePortNameMapping[connector.To.Point]
		}*/

		// 对弯管的连线端点进行特殊处理
		if componentTypeCADMap[connector.From.LegalName] == "弯管" {
			if connector.From.Point == "Point1" {
				fromPoint = connector.From.LegalName + "." + portNameMapping["Point1"]
			} else {
				fromPoint = connector.From.LegalName + "." + portNameMapping["Point2"]
			}
		}

		if componentTypeCADMap[connector.To.LegalName] == "弯管" {
			if connector.To.Point == "Point1" {
				toPoint = connector.To.LegalName + "." + portNameMapping["Point1"]
			} else {
				toPoint = connector.To.LegalName + "." + portNameMapping["Point2"]
			}
		}

		modelicaCode = modelicaCode + "connect(" + fromPoint + ", " + toPoint + ") annotation(Line(color = {0, 0, 127}));\n"
	}

	modelicaCode = modelicaCode + "annotation(uses(Modelica(version = \"4.0.0\")));\n" + "end " + modeName + ";"

	// 保存modelica源码
	if packageRecord.ID == "" {
		return false
	}
	nameList := strings.Split(modeNameAll, ".")
	if len(nameList) > 1 {
		parentName := strings.Join(nameList[:len(nameList)-1], ".")
		modelicaCode = "within " + parentName + ";" + modelicaCode
	}
	// oldCode := service.GetModelCode(item.ModelName)
	modelPath := packageRecord.FilePath
	if strings.HasSuffix(packageRecord.FilePath, "/package.mo") {
		modelPath = service.GetSourceFile(modeNameAll)
	}

	parseResult, ok := service.ParseCodeString(modelicaCode, modelPath)
	if ok && len(parseResult) > 0 {
		isExist := service.IsExistPackage(parseResult)
		if isExist && (modeNameAll != parseResult) {
			log.Println("模型名称重复")
			return false
		}
		loadResult := service.LoadCodeString(modelicaCode, modelPath)
		if loadResult {
			if parseResult != modeNameAll {
				// 判断是否是子模型
				if !strings.Contains(modeNameAll, ".") {
					DB.Model(DataBaseModel.YssimModels{}).Where("id = ?", packageRecord.ID).Update("package_name", parseResult)

					packageInformation := service.GetPackageInformation()
					packageInformationJson, _ := sonic.Marshal(packageInformation)
					DB.Model(DataBaseModel.YssimUserSpace{}).Where("id = ? AND username = ?", packageRecord.UserSpaceId, packageRecord.SysUser).Update("package_information", packageInformationJson)
				}
				service.DeleteLibrary(modeNameAll)
			}
			service.ModelSave(parseResult)
			log.Println("模型保存成功")
			return true
		}
	}

	log.Println("语法错误，请重新检查")
	return false
}
