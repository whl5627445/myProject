package service

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"
	"yssim-go/library/fileOperation"
	"yssim-go/library/omc"
	"yssim-go/library/xmlOperation"
)

var variableParameterTreeCache = map[string]xmlInit{}

// 初始化xml文件内容的缓存
func init() {
	go variableParameterCache()
}

// xml文件的缓存每5分钟清理一次
func variableParameterCache() {
	for {
		time.Sleep(time.Second * 300) // 每300秒清空一次tree缓存
		variableParameterTreeCache = map[string]xmlInit{}
	}
}

// GetVariableParameter 获取xml文件内变量的节点
func GetVariableParameter(path, parent string) []map[string]any {
	var result []map[string]any
	var filteredResult []map[string]any
	result = ratedConditionParameterResultTree(path, parent)
	for _, variable := range result {
		// 非节点不需要检查非空
		filteredResult = append(filteredResult, variable)

	}
	return filteredResult

}

// 购构造额定条件参数的结果树
func ratedConditionParameterResultTree(path, parent string) []map[string]any {
	v, ok := variableParameterTreeCache[path]
	if !ok {
		v = xmlInit{}
		err := xmlOperation.ParseXML(path, &v)
		variableParameterTreeCache[path] = v
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
	id := 0
	for _, variable := range scalarVariableList {
		name := variable.Name
		var splitNameList []string
		trimPrefixName := strings.TrimPrefix(name, parent+".")
		if strings.HasPrefix(name, parentName) {
			scalarVariableMap[name] = variable
			if !strings.HasPrefix(name, "der(") && !strings.HasPrefix(name, "$") {
				splitNameList = strings.Split(trimPrefixName, ".")
			} else {
				continue
			}
			if !nameMap[splitNameList[0]] {
				switch {
				case scalarVariableMap[name].IsValueChangeable == true && scalarVariableMap[name].HideResult == "false" && scalarVariableMap[name].IsProtected:
					dataList = append(dataList, getRatedConditionParameter(splitNameList, scalarVariableMap[name], id, nameMap))
				case scalarVariableMap[name].IsValueChangeable == true && scalarVariableMap[name].HideResult == "" && !scalarVariableMap[name].IsProtected:
					dataList = append(dataList, getRatedConditionParameter(splitNameList, scalarVariableMap[name], id, nameMap))
				}
			}
		}
	}
	return dataList
}

// 获取满足条件的额定条件参数节点
func getRatedConditionParameter(splitNameList []string, scalarVariableMap scalarVariable, id int, nameMap map[string]bool) map[string]any {
	data := map[string]any{
		"variables":           splitNameList[0],
		"description":         scalarVariableMap.Description,
		"is_value_changeable": scalarVariableMap.IsValueChangeable,
		"has_child":           false,
		"value":               scalarVariableMap.Real.Start,
		"unit":                scalarVariableMap.Real.Unit,
	}
	if len(splitNameList) > 1 {
		data["has_child"] = true
		data["description"] = ""
		data["is_value_changeable"] = false
		data["unit"] = ""
		data["value"] = ""
	}
	id += 1
	nameMap[splitNameList[0]] = true

	return data
}

type FormulaAnalysis struct {
	formulaStrList []string            // 在字符串中被解析出的完整公式
	variableMap    map[string]bool     // 解析后公式中的变量，存储的是原子单位， 不可再分隔
	variableList   []string            // 从variableMap中取出的变量
	formulaData    []map[string]string // 公式常量与完整公式的map切片 ，map包含"coefficient", "formula"两个字段
}

// GetFormulaList 获取公式数据列表与公式变量列表
func GetFormulaList(formulaStr string) ([]map[string]string, []string, error) {
	if formulaStr == "" {
		return nil, nil, errors.New("解析数据不能为空")
	}
	err := formulaStrVerify(formulaStr)
	if err != nil {
		return nil, nil, err
	}
	formulaStrList := []string{}
	index := strings.Index(formulaStr, "=")
	formulaStrList = append(formulaStrList, formulaStr[:index])
	formulaStrList = append(formulaStrList, strings.Split(formulaStr[index+1:], "+")...)
	f := FormulaAnalysis{
		formulaStrList: formulaStrList,
		variableMap:    make(map[string]bool, 0),
		variableList:   make([]string, 0),
		formulaData:    make([]map[string]string, 0),
	}

	f.formulaParse()
	f.getVariableList()
	return f.formulaData, f.variableList, nil
}

func formulaStrVerify(formulaStr string) error {
	if strings.Count(formulaStr, "=") > 1 {
		return errors.New("发现多个赋值操作，请检查数据后重新录入")
	}
	if strings.Count(formulaStr, "=") == 0 {
		return errors.New("缺少赋值操作，请检查数据后重新录入")
	}
	return nil
}

// 公式解析的入口函数
func (f *FormulaAnalysis) formulaParse() {

	for i := 0; i < len(f.formulaStrList); i++ {
		formulaList := strings.Split(f.formulaStrList[i], "*")
		data := map[string]string{"coefficient": "", "formula": ""}
		data["coefficient"] = formulaList[0]
		if len(formulaList) == 1 {
			data["formula"] = "1"
		} else {
			data["formula"] = strings.Join(formulaList[1:], " * ")
		}
		f.getVariable(formulaList[1:])
		f.formulaData = append(f.formulaData, data)
	}
}

// 将解析后的公式参数放入Map，顺便去重
func (f *FormulaAnalysis) getVariable(variableList []string) {
	for i := 0; i < len(variableList); i++ {
		pow := strings.Index(variableList[i], "^")
		if pow != -1 {
			f.variableMap[variableList[i][:pow]] = true
			continue
		}
		f.variableMap[variableList[i]] = true
	}
}

// 获取公式解析后的参数列表
func (f *FormulaAnalysis) getVariableList() {
	for k, _ := range f.variableMap {
		f.variableList = append(f.variableList, k)
	}
}

// GetCompileDependencies 获取库的版本，和所在文件
func GetCompileDependencies(packageName string) map[string]map[string]string {
	data := map[string]map[string]string{}
	loadedLibraries := omc.OMC.GetUses(packageName)
	for _, library := range loadedLibraries {
		name := library[0]
		version := library[1]
		sourceFile := omc.OMC.GetSourceFile(name)
		if sourceFile != "" {
			libraryVersion := version
			fileName := sourceFile
			data[name] = map[string]string{"version": libraryVersion, "file": fileName}
		}
	}
	return data
}

// GetPackagesSource 获取当前环境中所有库的版本，和所在文件
func GetPackagesSource() map[string]string {
	data := map[string]string{}
	loadedLibraries := omc.OMC.GetPackages()
	for _, library := range loadedLibraries {
		fileName := omc.OMC.GetSourceFile(library)
		data[library] = fileName
	}
	return data
}

// CopyPackage 将package所在文件夹复制到指定位置，返回package加载文件完整路径
func CopyPackage(src, dest string) (string, error) {
	packageDir, packageFile := filepath.Split(src)
	err := fileOperation.CopyDir(packageDir, dest)
	if err != nil {
		return "", errors.New(fmt.Sprintf("拷贝package文件夹失败： %s", err))
	}

	return dest + "/" + packageFile, nil
}
