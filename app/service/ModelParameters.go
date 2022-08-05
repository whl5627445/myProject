package service

import (
	"strconv"
	"strings"
	"yssim-go/library/omc"
)

type modelParameters struct {
	name                 string
	componentName        string
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

func (m modelParameters) getComponentModifierFixedValue(name string) string {
	data := omc.OMC.GetComponentModifierValue(m.modelName, name)
	return data
}

func (m modelParameters) getDerivedClassModifierValue() []string {
	DerivedClassModifierNames := omc.OMC.GetDerivedClassModifierNames(m.className)
	var DerivedClassModifierValue []string
	for i := 1; i < len(DerivedClassModifierNames); i++ {
		data := omc.OMC.GetDerivedClassModifierValue(m.className, DerivedClassModifierNames[i].(string))
		DerivedClassModifierValue = append(DerivedClassModifierValue, data)
	}
	return DerivedClassModifierValue
}

func (m modelParameters) getComponentModifierStartValue(name string, showStartAttribute bool) string {
	data := ""
	if showStartAttribute {
		data = omc.OMC.GetComponentModifierValue(m.modelName, name)
	} else {
		for i := 0; i < len(m.classAll); i++ {
			data = omc.OMC.GetComponentModifierValue(m.classAll[i], name)
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

func GetModelParameters(modelName string, name string, componentName string, packageName string) []interface{} {
	var m modelParameters
	var dataList []interface{}
	if name == "" || componentName == "" {
		m.name = modelName
		m.componentName = modelName
	}
	m.modelName = modelName
	m.classAll = omc.OMC.GetInheritedClassesListAll([]string{componentName})
	m.components = omc.OMC.GetElementsList(m.classAll)
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
		if p[2] != "-" {
			m.className = p[2].(string)
		} else {
			m.className = ""
		}
		varName := p[3].(string)
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
			if len(dList) <= 1 {
				continue
			}
			tab := dList[tabIndex].([]interface{})[0]
			group := dList[tabIndex].([]interface{})[1]
			dataDefault["tab"] = tab.(string)
			dataDefault["group"] = group.(string)
			showStartAttribute = dList[tabIndex].([]interface{})[3].(string)
		}
		ComponentModifierValue := omc.OMC.GetComponentModifierValue(modelName, name+"."+dataDefault["name"].(string))

		dataDefault["value"] = ComponentModifierValue
		if (p[10] != "parameter" && dataDefault["group"] != "Parameters" && p[9] != "True") || p[5] == "protected" || p[6] == "True" {
			continue
		}
		if p[10] == "parameter" || dataDefault["group"] != "Parameters" || p[9] != "True" {
			dataDefault["group"] = "Parameters"
			isEnumeration := omc.OMC.IsEnumeration(m.className)
			if isEnumeration == "true" {
				Literals := omc.OMC.GetEnumerationLiterals(m.className)
				dataDefault["options"] = func() []string {
					var oData []string
					for i := 0; i < len(Literals); i++ {
						if Literals[i] != "" {
							oData = append(oData, strings.TrimPrefix(m.className, ".")+"."+Literals[i])
						}
					}
					return oData
				}()
				dataDefault["type"] = "Enumeration"
			}
			parameterValue := m.getParameterValue(dataDefault["name"].(string))
			dataDefault["defaultvalue"] = parameterValue
			if p[2] == "Boolean" {
				dataDefault["type"] = "CheckBox"
				if ComponentModifierValue == "true" || ComponentModifierValue == "false" {
					dataDefault["value"] = ComponentModifierValue
					dataDefault["checked"] = ComponentModifierValue
					dataDefault["defaultvalue"] = ComponentModifierValue
				} else {
					if parameterValue == "true" || parameterValue == "false" {
						dataDefault["value"] = parameterValue
						dataDefault["checked"] = parameterValue
						dataDefault["defaultvalue"] = parameterValue
					}
				}
			}
		} else {
			componentModifierNames := omc.OMC.GetComponentModifierNamesList(m.classAll, varName)
			fixedValueString := m.getComponentModifierFixedValue(m.name + "." + varName + ".fixed")
			fixedValueBool, _ := strconv.ParseBool(fixedValueString)
			dataDefault["name"] = varName + ".start"
			dataDefault["unit"] = m.getDerivedClassModifierValue()
			dataDefault["group"] = "Initialization"
			fixed := map[string]interface{}{
				"tab":          dataDefault["tab"],
				"type":         "fixed",
				"group":        "Initialization",
				"name":         varName + ".fixed",
				"comment":      dataDefault["comment"],
				"defaultvalue": fixedValueBool,
				"value":        fixedValueBool,
				"unit":         m.getDerivedClassModifierValue(),
			}
			cmnStart := func() bool {
				for i := 0; i < len(componentModifierNames); i++ {
					if ("stateSelect") == componentModifierNames[i] || "start" == componentModifierNames[i] {
						return true
					}
				}
				return false
			}()
			switch {
			case showStartAttribute == "true":
				startValue := m.getComponentModifierStartValue(m.name+"."+varName+".start", true)
				dataDefault["defaultvalue"] = startValue
				dataList = append(dataList, fixed)
			case cmnStart == true:
				startValue := m.getComponentModifierStartValue(varName+".start", false)
				dataDefault["defaultvalue"] = startValue
				dataList = append(dataList, fixed)
			case true:
				continue
			}
		}
		dataDefault["unit"] = ""
		unit := m.getDerivedClassModifierValue()
		if len(unit) > 0 {
			dataDefault["unit"] = unit
		}
		dataList = append(dataList, dataDefault)
	}
	extendModifierName, extendModifierValue, extendModifierFinal := m.getExtendsModifierNameAndValue()
	if len(extendModifierName) > 0 && len(extendModifierValue) > 0 {
		for i := 0; i < len(extendModifierName); i++ {
			varName := strings.TrimSuffix(extendModifierName[i], ".start")
			m.className = m.componentsDict[varName].([]interface{})[2].(string)
			dataDefault := map[string]interface{}{
				"tab":          "General",
				"type":         "Normal",
				"group":        "Initialization",
				"name":         varName + ".start",
				"unit":         m.getDerivedClassModifierValue(),
				"comment":      m.componentsDict[varName].([]interface{})[3].(string),
				"defaultvalue": extendModifierValue[i],
				"value":        "",
			}
			fixedValueBool, _ := strconv.ParseBool(extendModifierFinal[i])
			fixed := map[string]interface{}{
				"tab":          dataDefault["tab"],
				"type":         "fixed",
				"group":        "Initialization",
				"name":         varName + ".fixed",
				"comment":      m.componentsDict[varName].([]interface{})[3].(string),
				"defaultvalue": fixedValueBool,
				"value":        fixedValueBool,
				"unit":         m.getDerivedClassModifierValue(),
			}
			dataList = append(dataList, fixed)
			dataList = append(dataList, dataDefault)
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
