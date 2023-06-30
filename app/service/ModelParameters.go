package service

import (
	"errors"
	"strings"
	"yssim-go/library/omc"
)

type modelParameters struct {
	componentClassName      string                            // 模型组件的模型名称
	componentName           string                            // 模型组件名称
	componentParameters     bool                              // 标记获取的是组件参数
	extend                  bool                              // 标记该参数是继承获得
	extendName              string                            // 该参数继承自哪个模型
	extendLevel2Name        string                            // 模型的第一个父类, 继承来的全局参数都由这个模型名称来设置标识符
	extendComponent         bool                              // 标记是组件参数,但该参数是继承来的
	extendsModifierNamesMap map[string]map[string]interface{} // 继承来的标识符map
	elementModifierNamesMap map[string]elementModifier        // 组件所属标识符,与继承来的不是一个东西
	parentAndChild          map[string]string                 // 记录模型名称的继承关系
	className               string                            // 记录某个当前正在使用的className
	modelName               string                            // 模型名称
	classAll                []string                          // 模型名称与父类们的集合
	level                   int                               // 记录当前参数在模型的第几层父类
	components              []interface{}                     // 当前class的组件列表
	componentAnnotations    []interface{}                     // 当前class的组件注解信息
	deduplicationMap        map[string]bool                   // 用来排除重复参数
}

type elementModifier struct {
	value      []string          // 用来记录值, 最多有两个, 一个用于value, 第二个位置保留
	className  string            // 当前标识符所在的class名称
	level      int               // 标记当前标识符在模型的第几层父类出现
	fixed      interface{}       // 参数的fixed值
	start      string            // 参数的start值
	startLevel int               // 标记当前参数的start标识符出现在模型的第几层父类出现
	other      map[string]string // 保留使用,用于记录参数的属性等信息,包括单位,显示单位,最大最小值等
}

func GetModelParameters(modelName, componentName, componentClassName string) []interface{} {
	var m modelParameters
	m.componentName = componentName
	m.componentClassName = componentClassName
	m.modelName = modelName
	m.classAll = []string{modelName}
	m.extendsModifierNamesMap = make(map[string]map[string]interface{}, 0)
	m.elementModifierNamesMap = make(map[string]elementModifier, 0)
	m.deduplicationMap = make(map[string]bool, 0)
	m.extend = false
	m.extendName = ""
	m.parentAndChild = map[string]string{}
	dataList := []interface{}{}
	if componentName == "" && componentClassName == "" { // 获取模型全局参数
		m.componentName = modelName
		m.componentClassName = modelName
		dataList = m.getClassParameters(modelName) // 执行获取参数逻辑的函数
	} else {
		m.getElementsModifierNamesAndValue(modelName, componentName) // 获取模型某个组件的参数, 预先获取这个组件的修饰符和值
		m.getComponentLevel(componentName, componentClassName)       // 找到这个组件是在模型的哪一层
		m.level += 1
		dataList = m.getClassParameters(componentClassName)
	}
	sortDataList := m.elementsToSort(dataList)
	return sortDataList
}

// getDerivedClassModifierValueALL 获取参数的单位, TODO 需要重构
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

// getUnit 获取参数的单位, TODO 需要重构
func getUnit(modelName, componentClassName, varName string) []string {
	modelNameAll := []string{modelName}
	for len(modelNameAll) > 0 {
		for _, name := range modelNameAll {
			value := omc.OMC.GetElementModifierValue(name, varName+"."+"unit")
			if value != "" {
				value = strings.ReplaceAll(value, "\"", "")
				value = strings.ReplaceAll(value, "\\", "")
				return []string{value}
			}
		}
		modelNameAll = omc.OMC.GetInheritedClassesList(modelNameAll)
	}

	return []string{getDerivedClassModifierValueALL(componentClassName)}
}

// getInherited 获取模型继承项, 并且继承等级加1
func (m *modelParameters) getInherited(modelName string) []string {
	if m.level == 1 {
		m.extendLevel2Name = modelName
	}
	classAll := []string{}
	inheritedList := omc.OMC.GetInheritedClasses(modelName)
	for i := 0; i < len(inheritedList); i++ {
		if inheritedList[i] != "" {
			classAll = append(classAll, inheritedList[i])
		}
	}
	m.level += 1
	return classAll
}

// geParent 获取模型的继承父类关系
func (m *modelParameters) geParent(modelName string) {
	inheritedList := omc.OMC.GetInheritedClasses(modelName)
	for i := 0; i < len(inheritedList); i++ {
		m.parentAndChild[inheritedList[i]] = modelName
	}
}

// getExtendsModifierNamesAndValue 获取模型继承过来的修饰符与值
func (m *modelParameters) getExtendsModifierNamesAndValue() {
	parentAndChild := m.parentAndChild
	for parent, child := range parentAndChild {
		extendsModifierNamesList := omc.OMC.GetExtendsModifierNames(child, parent)
		for _, name := range extendsModifierNamesList {
			_, ok := m.extendsModifierNamesMap[name]
			if !ok {
				extendsModifierValue := omc.OMC.GetExtendsModifierValue(child, parent, name)
				m.extendsModifierNamesMap[name] = map[string]interface{}{"parent": parent, "child": child, "value": extendsModifierValue, "level": m.level}
			}
		}
	}
}

// getElementsModifierNamesAndValue 获取组件的修饰符与值
func (m *modelParameters) getElementsModifierNamesAndValue(className, componentName string) {

	elementModifierNamesList := omc.OMC.GetElementModifierNames(className, componentName)
	for _, name := range elementModifierNamesList {
		elementModifierValue := omc.OMC.GetElementModifierValue(className, componentName+"."+name)
		modifierData := elementModifier{
			value:     []string{elementModifierValue},
			className: className,
			level:     m.level,
			fixed:     nil,
			start:     "",
			other:     map[string]string{},
		}

		switch {
		case strings.HasSuffix(name, ".fixed"):
			modifierName := componentName + "." + strings.TrimSuffix(name, ".fixed")
			elementModifierNamesMapData, ok := m.elementModifierNamesMap[modifierName]
			if !ok {
				elementModifierNamesMapData.fixed = elementModifierValue
				m.elementModifierNamesMap[modifierName] = elementModifierNamesMapData
			} else {
				modifierData.fixed = "true"
				m.elementModifierNamesMap[modifierName] = modifierData
			}
		case name == "fixed":
			modifierName := m.componentName + "." + componentName
			elementModifierNamesMapData, ok := m.elementModifierNamesMap[modifierName]
			if !ok {
				elementModifierNamesMapData.fixed = elementModifierValue
				m.elementModifierNamesMap[modifierName] = elementModifierNamesMapData
			} else {
				modifierData.fixed = "true"
				m.elementModifierNamesMap[modifierName] = modifierData
			}
		case strings.HasSuffix(name, ".start"):
			modifierName := componentName + "." + strings.TrimSuffix(name, ".start")
			elementModifierNamesMapData, ok := m.elementModifierNamesMap[modifierName]
			switch {
			case ok && elementModifierNamesMapData.start == "":
				elementModifierNamesMapData.start = elementModifierValue
				elementModifierNamesMapData.startLevel = m.level
				m.elementModifierNamesMap[modifierName] = elementModifierNamesMapData
			case !ok:
				modifierData.start = elementModifierValue
				m.elementModifierNamesMap[modifierName] = modifierData
			}
		case name == "start":
			modifierName := m.componentName + "." + componentName
			elementModifierNamesMapData, ok := m.elementModifierNamesMap[modifierName]
			switch {
			case ok && elementModifierNamesMapData.start == "":
				elementModifierNamesMapData.start = elementModifierValue
				elementModifierNamesMapData.startLevel = m.level
				m.elementModifierNamesMap[modifierName] = elementModifierNamesMapData
			case !ok:
				modifierData.start = elementModifierValue
				m.elementModifierNamesMap[modifierName] = modifierData
			}
		default:
			elementModifierNamesMapData, ok := m.elementModifierNamesMap[componentName+"."+name]
			switch {
			case ok && len(modifierData.value) < 2:
				elementModifierNamesMapData.value = append(elementModifierNamesMapData.value, elementModifierValue)
			case !ok:
				m.elementModifierNamesMap[componentName+"."+name] = modifierData
			}
		}
	}
}

// getComponentLevel 获取组件在模型父类的第几层, 0表示组件不是继承来的, 是当前模型自己的组件
func (m *modelParameters) getComponentLevel(componentName, componentClassName string) {
	classAll := []string{m.modelName}
	//classAll := m.getInherited(m.modelName)
Loop:
	for i := 0; i < len(classAll); i++ {
		m.geParent(classAll[i])
		m.getExtendsModifierNamesAndValue()
		m.components = omc.OMC.GetElements(classAll[i])
		m.componentAnnotations = omc.OMC.GetElementAnnotations(classAll[i])

		for index, c := range m.components {
			cAnnotations := m.componentAnnotations[index].([]interface{})
			componentList := c.([]interface{})
			cName := componentList[3].(string)
			cClassName := componentList[2].(string)
			if cName == componentName && cClassName == componentClassName && len(cAnnotations) > 0 && cAnnotations[0] == "Placement" {
				m.componentParameters = true
				m.getElementsModifierNamesAndValue(classAll[i], componentName)
				if m.level == 1 {
					m.extendLevel2Name = classAll[i]
				}
				if m.level > 0 {
					m.extendComponent = true
				}
				break Loop
			}
		}

		classes := m.getInherited(classAll[i])
		classAll = append(classAll, classes...)
	}
}

// getClassParameters 获取参数的主要逻辑
func (m *modelParameters) getClassParameters(className string) []interface{} {

	classAll := []string{className}
	dataList := []interface{}{}
	for i := 0; i < len(classAll); i++ {
		m.className = classAll[i]
		if (m.level > 1 && m.componentParameters) || (m.level > 0 && !m.componentParameters) {
			m.extend = true
			m.extendName = classAll[i]
		}
		if m.level == 1 {
			m.extendLevel2Name = classAll[i]
		}
		m.components = omc.OMC.GetElements(classAll[i])
		m.componentAnnotations = omc.OMC.GetElementAnnotations(classAll[i])
		m.geParent(classAll[i])
		m.getExtendsModifierNamesAndValue()
		m.getElementsModifierNamesAndValue(classAll[i], m.componentName)
		dataList = append(dataList, m.elementsAndAnnotations(className)...)
		classes := m.getInherited(classAll[i])
		classAll = append(classAll, classes...)
	}
	return dataList
}

// elementsToSort 组件排序, 参数类型的方到前面
func (m *modelParameters) elementsToSort(data []interface{}) []interface{} {
	parametersElements := []interface{}{}
	othersElements := []interface{}{}
	for _, d := range data {
		dElement := d.(map[string]interface{})
		if dElement["tab"].(string) == "General" {
			parametersElements = append(parametersElements, dElement)
			continue
		}
		othersElements = append(othersElements, dElement)
	}
	parametersElementsNew := []interface{}{}
	othersElementsNew := []interface{}{}
	for _, d := range parametersElements {
		dElement := d.(map[string]interface{})
		// "tab": "General"
		if dElement["group"].(string) == "Parameters" {
			parametersElementsNew = append(parametersElementsNew, dElement)
			continue
		}
		othersElementsNew = append(othersElementsNew, dElement)
	}
	parametersElementsNew = append(parametersElementsNew, othersElementsNew...)
	parametersElementsNew = append(parametersElementsNew, othersElements...)
	return parametersElementsNew
}

// elementsAndAnnotations 组件与注解的主要处理逻辑, 一般只包括 修饰符与parameter类型组件
func (m *modelParameters) elementsAndAnnotations(modelName string) []interface{} {
	dataList := []interface{}{}
	className := ""
	if m.extend == true {
		modelName = m.extendName
	}
	for i := 0; i < len(m.components); i++ {
		p := m.components[i].([]interface{})
		className = p[2].(string)
		varName := p[3].(string)
		parameterValue := omc.OMC.GetParameterValue(m.className, varName) // 获取当前组件的默认值

		m.getElementsModifierNamesAndValue(m.className, varName) // 获取当前组件的修饰符与值
		if m.deduplicationMap[varName] {
			continue
		}
		m.deduplicationMap[varName] = true
		isExtend := false
		if m.extendComponent {
			isExtend = m.extend
		}
		dataDefault := map[string]interface{}{"tab": "General", "type": "Normal", "group": "Parameters", "is_extend": isExtend, "extend_name": m.extendLevel2Name}

		IsExtendsModifierFinal := "false"
		if p[5] == "protected" || p[6] == true || p[8] == true { // 筛选模型, 部分受保护的,隐藏的需要被过滤
			continue
		}
		dataDefault["name"] = varName
		dataDefault["comment"] = p[4].(string)
		modifier := m.componentName + "." + varName

		elementModifierData := m.elementModifierNamesMap[modifier] // 查找有没有标识符标记该组件或参数
		elementModifierValue := elementModifierData.value          // 如果有标记的话, 取出值
		switch {                                                   // 根据参数是否被标识符标记, 参数所在模型的第几层父类判断是value还是默认值
		case m.level > 0 && len(elementModifierData.value) == 0:
			dataDefault["defaultvalue"] = parameterValue
		case len(elementModifierValue) != 0 && elementModifierData.level > 0:
			dataDefault["defaultvalue"] = elementModifierValue[0]
		case m.level == 0 && len(elementModifierValue) == 0:
			dataDefault["value"] = parameterValue
		case len(elementModifierValue) != 0 && elementModifierData.level == 0:
			dataDefault["value"] = elementModifierValue[0]
		default:
			dataDefault["defaultvalue"] = parameterValue
		}

		dList := m.componentAnnotations[i].([]interface{}) // 参数Dialog信息
		choices := func() map[string]interface{} {
			for index, d := range dList {
				if d == "choices" && index+1 <= len(dList)-1 { // 如果有choices关键字, 并包含true和false两个值,则表示该值以checkBox的形式出现, omc的固定返回格式
					if dList[index+1].([]interface{})[0] == "true" && dList[index+1].([]interface{})[1] == "false" {
						return map[string]interface{}{"checkBox": true} // 这里用checkBox标记,下
					}
					return map[string]interface{}{"value": dList[index+1], "checkBox": false}
				}
			}
			return nil
		}()

		DialogIndex, DialogIndexOk := func() (int, bool) { // 查找参数Dialog信息的位置,omc返回信息不是固定的位置
			for n := 0; n < len(dList); n++ {
				if dList[n] == "Dialog" {
					return n, true
				}
			}
			return 0, false
		}()
		showStartAttribute := ""

		if DialogIndexOk { // 处理参数Dialog信息
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
				modelNameList := strings.Split(m.modelName, ".")
				isPackage := omc.OMC.IsPackage(modelNameList[0])
				if dListTab[5].(string) != "-" && isPackage {
					dataDefault["type"] = "file"
					dataDefault["caption"] = dListTab[6].(string)
					dataDefault["filter"] = strings.Split(dListTab[5].(string), ";;")
				}
			}
		}
		emName := varName
		if m.extendsModifierNamesMap[emName] == nil {
			emName = m.componentName + "." + varName
		}
		if m.extendsModifierNamesMap[emName] != nil {
			extendsModifier := m.extendsModifierNamesMap[emName]
			IsExtendsModifierFinal = omc.OMC.IsExtendsModifierFinal(extendsModifier["child"].(string), extendsModifier["parent"].(string), varName)
			if IsExtendsModifierFinal == "true" { //判断参数是否是不可修改的, 如果是,则过滤该参数
				continue
			}
			extendsModifierValue := extendsModifier["value"] // 继承过来的标识符中如果有该参数的值,则根据level等级进行赋值
			if extendsModifier["level"].(int) > 0 {
				dataDefault["defaultvalue"] = extendsModifierValue
			} else {
				dataDefault["value"] = extendsModifierValue
			}
			dataDefault["name"] = strings.TrimPrefix(emName, m.componentName+".")
		}

		if p[9] == true { // 处理模板参数类型
			annotationBase := m.componentAnnotations[i].([]interface{})
			optionsBase := []string{}
			if len(annotationBase) > 1 && annotationBase[0] == "choices" {
				choicesData := annotationBase[1].([]interface{})
				if len(choicesData) > 2 {
				LOOP:
					for _, d := range choicesData[2].([]interface{}) {
						splitList := strings.Split(d.(string), " ")
						for index, s := range splitList {
							if s == "=" && index > 0 {
								newList := append(splitList[:index-1], varName)
								newList = append(newList, splitList[index:]...)
								optionsBase = append(optionsBase, strings.Join(newList, " "))
								continue LOOP
							}
						}
						optionsBase = append(optionsBase, d.(string))
					}
				}
			}
			options := []interface{}{}
			dataDefault["type"] = "Enumeration"
			if p[1] == "-" && p[13].(string) == "$Any" {
				dataDefault["value"] = ""
				dataDefault["options"] = options
				dataList = append(dataList, dataDefault)
				continue
			}
			switch {
			case p[13].(string) != "$Any":
				options = omc.OMC.GetAllSubtypeOf(p[13].(string), m.componentClassName)
				dataDefault["value"] = p[2].(string)
			case len(annotationBase) > 0 && annotationBase[0] == "choicesAllMatching=true":
				options = omc.OMC.GetAllSubtypeOf(p[2].(string), m.modelName)
			}
			classInformation := omc.OMC.GetClassInformation(p[2].(string))
			dataDefault["read_only"] = false
			if dataDefault["value"] == "" {
				dataDefault["value"] = omc.OMC.GetElementModifierValue(m.modelName, m.componentName+"."+dataDefault["name"].(string))
			}
			if m.level == 0 && dataDefault["value"] == "" {
				dataDefault["value"] = p[2].(string) + " - " + classInformation[1].(string)
			} else {
				dataDefault["read_only"] = true
			}
			if dataDefault["defaultvalue"] == "" {
				dataDefault["defaultvalue"] = p[2].(string) + " - " + classInformation[1].(string)
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

			dataList = append(dataList, dataDefault)
			continue
		}

		if p[10] == "parameter" { // 处理parameter
			isEnumeration := omc.OMC.IsEnumeration(className)
			if isEnumeration { // 处理枚举类型
				Literals := omc.OMC.GetEnumerationLiterals(className)
				dataDefault["options"] = func() []string {
					oData := []string{parameterValue}
					for _, literal := range Literals {
						literalValue := strings.TrimPrefix(className, ".") + "." + literal
						if literal != "" && literalValue != parameterValue {
							oData = append(oData, literalValue)
						}
					}
					return oData
				}()
				dataDefault["type"] = "Enumeration"
			}

			if p[2] == "Boolean" || choices != nil { // 处理bool类型, 该类型可能会是勾选,可能会是下拉的形式
				if choices["checkBox"] == true {
					dataDefault["type"] = "CheckBox"
					if len(elementModifierValue) != 0 {
						dataDefault["checked"] = elementModifierValue[0]
					} else {
						dataDefault["checked"] = parameterValue
					}
				} else {
					dataDefault["type"] = "Enumeration"
					dataDefault["options"] = []string{"", "true", "false"}
					dataDefault["defaultvalue"] = parameterValue
				}
			}

			unit := getUnit(m.componentClassName, className, varName)
			if len(unit) > 0 {
				dataDefault["unit"] = unit
			}
			dataList = append(dataList, dataDefault)
			continue
		}
		if elementModifierData.start != "" || elementModifierData.fixed != nil || showStartAttribute == "true" { // 处理 fixed类型参数
			dataDefault["type"] = "checkWrite"
			dataDefault["name"] = varName + ".start"
			dataDefault["group"] = "Initialization"
			isFixed := func() interface{} { // 标记该参数的fixed值
				switch {
				case elementModifierData.fixed == true || elementModifierData.fixed == "true":
					return true
				case elementModifierData.fixed == false || elementModifierData.fixed == "false":
					return false
				}
				return ""
			}()
			value := elementModifierData.start // 标记参数的start值
			if elementModifierData.startLevel > 0 || elementModifierData.level > 0 {
				dataDefault["defaultvalue"] = elementModifierData.start
				value = ""
			}
			dataDefault["value"] = map[string]interface{}{"isFixed": isFixed, "value": value}
			dataDefault["unit"] = ""
			unit := getUnit(m.componentClassName, className, varName)
			if len(unit) > 0 {
				dataDefault["unit"] = unit
			}
			dataList = append(dataList, dataDefault)
			continue
		}
	}
	return dataList
}

// SetComponentModifierValue 参数操作
func SetComponentModifierValue(className string, parameterValue map[string]string) bool {
	for k, v := range parameterValue {
		result := omc.OMC.SetElementModifierValue(className, k, v)
		if !result {
			return false
		}
	}
	return true
}

func SetElementModifierValue(className string, parameter, Value string) bool {
	result := omc.OMC.SetElementModifierValue(className, parameter, Value)
	return result
}

func SetExtendsModifierValue(className, extentsName, parameter, Value string) bool {
	result := omc.OMC.SetExtendsModifierValue(className, extentsName, parameter, Value)
	return result
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
