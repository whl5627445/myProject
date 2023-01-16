package service

import (
	"encoding/csv"
	"encoding/xml"
	"log"
	"os"
	"strconv"
	"strings"
	"yssim-go/library/fileOperation"
	"yssim-go/library/omc"
	"yssim-go/library/xmlOperation"
)

func ReadSimulationResult(varNameList []string, path string) ([][]float64, bool) {
	pwd, _ := os.Getwd()
	data, ok := omc.OMC.ReadSimulationResult(varNameList, pwd+"/"+path)
	return data, ok
}

func FilterSimulationResult(varNameList []string, path, newFileName string) bool {
	result, ok := ReadSimulationResult(varNameList, path)
	if ok {
		var csvData [][]string
		//for _, f := range result {
		//	var fData []string
		//	for _, s := range f {
		//		floatToStr := strconv.FormatFloat(s, 'f', -1, 64)
		//		fData = append(fData, floatToStr)
		//	}
		//	csvData = append(csvData, fData)
		//}
		for i := 0; i < len(result[0]); i++ {
			var fData []string
			for _, s := range result {
				floatToStr := strconv.FormatFloat(s[i], 'f', -1, 64)
				fData = append(fData, floatToStr)
			}
			csvData = append(csvData, fData)
		}
		nfs, ok := fileOperation.CreateFile(newFileName)
		if ok {
			defer nfs.Close()
			w := csv.NewWriter(nfs)
			w.Comma = ','
			w.UseCRLF = true
			row := append([]string{"time"}, varNameList...)
			err := w.Write(row)
			if err != nil {
				return false
			}
			w.Flush()
			err = w.WriteAll(csvData)
			if err != nil {
				return false
			}
			w.Flush()
			return true
		}
	}
	return false
}

type realType struct {
	XMLName     xml.Name `xml:"Real"`
	Start       string   `xml:"start,attr"`
	Fixed       string   `xml:"fixed,attr"`
	UseNominal  string   `xml:"useNominal,attr"`
	Unit        string   `xml:"unit,attr"`
	DisplayUnit string   `xml:"displayUnit,attr"`
}

type booleanType struct {
	XMLName    xml.Name `xml:"Boolean"`
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
	XMLName           xml.Name    `xml:"ScalarVariable"`
	Name              string      `xml:"name,attr"`
	ValueReference    string      `xml:"valueReference,attr"`
	Description       string      `xml:"description,attr"`
	Variability       string      `xml:"variability,attr"`
	IsDiscrete        bool        `xml:"isDiscrete,attr"`
	Causality         string      `xml:"causality,attr"`
	IsValueChangeable bool        `xml:"isValueChangeable,attr"`
	Alias             string      `xml:"alias,attr"`
	ClassIndex        string      `xml:"classIndex,attr"`
	ClassType         string      `xml:"classType,attr"`
	IsProtected       bool        `xml:"isProtected,attr"`
	HideResult        bool        `xml:"hideResult,attr"`
	FileName          string      `xml:"fileName,attr"`
	StartLine         string      `xml:"startLine,attr"`
	StartColumn       string      `xml:"startColumn,attr"`
	EndLine           string      `xml:"endLine,attr"`
	EndColumn         string      `xml:"endColumn,attr"`
	FileWritable      string      `xml:"fileWritable,attr"`
	Real              realType    `xml:"Real,omitempty"`
	Boolean           booleanType `xml:"Boolean,omitempty"`
}

type modelVariables struct {
	XMLName        xml.Name         `xml:"ModelVariables"`
	ScalarVariable []scalarVariable `xml:"ScalarVariable"`
}

type xmlInit struct {
	XMLName           xml.Name          `xml:"fmiModelDescription"`
	ModelVariables    modelVariables    `xml:"ModelVariables"`
	DefaultExperiment defaultExperiment `xml:"DefaultExperiment"`
}

func SimulationResultTree(path string, parent string) []map[string]interface{} {

	v := xmlInit{}
	err := xmlOperation.ParseXML(path, &v)
	if err != nil {
		log.Println(err)
	}
	parentName := parent
	if parent != "" {
		parentName = parent + "."
	}
	scalarVariableList := v.ModelVariables.ScalarVariable
	var nameList []string
	scalarVariableMap := make(map[string]scalarVariable, 0)
	for _, variable := range scalarVariableList {
		nameList = append(nameList, variable.Name)
		scalarVariableMap[variable.Name] = variable
	}
	var dataList []map[string]interface{}
	nameMap := map[string]bool{}
	id := 0
	for _, name := range nameList {
		trimPrefixName := strings.TrimPrefix(name, parent+".")
		var splitName []string
		if !strings.HasPrefix(name, "der(") && !strings.HasPrefix(name, "$") {
			splitName = strings.Split(trimPrefixName, ".")
		} else {
			//splitName = []string{trimPrefixName}
			continue
		}
		if strings.HasPrefix(name, parentName) && !nameMap[splitName[0]] && scalarVariableMap[name].HideResult == false && scalarVariableMap[name].IsProtected == false {
			data := map[string]interface{}{
				"variables":    splitName[0],
				"description":  scalarVariableMap[name].Description,
				"display_unit": scalarVariableMap[name].Real.DisplayUnit,
				"has_child":    false,
				"id":           id,
				"start":        scalarVariableMap[name].Real.Start,
				"unit":         scalarVariableMap[name].Real.Unit,
			}
			if len(splitName) > 1 {
				data["has_child"] = true
				data["unit"] = ""
				data["display_unit"] = ""
			}
			id += 1
			nameMap[splitName[0]] = true
			dataList = append(dataList, data)
		}
	}
	return dataList
}
