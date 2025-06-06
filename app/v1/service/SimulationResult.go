package service

import (
	"encoding/csv"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"yssim-go/app/DataBaseModel"
	"yssim-go/library/fileOperation"
	"yssim-go/library/omc"
	"yssim-go/library/xmlOperation"

	"github.com/beevik/etree"
)

// mergeArrays 合并两个二维数组
func mergeArrays(arr1, arr2 [][]string) [][]string {
	var merged [][]string

	maxLen := len(arr1)
	if len(arr2) > maxLen {
		maxLen = len(arr2)
	}

	for i := 0; i < maxLen; i++ {
		var subArr []string

		if i < len(arr1) {
			subArr = append(subArr, arr1[i]...)
		} else {
			subArr = append(subArr, make([]string, len(arr1[0]))...)
		}

		if i < len(arr2) {
			subArr = append(subArr, arr2[i]...)
		} else {
			subArr = append(subArr, make([]string, len(arr2[0]))...)
		}

		merged = append(merged, subArr)
	}

	return merged
}

// ReadSimulationResult 原读取mat文件中单个变量的值
func ReadSimulationResult(varNameList []string, path string) ([][]float64, bool) {
	pwd, _ := os.Getwd()
	data, ok := omc.OMC.ReadSimulationResult(varNameList, pwd+"/"+path)
	return data, ok
}

// ReadSimulationResultList 返回给定所有变量的值
func ReadSimulationResultList(varNameList []string, path string) ([][]float64, bool) {
	pwd, _ := os.Getwd()
	data, ok := omc.OMC.ReadSimulationResultList(varNameList, pwd+"/"+path)
	return data, ok
}

// ReadSimulationResultFromGrpc 读取单个mat文件中的单个变量
func ReadSimulationResultFromGrpc(path string, varName string) ([][]float64, bool) {
	var data [][]float64
	replyTime, err := GrpcGetResult(path, "time")
	replyVar, err := GrpcGetResult(path, varName)
	if err != nil {
		fmt.Println("调用grpc服务(GrpcGetResult)出错：", err)
		return nil, false
	}
	if replyVar.Log == "true" {
		reply1Data := make([]float64, len(replyTime.Data))
		for i, v := range replyTime.Data {
			reply1Data[i] = v
		}
		reply2Data := make([]float64, len(replyVar.Data))
		for i, v := range replyVar.Data {
			reply2Data[i] = v
		}
		data = append(data, reply1Data)
		data = append(data, reply2Data)
		return data, true
	} else {
		fmt.Println(replyVar.Log)
		return nil, false
	}
}

// FilterSimulationResult 保存多个记录多个过滤变量到csv文件
func FilterSimulationResult(items map[string][]string, recordDict map[string]DataBaseModel.YssimSimulateRecord, newFileName string) bool {

	var csvData [][]string
	headFlag := 1
	for key, value := range items {
		var result [][]float64
		// 定义csv的第一行
		headFlagName := recordDict[key].AnotherName + "." // 标记
		headRow := []string{headFlagName + "time"}
		for i := 0; i < len(value); i++ {
			headRow = append(headRow, headFlagName+value[i])
		}
		// 获取结果数据
		for i := 0; i < len(value); i++ {
			singleRes, ok := ReadSimulationResult([]string{value[i]}, recordDict[key].SimulateModelResultPath+"result_res.mat")
			if !ok {
				return false
			}
			if i == 0 {
				// 一个结果只插入一列时间
				result = append(result, singleRes[0])
			}
			result = append(result, singleRes[1])
		}

		var csvDataOne [][]string
		csvDataOne = append(csvDataOne, headRow) // 先写入第一行变量名
		for i := 0; i < len(result[0]); i++ {    // 逐行写入
			var fData []string
			for _, s := range result {
				floatToStr := strconv.FormatFloat(s[i], 'f', -1, 64)
				fData = append(fData, floatToStr)
			}
			csvDataOne = append(csvDataOne, fData)
		}
		if headFlag == 1 {
			csvData = csvDataOne
		} else {
			csvData = mergeArrays(csvData, csvDataOne)
		}
		headFlag++
	}
	// 将最后生成的csvData保存为csv文件
	nfs, ok := fileOperation.CreateFile(newFileName)
	if ok {
		defer nfs.Close()
		bom := []byte{0xef, 0xbb, 0xbf} // 写入 UTF-8 BOM 防止中文乱码
		nfs.Write(bom)
		w := csv.NewWriter(nfs)
		w.Comma = ','
		w.UseCRLF = true
		err := w.WriteAll(csvData)
		if err != nil {
			return false
		}
		w.Flush()
		return true
	}

	return false
}

type realType struct {
	XMLName      xml.Name `xml:"Real"`
	Start        string   `xml:"start,attr"`
	Fixed        string   `xml:"fixed,attr"`
	UseNominal   string   `xml:"useNominal,attr"`
	Unit         string   `xml:"unit,attr"`
	DisplayUnit  string   `xml:"displayUnit,attr"`
	DeclaredType string   `xml:"declaredType,attr"` // DeclaredType只对dymola的xml生效
}

type booleanType struct {
	XMLName    xml.Name `xml:"Boolean"`
	Start      string   `xml:"start,attr"`
	Fixed      string   `xml:"fixed,attr"`
	UseNominal string   `xml:"useNominal,attr"`
	Unit       string   `xml:"unit,attr"`
}

type integerType struct {
	XMLName    xml.Name `xml:"Integer"`
	Start      string   `xml:"start,attr"`
	Fixed      string   `xml:"fixed,attr"`
	UseNominal string   `xml:"useNominal,attr"`
	Unit       string   `xml:"unit,attr"`
}

type defaultExperiment struct {
	XMLName        xml.Name `xml:"DefaultExperiment"`
	StartTime      string   `xml:"startTime,attr"`
	StopTime       string   `xml:"stopTime,attr"`
	StepSize       string   `xml:"stepSize,attr"`
	Tolerance      string   `xml:"tolerance,attr"`
	Solver         string   `xml:"solver,attr"`
	OutputFormat   string   `xml:"outputFormat,attr"`
	VariableFilter string   `xml:"variableFilter,attr"`
}

type scalarVariable struct {
	XMLName           xml.Name `xml:"ScalarVariable"`
	Name              string   `xml:"name,attr"`
	ValueReference    string   `xml:"valueReference,attr"`
	Description       string   `xml:"description,attr"`
	Variability       string   `xml:"variability,attr"`
	IsDiscrete        bool     `xml:"isDiscrete,attr"`
	IsValueChangeable bool     `xml:"isValueChangeable,attr"`
	Alias             string   `xml:"alias,attr"`
	ClassIndex        string   `xml:"classIndex,attr"`
	ClassType         string   `xml:"classType,attr"`
	IsProtected       bool     `xml:"isProtected,attr"`
	HideResult        string   `xml:"hideResult,attr"`
	FileName          string   `xml:"fileName,attr"`
	StartLine         string   `xml:"startLine,attr"`
	StartColumn       string   `xml:"startColumn,attr"`
	EndLine           string   `xml:"endLine,attr"`
	EndColumn         string   `xml:"endColumn,attr"`
	FileWritable      string   `xml:"fileWritable,attr"`
	// 用于解析dymola的xml文件中的输入变量
	Causality string      `xml:"causality,attr"`
	Initial   string      `xml:"initial,attr"`
	Real      realType    `xml:"Real,omitempty"`
	Boolean   booleanType `xml:"Boolean,omitempty"`
	Integer   integerType `xml:"Integer,omitempty"`
}

type modelVariables struct {
	XMLName        xml.Name         `xml:"ModelVariables"`
	ScalarVariable []scalarVariable `xml:"ScalarVariable"`
}

type dmRealType struct {
	// 只对dymola的xml生效
	XMLName     xml.Name `xml:"Real"`
	Unit        string   `xml:"unit,attr"`
	DisplayUnit string   `xml:"displayUnit,attr"`
}
type DmType struct {
	// 只对dymola的xml生效
	XMLName xml.Name   `xml:"SimpleType"`
	Name    string     `xml:"name,attr"`
	Real    dmRealType `xml:"Real,omitempty"`
}
type typeDefinitions struct {
	// 只对dymola的xml生效
	XMLName    xml.Name `xml:"TypeDefinitions"`
	SimpleType []DmType `xml:"SimpleType"`
}
type xmlInit struct {
	XMLName           xml.Name          `xml:"fmiModelDescription"`
	ModelVariables    modelVariables    `xml:"ModelVariables"`
	DefaultExperiment defaultExperiment `xml:"DefaultExperiment"`
	TypeDefinitions   typeDefinitions   `xml:"TypeDefinitions"` // TypeDefinitions只对dymola的xml生效
}

var treeCache = map[string]xmlInit{}

func init() {
	go clearTreeCache()
}

func clearTreeCache() {
	for {
		time.Sleep(time.Second * 300) // 每300秒清空一次tree缓存
		treeCache = map[string]xmlInit{}
	}
}
func CheckNodeEmpty(path, parent string) bool {
	res := false
	v, ok := treeCache[path]
	if !ok {
		v = xmlInit{}
		err := xmlOperation.ParseXML(path, &v)
		treeCache[path] = v
		if err != nil {
			log.Println(err)
		}
	}
	parentName := ""
	if parent != "" {
		parentName = parent + "."
	}
	scalarVariableList := v.ModelVariables.ScalarVariable
	scalarVariableMap := make(map[string]scalarVariable, 0)
	for _, variable := range scalarVariableList {
		name := variable.Name
		if strings.HasPrefix(name, parentName) {
			scalarVariableMap[name] = variable
			// omc的xml判断
			if scalarVariableMap[name].HideResult != "true" && scalarVariableMap[name].IsValueChangeable {
				res = true
				break
			}
			// dymola的xml判断
			if scalarVariableMap[name].Causality == "parameter" || scalarVariableMap[name].Initial == "exact" {
				res = true
				break
			}
		}
	}
	return res

}

func SimulationResultTree(path, parent, keyWords string) []map[string]any {
	v, ok := treeCache[path+"result_init.xml"]
	if !ok {
		v = xmlInit{}
		err := xmlOperation.ParseXML(path+"result_init.xml", &v)
		treeCache[path+"result_init.xml"] = v
		if err != nil {
			log.Println(err)
		}
	}
	parentName := ""
	if parent != "" {
		parentName = parent + "."
	}
	scalarVariableList := v.ModelVariables.ScalarVariable
	scalarVariableMap := make(map[string]scalarVariable, 0)
	var dataList []map[string]any
	nameMap := map[string]bool{}
	var resultNameList, splitName []string
	var name, trimPrefixName string
	id := 0
	for _, variable := range scalarVariableList {
		name = variable.Name
		trimPrefixName = strings.TrimPrefix(name, parent+".")
		if strings.HasPrefix(name, parentName) && strings.Contains(strings.ToLower(name), strings.ToLower(keyWords)) {
			scalarVariableMap[name] = variable
			if !strings.HasPrefix(name, "der(") && !strings.HasPrefix(name, "$") {
				splitName = strings.Split(trimPrefixName, ".")
			} else {
				continue
			}
			if !nameMap[splitName[0]] {
				switch {
				case scalarVariableMap[name].HideResult != "true":
					if !scalarVariableMap[name].IsValueChangeable && len(splitName) == 1 {
						resultNameList = append(resultNameList, scalarVariableMap[name].Name)
					}
					dataList = append(dataList, SetResultTree(splitName, scalarVariableMap[name], id, nameMap, path))
				}
			}
		}
	}
	if len(resultNameList) > 0 {
		result, ok := ReadSimulationResultList(resultNameList, path+"result_res.mat")
		if ok {
			indexNameMap := make(map[string]int) // 用来存储datalist中的变量名称和下标
			for dataListIndex, variable := range dataList {
				indexNameMap[variable["variables"].(string)] = dataListIndex
			}
			for index := range resultNameList {
				ordinate := result[index+1]
				ordinateLength := len(ordinate)
				// 避免读取结果失败，ordinate下标越界异常
				if ordinateLength >= 1 {
					trimPrefixName = strings.TrimPrefix(resultNameList[index], parent+".")
					splitName = strings.Split(trimPrefixName, ".")
					dataListIndex, dataOk := indexNameMap[splitName[0]]
					if dataOk {
						dataList[dataListIndex]["start"] = strconv.FormatFloat(ordinate[ordinateLength-1], 'f', -1, 64)
					}
				}
			}
		}
	}
	return dataList
}

func SimulateResultParameters(path, parent, keyWords string) []any {
	v, ok := treeCache[path+"result_init.xml"]
	if !ok {
		v = xmlInit{}
		err := xmlOperation.ParseXML(path+"result_init.xml", &v)
		treeCache[path+"result_init.xml"] = v
		if err != nil {
			log.Println(err)
		}
	}
	parentName := ""
	if parent != "" {
		parentName = parent + "."
	}
	scalarVariableList := v.ModelVariables.ScalarVariable
	scalarVariableMap := make(map[string]scalarVariable, 0)
	var dataList []any
	nameMap := map[string]bool{}
	var splitName []string
	var name, trimPrefixName string
	for _, variable := range scalarVariableList {
		name = variable.Name
		trimPrefixName = strings.TrimPrefix(name, parent+".")
		if strings.HasPrefix(name, parentName) && strings.Contains(strings.ToLower(name), strings.ToLower(keyWords)) {
			scalarVariableMap[name] = variable
			if !strings.HasPrefix(name, "der(") && !strings.HasPrefix(name, "$") {
				splitName = strings.Split(trimPrefixName, ".")
			} else {
				continue
			}
			if !nameMap[splitName[0]] {
				dataList = append(dataList, GetSimulateResultParameters(splitName, scalarVariableMap[name], nameMap))
			}
		}
	}
	return dataList
}

func SetResultTree(splitName []string, scalarVariableMap scalarVariable, id int, nameMap map[string]bool, path string) map[string]any {
	data := map[string]any{
		"variables":           splitName[0],
		"description":         scalarVariableMap.Description,
		"display_unit":        scalarVariableMap.Real.DisplayUnit,
		"has_child":           false,
		"id":                  id,
		"is_value_changeable": scalarVariableMap.IsValueChangeable,
		"unit":                scalarVariableMap.Real.Unit,
	}
	if scalarVariableMap.IsValueChangeable {
		if scalarVariableMap.Boolean.XMLName.Local == "Boolean" {
			data["start"] = scalarVariableMap.Boolean.Start
		} else {
			data["start"] = scalarVariableMap.Real.Start
		}
	}
	if len(splitName) > 1 {
		data["has_child"] = true
		data["unit"] = ""
		data["is_value_changeable"] = false
		data["display_unit"] = ""
		data["start"] = ""
	}
	id += 1
	nameMap[splitName[0]] = true
	return data
}

func GetSimulateResultParameters(splitName []string, scalarVariableMap scalarVariable, nameMap map[string]bool) map[string]any {
	data := map[string]any{}

	if scalarVariableMap.Variability == "continuous" {
		data["name"] = splitName[0] + ".start"
		if scalarVariableMap.Boolean.XMLName.Local == "Boolean" {
			data["value"] = map[string]any{"value": scalarVariableMap.Boolean.Start, "isFixed": scalarVariableMap.Boolean.Fixed}
		} else if scalarVariableMap.Real.XMLName.Local == "Real" {
			data["value"] = map[string]any{"value": scalarVariableMap.Real.Start, "isFixed": scalarVariableMap.Real.Fixed}
		} else if scalarVariableMap.Integer.XMLName.Local == "Integer" {
			data["value"] = map[string]any{"value": scalarVariableMap.Integer.Start, "isFixed": scalarVariableMap.Integer.Fixed}
		}
	} else {
		data["name"] = splitName[0]
		if scalarVariableMap.Boolean.XMLName.Local == "Boolean" {
			data["value"] = scalarVariableMap.Boolean.Start
		} else if scalarVariableMap.Real.XMLName.Local == "Real" {
			data["value"] = scalarVariableMap.Real.Start
		} else if scalarVariableMap.Integer.XMLName.Local == "Integer" {
			data["value"] = scalarVariableMap.Integer.Start
		}
	}

	nameMap[splitName[0]] = true
	return data
}

func SetCpuTimeResultTree() map[string]any {

	data := map[string]any{
		"variables":           "$cpuTime",
		"description":         "",
		"display_unit":        "",
		"has_child":           false,
		"id":                  0,
		"is_value_changeable": false,
		"unit":                "",
		"start":               "",
		"hide_dollar_symbol":  true,
	}

	return data
}

func AppInputTree(compileType, path, parent, keyWords string) []map[string]any {
	var result []map[string]any
	if compileType == "OM" {
		result = SimulationResultTree(path, parent, keyWords)
	}
	if compileType == "DM" {
		result = DymolaSimulationResultTree(path, parent, keyWords)
	}
	var filteredResult []map[string]any
	parentName := ""
	if parent != "" {
		parentName = parent + "."
	}
	for _, variable := range result {
		// 非节点不需要检查非空
		if variable["has_child"] == false {
			if variable["is_value_changeable"] == true {
				filteredResult = append(filteredResult, variable)
			}
		} else { // 如果是节点，判断是不是空节点
			parent_ := parentName + variable["variables"].(string)
			var result_ bool
			result_ = CheckNodeEmpty(path, parent_)
			if result_ {
				filteredResult = append(filteredResult, variable)
			}
		}
	}
	return filteredResult

}

func DymolaSimulationResultTree(path, parent, keyWords string) []map[string]any {
	// 读取xml文件
	v := xmlInit{}
	err := xmlOperation.ParseXML(path, &v)
	if err != nil {
		log.Println(err)
	}
	parentName := ""
	if parent != "" {
		parentName = parent + "."
	}
	// 将所有的单位类型保存为typeDefinitionsMap
	typeDefinitionsList := v.TypeDefinitions.SimpleType
	typeDefinitionsMap := make(map[string]DmType, 0)
	for _, variable := range typeDefinitionsList {
		name := variable.Name
		typeDefinitionsMap[name] = variable
	}

	scalarVariableList := v.ModelVariables.ScalarVariable
	scalarVariableMap := make(map[string]scalarVariable, 0)
	var dataList []map[string]any
	nameMap := map[string]bool{}
	id := 0
	for _, variable := range scalarVariableList {
		name := variable.Name
		var splitName []string
		trimPrefixName := strings.TrimPrefix(name, parent+".")
		if strings.HasPrefix(name, parentName) && strings.Contains(strings.ToLower(name), strings.ToLower(keyWords)) {
			scalarVariableMap[name] = variable
			if !strings.HasPrefix(name, "der(") && !strings.HasPrefix(name, "$") {
				splitName = strings.Split(trimPrefixName, ".")
			} else {
				continue
			}
			// if !nameMap[splitName[0]] && !scalarVariableMap[name].HideResult && !scalarVariableMap[name].IsProtected {
			if !nameMap[splitName[0]] {
				unit := scalarVariableMap[name].Real.Unit
				displayUnit := scalarVariableMap[name].Real.DisplayUnit
				declaredTypeName := scalarVariableMap[name].Real.DeclaredType

				// 如果scalarVariable节点中的单位不存在并且DeclaredType存在，则从typeDefinitionsMap中获取Unit和DisplayUnit
				if unit == "" && declaredTypeName != "" {
					unit = typeDefinitionsMap[declaredTypeName].Real.Unit
					displayUnit = typeDefinitionsMap[declaredTypeName].Real.DisplayUnit
				}
				isValueChangeable := false
				if scalarVariableMap[name].Causality == "parameter" || scalarVariableMap[name].Initial == "exact" {
					isValueChangeable = true
				}
				data := map[string]any{
					"variables":           splitName[0],
					"description":         scalarVariableMap[name].Description,
					"display_unit":        displayUnit,
					"has_child":           false,
					"id":                  id,
					"is_value_changeable": isValueChangeable,
					"unit":                unit,
					"start":               scalarVariableMap[name].Real.Start,
				}
				if scalarVariableMap[name].Boolean.XMLName.Local == "Boolean" {
					data["start"] = scalarVariableMap[name].Boolean.Start
				}
				if len(splitName) > 1 {
					data["has_child"] = true
					data["is_value_changeable"] = false
					data["unit"] = ""
					data["display_unit"] = ""
					data["start"] = ""
				}
				id += 1
				nameMap[splitName[0]] = true
				dataList = append(dataList, data)

			}

		}
	}
	return dataList
}

func FmpySimulationResultTree(modelName, path, parent, keyWords string) []map[string]any {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		if res := omc.OMC.DumpXMLDAE(modelName); res[0].(string) == "true" {
			os.Rename(res[1].(string), path)
		}
	}
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(path); err != nil {
		log.Printf("解析%v失败, Error------>%v", path, err)
		return nil
	}
	id := 0
	var dataList []map[string]any
	nameMap := map[string]bool{}
	variables := doc.SelectElement("dae").SelectElement("variables")
	// 解析orderedVariables节点
	if orderedVariables := variables.SelectElement("orderedVariables"); orderedVariables != nil {
		if variablesList := orderedVariables.SelectElement("variablesList"); variablesList != nil {
			dataList1, id1, nameMap1 := xmlOperation.GetVarXml(variablesList, parent, keyWords, id, nameMap)
			dataList = append(dataList, dataList1...)
			id = id1
			nameMap = nameMap1
		}
	}
	// 解析knownVariables节点
	if knownVariables := variables.SelectElement("knownVariables"); knownVariables != nil {
		if variablesList := knownVariables.SelectElement("variablesList"); variablesList != nil {
			dataList2, id2, nameMap2 := xmlOperation.GetVarXml(variablesList, parent, keyWords, id, nameMap)
			dataList = append(dataList, dataList2...)
			id = id2
			nameMap = nameMap2
		}
	}
	// 解析aliasVariables节点
	if aliasVariables := variables.SelectElement("aliasVariables"); aliasVariables != nil {
		if variablesList := aliasVariables.SelectElement("variablesList"); variablesList != nil {
			dataList3, _, _ := xmlOperation.GetVarXml(variablesList, parent, keyWords, id, nameMap)
			dataList = append(dataList, dataList3...)
		}
	}
	// 获取没有子节点的变量名
	var dataList2 []map[string]any
	var dataNameList []string
	for i := 0; i < len(dataList); i++ {
		if dataList[i]["has_child"] == false {
			if parent == "" {
				dataNameList = append(dataNameList, dataList[i]["variables"].(string))
			} else {
				dataNameList = append(dataNameList, parent+"."+dataList[i]["variables"].(string))
			}

		}
	}
	// 调用grpc判断变量名（list）是否存在值
	GrpcCheckVarExistRes := GrpcCheckVarExist(path, dataNameList)
	// dataList去掉不存在值的元素
	for i := 0; i < len(dataList); i++ {
		if dataList[i]["has_child"] == false {
			if parent == "" {
				if !GrpcCheckVarExistRes[dataList[i]["variables"].(string)] {
					continue
				}
			} else {
				if !GrpcCheckVarExistRes[parent+"."+dataList[i]["variables"].(string)] {
					continue
				}
			}
		}
		dataList2 = append(dataList2, dataList[i])
	}
	return dataList2
}

func AppSimulateResult(appPageId string, varNameList []string) ([]map[string]any, error) {
	var appPageRecord DataBaseModel.AppPage
	var resData []map[string]any
	// 查询appPageId是否存在
	DB.Where("id = ? ", appPageId).First(&appPageRecord)
	var appDataSourceRecord DataBaseModel.AppDataSource
	DB.Where("id = ? ", appPageRecord.DataSourceId).First(&appDataSourceRecord)
	if appPageRecord.ID == "" || appDataSourceRecord.ID == "" {
		return nil, errors.New("not found")
	}

	for i := 0; i < len(varNameList); i++ {
		data, ok := ReadSimulationResultFromGrpc(appDataSourceRecord.CompilePath+"result_res_single.mat", varNameList[i])
		if ok {
			ordinate := data[1]
			abscissa := data[0]
			length := len(ordinate)
			if length > 1000 {
				step := length / 1000
				o := []float64{}
				a := []float64{}
				for s := 0; s < length; s++ {
					index := s * step
					if index > length || len(o) > 999 {
						break
					}
					o = append(o, data[1][index])
					a = append(a, data[0][index])
				}
				ordinate = o
				abscissa = a
			}
			oneData := map[string]any{
				"variable": varNameList[i],
				"abscissa": abscissa,
				"ordinate": ordinate,
			}
			resData = append(resData, oneData)
		}

	}
	return resData, nil

}

func AppReleaseResult(appPageId string) (map[string]any, error) {
	var appPageRecord DataBaseModel.AppPage
	resData := make(map[string]any)
	csvData := make(map[string]any)
	// 查询appPageId是否存在
	DB.Where("id = ? ", appPageId).First(&appPageRecord)
	var appDataSourceRecord DataBaseModel.AppDataSource
	DB.Where("id = ? ", appPageRecord.DataSourceId).First(&appDataSourceRecord)
	if appPageRecord.ID == "" || appDataSourceRecord.ID == "" {
		return nil, errors.New("not found")
	}

	if appPageRecord.MulResultPath == "" && !fileOperation.Exists(appPageRecord.MulResultPath+"release/") {
		return nil, errors.New("not found")
	} else {
		// 获取appPageRecord.MultiSimulationResultsPath下的所有csv文件
		var csvFileNames []string
		err := filepath.Walk(appPageRecord.MulResultPath+"release/", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && strings.HasSuffix(info.Name(), ".csv") {
				csvFileNames = append(csvFileNames, info.Name())
			}
			return nil
		})
		if err != nil {
			fmt.Println(err)
		}
		for i := 0; i < len(csvFileNames); i++ {
			// 读取csv数据
			file, err := os.Open(appPageRecord.MulResultPath + "release/" + csvFileNames[i])
			if err != nil {
				return nil, err
			}
			defer file.Close()
			reader := csv.NewReader(file)
			records, err := reader.ReadAll()
			if err != nil {
				return nil, err
			}
			// 遍历每一列数据
			resultMap := make(map[string]any)
			for n := 0; n < len(records[0]); n++ {
				var column []string
				for _, record := range records {
					column = append(column, record[n])
				}
				var floatArr []float64
				// 将字符串数组转换为 float 数组
				for _, s := range column[1:] {
					if s != "" {
						f, err := strconv.ParseFloat(s, 64)
						if err != nil {
							fmt.Println(err)
						}
						floatArr = append(floatArr, f)
					}
				}
				resultMap[column[0]] = floatArr
			}
			csvDataKey := csvFileNames[i][:len(csvFileNames[i])-4]
			csvData[csvDataKey] = resultMap
		}
	}
	resData["mul_simulate_data"] = csvData
	resData["naming_order"] = appPageRecord.NamingOrder
	return resData, nil

}

func AppPreviewResult(appPageId string) (map[string]any, error) {
	var appPageRecord DataBaseModel.AppPage
	resData := make(map[string]any)
	csvData := make(map[string]any)
	// 查询appPageId是否存在
	DB.Where("id = ? ", appPageId).First(&appPageRecord)
	var appDataSourceRecord DataBaseModel.AppDataSource
	DB.Where("id = ? ", appPageRecord.DataSourceId).First(&appDataSourceRecord)
	if appPageRecord.ID == "" || appDataSourceRecord.ID == "" {
		return nil, errors.New("not found")
	}

	if appPageRecord.MulResultPath == "" && !fileOperation.Exists(appPageRecord.MulResultPath+"preview/") {
		return nil, errors.New("not found")
	} else {
		// 获取appPageRecord.MultiSimulationResultsPath下的所有csv文件
		var csvFileNames []string
		err := filepath.Walk(appPageRecord.MulResultPath+"preview/", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && strings.HasSuffix(info.Name(), ".csv") {
				csvFileNames = append(csvFileNames, info.Name())
			}
			return nil
		})
		if err != nil {
			fmt.Println(err)
		}
		for i := 0; i < len(csvFileNames); i++ {
			// 读取csv数据
			file, err := os.Open(appPageRecord.MulResultPath + "preview/" + csvFileNames[i])
			if err != nil {
				return nil, err
			}
			defer file.Close()
			reader := csv.NewReader(file)
			records, err := reader.ReadAll()
			if err != nil {
				return nil, err
			}
			// 遍历每一列数据
			resultMap := make(map[string]any)
			for n := 0; n < len(records[0]); n++ {
				var column []string
				for _, record := range records {
					column = append(column, record[n])
				}
				var floatArr []float64
				// 将字符串数组转换为 float 数组
				for _, s := range column[1:] {
					if s != "" {
						f, err := strconv.ParseFloat(s, 64)
						if err != nil {
							fmt.Println(err)
						}
						floatArr = append(floatArr, f)
					}
				}
				resultMap[column[0]] = floatArr
			}
			csvDataKey := csvFileNames[i][:len(csvFileNames[i])-4]
			csvData[csvDataKey] = resultMap
		}
	}
	resData["mul_simulate_data"] = csvData
	resData["naming_order"] = appPageRecord.NamingOrder
	return resData, nil

}
