package service

import (
	"strconv"
	"strings"
	"yssim-go/library/omc"
)

//var ParameterTranslation = config.ParameterTranslation

type modelParameters struct {
	name                 string
	componentClassName   string
	packageName          string
	className            string
	modelName            string
	classAll             []string
	components           [][]interface{}
	componentAnnotations [][]interface{}
	componentsDict       map[string]interface{}
}

func (m modelParameters) getParameterValue(name string) string {
	data := ""
	for i := 0; i < len(m.classAll); i++ {
		data = omc.OMC.GetParameterValue(m.classAll[i], name)
		if data != "" {
			data = strings.ReplaceAll(data, "\"", "")
			data = strings.ReplaceAll(data, "\n", "")
			return data
		}
	}

	return data
}

func (m modelParameters) getElementModifierFixedValue(name string) string {
	data := omc.OMC.GetElementModifierValue(m.modelName, name)
	return data
}

func getDerivedClassModifierValueALL(className string) string {

	classAll := GetICList(className)
	var DerivedClassModifierValue string
	for p := 0; p < len(classAll); p++ {
		data := omc.OMC.GetDerivedClassModifierValue(classAll[p], "unit")
		if data != "" {
			return data
		}
	}
	return DerivedClassModifierValue
}

func getUnit(modelName, componentClassName, varName string) []string {
	modelNameAll := []string{modelName}
	for len(modelNameAll) > 0 {
		for _, name := range modelNameAll {
			value := omc.OMC.GetElementModifierValue(name, varName+"."+"unit")
			if value != "" {
				return []string{value}
			}
		}
		modelNameAll = omc.OMC.GetInheritedClassesList(modelNameAll)
	}

	return []string{getDerivedClassModifierValueALL(componentClassName)}
}

func (m modelParameters) getElementModifierStartValue(name string, showStartAttribute bool) string {
	data := ""
	if showStartAttribute {
		data = omc.OMC.GetElementModifierValue(m.modelName, name)
	} else {
		for i := 0; i < len(m.classAll); i++ {
			data = omc.OMC.GetElementModifierValue(m.classAll[i], name)
			if data != "" {
				break
			}
		}
	}
	data = strings.ReplaceAll(data, "\"", "")
	data = strings.ReplaceAll(data, "\n", "")
	return data
}

func (m modelParameters) getExtendsModifierNameAndValue() ([]string, []string, []string) {
	var dataNameList []string
	var dataValueList []string
	var dataFinalList []string
	n := 0
	f := n + 1
	for i := 0; i < len(m.classAll)-1; i++ {
		nameData := omc.OMC.GetExtendsModifierNames(m.classAll[n], m.classAll[f])
		if len(nameData) > 0 {
			dataNameList = append(dataNameList, nameData...)
			valueData := omc.OMC.GetExtendsModifierValue(m.classAll[n], m.classAll[f], nameData[0])
			dataValueList = append(dataValueList, valueData)
			finalData := omc.OMC.IsExtendsModifierFinal(m.classAll[n], m.classAll[f], strings.Split(nameData[0], ".")[0])
			dataFinalList = append(dataFinalList, finalData)
		}
		f += 1
		if f == len(m.classAll) {
			n += 1
			f = n + 1
		}
	}
	return dataNameList, dataValueList, dataFinalList
}

func GetModelParameters(modelName, componentName, componentClassName string) []interface{} {
	var m modelParameters
	var dataList []interface{}
	m.name = componentName
	m.componentClassName = componentClassName
	m.modelName = modelName
	m.classAll = omc.OMC.GetInheritedClassesListAll([]string{modelName})
	for i := 1; i < len(m.classAll); i++ {
		elementsData := omc.OMC.GetElements(m.classAll[i])
		annotationsData := omc.OMC.GetElementAnnotations(m.classAll[i])
		for n, d := range elementsData {
			if d.([]interface{})[3] == componentName && (len(annotationsData[n].([]interface{})) != 0 || annotationsData[n].([]interface{})[0].(string) == "Placement") {
				m.modelName = m.classAll[i]
				inheritedClassAll := omc.OMC.GetInheritedClassesListAll([]string{d.([]interface{})[2].(string)})
				m.classAll = inheritedClassAll
				break
			}
		}
	}
	if m.modelName == modelName {
		m.classAll = omc.OMC.GetInheritedClassesListAll([]string{m.componentClassName})
	}
	for _, c := range m.classAll {
		data := omc.OMC.GetElements(c)
		for _, d := range data {
			dd := append(d.([]interface{}), c)
			m.components = append(m.components, dd)
		}
	}
	m.componentAnnotations = omc.OMC.GetElementAnnotationsList(m.classAll)
	if len(m.components) == 0 {
		return dataList
	}
	m.componentsDict = map[string]interface{}{}
	for i := 0; i < len(m.components); i++ {
		m.components[i] = append(m.components[i], m.componentAnnotations[i])
		m.componentsDict[m.components[i][3].(string)] = m.components[i]
	}
	for i := 0; i < len(m.components); i++ {
		dataDefault := map[string]interface{}{"tab": "General", "type": "Normal", "group": ""}
		p := m.componentsDict[m.components[i][3].(string)].([]interface{})

		varName := p[3].(string)
		if p[2] != "-" {
			m.className = p[2].(string)
		} else {
			m.className = ""
		}
		IsExtendsModifierFinal := omc.OMC.IsExtendsModifierFinal(componentClassName, p[len(p)-2].(string), varName)
		if p[5] == "protected" || IsExtendsModifierFinal == "true" || p[6] == "true" {
			continue
		}
		dataDefault["name"] = varName
		dataDefault["comment"] = p[4].(string)

		dList := p[len(p)-1].([]interface{})

		DialogIndex, DialogIndexOk := func() (int, bool) {
			for n := 0; n < len(dList); n++ {
				if dList[n] == "Dialog" {
					return n, true
				}
			}
			return 0, false
		}()
		showStartAttribute := ""
		if DialogIndexOk {
			tabIndex := DialogIndex + 1
			if len(dList) <= 1 || dList[tabIndex].([]interface{})[len(dList[tabIndex].([]interface{}))-1] == "true" {
				continue
			}
			tab := dList[tabIndex].([]interface{})[0]
			group := dList[tabIndex].([]interface{})[1]
			dataDefault["tab"] = tab.(string)
			dataDefault["group"] = group.(string)
			showStartAttribute = dList[tabIndex].([]interface{})[3].(string)
		}
		if showStartAttribute == "true" {
			fixedValueString := m.getElementModifierFixedValue(m.name + "." + varName + ".fixed")
			startValueString := m.getElementModifierFixedValue(m.name + "." + varName + ".start")
			value, _ := strconv.ParseBool(fixedValueString)
			var fixedValueBool interface{}
			if fixedValueString == "" {
				fixedValueBool = ""
			} else {
				fixedValueBool = value
			}
			fixed := map[string]interface{}{
				"type":         "fixed",
				"name":         varName + ".fixed",
				"comment":      dataDefault["comment"],
				"tab":          dataDefault["tab"],
				"group":        dataDefault["group"],
				"defaultvalue": "",
				"value":        fixedValueBool,
				"unit":         getUnit(componentClassName, m.className, varName),
			}
			start := map[string]interface{}{
				"type":         "Normal",
				"name":         varName + ".start",
				"comment":      dataDefault["comment"],
				"tab":          dataDefault["tab"],
				"group":        dataDefault["group"],
				"defaultvalue": "",
				"value":        startValueString,
				"unit":         getUnit(componentClassName, m.className, varName),
			}
			dataList = append(dataList, fixed)
			dataList = append(dataList, start)
			continue
		}

		if p[9] == "true" {
			dataDefault["type"] = "Enumeration"
			dataDefault["defaultvalue"] = p[2]
			dataDefault["disable"] = true
			dataDefault["group"] = "参数"
			oData := make([]string, 1)
			if p[13].(string) != "$Any" {
				// 模板参数获取有内存泄露问题， 暂时不用
				//options := omc.OMC.GetAllSubtypeOf(p[13].(string), componentClassName)
				//for _, option := range options {
				//	optionData := "redeclare "+ option.(string) + " " + componentName+"."+dataDefault["name"].(string)
				//	oData = append(oData, optionData)
				//}
				dataDefault["disable"] = false
			}
			dataDefault["options"] = oData
			dataList = append(dataList, dataDefault)
			continue
		}
		componentModifierValue := omc.OMC.GetElementModifierValue(m.modelName, componentName+"."+dataDefault["name"].(string))

		dataDefault["value"] = componentModifierValue
		if p[10] == "parameter" || componentModifierValue != "" || dataDefault["group"] != "" {
			if dataDefault["group"] == "" || dataDefault["group"] == "Parameters" {
				dataDefault["group"] = "参数"
			}
			isEnumeration := omc.OMC.IsEnumeration(m.className)
			if isEnumeration {
				Literals := omc.OMC.GetEnumerationLiterals(m.className)
				dataDefault["options"] = func() []string {
					var oData []string
					for _, literal := range Literals {
						if literal != "" {
							oData = append(oData, strings.TrimPrefix(m.className, ".")+"."+literal)
						}
					}
					return oData
				}()
				dataDefault["type"] = "Enumeration"
			}
			parameterValue := omc.OMC.GetParameterValue(p[len(p)-2].(string), dataDefault["name"].(string))
			dataDefault["defaultvalue"] = parameterValue
			if p[2] == "Boolean" {
				dataDefault["type"] = "CheckBox"
				if componentModifierValue == "true" || componentModifierValue == "false" {
					dataDefault["value"] = componentModifierValue
					dataDefault["checked"] = componentModifierValue
					dataDefault["defaultvalue"] = componentModifierValue
				} else {
					if parameterValue == "true" || parameterValue == "false" {
						dataDefault["value"] = parameterValue
						dataDefault["checked"] = parameterValue
						dataDefault["defaultvalue"] = parameterValue
					}
				}
			}
			dataDefault["unit"] = ""
			unit := getUnit(componentClassName, m.className, varName)
			if len(unit) > 0 {
				dataDefault["unit"] = unit
			}
			dataList = append(dataList, dataDefault)
			continue
		}

		fixedValueString := m.getElementModifierFixedValue(m.name + "." + varName + ".fixed")
		startValueString := m.getElementModifierFixedValue(m.name + "." + varName + ".start")
		if fixedValueString == "" && startValueString == "" {
			name := p[len(p)-2].(string)
			fixedValueString = omc.OMC.GetElementModifierValue(name, varName+".fixed")
			startValueString = omc.OMC.GetElementModifierValue(name, varName+".start")
		}
		value, _ := strconv.ParseBool(fixedValueString)
		var fixedValueBool interface{}
		if fixedValueString == "" {
			fixedValueBool = ""
		} else {
			fixedValueBool = value
		}
		if fixedValueString != "" || startValueString != "" {
			fixed := map[string]interface{}{
				"type":         "fixed",
				"name":         varName + ".fixed",
				"comment":      dataDefault["comment"],
				"defaultvalue": "",
				"value":        fixedValueBool,
				"unit":         getUnit(componentClassName, m.className, varName),
				"tab":          "General",
				"group":        "Initialization",
			}
			start := map[string]interface{}{
				"type":         "Normal",
				"name":         varName + ".start",
				"comment":      dataDefault["comment"],
				"defaultvalue": "",
				"value":        startValueString,
				"unit":         getUnit(componentClassName, m.className, varName),
				"tab":          "General",
				"group":        "Initialization",
			}
			dataList = append(dataList, fixed)
			dataList = append(dataList, start)
		}
	}
	return dataList
}

func SetComponentModifierValue(className string, parameterValue map[string]string) bool {
	for k, v := range parameterValue {
		result := omc.OMC.SetComponentModifierValue(className, k, v)

		if result != "Ok" {
			return false
		}
	}
	return true
}
