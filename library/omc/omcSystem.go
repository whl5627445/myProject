package omc

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"regexp"
	"sync"
	"yssim-go/app/serviceType"
	"yssim-go/config"

	"github.com/bytedance/sonic"

	"os"
	"strconv"
	"strings"

	"github.com/go-zeromq/zmq4"
)

type ZmqObject struct {
	zmq4.Socket
	sync.Mutex
}

var cacheRefresh = false
var redisCacheKey = &config.RedisCacheKey
var userName = config.USERNAME

var allModelCache = config.R

// SendExpression 发送指令，获取数据
func (o *ZmqObject) SendExpression(cmd string) ([]any, bool) {
	//s := time.Now().UnixNano() / 1e6
	var msg []byte
	ctx := context.Background()
	msg, err := allModelCache.HGet(ctx, *redisCacheKey, cmd).Bytes()
	if len(msg) == 0 || string(msg) == "null" {
		o.Lock()
		_ = o.Send(zmq4.NewMsgString(cmd))
		data, _ := o.Recv()
		msg = data.Bytes()
		o.Unlock()
		if cacheRefresh && len(msg) > 0 {
			allModelCache.HSet(ctx, *redisCacheKey, cmd, msg)
		}
	}
	parseData, err := dataToGo(msg)
	if err != nil {
		log.Println("cmd: ", cmd)
		return nil, false
	}
	if len(parseData) == 0 {
		return nil, false
	}
	//log.Println("cmd", cmd)
	//if time.Now().UnixNano()/1e6-s > 20 {
	//	log.Println("消耗时间: ", time.Now().UnixNano()/1e6-s)
	//}
	return parseData, true
}

// SendExpressionNoParsed 发送指令，获取数据，但是不进行数据转换
func (o *ZmqObject) SendExpressionNoParsed(cmd string) ([]byte, bool) {

	var msg []byte
	//s := time.Now().UnixNano() / 1e6
	ctx := context.Background()
	msg, _ = allModelCache.HGet(ctx, *redisCacheKey, cmd).Bytes()
	if len(msg) == 0 || string(msg) == "null" {
		o.Lock()
		_ = o.Send(zmq4.NewMsgString(cmd))
		data, _ := o.Recv()
		msg = data.Bytes()
		o.Unlock()
		if cacheRefresh && len(msg) > 0 {
			allModelCache.HSet(ctx, *redisCacheKey, cmd, msg)
		}
	}
	if string(msg) == "\"\"\n" {
		return []byte{}, false
	}
	//msg = bytes.ReplaceAll(msg, []byte("\"\"\n"), []byte(""))
	msg = bytes.ReplaceAll(msg, []byte("\"false\""), []byte("false"))
	msg = bytes.ReplaceAll(msg, []byte("\"true\""), []byte("true"))
	if len(msg) == 0 {
		return nil, false
	}
	//if time.Now().UnixNano()/1e6-s > 10 {
	//	log.Println("cmd: ", cmd)
	//	log.Println("消耗时间: ", time.Now().UnixNano()/1e6-s)
	//}
	//log.Println("cmd", cmd)
	return msg, true
}

// BuildModel 编译模型为可执行文件
func (o *ZmqObject) BuildModel(className, fileNamePrefix string, simulateParametersData map[string]string) bool {
	cmd := className + ", fileNamePrefix = \"" + fileNamePrefix + "result\""
	for k, v := range simulateParametersData {
		if k != "" {
			cmd = cmd + "," + k + "=" + v
		}
	}
	cmd = "buildModel(" + cmd + ")"
	buildModelData, ok := o.SendExpressionNoParsed(cmd)
	buildModelData = bytes.TrimSuffix(buildModelData, []byte("\n"))
	buildModelData = bytes.ReplaceAll(buildModelData, []byte("\"false\""), []byte("false"))
	buildModelData = bytes.ReplaceAll(buildModelData, []byte("\"true\""), []byte("true"))
	if ok && string(buildModelData) != "{\"\",\"\"}" {
		return true
	}
	return false
}

// SetOptions 清空加载的模型库， 设置OMC的命令行选项
func (o *ZmqObject) SetOptions() {
	//o.SendExpressionNoParsed("clearCommandLineOptions()")
	o.SendExpressionNoParsed("clearMessages()")
	//o.SendExpressionNoParsed("clear()")
	//o.SendExpressionNoParsed("clearVariables()")
	//o.SendExpressionNoParsed("clearProgram()")
	//o.SendExpressionNoParsed("setCommandLineOptions(\"-d=nfAPI,execstat,rml,nfAPIDynamicSelect=false\")")
	//o.SendExpressionNoParsed("setCommandLineOptions(\"-d=initialization,NLSanalyticJacobian\")")
	o.SendExpressionNoParsed("setCommandLineOptions(\"+ignoreSimulationFlagsAnnotation=false\")")
	o.SendExpressionNoParsed("setCommandLineOptions(\"+ignoreCommandLineOptionsAnnotation=false\")")
	o.SendExpressionNoParsed("setCommandLineOptions(\"--simCodeTarget=C\")")
	//o.SendExpressionNoParsed("setCommandLineOptions(\"--matchingAlgorithm=PFPlusExt --indexReductionMethod=dynamicStateSelection\")")
	o.SendExpressionNoParsed("setModelicaPath(\"/usr/lib/omlibrary\")")
	o.SendExpressionNoParsed("setCommandLineOptions(\"--matchingAlgorithm=PFPlusExt \")")
	o.SendExpressionNoParsed("setCommandLineOptions(\"--indexReductionMethod=dynamicStateSelection\")")
	o.SendExpressionNoParsed("setCommandLineOptions(\"-d=initialization,NLSanalyticJacobian\")")
	//o.SendExpressionNoParsed("setCommandLineOptions(\"--NAPI=true\")")
	//o.SendExpressionNoParsed("setCompiler(\"clang\")")
	//o.SendExpressionNoParsed("setCXXCompiler(\"clang++\")")
}

func (o *ZmqObject) GetCommandLineOptions() string {
	cmd := "getCommandLineOptions()"
	data, ok := o.SendExpressionNoParsed(cmd)
	data = bytes.TrimSuffix(data, []byte("\n"))
	if ok {
		return string(data)
	}
	return ""
}

// CacheRefreshSet 改变缓存策略
func (o *ZmqObject) CacheRefreshSet(cache bool) {
	cacheRefresh = cache
}

// GetInheritedClasses 获取给定切片当中所有模型的继承项
func (o *ZmqObject) GetInheritedClasses(className string) []string {
	cmd := "getInheritedClasses(" + className + ")"
	inheritedClassesData, ok := o.SendExpression(cmd)
	var dataList []string
	if ok {
		for i := 0; i < len(inheritedClassesData); i++ {
			if inheritedClassesData[i].(string) != className {
				dataList = append(dataList, inheritedClassesData[i].(string))
			}
		}
	}
	return dataList
}

// GetInheritedClassesList 获取给定切片当中所有模型的继承项
func (o *ZmqObject) GetInheritedClassesList(classNameList []string) []string {
	var dataList []string
	for i := 0; i < len(classNameList); i++ {
		cmd := "getInheritedClasses(" + classNameList[i] + ")"
		inheritedClassesData, ok := o.SendExpression(cmd)
		if ok {
			for p := 0; p < len(inheritedClassesData); p++ {
				if inheritedClassesData[p].(string) == classNameList[i] {
					continue
				}
				dataList = append(dataList, inheritedClassesData[p].(string))
			}
		}
	}
	return dataList
}

// GetInheritedClassesListAll 获取给定切片当中所有模型的继承项，包含原始数据, 迭代到最顶层的继承
func (o *ZmqObject) GetInheritedClassesListAll(classNameList []string) []string {
	var dataList = classNameList
	var nameList = classNameList
	for {
		InheritedClassesData := o.GetInheritedClassesList(nameList)
		if len(InheritedClassesData) > 0 {
			dataList = append(dataList, InheritedClassesData...)
			nameList = InheritedClassesData
		} else {
			break
		}
	}
	return dataList
}

// GetDiagramAnnotationList 获取给定切片当中模型的本身视图数据，多个模型的数据一起返回
func (o *ZmqObject) GetDiagramAnnotationList(classNameList []string) []any {
	var dataList []any
	for i := 0; i < len(classNameList); i++ {
		cmd := "getDiagramAnnotation(" + classNameList[i] + ")"
		diagramAnnotationData, ok := o.SendExpression(cmd)
		if ok && len(diagramAnnotationData) > 8 {
			for di := 0; di < len(diagramAnnotationData); di++ {
				dataList = append(dataList, diagramAnnotationData[di])
			}
		}
	}
	return dataList
}

// GetDiagramAnnotation 获取模型的diagram注释信息
func (o *ZmqObject) GetDiagramAnnotation(className string) []any {
	cmd := "getDiagramAnnotation(" + className + ")"
	diagramAnnotationData, ok := o.SendExpression(cmd)
	if ok {
		return diagramAnnotationData
	}
	return nil
}

// GetConnectionCountList 获取切片给定模型当中有多少个连接线，一个数字，有多少个模型名称，就返回多少个数字
func (o *ZmqObject) GetConnectionCountList(classNameList []string) []int {
	var dataList []int
	for i := 0; i < len(classNameList); i++ {
		cmd := "getConnectionCount(" + classNameList[i] + ")"
		ConnectionCountNum, ok := o.SendExpressionNoParsed(cmd)
		ConnectionCountNum = bytes.ReplaceAll(ConnectionCountNum, []byte("\n"), []byte(""))
		if ok {
			num, _ := strconv.Atoi(string(ConnectionCountNum))
			dataList = append(dataList, num)
		}
	}
	return dataList
}

// GetNthConnection 获取给定模型与指定数字的连接线段其实位置与终点位置，返回接口的名称列表
func (o *ZmqObject) GetNthConnection(className string, num int) []string {
	var dataList []string
	cmd := "getNthConnection(" + className + "," + strconv.Itoa(num) + ")"
	ConnectionCountNum, ok := o.SendExpression(cmd)
	if ok {
		for i := 0; i < len(ConnectionCountNum); i++ {
			dataList = append(dataList, ConnectionCountNum[i].(string))
		}
	}
	return dataList
}

// GetNthConnectionAnnotation 获取模型与数字对应连接线的画图数据
func (o *ZmqObject) GetNthConnectionAnnotation(className string, num int) []any {
	var data []any
	cmd := "getNthConnectionAnnotation(" + className + "," + strconv.Itoa(num) + ")"
	NthConnectionAnnotationData, _ := o.SendExpression(cmd)
	data = append(data, NthConnectionAnnotationData...)
	return data
}

// GetComponents 获取给定模型的组成部分，包含组件信息，返回列表
func (o *ZmqObject) GetComponents(className string) []any {
	cmd := "getComponents(" + className + ")"
	components, _ := o.SendExpression(cmd)
	return components
}

// GetComponentsList 获取切片给定模型的组成部分，包含组件信息，返回二维列表
func (o *ZmqObject) GetComponentsList(classNameList []string) [][]any {
	var dataList [][]any
	for i := 0; i < len(classNameList); i++ {
		cmd := "getComponents(" + classNameList[i] + ")"
		components, ok := o.SendExpression(cmd)
		if ok {
			for p := 0; p < len(components); p++ {
				dataList = append(dataList, components[p].([]any))
			}
		}
	}
	return dataList
}

// GetDefaultComponentName 获取指定模型名称的默认组件名， 可能为空
func (o *ZmqObject) GetDefaultComponentName(className string) string {
	cmd := "getDefaultComponentName(" + className + ")"
	name, ok := o.SendExpressionNoParsed(cmd)
	if ok {
		name = name[1 : len(name)-2]
		return string(name)
	}
	return ""
}

// GetElements 获取给定模型的组成部分，包含组件信息,新API
func (o *ZmqObject) GetElements(className string) []any {
	if className == "Real" || className == "" {
		return nil
	}
	cmd := "getElements(" + className + ", useQuotes = true)"
	components, _ := o.SendExpressionNoParsed(cmd)
	components = bytes.ReplaceAll(components, []byte("{"), []byte("["))
	components = bytes.ReplaceAll(components, []byte("}"), []byte("]"))
	components = bytes.TrimSuffix(components, []byte("\n"))
	components = bytes.ReplaceAll(components, []byte("\n"), []byte("\\n"))
	re, _ := regexp.Compile(`"\$Any","\[.*"],\["`)
	rResultList := re.FindAll(components, -1)
	for _, r := range rResultList {
		noSuffixAndPrefix := bytes.TrimSuffix(r, []byte("\"],[\""))
		noSuffixAndPrefix = bytes.TrimPrefix(noSuffixAndPrefix, []byte("\"$Any\",\"["))
		noSuffixAndPrefix = bytes.Replace(noSuffixAndPrefix, []byte("\""), []byte("\\\""), -1)
		newStr := bytes.Join([][]byte{[]byte("\"$Any\",\"["), noSuffixAndPrefix, []byte("\"],[\"")}, []byte(""))
		components = bytes.Replace(components, r, newStr, -1)
	}
	var resData []any
	if string(components) != "Error" && len(components) > 0 {
		err := sonic.Unmarshal(components, &resData)
		if err != nil {
			log.Println("getElements err: ", err)
			log.Println("components: ", string(components))
			log.Println("cmd: ", cmd)
			return nil
		}
	}
	return resData
}

// GetElementsList 获取切片给定模型的组成部分，包含组件信息,新API
func (o *ZmqObject) GetElementsList(classNameList []string) [][]any {
	var dataList [][]any
	for i := 0; i < len(classNameList); i++ {
		if classNameList[i] == "Real" || classNameList[i] == "" {
			continue
		}
		cmd := "getElements(" + classNameList[i] + ", useQuotes = true)"
		components, ok := o.SendExpressionNoParsed(cmd)
		components = bytes.ReplaceAll(components, []byte("{"), []byte("["))
		components = bytes.ReplaceAll(components, []byte("}"), []byte("]"))
		components = bytes.TrimSuffix(components, []byte("\n"))
		re, _ := regexp.Compile(`"\$Any","\[.*"],\["`)
		rResultList := re.FindAll(components, -1)
		for _, r := range rResultList {
			noSuffixAndPrefix := bytes.TrimSuffix(r, []byte("\"],[\""))
			noSuffixAndPrefix = bytes.TrimPrefix(noSuffixAndPrefix, []byte("\"$Any\",\"["))
			noSuffixAndPrefix = bytes.Replace(noSuffixAndPrefix, []byte("\""), []byte("\\\""), -1)
			newStr := bytes.Join([][]byte{[]byte("\"$Any\",\"["), noSuffixAndPrefix, []byte("\"],[\"")}, []byte(""))
			components = bytes.Replace(components, r, newStr, -1)
		}
		var resData []any
		err := sonic.Unmarshal(components, &resData)
		if err != nil {
			log.Println("getElements err: ", err)
			log.Println("components: ", string(components))
			log.Println("cmd: ", cmd)
			return nil
		}
		if ok && len(resData) > 0 {
			for p := 0; p < len(resData); p++ {
				dataList = append(dataList, resData[p].([]any))
			}
		}
	}
	return dataList
}

// GetComponentAnnotationsList 获取切片给定模型的组件注释信息列表
func (o *ZmqObject) GetComponentAnnotationsList(classNameList []string) [][]any {
	var dataList [][]any
	for i := 0; i < len(classNameList); i++ {
		cmd := "getComponentAnnotations(" + classNameList[i] + ")"
		componentAnnotations, ok := o.SendExpression(cmd)
		if ok {
			for p := 0; p < len(componentAnnotations); p++ {
				dataList = append(dataList, componentAnnotations[p].([]any))
			}
		}
	}
	return dataList
}

// GetElementAnnotations 获取给定模型名字模型的组件注释信息,新API
func (o *ZmqObject) GetElementAnnotations(className string) []any {
	var componentAnnotations []any
	if className == "Real" {
		return componentAnnotations
	}
	cmd := "getElementAnnotations(" + className + ", useQuotes = true)"
	componentAnnotations, _ = o.SendExpression(cmd)
	//annotation(Placement(visible = true,transformation(origin = {16,-87},extent = {{-10,  -10},{10, 10}},rotation = 0,iconTransformation(origin = {168.94117431640626,-87.15294189453125},extent = {{-10,-10},{10,10}},rotation = 0))))
	//annotation(Placement(visible = true,transformation(origin = {-8, 40}, extent = {{-10, -10},{10, 10}}, rotation = 0), iconTransformation(origin = {-36, 20}, extent = {{-10, -10}, {10, 10}}, rotation = 0)));
	return componentAnnotations
}

// GetElementAnnotationsList 获取切片给定模型的组件注释信息列表,新API
func (o *ZmqObject) GetElementAnnotationsList(classNameList []string) [][]any {
	var dataList [][]any

	for i := 0; i < len(classNameList); i++ {
		if classNameList[i] == "Real" {
			continue
		}
		cmd := "getElementAnnotations(" + classNameList[i] + ", useQuotes = true)"
		componentAnnotations, ok := o.SendExpression(cmd)
		if ok {
			for p := 0; p < len(componentAnnotations); p++ {
				dataList = append(dataList, componentAnnotations[p].([]any))
			}
		}
	}
	return dataList
}

// GetIconAnnotation 获取给定模型的图标注释数据
func (o *ZmqObject) GetIconAnnotation(className string) []any {
	var dataList []any
	cmd := "getIconAnnotation(" + className + ")"
	iconAnnotationData, ok := o.SendExpression(cmd)
	if ok && len(iconAnnotationData) > 8 {
		return iconAnnotationData
	}
	return dataList
}

// GetIconAnnotationLineData 获取给定模型的图标数据
func (o *ZmqObject) GetIconAnnotationLineData(className string) []any {
	var dataList []any
	cmd := "getIconAnnotation(" + className + ")"
	iconAnnotationData, ok := o.SendExpression(cmd)
	if ok && len(iconAnnotationData) > 8 {
		dataList = iconAnnotationData
	}
	return dataList
}

// GetIconAnnotationList 获取给定切片模型的图标注释信息
func (o *ZmqObject) GetIconAnnotationList(classNameList []string) []any {
	var dataList []any
	for i := 0; i < len(classNameList); i++ {
		cmd := "getIconAnnotation(" + classNameList[i] + ")"
		iconAnnotationData, ok := o.SendExpression(cmd)
		if ok && len(iconAnnotationData) > 8 {
			data := iconAnnotationData[8]
			dataList = append(dataList, data.([]any)...)
		}
	}
	return dataList
}

// GetPackages 获取已加载的的包名称列表
func (o *ZmqObject) GetPackages() []string {
	var dataList []string
	cmd := "getClassNames()"

	classNamesData, _ := o.SendExpression(cmd)
	for i := 0; i < len(classNamesData); i++ {
		dataList = append(dataList, classNamesData[i].(string))
	}
	return dataList
}

// GetClassNames 获取给定模型包含的子节点， all参数为true时，递归查询所有子节点，返回切片形式
func (o *ZmqObject) GetClassNames(className string, all bool) []string {
	var dataList []string
	var cmd string
	if all == true {
		cmd = "getClassNames(" + className + ",true,true,false,false,false,false)"
	} else {
		cmd = "getClassNames(" + className + ",false,false,false,false,false,false)"
	}

	classNamesData, _ := o.SendExpression(cmd)
	for i := 0; i < len(classNamesData); i++ {
		dataList = append(dataList, classNamesData[i].(string))
	}
	return dataList
}

// ListFile 返回给定模型的文件源码
func (o *ZmqObject) ListFile(className string) string {
	code := ""
	cmd := "listFile(" + className + ",nestedClasses=true)"
	codeData, ok := o.SendExpressionNoParsed(cmd)
	if ok {
		code = string(codeData)
	}
	return code
}

// List 返回给定模型的源码
func (o *ZmqObject) List(className string) string {
	code := ""
	cmd := "list(" + className + ")"
	codeData, ok := o.SendExpressionNoParsed(cmd)
	if ok {
		code = string(codeData)
	}
	return code
}

// GetElementModifierNames 获取模型组件的修饰符名称
func (o *ZmqObject) GetElementModifierNames(className string, componentName string) []string {
	var modifierNames = []string{}
	cmd := "getElementModifierNames(" + className + ",\"" + componentName + "\")"
	data, ok := o.SendExpressionNoParsed(cmd)
	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))
	//data = bytes.ReplaceAll(data, []byte("\""), []byte(""))
	//data = bytes.ReplaceAll(data, []byte("\\"), []byte(""))
	if ok && len(data) > 0 {
		data = data[1 : len(data)-1]
		data = append([]byte{'['}, data...)
		data = append(data, ']')
		sonic.Unmarshal(data, &modifierNames)
	}
	return modifierNames
}

// GetElementModifierValue 获取模型组件对应修饰符名称的值
func (o *ZmqObject) GetElementModifierValue(className string, modifierName string) string {
	cmd := "getElementModifierValue(" + className + "," + modifierName + ")"
	data, _ := o.SendExpressionNoParsed(cmd)
	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))
	data = bytes.ReplaceAll(data, []byte("\""), []byte(""))
	data = bytes.ReplaceAll(data, []byte("\\"), []byte(""))
	return string(data)
}

// IsEnumeration 判断给定名称是否是枚举类型
func (o *ZmqObject) IsEnumeration(className string) bool {
	cmd := "isEnumeration(" + className + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	result = bytes.ReplaceAll(result, []byte("\""), []byte(""))
	if ok && string(result) == "true" {
		return true
	}
	return false
}

// GetEnumerationLiterals 获取枚举类型参数名的文字
func (o *ZmqObject) GetEnumerationLiterals(parameterName string) []string {
	var dataList []string
	cmd := "getEnumerationLiterals(" + parameterName + ")"
	data, _ := o.SendExpression(cmd)
	for i := 0; i < len(data); i++ {
		dataList = append(dataList, data[i].(string))
	}
	return dataList
}

// GetParameterValue 获取参数的默认值
func (o *ZmqObject) GetParameterValue(className string, modifierName string) string {
	cmd := "getParameterValue(" + className + ",\"" + modifierName + "\")"
	data, _ := o.SendExpressionNoParsed(cmd)
	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))
	data = bytes.ReplaceAll(data, []byte("\""), []byte(""))
	data = bytes.ReplaceAll(data, []byte("\\"), []byte(""))
	return string(data)
}

// GetDerivedClassModifierValue 获取模型派生类修饰符的值
func (o *ZmqObject) GetDerivedClassModifierValue(className string, modifierName string) string {
	cmd := "getDerivedClassModifierValue(" + className + "," + modifierName + ")"
	data, _ := o.SendExpressionNoParsed(cmd)
	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))
	data = bytes.ReplaceAll(data, []byte("\""), []byte(""))
	data = bytes.ReplaceAll(data, []byte("\\"), []byte(""))
	return string(data)
}

// GetDerivedClassModifierNames 获取模型派生类修饰符的值
func (o *ZmqObject) GetDerivedClassModifierNames(className string) []any {
	cmd := "getDerivedClassModifierNames(" + className + ")"
	data, _ := o.SendExpression(cmd)
	//data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))
	//data = bytes.ReplaceAll(data, []byte("\""), []byte(""))
	//data = bytes.ReplaceAll(data, []byte("\\"), []byte(""))
	return data
}

// GetExtendsModifierNames 获取模型继承类修饰符的名称列表
func (o *ZmqObject) GetExtendsModifierNames(classNameOne string, classNameTwo string) []string {
	var dataList []string
	cmd := "getExtendsModifierNames(" + classNameOne + "," + classNameTwo + ", useQuotes = true)"
	data, ok := o.SendExpression(cmd)
	if ok {
		for p := 0; p < len(data); p++ {
			dataList = append(dataList, data[p].(string))
		}
	}
	return dataList
}

// GetExtendsModifierValue 获取模型与被模型的修饰符值
func (o *ZmqObject) GetExtendsModifierValue(classNameOne string, classNameTwo string, name string) string {
	cmd := "getExtendsModifierValue(" + classNameOne + "," + classNameTwo + "," + name + ")"
	data, _ := o.SendExpressionNoParsed(cmd)
	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))
	data = bytes.ReplaceAll(data, []byte("\""), []byte(""))
	return string(data)
}

// IsExtendsModifierFinal 判断是否是继承的修饰符，且不可被继承者修改
func (o *ZmqObject) IsExtendsModifierFinal(classNameOne string, classNameTwo string, name string) string {
	cmd := "isExtendsModifierFinal(" + classNameOne + "," + classNameTwo + "," + name + ")"
	data, _ := o.SendExpressionNoParsed(cmd)
	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))
	data = bytes.ReplaceAll(data, []byte("\""), []byte(""))
	return string(data)
}

// SetComponentModifierValue 设置组件修饰符的值
func (o *ZmqObject) SetComponentModifierValue(className string, parameter string, value string) bool {
	code := "=" + value + ""
	if value == "" {
		code = "()"
	}
	cmd := "setComponentModifierValue(" + className + ", " + parameter + ", $Code(" + code + "))"
	data, ok := o.SendExpressionNoParsed(cmd)
	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))
	if ok && string(data) == "Ok" {
		return true
	}
	return false
}

// SetElementModifierValue 设置组件修饰符的值
func (o *ZmqObject) SetElementModifierValue(className string, parameter string, value string) bool {
	code := "=" + value + ""
	if strings.HasPrefix(value, "redeclare") {
		code = "(" + value + ")"
	}
	if value == "" {
		code = "()"
	}
	cmd := "setElementModifierValue(" + className + ", " + parameter + ", $Code(" + code + "))"
	data, ok := o.SendExpressionNoParsed(cmd)
	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))
	if ok && string(data) == "Ok" {
		return true
	}
	return false
}

// SetExtendsModifierValue  设置组件修饰符的值
func (o *ZmqObject) SetExtendsModifierValue(className, extendsName, parameter, value string) bool {
	// setExtendsModifierValue(test12345, Modelica.Blocks.Examples.PID_Controller, kinematicPTP.startTime, $Code(=10))
	value = strings.ReplaceAll(value, "\"", "\\\"")
	code := "=" + value + ""
	if strings.HasPrefix(value, "redeclare") {
		code = "=\"" + value + "\""
	}
	//if value == "" {
	//	code = "()"
	//}
	cmd := "setExtendsModifierValue(" + className + ", " + extendsName + ", " + parameter + ", $Code(" + code + "))"
	data, ok := o.SendExpressionNoParsed(cmd)
	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))
	if ok && string(data) == "Ok" {
		return true
	}
	return false
}

// RenameComponentInClass 重新修改模型组件的名称。 此函数不返回正确与错误，不报错均视为执行成功
func (o *ZmqObject) RenameComponentInClass(className string, oldComponentName string, newComponentName string) bool {
	cmd := "renameComponentInClass(" + className + "," + oldComponentName + ", " + newComponentName + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) != "Error" {
		return true
	}
	return false
}

// SetComponentProperties 设置模型组件的属性
func (o *ZmqObject) SetComponentProperties(className string, newComponentName string, final string, protected string, replaceable string, variability string, inner string, outer string, causality string) bool {
	// setComponentProperties(PID_Controller,PI,{true,false,true,false}, {""}, {false,false}, {""})
	cmdParameterList := []string{className, ",", newComponentName, ",{", final, ",false,", protected, ",", replaceable,
		"},{\"", variability, "\"}", ",{", inner, ",", outer, "},{\"", causality, "\"}"}
	cmd := "setComponentProperties(" + strings.Join(cmdParameterList, "") + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "Ok" {
		return true
	}
	return false
}

// ExistClass 判断模型名称是否已经存在
func (o *ZmqObject) ExistClass(className string) bool {
	cmd := "existClass(" + className + ")"
	result, _ := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if string(result) == "false" {
		return false
	}
	return true
}

// GetClassInformation 获取模型的基本信息
func (o *ZmqObject) GetClassInformation(className string) []any {
	cmd := "getClassInformation(" + className + ")"
	data, ok := o.SendExpression(cmd)
	if ok {
		return data
	}
	return nil
}

// CopyClass 复制模型
func (o *ZmqObject) CopyClass(className string, copiedClassName string, parentName string) bool {
	cmd := "copyClass(" + className + ",\"" + copiedClassName + "\"," + parentName + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	if ok && string(result) == "false" {
		return false
	}
	return true
}

// DeleteClass 删除模型
func (o *ZmqObject) DeleteClass(className string) bool {
	cmd := "deleteClass(" + className + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	if ok && string(result) == "false" {
		return false
	}
	return true
}

// AddComponent 新增模型组件
func (o *ZmqObject) AddComponent(newComponentName, oldComponentName, className, origin, rotation string, extent []string) bool {
	annotate := "annotate=Placement(visible=true, transformation=transformation(origin={" + origin + "}, extent={{" + extent[0] + "},{" + extent[1] + "}}, rotation=" + rotation + "))"
	cmd := "addComponent(" + newComponentName + "," + oldComponentName + "," + className + "," + annotate + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "false" {
		return false
	}
	return true
}

// AddInterfacesComponent 新增connector类型的模型组件
func (o *ZmqObject) AddInterfacesComponent(newComponentName, oldComponentName, className, origin, rotation string, extent []string) bool {
	// addComponent(Modelica.Blocks.Interfaces.RealInput,t,realinput,annotate=Placement(visible=true, transformation=transformation(origin={168.94117431640626,-87.15294189453125}, extent={{-10,-10},{10,10}}, rotation=0,iconTransformation=transformation(origin={168.94117431640626,-87.15294189453125}, extent={{-10,-10},{10,10}}, rotation=0))))
	//addComponent(y, Modelica.Blocks.Interfaces.RealVectorOutput,q,annotate=Placement(visible=true, transformation=transformation(origin={-36,-22}, extent={{-20,-20},{20,20}}, rotation=0), iconTransformation=transformation(origin={-36,-22}, extent={{-20,-20},{20,20}}, rotation=0))) 09:42:19:196
	annotate := "annotate=Placement(visible=true, transformation=transformation(origin={" + origin + "}, extent={{" + extent[0] + "},{" + extent[1] + "}}, rotation=" + rotation + "),iconTransformation=transformation(origin={" + origin + "}, extent={{" + extent[0] + "},{" + extent[1] + "}}, rotation=" + rotation + "))"
	cmd := "addComponent(" + newComponentName + "," + oldComponentName + "," + className + "," + annotate + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "false" {
		return false
	}
	return true
}

// DeleteComponent  删除模型组件
func (o *ZmqObject) DeleteComponent(componentName, className string) bool {
	cmd := "deleteComponent(" + componentName + "," + className + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "false" {
		return false
	}
	return true
}

// UpdateComponent  更新模型组件
func (o *ZmqObject) UpdateComponent(newComponentName, oldComponentName, className, origin, rotation string, extent []string) bool {
	annotate := "annotate=Placement(visible=true, transformation=transformation(origin={" + origin + "}, extent={{" + extent[0] + "},{" + extent[1] + "}}, rotation=" + rotation + "))"
	cmd := "updateComponent(" + newComponentName + "," + oldComponentName + "," + className + "," + annotate + ")"
	// updateComponent(cos,Modelica.Blocks.Math.Cos,test,annotate=Placement(visible=true, transformation=transformation(origin={4,-16}, extent={{-10,-10},{10,10}}, rotation=0)))
	// updateComponent(CriticalDamping,Modelica.Blocks.Continuous.Filter,test.Filter,annotate=Placement(visible=true, transformation=transformation(origin={34,0}, extent={{-20.0,40.0},{0.0,60.0}}, rotation=0)))
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "true" {
		return true
	}
	return false
}

// UpdateInterfacesComponent  更新connector类型的模型组件
func (o *ZmqObject) UpdateInterfacesComponent(newComponentName, oldComponentName, className, origin, rotation string, extent []string) bool {
	// updateComponent(y,Modelica.Blocks.Interfaces.RealVectorOutput,q,        annotate=Placement(visible=true, transformation=transformation(origin={-60,-20}, extent={{-20,-20},{20,20}}, rotation=0),         iconTransformation=transformation(origin={-36,-22}, extent={{-20,-20},{20,20}}, rotation=0)))
	// updateComponent(u2,Modelica.Blocks.Interfaces.IntegerInput,FullRobot_ng,annotate=Placement(visible=true, transformation=transformation(origin={84,-56}, extent={{-20.0,-20.0},{20.0,20.0}}, rotation=0.0),iconTransformation=transformation(origin={84,-56}, extent={{-20.0,-20.0},{20.0,20.0}}, rotation=0.0)))
	annotate := "annotate=Placement(visible=true, transformation=transformation(origin={" + origin + "}, extent={{" + extent[0] + "},{" + extent[1] + "}}, rotation=" + rotation + "),iconTransformation=transformation(origin={" + origin + "}, extent={{" + extent[0] + "},{" + extent[1] + "}}, rotation=" + rotation + "))"
	cmd := "updateComponent(" + newComponentName + "," + oldComponentName + "," + className + "," + annotate + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "true" {
		return true
	}
	return false
}

// AddConnection  新增模型组件之间的连线
func (o *ZmqObject) AddConnection(classNameAll, connectStart, connectEnd, color string, linePoints []string) bool {
	var linePointsList []string
	for _, point := range linePoints {
		linePointsList = append(linePointsList, "{"+point+"}")
	}
	points := strings.Join(linePointsList, ",")
	annotate := "annotate=Line(points={" + points + "},color={" + color + "}))"
	cmd := "addConnection(" + connectStart + "," + connectEnd + "," + classNameAll + "," + annotate
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "Ok" {
		return true
	}
	return false
}

// DeleteConnection  删除模型组件之间的连线
func (o *ZmqObject) DeleteConnection(classNameAll, connectStart, connectEnd string) bool {
	cmd := "deleteConnection(" + connectStart + "," + connectEnd + "," + classNameAll + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "Ok" {
		return true
	}
	return false
}

// UpdateConnectionNames  更新模型组件之间的连线的名称
func (o *ZmqObject) UpdateConnectionNames(classNameAll, fromName, toName, fromNameNew, toNameNew string) bool {
	cmd := "updateConnectionNames(\"" + classNameAll + "\",\"" + fromName + "\",\"" + toName + "\",\"" + fromNameNew + "\",\"" + toNameNew + "\")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "Ok" {
		return true
	}
	return false
}

// UpdateConnectionAnnotation  更新模型组件之间的连线相关属性
func (o *ZmqObject) UpdateConnectionAnnotation(classNameAll, connectStart, connectEnd, color string, linePoints []string) bool {
	var linePointsList []string
	for _, point := range linePoints {
		linePointsList = append(linePointsList, "{"+point+"}")
	}
	points := strings.Join(linePointsList, ",")
	annotate := "annotate=$annotation(Line(points={" + points + "},color={" + color + "}))"
	cmd := "updateConnectionAnnotation(" + classNameAll + ",\"" + connectStart + "\",\"" + connectEnd + "\",\"" + annotate + "\")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "true" {
		return true
	}
	return false
}

// CheckModel  检查模型
func (o *ZmqObject) CheckModel(className string) string {
	cmd := "checkModel(" + className + ")"
	data, ok := o.SendExpressionNoParsed(cmd)
	//data = bytes.TrimSuffix(data, []byte("\n"))
	if ok {
		return string(data[1 : len(data)-1])
	}
	return ""
}

// GetMessagesStringInternal  获取消息文字字符串，所有类型，包含正常消息，警告，错误
func (o *ZmqObject) GetMessagesStringInternal() string {
	cmd := "getMessagesStringInternal()"
	data, ok := o.SendExpressionNoParsed(cmd)
	if ok && len(data) > 3 {
		return string(data[1 : len(data)-1])
	}
	return ""
}

// GetDocumentationAnnotation 获取模型文档注释
func (o *ZmqObject) GetDocumentationAnnotation(className string) []string {
	var docList = []string{"", "", ""}
	cmd := "getDocumentationAnnotation(" + className + ")"
	data, ok := o.SendExpressionNoParsed(cmd)
	data = bytes.TrimSuffix(data, []byte("\n"))
	data = bytes.ReplaceAll(data, []byte("\r"), []byte("\\r"))
	data = bytes.ReplaceAll(data, []byte("\n"), []byte("\\n"))
	if ok && len(data) > 0 {
		data = data[1 : len(data)-1]
		data = append([]byte{'['}, data...)
		data = append(data, ']')
		var docData = []any{"", "", ""}
		err := sonic.Unmarshal(data, &docData)
		if err != nil {
			log.Println("docData: ", string(data))
			log.Println("err: ", err)
		}
		docList[0] = docData[0].(string)
		docList[1] = docData[1].(string)
		docList[2] = docData[2].(string)
	}
	return docList
}

// SetDocumentationAnnotation 更新模型文档注释
func (o *ZmqObject) SetDocumentationAnnotation(className, info, revisions string) bool {
	info = strings.ReplaceAll(info, "\"", "\\\"")
	cmd := "setDocumentationAnnotation(" + className + ",\"" + info + "\",\"" + revisions + "\")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "true" {
		return true
	}
	return false
}

// UriToFilename 将ModelicaUri转换成本机资源路径
func (o *ZmqObject) UriToFilename(uri string) string {
	cmd := "uriToFilename(\"" + uri + "\")"
	data, ok := o.SendExpressionNoParsed(cmd)
	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))
	data = bytes.ReplaceAll(data, []byte("\""), []byte(""))
	if ok {
		return string(data)
	}
	return ""
}

// ConvertUnits 单位转换
func (o *ZmqObject) ConvertUnits(s1, s2 string) []any {
	cmd := "convertUnits(\"" + s1 + "\",\"" + s2 + "\")"
	data, _ := o.SendExpression(cmd)
	return data
}

// GetAnnotationModifierValue 获取注释的变量值
func (o *ZmqObject) GetAnnotationModifierValue(className, vendorAnnotation, modifierName string) string {
	cmd := "getAnnotationModifierValue(" + className + ",\"" + vendorAnnotation + "\",\"" + modifierName + "\")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\""), []byte(""))
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) != "" {
		return string(result)
	}
	return ""
}

// GetSimulationOptions 获取模型仿真设置
func (o *ZmqObject) GetSimulationOptions(className string) []string {
	var dataList = []string{"", "", "", "", "", "", ""}
	cmd := "getSimulationOptions(" + className + ")"
	data, ok := o.SendExpression(cmd)
	if ok && len(data) > 4 {
		dataList[0] = data[0].(string)
		dataList[1] = data[1].(string)
		dataList[2] = data[2].(string)
		dataList[3] = data[3].(string)
		dataList[4] = data[4].(string)
	}
	// 获取求解其类型的注释
	solver := o.GetAnnotationModifierValue(className, "__OpenModelica_simulationFlags", "solver")
	//log.Println("solver:", solver)
	if strings.Contains(solver, "not Found") {
		solver = "dassl"
	}
	dataList[5] = solver
	// 获取仿真类型
	simulateType := o.GetAnnotationModifierValue(className, "simulate_type", "solver")
	if strings.Contains(simulateType, "not Found") {
		simulateType = "OM"
	}
	dataList[6] = simulateType
	return dataList
}

// AddClassAnnotation 新增模型注释
func (o *ZmqObject) AddClassAnnotation(className, annotate string) bool {
	cmd := "addClassAnnotation(" + className + ", annotate=" + annotate + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "true" {
		return true
	}
	return false
}

// ReadSimulationResult 读取仿真结果，给定需要读取的变量列表与结果文件路径，返回二维列表
func (o *ZmqObject) ReadSimulationResult(varNameList []string, path string) ([][]float64, bool) {
	// readSimulationResult("result_res.mat", {time,Bessel.a[1]}, 0)
	varNameStr := "time," + strings.Join(varNameList, ",")
	cmd := "readSimulationResult(\"" + path + "\",{" + varNameStr + "},0)"
	data, _ := o.SendExpressionNoParsed(cmd)
	data = bytes.ReplaceAll(data, []byte("{"), []byte("["))
	data = bytes.ReplaceAll(data, []byte("}"), []byte("]"))
	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))

	dataUnmarshal := make([][]float64, 0, 2)
	dataUnmarshal = append(dataUnmarshal, make([]float64, 0, 1))
	dataUnmarshal = append(dataUnmarshal, make([]float64, 0, 1))
	_ = sonic.Unmarshal(data, &dataUnmarshal)
	dataList := make([][]float64, 0, 2)
	dataList = append(dataList, make([]float64, 0, 1))
	dataList = append(dataList, make([]float64, 0, 1))
	d := dataUnmarshal[0]
	for index, _ := range d {
		// 时间和数据可能存在重复，循环将时间相同的部分移除
		if index != 0 && (dataUnmarshal[0][index] == dataUnmarshal[0][index-1]) {
			continue
		}
		dataList[0] = append(dataList[0], dataUnmarshal[0][index])
		dataList[1] = append(dataList[1], dataUnmarshal[1][index])
	}
	return dataList, true
}

// ParseFile 解析mo文件
func (o *ZmqObject) ParseFile(path string) (string, bool) {
	pwd, _ := os.Getwd()
	cmd := "parseFile(\"" + pwd + "/" + path + "\",\"UTF-8\")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && len(result) > 0 {
		result = result[1 : len(result)-1]
		return string(result), true
	}
	return "", false
}

// LoadFile 加载mo文件
func (o *ZmqObject) LoadFile(path string) bool {
	pwd, _ := os.Getwd()
	cmd := "loadFile(\"" + pwd + "/" + path + "\")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "true" {
		return true
	}
	return false
}

// LoadFileNoPwd LoadFile 加载mo文件
func (o *ZmqObject) LoadFileNoPwd(path string) bool {
	cmd := "loadFile(\"" + "/" + path + "\")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "true" {
		return true
	}
	return false
}

// ParseString 解析Modelica字符串
func (o *ZmqObject) ParseString(code, path string) (string, bool) {
	jsonCode, _ := sonic.Marshal(code)
	jsonCode = bytes.ReplaceAll(jsonCode, []byte("\\u003c"), []byte("<"))
	jsonCode = bytes.ReplaceAll(jsonCode, []byte("\\u003e"), []byte(">"))
	jsonCode = bytes.ReplaceAll(jsonCode, []byte("\\u0026"), []byte("&"))
	jsonCode = bytes.ReplaceAll(jsonCode, []byte("\\u2028"), []byte("U+2028"))
	jsonCode = bytes.ReplaceAll(jsonCode, []byte("\\u2029"), []byte("U+2029"))
	cmd := "parseString(" + string(jsonCode) + ",\"" + path + "\")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	resultStr := string(result)
	if ok && strings.HasPrefix(resultStr, "{") && strings.HasSuffix(resultStr, "}") {
		resultStr = resultStr[1 : len(resultStr)-1]
		return resultStr, true
	}
	return resultStr, false
}

// LoadString 加载Modelica字符串
func (o *ZmqObject) LoadString(code, path string) bool {
	pwd, _ := os.Getwd()
	jsonCode, _ := sonic.Marshal(code)
	jsonCode = bytes.ReplaceAll(jsonCode, []byte("\\u003c"), []byte("<"))
	jsonCode = bytes.ReplaceAll(jsonCode, []byte("\\u003e"), []byte(">"))
	jsonCode = bytes.ReplaceAll(jsonCode, []byte("\\u0026"), []byte("&"))
	jsonCode = bytes.ReplaceAll(jsonCode, []byte("\\u2028"), []byte("U+2028"))
	jsonCode = bytes.ReplaceAll(jsonCode, []byte("\\u2029"), []byte("U+2029"))
	cmd := "loadString(" + string(jsonCode) + ",\"" + pwd + "/" + path + "\",\"UTF-8\")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "true" {
		return true
	}
	return false
}

func (o *ZmqObject) LoadModel(modelName, version string) bool {

	cmd := fmt.Sprintf("loadModel(%s, {\"%s\"},true,\"\",false)", modelName, version)
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "true" {
		return true
	}
	return false
}

// CopyLoadString 加载复制后的Modelica字符串
func (o *ZmqObject) CopyLoadString(code, modelName string) bool {
	cmd := "loadString(" + code + ",\"" + modelName + "\",\"UTF-8\",false)"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "true" {
		return true
	}
	return false
}

// GetClassRestriction 获取给定模型名称的类型
func (o *ZmqObject) GetClassRestriction(className string) string {
	cmd := "getClassRestriction(" + className + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	result = bytes.ReplaceAll(result, []byte("\""), []byte(""))
	if ok && len(result) > 0 {
		return string(result)
	}
	return ""
}

// GetModelInstance 获取给定模型名称的实例化json数据字符串
func (o *ZmqObject) GetModelInstance(className string) string {
	cmd := "getModelInstance(" + className + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	result = result[1 : len(result)-1]
	if ok && len(result) > 0 {
		return string(result)
	}
	return ""
}

// Save 保存模型源码到文件，文件路径由omc查找
func (o *ZmqObject) Save(className string) bool {
	cmd := "save(" + className + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "true" {
		return true
	}
	return false
}

// SaveModel 保存模型源码到指定文件
func (o *ZmqObject) SaveModel(fileName, className string) bool {
	pwd, _ := os.Getwd()
	cmd := "saveModel(\"" + pwd + "/" + fileName + "\"," + className + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "true" {
		return true
	}
	return false
}

// GetAvailableLibraries 获取可用库
func (o *ZmqObject) GetAvailableLibraries() []any {
	cmd := "getAvailableLibraries()"
	result, ok := o.SendExpression(cmd)
	if ok {
		return result
	}
	return nil
}

// GetAvailableLibraryVersions 获取可用库版本
func (o *ZmqObject) GetAvailableLibraryVersions(packageName string) []string {
	cmd := fmt.Sprintf("getAvailableLibraryVersions(%s)", packageName)
	result, ok := o.SendExpression(cmd)
	if ok {
		var data []string
		for i := 0; i < len(result); i++ {
			data = append(data, result[i].(string))
		}
		return data
	}
	return nil
}

// GetLoadedLibraries 获取已加载库
func (o *ZmqObject) GetLoadedLibraries() []string {
	cmd := "getLoadedLibraries()"
	result, ok := o.SendExpression(cmd)
	if ok {
		var data []string
		for i := 0; i < len(result); i++ {
			data = append(data, result[i].([]any)[0].(string))
		}
		return data
	}
	return nil
}

// GetUses 获取库使用了哪些依赖
func (o *ZmqObject) GetUses(packageName string) [][]string {
	cmd := fmt.Sprintf("getUses(%s)", packageName)
	result, ok := o.SendExpression(cmd)
	if ok {
		var data [][]string
		for i := 0; i < len(result); i++ {
			var d []string
			for _, dd := range result[i].([]any) {
				d = append(d, dd.(string))
			}
			data = append(data, d)
		}
		return data
	}
	return nil
}

// GetAllSubtypeOf 获取模板数据
func (o *ZmqObject) GetAllSubtypeOf(baseClassName, className string) []any {
	cmd := "getAllSubtypeOf(" + baseClassName + "," + className + ",false,false,false)"
	result, ok := o.SendExpression(cmd)
	if ok && len(result) > 0 {
		return result
	}
	return make([]any, 0)
}

// GcSetMaxHeapSize 设置使用的最大内存上限
func (o *ZmqObject) GcSetMaxHeapSize(size string) []any {
	cmd := "GC_set_max_heap_size(" + size + ")"
	result, ok := o.SendExpression(cmd)
	if ok && len(result) > 0 {
		return result
	}
	return make([]any, 1)
}

// IsPackage 判断是否是包类型
func (o *ZmqObject) IsPackage(packageName string) bool {
	cmd := "isPackage(" + packageName + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "true" {
		return true
	}
	return false
}

// GetSourceFile 获取给定包或者模型的源文件
func (o *ZmqObject) GetSourceFile(packageName string) string {
	cmd := "getSourceFile(" + packageName + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && len(result) > 1 {
		return string(result)[1 : len(result)-1]
	}
	return ""
}

func (o *ZmqObject) SetSourceFile(packageName, path string) bool {
	cmd := "setSourceFile(" + packageName + ",\"" + path + "\")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "true" {
		return true
	}
	return false
}

// BuildModelFMU 构建FMU文件
func (o *ZmqObject) BuildModelFMU(className string, fmuFileNameId string) string {
	//fileNamePrefix := "/home/xuqingda/GolandProjects/YssimGoService/"
	//(Modelica.Blocks.Examples.PID_Controller,"2.0","me_cs","<default>",{"static"},false)
	//cmd := "buildModelFMU(" + className + ",\"2.0\",\"me_cs\",\"" + fmuFileNameId + "\",{\"static\"},false" + ")"
	cmd := "buildModelFMU(" + className + ",\"2.0\",\"me_cs\",\"" + fmuFileNameId + "\",{\"static\"},false" + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	if ok {
		return string(result)
	}
	return ""
}

// ModelInstance 将模型实例化后解析到给定的指针地址变量
func (o *ZmqObject) ModelInstance(modelName string, ModelInstance *serviceType.ModelInstance) bool {
	cmd := "getModelInstance(" + modelName + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	if ok && len(result) > 1 {
		result = bytes.ReplaceAll(result, []byte("\\\""), []byte("\""))
		result = bytes.ReplaceAll(result, []byte("\\\\"), []byte("\\"))
		result = result[1 : len(result)-2]
		err := sonic.Unmarshal(result, ModelInstance)

		if err != nil {
			log.Println(err)
		}
		return true
	}
	return false
}

// DumpXMLDAE  生成result_init.xml文件
func (o *ZmqObject) DumpXMLDAE(className string) []any {
	cmd := "dumpXMLDAE(" + className + ")"
	result, ok := o.SendExpression(cmd)
	if ok {
		return result
	}
	return make([]any, 2)

}

// AddComponentParameter 新建一个以varType为类型的组件
func (o *ZmqObject) AddComponentParameter(varName, varType, className, defaultValue string) bool {
	addComponentParameterCmd := "addComponent(" + varName + "," + varType + "," + className + ")"
	// addComponent(varName, varType,className)
	// 将className的varName组件设置成参数类型，通过修改属性实现
	// setComponentProperties(className,varName,{false,false,false,false}, {"parameter"}, {false,false}, {""})
	// 删除全局变量参数
	// deleteComponent(varName,className)
	resultAdd, ok := o.SendExpressionNoParsed(addComponentParameterCmd)
	resultAdd = bytes.ReplaceAll(resultAdd, []byte("\n"), []byte(""))
	if ok && string(resultAdd) == "true" {
		setComponentPropertiesCmd := "setComponentProperties(" + className + "," + varName + ",{false,false,false,false}, {\"parameter\"}, {false,false}, {\"\"})"
		result, ok := o.SendExpressionNoParsed(setComponentPropertiesCmd)
		result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
		if ok && string(result) == "Ok" {
			ok = o.SetComponentModifierValue(className, varName, defaultValue)
			if ok {
				return true
			}
		}
	}
	cmd := "deleteComponent(" + varName + "," + className + ")"
	_, _ = o.SendExpressionNoParsed(cmd)
	return false
}

// DeleteComponentParameter  删除组件参数
func (o *ZmqObject) DeleteComponentParameter(varName, className string) bool {
	cmd := "deleteComponent(" + varName + "," + className + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "false" {
		return false
	}
	return true
}

func (o *ZmqObject) GetIconAndDiagramAnnotations(classNameList []string, isIcon bool) []any {
	var data []any
	//ctx := context.Background()
	//var msg []byte
	for _, name := range classNameList {
		nType := o.GetClassRestriction(classNameList[len(classNameList)-1])
		//if nType != "connector" && nType != "expandable connector" {
		//	msg, _ = allModelCache.HGet(ctx, userName+"-yssim-componentGraphicsData", name).Bytes()
		//}
		//if len(msg) > 0 && string(msg) != "null" {
		//	var d []any
		//	err := sonic.Unmarshal(msg, &d)
		//	if err == nil && len(d) > 8 {
		//		data = append(data, d[8].([]any)...)
		//	}
		//}
		result := make([]any, 0)
		if (nType == "connector" || nType == "expandable connector") && !isIcon {
			result = o.GetDiagramAnnotation(name)
			if len(result) < 8 {
				result = o.GetIconAnnotationLineData(name)
			}
		} else {
			result = o.GetIconAnnotationLineData(name)
			//setData, _ := sonic.Marshal(result)
			//allModelCache.HSet(ctx, userName+"-yssim-componentGraphicsData", name, setData)
		}
		if len(result) > 8 {
			result = result[8].([]any)
			data = append(data, result...)
		}
	}
	return data
}
func (o *ZmqObject) GetCoordinateSystem(className string, isIcon bool) []any {

	nType := o.GetClassRestriction(className)

	result := make([]any, 0)
	if (nType == "connector" || nType == "expandable connector") && !isIcon {
		result = o.GetDiagramAnnotation(className)
		if len(result) < 8 {
			result = o.GetIconAnnotationLineData(className)
		}
	} else {
		result = o.GetIconAnnotationLineData(className)
	}

	if len(result) > 0 {
		return result[:8]
	}
	return nil
}

func (o *ZmqObject) GetIconAnnotations(className string) []any {
	var data []any
	//ctx := context.Background()
	//var msg []byte
	//msg, _ = allModelCache.HGet(ctx, userName+"-yssim-IconGraphicsData", className).Bytes()
	//if len(msg) > 0 && string(msg) != "null" {
	//	err := sonic.Unmarshal(msg, &data)
	//	if err != nil {
	//		log.Println("err", err)
	//		return nil
	//	}
	//	return data
	//}
	data = o.GetIconAnnotationLineData(className)
	//setData, _ := sonic.Marshal(data)
	//allModelCache.HSet(ctx, userName+"-yssim-IconGraphicsData", className, setData)
	return data
}
