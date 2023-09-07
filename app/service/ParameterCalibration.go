package service

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"yssim-go/library/fileOperation"
	"yssim-go/library/omc"
	"yssim-go/library/xmlOperation"
)

func GetVariableParameter(path, parent string) []map[string]any {
	var result []map[string]any
	result = ratedConditionParameterResultTree(path, parent)

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

func ratedConditionParameterResultTree(path, parent string) []map[string]any {
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
				switch scalarVariableMap[name].IsValueChangeable {
				case scalarVariableMap[name].HideResult == "false" && scalarVariableMap[name].IsProtected:
					dataList = append(dataList, getRatedConditionParameter(splitNameList, scalarVariableMap[name], id, nameMap))
				case scalarVariableMap[name].HideResult == "" && !scalarVariableMap[name].IsProtected:
					dataList = append(dataList, getRatedConditionParameter(splitNameList, scalarVariableMap[name], id, nameMap))
				}
			}
		}
	}
	return dataList
}

func getRatedConditionParameter(splitNameList []string, scalarVariableMap scalarVariable, id int, nameMap map[string]bool) map[string]any {
	data := map[string]any{
		"variables":           splitNameList[0],
		"description":         scalarVariableMap.Description,
		"is_value_changeable": scalarVariableMap.IsValueChangeable,
		"has_child":           false,
		"id":                  id,
		"value":               scalarVariableMap.Real.Start,
		"unit":                scalarVariableMap.Real.Unit,
	}
	if len(splitNameList) > 1 {
		data["has_child"] = true
	}
	id += 1
	nameMap[splitNameList[0]] = true

	return data
}

type FormulaAnalysis struct {
	formulaStrList []string
	variableMap    map[string]bool
	variableList   []string
	formulaData    []map[string]string
}

func GetFormulaList(formulaStr string) ([]map[string]string, []string) {
	formulaStr = strings.TrimSpace(formulaStr)
	formulaStrList := strings.Split(formulaStr, "+")
	f := FormulaAnalysis{
		formulaStrList: formulaStrList,
		variableMap:    make(map[string]bool, 0),
		variableList:   make([]string, 0),
		formulaData:    make([]map[string]string, 0),
	}
	if formulaStr == "" {
		return f.formulaData, f.variableList
	}
	f.formulaParse()
	f.getVariableList()
	return f.formulaData, f.variableList
}

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

func (f *FormulaAnalysis) getVariable(variableList []string) {
	for i := 0; i < len(variableList); i++ {
		f.variableMap[variableList[i]] = true
	}
}

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

func GetPackagesSource() map[string]string {
	// 获取当前环境中所有库的版本，和所在文件
	data := map[string]string{}
	loadedLibraries := omc.OMC.GetPackages()
	for _, library := range loadedLibraries {
		fileName := omc.OMC.GetSourceFile(library)
		data[library] = fileName
	}
	return data
}

func CopyPackage(src, dest string) (string, error) {
	packageDir, packageFile := filepath.Split(src)
	err := fileOperation.CopyDir(packageDir, dest)
	if err != nil {
		return "", errors.New(fmt.Sprintf("拷贝package文件夹失败： %s", err))
	}

	return dest + "/" + packageFile, nil
}
