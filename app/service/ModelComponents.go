package service

import (
	"strconv"
	"strings"
	"yssim-go/app/DataType"
	"yssim-go/config"
	"yssim-go/library/omc"
	"yssim-go/library/stringOperation"
)

func GetElements(className, componentName string) []any {
	classNameList := GetICList(className)
	var componentsData []any
	var annotationsData []any
	for i := 0; i < len(classNameList); i++ {
		classnameData := omc.OMC.GetElements(classNameList[i])
		classnameAnnotationsData := omc.OMC.GetElementAnnotations(classNameList[i])
		componentsData = append(componentsData, classnameData...)
		annotationsData = append(annotationsData, classnameAnnotationsData...)

	}

	var componentData []any
	for i := 0; i < len(componentsData); i++ {
		cData := componentsData[i].([]any)
		switch {
		case componentName != "" && cData[3] == componentName:
			return cData
		case !(cData[6] == "true" || len(annotationsData[i].([]any)) == 0 || annotationsData[i].([]any)[0].(string) != "Placement") && componentName == "":
			componentData = append(componentData, cData)
		}
	}

	return componentData
}

func getDefaultComponentName(className string) string {
	return omc.OMC.GetDefaultComponentName(className)
}

func GetComponentName(modelName, className string) string {
	defaultComponentName := getDefaultComponentName(className)
	name := ""
	if defaultComponentName != "" {
		name = defaultComponentName
	} else {
		nameList := strings.Split(className, ".")
		name = strings.ToLower(nameList[len(nameList)-1])
	}
	modelNameList := GetICList(modelName)
	componentsData := omc.OMC.GetElementsList(modelNameList)
	nameNum := 0
	nameMap := map[string]bool{}
	if _, ok := config.ModelicaKeywords[name]; ok {
		nameNum += 1
	}
	for _, c := range componentsData {
		cList := c
		n := cList[3].(string)
		if len(cList) >= 2 && strings.HasPrefix(n, name) {
			nameMap[n] = true
			nameNum += 1
		}
	}
	for i := 1; i < nameNum+1; i++ {
		n := name + strconv.Itoa(i)
		_, ok := config.ModelicaKeywords[n]
		_, mapOk := nameMap[n]
		if !ok && !mapOk {
			return n
		}
	}
	return name
}

func addComponentVerification(oldComponentName, newComponentName, modelName string) (bool, string) {
	classType := omc.OMC.GetClassRestriction(oldComponentName)
	modelType := omc.OMC.GetClassRestriction(modelName)
	if modelType == "package" {
		return false, " package类型不能新增组件"
	}
	noType := config.ClassTypeAll[classType]
	if !noType {
		return false, "不能插入：" + oldComponentName + ", 这是一个 \"" + classType + " \"类型。组件视图层只允许有model、record、expandable connector、class、connector、function或者block。"
	}
	elementsData := omc.OMC.GetElements(modelName)
	for _, e := range elementsData {
		if e.([]any)[3] == newComponentName {
			return false, "新增组件失败，名称 \"" + newComponentName + "\" 已经存在或是 Modelica 关键字。 请选择其他名称。"
		}
	}
	return true, ""
}

func AddComponent(componentName, componentClassName, modelNameAll, origin, rotation string, extent []string) (bool, string) {
	v, msg := addComponentVerification(componentClassName, componentName, modelNameAll)
	if !v {
		return false, msg
	}
	result := false
	if omc.OMC.GetClassRestriction(componentClassName) != "connector" {
		result = omc.OMC.AddComponent(componentName, componentClassName, modelNameAll, origin, rotation, extent)
	} else {
		result = omc.OMC.AddInterfacesComponent(componentName, componentClassName, modelNameAll, origin, rotation, extent)
	}
	if !result {
		msg = "新增组件失败"
	}
	return result, msg
}

func DeleteComponent(componentName, modelNameAll string) bool {
	result := false
	components := omc.OMC.GetComponents(modelNameAll)
	for _, component := range components {
		if componentName == component.([]any)[1].(string) {
			result = omc.OMC.DeleteComponent(componentName, modelNameAll)
			break
		}
	}
	return result
}

func UpdateComponent(componentName, componentClassName, modelNameAll, origin, rotation string, extent []string) bool {
	result := false
	if omc.OMC.GetClassRestriction(componentClassName) != "connector" {
		result = omc.OMC.UpdateComponent(componentName, componentClassName, modelNameAll, origin, rotation, extent)
	} else {
		result = omc.OMC.UpdateInterfacesComponent(componentName, componentClassName, modelNameAll, origin, rotation, extent)
	}
	return result
}

func GetICList(name string) []string {
	nameList := []string{name}
	dataList := []string{name}
	for {
		var data []string
		for _, n := range nameList {
			data = append(data, omc.OMC.GetInheritedClasses(n)...)

		}
		if len(data) == 0 {
			break
		}
		dataList = append(data, dataList...)
		nameList = data
	}
	// dataList去重
	var datalistLen = len(dataList)
	for i := 0; i < datalistLen; i++ {
		for j := i + 1; j < datalistLen; j++ {
			if dataList[i] == dataList[j] {
				dataList = append(dataList[:i], dataList[i+1:]...)
				datalistLen--
				i--
				break
			}
		}
	}
	return dataList
}

// GetExtendedModel 获取父类列表
func GetExtendedModel(className string) []string {
	dataList := omc.OMC.GetInheritedClasses(className)
	if dataList != nil {
		return dataList
	} else {
		return nil
	}
}

// GetUMLElements 获取模型里的组件
func GetUMLElements(className string) []any {
	classnameData := omc.OMC.GetElements(className)
	return classnameData
}

type getUMLData struct {
	classNameList       []string
	secondClassNameList []string
	finalResultData     []map[string]*DataType.GetUMLData
	rootUmlData         DataType.GetUMLData
}

func GetModelUMLData(className string) []map[string]*DataType.GetUMLData {
	var umlData = getUMLData{}
	umlData.classNameList = append(umlData.classNameList, className)
	rootInformation := GetClassInformation(className)
	rootUmlData := &DataType.GetUMLData{
		ClassName:   GetLastModelName(className),
		ModelType:   "root",
		Description: rootInformation[1].(string),
		Level:       1,
		Library:     GetReferenceLibraries(className),
	}
	var finalResultData []map[string]*DataType.GetUMLData
	rootMap := map[string]*DataType.GetUMLData{
		rootUmlData.ClassName: rootUmlData,
	}
	finalResultData = append(finalResultData, rootMap)
	umlData.rootUmlData = *rootUmlData
	umlData.finalResultData = finalResultData
	//获取子节点
	umlData.GetSubUMLData(className, className)
	//获取父节点
	umlData.GetExtendUml(className, className)
	return umlData.finalResultData
}

// GetSubUMLData 获取子节点
func (umlData *getUMLData) GetSubUMLData(className, initialName string) {
	componentDataList := GetUMLElements(className)
	for i := 0; i < len(componentDataList); i++ {
		cData := componentDataList[i].([]any)
		subInformation := GetClassInformation(cData[2].(string))
		//节点类型是type或者“”，过滤掉
		if subInformation[0].(string) == "type" || subInformation[0].(string) == "" {
			continue
		}
		//节点属性是protected但不是expandable过滤掉
		if cData[5].(string) == "protected" {
			if !strings.Contains(subInformation[0].(string), "expandable") {
				continue
			}
		}
		var extendsModelData []DataType.ExtendsModelData
		rootExtendsModelData := DataType.ExtendsModelData{
			ClassName: umlData.rootUmlData.ClassName,
			Count:     1,
		}

		//获取子节点名
		var subClassName string
		if stringOperation.ContainsString(subInformation[0].(string)) {
			subClassName = className + "." + cData[3].(string)
			rootExtendsModelData.RelationShip = "relevance"
			if className != initialName {
				continue
			}
		} else {
			subClassName = umlData.distinct(cData[2].(string))
			rootExtendsModelData.RelationShip = "polymerization"
		}
		extendsModelData = append(extendsModelData, rootExtendsModelData)
		subUmlData := &DataType.GetUMLData{
			ClassName:        subClassName,
			Level:            umlData.rootUmlData.Level - 1,
			ExtendsModelData: extendsModelData,
		}
		//模型类型为model，则无修饰，多个形容词取最后一个，若组件信息第三个值为true，则修饰词为partial
		if !(subInformation[0].(string) == "model") {
			if strings.ContainsRune(subInformation[0].(string), ' ') {
				subUmlData.ModelType = stringOperation.GetComponentType(subInformation[0].(string))
			} else {
				subUmlData.ModelType = subInformation[0].(string)
			}
		}
		if subInformation[2].(string) == "true" {
			subUmlData.ModelType = "partial"
		}

		//根据子节点的名称判断finalResult是否存在该子节点，存在则在子节点的被指向的节点的数量加1或增加被指向的节点
		//不存在则直接在finalResult中添加该子节点
		ok := umlData.IsExistKey(subUmlData.ClassName, rootExtendsModelData)
		if !ok {
			subUmlResultData := map[string]*DataType.GetUMLData{
				subUmlData.ClassName: subUmlData,
			}
			umlData.finalResultData = append(umlData.finalResultData, subUmlResultData)
		}
	}
}

// GetExtendUml 获取父类节点
func (umlData *getUMLData) GetExtendUml(className, initialName string) {
	rootExtendModelNameList := GetExtendedModel(className)
	var extendsModelData []DataType.ExtendsModelData
	for _, extendModelName := range rootExtendModelNameList {
		extendsModelInformation := GetClassInformation(extendModelName)
		extendClassName := umlData.distinct(extendModelName)
		rootExtendModel := DataType.ExtendsModelData{
			ClassName:    extendClassName,
			RelationShip: "inherit",
		}
		extendsModelData = append(extendsModelData, rootExtendModel)
		for _, m := range umlData.finalResultData {
			// 检查是否存在指定的键
			if value, ok := m[umlData.rootUmlData.ClassName]; ok {
				value.ExtendsModelData = extendsModelData
			}
		}
		extendsModelUmlData := &DataType.GetUMLData{
			ClassName: extendClassName,
			Level:     umlData.rootUmlData.Level + 1,
		}
		//模型类型为model，则无修饰，多个形容词取最后一个，若组件信息第三个值为true，则修饰词为partial
		if !(extendsModelInformation[0].(string) == "model") {
			if strings.ContainsRune(extendsModelInformation[0].(string), ' ') {
				extendsModelUmlData.ModelType = stringOperation.GetComponentType(extendsModelInformation[0].(string))
			} else {
				extendsModelUmlData.ModelType = extendsModelInformation[0].(string)
			}
		}
		if extendsModelInformation[2].(string) == "true" {
			extendsModelUmlData.ModelType = "partial"
		}

		extendsModelUmlDataMap := map[string]*DataType.GetUMLData{
			extendsModelUmlData.ClassName: extendsModelUmlData,
		}
		umlData.finalResultData = append(umlData.finalResultData, extendsModelUmlDataMap)
		umlData.GetSubUMLData(extendModelName, initialName)
		umlData.GetExtendUml(extendModelName, initialName)
	}
}

func GetSecondLastModelName(className string) string {
	name := strings.Split(className, ".")
	return strings.Join(name[len(name)-2:], ".")
}

func GetThirdLastModelName(className string) string {
	name := strings.Split(className, ".")
	return strings.Join(name[len(name)-3:], ".")
}

// distinct  解决节点名重复的问题
func (umlData *getUMLData) distinct(target string) string {

	//Modelica.Blocks.Examples.PID_Controller 字符串存在数组strArray中，则截取最后一位小数点后的内容
	if stringOperation.SliceIndexString(target, umlData.classNameList) {
		return GetLastModelName(target)
	}
	//不存在，则判断数组中是否有以GetLastModelName(target)结尾的字符串
	for _, s := range umlData.classNameList {
		if strings.HasSuffix(s, GetLastModelName(target)) {
			//Modelica.Blocks.Examples.PID_Controller 字符串存在数组strArray中，则截取最后第二位小数点后的内容
			if stringOperation.SliceIndexString(target, umlData.secondClassNameList) {
				return GetSecondLastModelName(target)
			}
			//不存在，则判断数组secondStrArray中是否有以GetLastModelName(target)结尾的字符串
			for _, s2 := range umlData.secondClassNameList {
				if strings.HasSuffix(s2, GetSecondLastModelName(target)) {
					//若有以GetLastModelName(target)结尾的字符串，则返回GetThirdLastModelName
					return GetThirdLastModelName(target)
				}
			}
			//没有以GetSecondLastModelName(target)结尾的，则将字符串添加到数组中，并返回GetSecondLastModelName
			umlData.secondClassNameList = append(umlData.secondClassNameList, target)
			return GetSecondLastModelName(target)
		}
	}
	//没有以GetLastModelName(target)结尾的，则将字符串添加到数组中，并返回GetLastModelName
	umlData.classNameList = append(umlData.classNameList, target)
	return GetLastModelName(target)
}

// GetLastModelName 获取Modelica.Blocks.Examples.PID_Controller中PID_Controller
func GetLastModelName(className string) string {
	return className[strings.LastIndex(className, ".")+1:]
}

// GetReferenceLibraries 获取模型依赖库
func GetReferenceLibraries(className string) []string {
	var referencedLibraries []string
	var packageUse [][]string
	if strings.Contains(className, ".") {
		packageUse = GetPackageUses(className[:strings.Index(className, ".")])
	} else {
		packageUse = GetPackageUses(className)
	}
	for _, modelLibrary := range packageUse {
		if modelLibrary[0] != "Complex" && modelLibrary[0] != "ModelicaServices" {
			referencedLibraries = append(referencedLibraries, modelLibrary[0])
		}
	}
	return referencedLibraries
}

// IsExistKey 判断最终结果finalResultData是否存在这个key
func (umlData *getUMLData) IsExistKey(className string, rootExtendsModelData DataType.ExtendsModelData) bool {

	for _, m := range umlData.finalResultData {
		if value, ok := m[className]; ok {
			extendsModelList := value.ExtendsModelData
			flag := false
			index := 0
			for i := 0; i < len(extendsModelList); i++ {
				if extendsModelList[i].ClassName == rootExtendsModelData.ClassName {
					flag = true
					index = i
				}
			}
			if flag {
				extendsModelList[index].Count = extendsModelList[index].Count + 1
			} else {
				value.ExtendsModelData = append(value.ExtendsModelData, rootExtendsModelData)
			}
			return ok
		}
	}
	return false
}
