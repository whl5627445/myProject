package service

import (
	"errors"
	"strconv"
	"strings"
	"yssim-go/library/omc"
)

type modelParameters struct {
	name               string
	componentClassName string
	//packageName          string
	className            string
	modelName            string
	classAll             []string
	components           [][]interface{}
	componentAnnotations [][]interface{}
	componentsDict       map[string]interface{}
}

func (m modelParameters) getParameterValue(name string) string {
	data := ""
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
	return data
}

func (m modelParameters) getExtendsModifierNameAndValue() ([]string, []string, []string) {
	var dataNameList []string
	var dataValueList []string
	var dataFinalList []string
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
	modelInheritedClasses := map[string]bool{}
	extendsModifierNamesList := []string{}
	extendsModifierNamesMap := make(map[string]map[string]string, 0)
	for i := 1; i < len(m.classAll) && !bEnd; i++ {
		modelInheritedClasses[m.classAll[i]] = true
		extendsModifierNamesList = append(extendsModifierNamesList, omc.OMC.GetExtendsModifierNames(modelName, m.classAll[i])...)
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
	for _, emName := range extendsModifierNamesList {
		extendsModifierNamesMap[emName] = make(map[string]string, 0)
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
			dListTab := dList[tabIndex].([]interface{})
			if tabIndex > 0 && len(dListTab) > 3 {
				if len(dList) <= 1 || dListTab[len(dListTab)-1] == "true" {
					continue
				}
				tab := dListTab[0]
				group := dListTab[1]
				dataDefault["tab"] = tab.(string)
				dataDefault["group"] = group.(string)
				showStartAttribute = dListTab[3].(string)
				modelNameList := strings.Split(modelName, ".")
				isPackage := omc.OMC.IsPackage(modelNameList[0])
				if dListTab[5].(string) != "-" && isPackage {
					dataDefault["type"] = "file"
					dataDefault["caption"] = dListTab[6].(string)
					dataDefault["filter"] = strings.Split(dListTab[5].(string), ";;")
				}
			}
		}

		emName := componentName + "." + varName
		if extendsModifierNamesMap[emName] != nil {
			extendsModifierValue := omc.OMC.GetExtendsModifierValue(modelName, m.modelName, emName)
			dataDefault["value"] = extendsModifierValue
			dataDefault["name"] = strings.TrimPrefix(emName, componentName+".")
			dataDefault["unit"] = getUnit(componentClassName, m.className, varName)
			dataList = append(dataList, dataDefault)
			continue
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

		if p[9] == true {
			annotationBase := p[16].([]interface{})
			optionsBase := []string{}
			if len(annotationBase) > 1 && annotationBase[0] == "choices" {
				choicesData := annotationBase[1].([]interface{})
				if len(choicesData) > 2 {
					for _, d := range choicesData[2].([]interface{}) {
						optionsBase = append(optionsBase, d.(string))
					}
				}
			}
			options := []interface{}{}
			dataDefault["defaultvalue"] = p[2]
			dataDefault["type"] = "Enumeration"
			if p[1] == "-" && p[13].(string) == "$Any" {
				dataDefault["defaultvalue"] = "replaceable" + " " + p[2].(string) + " " + p[3].(string)
				dataDefault["value"] = ""
				dataDefault["options"] = options
				dataList = append(dataList, dataDefault)
				continue
			}
			if p[13].(string) != "$Any" {
				options = omc.OMC.GetAllSubtypeOf(p[13].(string), componentClassName)
			} else {
				options = omc.OMC.GetAllSubtypeOf(p[2].(string), m.modelName)
			}

			oData := make([]string, 1)
			oData = append(oData, optionsBase...)
			for _, option := range options {
				optionData := ""
				if p[1] != "-" {
					optionData = "redeclare " + p[1].(string) + " " + p[3].(string) + " = " + option.(string)
				} else {
					optionData = "replaceable " + p[2].(string) + " " + p[3].(string)
				}
				oData = append(oData, optionData)
			}
			dataDefault["options"] = oData
			dataDefault["value"] = omc.OMC.GetElementModifierValue(m.modelName, componentName+"."+dataDefault["name"].(string))
			dataList = append(dataList, dataDefault)
			continue
		}

		if p[10] == "parameter" || dataDefault["group"] != "Parameters" || DialogIndexOk {
			componentModifierValue := omc.OMC.GetElementModifierValue(m.modelName, componentName+"."+dataDefault["name"].(string))
			dataDefault["value"] = componentModifierValue

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

	return dataList
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
		result := omc.OMC.SetElementModifierValue(className, k, v)
		if !result {
			return false
		}
	}
	return true
}

func AddComponentParameters(varName, varType, className string) (bool, error) {
	check := checkComponentParameter(className, varName)
	if check != nil {
		return false, check
	}
	var defaultValue string
	switch {
	case varType == "Boolean":
		defaultValue = "false"
	case varType == "Real":
		defaultValue = ""
	case varType == "Integer":
		defaultValue = "0"
	}
	ok := omc.OMC.AddComponentParameter(varName, varType, className, defaultValue)
	if ok {
		return true, nil
	}
	return false, errors.New("新增参数失败")
}

func DeleteComponentParameters(varName, className string) (bool, error) {
	components := omc.OMC.GetComponents(className)
	for i := 0; i < len(components); i++ {
		name := components[i].([]interface{})[1]
		if name == varName {
			ok := omc.OMC.DeleteComponentParameter(varName, className)
			if ok {
				return true, nil
			}
		}
	}
	return false, errors.New("只能删除本模型下的参数")
}

func checkComponentParameter(className, varName string) error {
	components := GetElements(className, varName)
	if len(components) > 0 {
		return errors.New("参数名已存在")
	}
	return nil
}
