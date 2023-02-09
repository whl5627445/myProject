package service

import (
	"strconv"
	"strings"
	"yssim-go/library/omc"
)

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
	bEnd := false
	for i := 1; i < len(m.classAll) && !bEnd; i++ {
		elementsData := omc.OMC.GetElements(m.classAll[i])
		annotationsData := omc.OMC.GetElementAnnotations(m.classAll[i])
		for n, d := range elementsData {
			if d.([]interface{})[3] == componentName && (len(annotationsData[n].([]interface{})) != 0 || annotationsData[n].([]interface{})[0].(string) == "Placement") {
				m.modelName = m.classAll[i]
				inheritedClassAll := omc.OMC.GetInheritedClassesListAll([]string{d.([]interface{})[2].(string)})
				m.classAll = inheritedClassAll
				bEnd = true
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
		dataDefault := map[string]interface{}{"tab": "General", "type": "Normal", "group": "Parameters"}
		p := m.componentsDict[m.components[i][3].(string)].([]interface{})

		varName := p[3].(string)
		if p[2] != "-" {
			m.className = p[2].(string)
		} else {
			m.className = ""
		}
		IsExtendsModifierFinal := "false"
		for _, c := range m.classAll {
			IsExtendsModifierFinal = omc.OMC.IsExtendsModifierFinal(c, p[len(p)-2].(string), varName)
			if IsExtendsModifierFinal == "true" {
				break
			}
		}
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
			fixedValue, _ := strconv.ParseBool(fixedValueString)
			var fixedValueBool interface{}
			if fixedValueString == "" {
				fixedValueBool = ""
			} else {
				fixedValueBool = fixedValue
			}
			value := map[string]interface{}{"isFixed": fixedValueBool, "value": startValueString}
			data := map[string]interface{}{
				"type":    "checkWrite",
				"name":    varName + ".start",
				"comment": dataDefault["comment"],
				"tab":     dataDefault["tab"],
				"group":   dataDefault["group"],
				"value":   value,
				"unit":    getUnit(componentClassName, m.className, varName),
			}
			dataList = append(dataList, data)
			continue
		}

		if p[9] == "true" {
			continue
			//dataDefault["type"] = "Enumeration"
			//dataDefault["defaultvalue"] = p[2]
			//dataDefault["disable"] = true
			//dataDefault["group"] = "参数"
			//oData := make([]string, 1)
			//if p[13].(string) != "$Any" {
			// 模板参数获取有内存泄露问题， 暂时不用
			//options := omc.OMC.GetAllSubtypeOf(p[13].(string), componentClassName)
			//for _, option := range options {
			//	optionData := "redeclare "+ option.(string) + " " + componentName+"."+dataDefault["name"].(string)
			//	oData = append(oData, optionData)
			//}
			//	dataDefault["disable"] = false
			//}
			//dataDefault["options"] = oData
			//dataList = append(dataList, dataDefault)
			//continue
		}

		if p[10] == "parameter" || dataDefault["group"] != "Parameters" || DialogIndexOk == true {
			componentModifierValue := omc.OMC.GetElementModifierValue(m.modelName, componentName+"."+dataDefault["name"].(string))
			dataDefault["value"] = componentModifierValue
			if componentModifierValue == "" {
				for _, n := range m.classAll {
					componentExtendsModifierValue := omc.OMC.GetExtendsModifierValue(m.componentClassName, n, varName)
					if componentExtendsModifierValue != "" {
						dataDefault["value"] = componentExtendsModifierValue
						break
					}
				}
			}

			isEnumeration := omc.OMC.IsEnumeration(m.className)
			if isEnumeration {
				Literals := omc.OMC.GetEnumerationLiterals(m.className)
				dataDefault["options"] = func() []string {
					oData := []string{componentModifierValue}
					for _, literal := range Literals {
						literalValue := strings.TrimPrefix(m.className, ".") + "." + literal
						if literal != "" && literalValue != componentModifierValue {
							oData = append(oData, literalValue)
						}
					}
					return oData
				}()
				dataDefault["type"] = "Enumeration"
			}
			parameterValue := omc.OMC.GetParameterValue(p[len(p)-2].(string), dataDefault["name"].(string))
			dataDefault["defaultvalue"] = parameterValue

			if p[2] == "Boolean" && (componentModifierValue != "" || parameterValue != "") {
				dataDefault["type"] = "CheckBox"
				if componentModifierValue != "" {
					if componentModifierValue == "true" || componentModifierValue == "false" {
						dataDefault["checked"] = componentModifierValue
					} else {
						dataDefault["type"] = "Enumeration"
						dataDefault["options"] = []string{componentModifierValue, "true", "false"}
					}
					dataDefault["defaultvalue"] = componentModifierValue
					dataDefault["value"] = componentModifierValue
				} else {
					if parameterValue == "true" || parameterValue == "false" {
						dataDefault["checked"] = parameterValue
					} else {
						dataDefault["type"] = "Enumeration"
						dataDefault["options"] = []string{componentModifierValue, "true", "false"}
					}
					dataDefault["value"] = parameterValue
					dataDefault["defaultvalue"] = parameterValue
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
		name := p[len(p)-2].(string)
		if fixedValueString == "" {
			fixedValueString = m.getStartAndFixedValue(name, varName, "fixed")
		}
		if startValueString == "" {
			startValueString = m.getStartAndFixedValue(name, varName, "start")
		}
		if fixedValueString != "" || startValueString != "" {
			fixedValue, _ := strconv.ParseBool(fixedValueString)
			var fixedValueBool interface{}
			if fixedValueString == "" {
				fixedValueBool = ""
			} else {
				fixedValueBool = fixedValue
			}
			value := map[string]interface{}{"isFixed": fixedValueBool, "value": startValueString}
			data := map[string]interface{}{
				"type":    "checkWrite",
				"name":    varName + ".start",
				"comment": dataDefault["comment"],
				"tab":     dataDefault["tab"],
				"group":   "Initialization",
				"value":   value,
				"unit":    getUnit(componentClassName, m.className, varName),
			}
			dataList = append(dataList, data)
		}
	}
	nameMap := make(map[string]int, 1)
	var data []interface{}
	extendsModifierNamesList := omc.OMC.GetExtendsModifierNames(modelName, m.modelName)
	extendsNameMap := map[string]map[string]string{}
	for _, nameAll := range extendsModifierNamesList {
		if strings.HasPrefix(nameAll, componentName) {
			nameList := strings.Split(nameAll, ".")
			cName := nameList[0]
			pName := nameList[1]
			extendsNameMap[pName] = map[string]string{"cName": cName, "nameAll": nameAll}
		}
	}
	for n := 0; n < len(dataList); n++ {
		pData := dataList[n].(map[string]interface{})
		name := pData["name"].(string)
		_, ok := nameMap[name]
		if !ok {
			extendsModifierData, ok := extendsNameMap[name]
			if ok {
				extendsModifierValue := omc.OMC.GetExtendsModifierValue(modelName, m.modelName, extendsModifierData["nameAll"])
				pData["value"] = extendsModifierValue
			}
			data = append(data, pData)
			nameMap[name] = n
		}
	}
	return data
}

func (m modelParameters) getStartAndFixedValue(name, varName, varType string) string {
	valueString := omc.OMC.GetElementModifierValue(name, varName+"."+varType)
	if valueString == "" {
		extendName := []string{m.componentClassName}
		for extendName != nil {
			extendName = omc.OMC.GetInheritedClassesList(extendName)
			for _, n := range extendName {
				valueString = omc.OMC.GetExtendsModifierValue(m.componentClassName, n, varName+"."+varType)
				if valueString != "" {
					extendName = nil
					break
				}
			}
		}
	}
	return valueString
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
