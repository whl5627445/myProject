package omc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
	"yssim-go/config"

	//"github.com/go-zeromq/zmq4"

	"github.com/go-zeromq/zmq4"
	"os"
	"strconv"
	"strings"
)

type ZmqObject struct {
	zmq4.Socket
	sync.Mutex
}

var cacheRefresh = false

//var AllModelCache = make(map[string][]byte, 1000)

var AllModelCache = config.R

// SendExpression 发送指令，获取数据
func (o *ZmqObject) SendExpression(cmd string) ([]interface{}, bool) {
	s := time.Now().UnixNano() / 1e6
	o.Lock()
	defer o.Unlock()
	var msg []byte

	//msg, ok := AllModelCache[cmd]
	ctx := context.Background()
	msg, err := AllModelCache.HGet(ctx, "yssim-GraphicsData", cmd).Bytes()
	//if !ok {
	if err != nil {
		_ = o.Send(zmq4.NewMsgString(cmd))

		data, _ := o.Recv()
		msg = data.Bytes()
		if cacheRefresh && len(msg) > 0 {
			//AllModelCache[cmd] = msg
			AllModelCache.HSet(ctx, "yssim-GraphicsData", cmd, msg)
		}
	}
	parseData, _ := DataToGo(msg)
	if len(parseData) == 0 {
		return nil, false
	}
	if time.Now().UnixNano()/1e6-s > 10 {
		fmt.Println("cmd: ", cmd)
		fmt.Println("消耗时间: ", time.Now().UnixNano()/1e6-s)
	}
	return parseData, true
}

// SendExpressionNoParsed 发送指令，获取数据，但是不进行数据转换
func (o *ZmqObject) SendExpressionNoParsed(cmd string) ([]byte, bool) {
	s := time.Now().UnixNano() / 1e6
	o.Lock()
	defer o.Unlock()
	var msg []byte
	//msg, ok := AllModelCache[cmd]
	ctx := context.Background()
	msg, err := AllModelCache.HGet(ctx, "yssim-GraphicsData", cmd).Bytes()
	//if !ok {
	if err != nil {
		_ = o.Send(zmq4.NewMsgString(cmd))
		data, _ := o.Recv()
		msg = data.Bytes()
		if cacheRefresh && len(msg) > 0 {
			//AllModelCache[cmd] = msg
			AllModelCache.HSet(ctx, "yssim-GraphicsData", cmd, msg)
		}
	}
	msg = bytes.ReplaceAll(msg, []byte("\"\"\n"), []byte(""))
	msg = bytes.ReplaceAll(msg, []byte("\"false\""), []byte("false"))
	msg = bytes.ReplaceAll(msg, []byte("\"true\""), []byte("true"))
	if len(msg) == 0 {
		return nil, false
	}
	if time.Now().UnixNano()/1e6-s > 10 {
		fmt.Println("cmd: ", cmd)
		fmt.Println("消耗时间: ", time.Now().UnixNano()/1e6-s)
	}
	return msg, true
}

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

// 清空加载的模型库
func (o *ZmqObject) Clear() {
	o.SendExpressionNoParsed("clearCommandLineOptions()")
	//o.SendExpressionNoParsed("clear()")
	//o.SendExpressionNoParsed("clearVariables()")
	//o.SendExpressionNoParsed("clearProgram()")
	o.SendExpressionNoParsed("setCommandLineOptions(\"+ignoreSimulationFlagsAnnotation=false\")")
	//o.SendExpressionNoParsed("setCommandLineOptions(\"-d=nfAPI\")")
	o.SendExpressionNoParsed("setCommandLineOptions(\"-d=nogen,noevalfunc,newInst,nfAPI\")")
	//o.SendExpressionNoParsed("setCommandLineOptions({\"-g=Modelica\",\"-d=nogen,noevalfunc,newInst,nfAPI\"})")
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

//改变缓存策略
func (o *ZmqObject) CacheRefreshSet(cache bool) {
	cacheRefresh = cache
}

// 获取给定切片当中所有模型的继承项，包含原始数据
func (o *ZmqObject) GetInheritedClassesList(classNameList []string) []string {
	var dataList []string
	for i := 0; i < len(classNameList); i++ {
		cmd := "getInheritedClasses(" + classNameList[i] + ")"
		InheritedclassesData, ok := o.SendExpression(cmd)
		if ok {
			for p := 0; p < len(InheritedclassesData); p++ {
				if InheritedclassesData[p].(string) == classNameList[i] {
					continue
				}
				dataList = append(dataList, InheritedclassesData[p].(string))
			}
		}
		//cmd := "getInheritedClasses(" + classNameList[i] + ")"
		//InheritedclassesData, ok := AllModelCache[cmd].([]interface{})
		//if !ok {
		//	InheritedclassesData, ok = o.SendExpression(cmd)
		//}
		//if ok {
		//	for p := 0; p < len(InheritedclassesData); p++ {
		//		dataList = append(dataList, InheritedclassesData[p].(string))
		//		AllModelCache[cmd] = InheritedclassesData
		//	}
		//
		//}
	}
	return dataList
}

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

// 获取给定切片当中模型的本身视图数据
func (o *ZmqObject) GetDiagramAnnotationList(classNameList []string) []interface{} {
	var dataList []interface{}
	for i := 0; i < len(classNameList); i++ {
		cmd := "getDiagramAnnotation(" + classNameList[i] + ")"
		//diagramannotationData, ok := AllModelCache[cmd].([]interface{})
		diagramannotationData, ok := o.SendExpression(cmd)
		//if !ok {
		//	diagramannotationData, ok = o.SendExpression(cmd)
		//}
		if ok && len(diagramannotationData) > 8 {
			for di := 0; di < len(diagramannotationData); di++ {
				dataList = append(dataList, diagramannotationData[di])
				//AllModelCache[cmd] = diagramannotationData
			}
		}
	}
	return dataList
}

// 获取切片给定模型当中有多少个连接线，一个数字
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

// 获取给定模型与指定数字的连接线段其实位置与终点位置，返回接口的名称
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

// 获取模型与数字对应连接线的画图数据
func (o *ZmqObject) GetNthConnectionAnnotation(className string, num int) []interface{} {
	var data []interface{}
	cmd := "getNthConnectionAnnotation(" + className + "," + strconv.Itoa(num) + ")"
	NthConnectionAnnotationData, _ := o.SendExpression(cmd)
	data = append(data, NthConnectionAnnotationData...)
	return data
}

// 获取给定模型的组成部分，包含组件信息
func (o *ZmqObject) GetComponents(className string) []interface{} {
	cmd := "getComponents(" + className + ")"
	components, _ := o.SendExpression(cmd)
	return components
}

// 获取切片给定模型的组成部分，包含组件信息
func (o *ZmqObject) GetComponentsList(classNameList []string) [][]interface{} {
	var dataList [][]interface{}
	for i := 0; i < len(classNameList); i++ {
		cmd := "getComponents(" + classNameList[i] + ")"
		components, ok := o.SendExpression(cmd)
		if ok {
			for p := 0; p < len(components); p++ {
				dataList = append(dataList, components[p].([]interface{}))
			}
		}

		//cmd := "getComponents(" + classNameList[i] + ")"
		//components, ok := AllModelCache[cmd].([]interface{})
		//if !ok {
		//
		//	components, ok = o.SendExpression(cmd)
		//}
		//if ok {
		//	for p := 0; p < len(components); p++ {
		//		dataList = append(dataList, components[p].([]interface{}))
		//		AllModelCache[cmd] = components
		//	}
		//}
	}
	return dataList
}

// 获取给定模型的组成部分，包含组件信息,新API
func (o *ZmqObject) GetElements(className string) []interface{} {
	if className == "Real" {
		return nil
	}
	cmd := "getElements(" + className + ", useQuotes = false)"
	components, _ := o.SendExpression(cmd)
	return components
}

// 获取切片给定模型的组成部分，包含组件信息,新API
func (o *ZmqObject) GetElementsList(classNameList []string) [][]interface{} {
	var dataList [][]interface{}
	for i := 0; i < len(classNameList); i++ {
		if classNameList[i] == "Real" {
			continue
		}
		cmd := "getElements(" + classNameList[i] + ", useQuotes = false)"
		components, ok := o.SendExpression(cmd)
		if ok {
			for p := 0; p < len(components); p++ {
				dataList = append(dataList, components[p].([]interface{}))
			}
		}
	}
	return dataList
}

// 获取切片给定模型的组件注释信息
func (o *ZmqObject) GetComponentAnnotationsList(classNameList []string) [][]interface{} {
	var dataList [][]interface{}
	for i := 0; i < len(classNameList); i++ {
		cmd := "getComponentAnnotations(" + classNameList[i] + ")"
		componentAnnotations, ok := o.SendExpression(cmd)
		if ok {
			for p := 0; p < len(componentAnnotations); p++ {
				dataList = append(dataList, componentAnnotations[p].([]interface{}))
			}
		}

		//cmd := "getComponentAnnotations(" + classNameList[i] + ")"
		//componentAnnotations, ok := AllModelCache[cmd].([]interface{})
		//if !ok {
		//
		//	componentAnnotations, ok = o.SendExpression(cmd)
		//
		//}
		//if ok {
		//	for p := 0; p < len(componentAnnotations); p++ {
		//		dataList = append(dataList, componentAnnotations[p].([]interface{}))
		//		AllModelCache[cmd] = componentAnnotations
		//	}
		//}
	}
	return dataList
}

// 获取切片给定模型的组件注释信息,新API
func (o *ZmqObject) GetElementAnnotationsList(classNameList []string) [][]interface{} {
	var dataList [][]interface{}

	for i := 0; i < len(classNameList); i++ {
		if classNameList[i] == "Real" {
			continue
		}
		cmd := "getElementAnnotations(" + classNameList[i] + ", useQuotes = false)"
		componentAnnotations, ok := o.SendExpression(cmd)
		if ok {
			for p := 0; p < len(componentAnnotations); p++ {
				dataList = append(dataList, componentAnnotations[p].([]interface{}))
			}
		}
	}
	return dataList
}

// 获取给定模型的图标数据
func (o *ZmqObject) GetIconAnnotation(className string) []interface{} {
	var dataList []interface{}
	cmd := "getIconAnnotation(" + className + ")"
	iconAnnotationData, ok := o.SendExpression(cmd)
	if ok {
		dataList = append(dataList, iconAnnotationData...)
	}

	//cmd := "getIconAnnotation(" + className + ")"
	//iconAnnotationData, ok := AllModelCache[cmd].([]interface{})
	//if !ok {
	//
	//	iconAnnotationData, ok = o.SendExpression(cmd)
	//}
	//if ok {
	//	dataList = append(dataList, iconAnnotationData...)
	//	AllModelCache[cmd] = iconAnnotationData
	//}
	return dataList
}

// 获取给定切片模型的图标注释信息
func (o *ZmqObject) GetIconAnnotationList(classNameList []string) []interface{} {
	var dataList []interface{}
	for i := 0; i < len(classNameList); i++ {
		cmd := "getIconAnnotation(" + classNameList[i] + ")"
		iconAnnotationData, ok := o.SendExpression(cmd)
		if ok && len(iconAnnotationData) > 8 {
			data := iconAnnotationData[8]
			dataList = append(dataList, data.([]interface{})...)
		}

		//cmd := "getIconAnnotation(" + classNameList[i] + ")"
		//iconAnnotationData, ok := AllModelCache[cmd].([]interface{})
		//if !ok {
		//
		//	iconAnnotationData, ok = o.SendExpression(cmd)
		//}
		//if ok && len(iconAnnotationData) > 8 {
		//	data := iconAnnotationData[8]
		//	dataList = append(dataList, data.([]interface{})...)
		//	AllModelCache[cmd] = iconAnnotationData
		//}
	}
	return dataList
}

func (o *ZmqObject) GetPackages() []string {
	var dataList []string
	cmd := "getClassNames()"

	classNamesData, _ := o.SendExpression(cmd)
	for i := 0; i < len(classNamesData); i++ {
		dataList = append(dataList, classNamesData[i].(string))
	}
	return dataList
}

// 获取给定模型包含的子节点， all参数为true时，递归查询所有子节点，返回切片形式
func (o *ZmqObject) GetClassNames(className string, all bool) []string {
	var dataList []string
	var cmd string
	if all == true {
		cmd = "getClassNames(" + className + ",true,true,false,false,true,false)"
	} else {
		cmd = "getClassNames(" + className + ",false,false,false,false,true,false)"
	}

	classNamesData, _ := o.SendExpression(cmd)
	for i := 0; i < len(classNamesData); i++ {
		dataList = append(dataList, classNamesData[i].(string))
	}
	return dataList
}

// 返回给定模型的源码
func (o *ZmqObject) List(className string) string {
	code := ""
	cmd := "list(" + className + ")"
	codeData, ok := o.SendExpressionNoParsed(cmd)
	if ok {
		code = string(codeData)
	}
	return code
}

//
func (o *ZmqObject) GetElementModifierValue(className string, modifierName string) string {
	cmd := "getElementModifierValue(" + className + "," + modifierName + ")"
	data, _ := o.SendExpressionNoParsed(cmd)
	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))
	data = bytes.ReplaceAll(data, []byte("\""), []byte(""))
	return string(data)
}

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

func (o *ZmqObject) GetEnumerationLiterals(parameterName string) []string {
	var dataList []string
	cmd := "getEnumerationLiterals(" + parameterName + ")"
	data, _ := o.SendExpression(cmd)
	for i := 0; i < len(data); i++ {
		dataList = append(dataList, data[i].(string))
	}
	return dataList
}

func (o *ZmqObject) GetParameterValue(className string, modifierName string) string {
	cmd := "getParameterValue(" + className + ",\"" + modifierName + "\")"
	data, _ := o.SendExpressionNoParsed(cmd)
	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))
	data = bytes.ReplaceAll(data, []byte("\""), []byte(""))
	data = bytes.ReplaceAll(data, []byte("\\"), []byte(""))
	return string(data)
}

func (o *ZmqObject) GetElementModifierNamesList(classNameList []string, componentName string) []string {
	var dataList []string
	for i := 0; i < len(classNameList); i++ {
		cmd := "getElementModifierNames(" + classNameList[i] + ",\"" + componentName + "\")"
		data, ok := o.SendExpression(cmd)
		if ok {
			for p := 0; p < len(data); p++ {
				dataList = append(dataList, data[p].(string))
			}
		}
	}
	return dataList
}

func (o *ZmqObject) GetDerivedClassModifierNames(className string) []interface{} {
	cmd := "getDerivedClassModifierNames(" + className + ")"
	data, _ := o.SendExpression(cmd)
	return data
}

func (o *ZmqObject) GetDerivedClassModifierValue(className string, modifierName string) string {
	cmd := "getDerivedClassModifierValue(" + className + "," + modifierName + ")"
	data, _ := o.SendExpressionNoParsed(cmd)
	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))
	data = bytes.ReplaceAll(data, []byte("\""), []byte(""))
	data = bytes.ReplaceAll(data, []byte("\\"), []byte(""))
	return string(data)
}

func (o *ZmqObject) GetExtendsModifierNames(classNameOne string, classNameTwo string) []string {
	var dataList []string
	cmd := "getExtendsModifierNames(" + classNameOne + "," + classNameTwo + ", useQuotes = false)"
	data, ok := o.SendExpression(cmd)
	if ok {
		for p := 0; p < len(data); p++ {
			dataList = append(dataList, data[p].(string))
		}
	}
	return dataList
}

func (o *ZmqObject) GetExtendsModifierValue(classNameOne string, classNameTwo string, name string) string {
	cmd := "getExtendsModifierValue(" + classNameOne + "," + classNameTwo + "," + name + ")"
	data, _ := o.SendExpressionNoParsed(cmd)
	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))
	data = bytes.ReplaceAll(data, []byte("\""), []byte(""))
	return string(data)
}

func (o *ZmqObject) IsExtendsModifierFinal(classNameOne string, classNameTwo string, name string) string {
	cmd := "isExtendsModifierFinal(" + classNameOne + "," + classNameTwo + "," + name + ")"
	data, _ := o.SendExpressionNoParsed(cmd)
	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))
	data = bytes.ReplaceAll(data, []byte("\""), []byte(""))
	return string(data)
}

func (o *ZmqObject) SetComponentModifierValue(className string, parameter string, value string) string {
	code := "=" + value + ""
	if value == "" {
		code = "()"
	}
	cmd := "setComponentModifierValue(" + className + ", " + parameter + ", $Code(" + code + "))"
	data, _ := o.SendExpressionNoParsed(cmd)
	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))
	return string(data)
}

func (o *ZmqObject) RenameComponentInClass(className string, oldComponentName string, newComponentName string) bool {
	cmd := "renameComponentInClass(" + className + "," + oldComponentName + ", " + newComponentName + ")"
	_, ok := o.SendExpression(cmd)
	if ok {
		return true
	}
	return false
}

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

func (o *ZmqObject) ExistClass(className string) bool {
	cmd := "existClass(" + className + ")"
	result, _ := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if string(result) == "false" {
		return false
	}
	return true
}

func (o *ZmqObject) GetClassInformation(className string) []interface{} {
	cmd := "getClassInformation(" + className + ")"
	data, ok := o.SendExpression(cmd)
	if ok {
		return data
	}
	return nil
}

func (o *ZmqObject) CopyClass(className string, copiedClassName string, parentName string) bool {
	cmd := "copyClass(" + className + ",\"" + copiedClassName + "\"," + parentName + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	if ok && string(result) == "false" {
		return false
	}
	return true
}

func (o *ZmqObject) DeleteClass(className string) bool {
	cmd := "deleteClass(" + className + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	if ok && string(result) == "false" {
		return false
	}
	return true
}

func (o *ZmqObject) AddComponent(className, newComponentName, oldComponentName, origin, rotation string, extent []string) bool {
	annotate := "annotate=Placement(visible=true, transformation=transformation(origin={" + origin + "}, extent={{" + extent[0] + "},{" + extent[1] + "}}, rotation=" + rotation + "))"
	cmd := "addComponent(" + newComponentName + "," + oldComponentName + "," + className + "," + annotate + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	if ok && string(result) == "false" {
		return false
	}
	return true
}

func (o *ZmqObject) DeleteComponent(componentName, className string) bool {
	cmd := "deleteComponent(" + componentName + "," + className + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	if ok && string(result) == "false" {
		return false
	}
	return true
}

func (o *ZmqObject) UpdateComponent(componentName, ComponentClassName, modelNameAll, origin, rotation string, extent []string) bool {
	// updateComponent(add3,Modelica.Blocks.Math.Add3,test,annotate=Placement(visible=true, transformation=transformation(origin={-68,24}, extent={{-10,-10},{10,10}}, rotation=0)))
	//updateComponent(tan1,Modelica.Blocks.Math.Tan,test,annotate=Placement(visible=true, transformation=transformation(origin={-,-}, extent={{120,-52},{140,-72}}, rotation=0.0)))
	annotate := "annotate=Placement(visible=true, transformation=transformation(origin={" + origin + "}, extent={{" + extent[0] + "},{" + extent[1] + "}}, rotation=" + rotation + "))"
	cmd := "updateComponent(" + componentName + "," + ComponentClassName + "," + modelNameAll + "," + annotate + ")"
	// updateComponent(cos,Modelica.Blocks.Math.Cos,test,annotate=Placement(visible=true, transformation=transformation(origin={4,-16}, extent={{-10,-10},{10,10}}, rotation=0)))
	// updateComponent(CriticalDamping,Modelica.Blocks.Continuous.Filter,test.Filter,annotate=Placement(visible=true, transformation=transformation(origin={34,0}, extent={{-20.0,40.0},{0.0,60.0}}, rotation=0)))
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "true" {
		return true
	}
	return false
}

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

func (o *ZmqObject) DeleteConnection(classNameAll, connectStart, connectEnd string) bool {
	cmd := "deleteConnection(" + connectStart + "," + connectEnd + "," + classNameAll + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "Ok" {
		return true
	}
	return false
}

func (o *ZmqObject) UpdateConnectionNames(classNameAll, fromName, toName, fromNameNew, toNameNew string) bool {
	cmd := "updateConnectionNames(\"" + classNameAll + "\",\"" + fromName + "\",\"" + toName + "\",\"" + fromNameNew + "\",\"" + toNameNew + "\")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "Ok" {
		return true
	}
	return false
}

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

func (o *ZmqObject) CheckModel(className string) string {
	cmd := "checkModel(" + className + ")"
	data, ok := o.SendExpressionNoParsed(cmd)
	//data = bytes.TrimSuffix(data, []byte("\n"))
	if ok {
		return string(data[1 : len(data)-1])
	}
	return ""
}

func (o *ZmqObject) GetMessagesStringInternal() string {
	cmd := "getMessagesStringInternal()"
	data, ok := o.SendExpressionNoParsed(cmd)
	if ok && len(data) > 3 {
		return string(data[1 : len(data)-1])
	}
	return ""
}

func (o *ZmqObject) GetDocumentationAnnotation(className string) []string {
	var docList = []string{"", "", ""}
	cmd := "getDocumentationAnnotation(" + className + ")"
	data, ok := o.SendExpressionNoParsed(cmd)
	data = bytes.TrimSuffix(data, []byte("\n"))
	data = bytes.ReplaceAll(data, []byte("\n"), []byte("\\n"))
	if ok && len(data) > 0 {
		data = data[1 : len(data)-1]
		data = append([]byte{'['}, data...)
		data = append(data, ']')
		var docData []interface{}
		_ = json.Unmarshal(data, &docData)
		docList[0] = docData[0].(string)
		docList[1] = docData[1].(string)
		docList[2] = docData[2].(string)
		return docList
	}
	return []string{"", "", ""}
}

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

func (o *ZmqObject) ConvertUnits(s1, s2 string) []interface{} {
	cmd := "convertUnits(\"" + s1 + "\",\"" + s2 + "\")"
	data, _ := o.SendExpression(cmd)
	return data
}

func (o *ZmqObject) GetSimulationOptions(className string) []string {
	var dataList = []string{"", "", "", "", ""}
	cmd := "getSimulationOptions(" + className + ")"
	data, ok := o.SendExpression(cmd)
	if ok && len(data) > 4 {
		dataList[0] = data[0].(string)
		dataList[1] = data[1].(string)
		dataList[2] = data[2].(string)
		dataList[3] = data[3].(string)
		dataList[4] = data[4].(string)
	}
	return dataList
}

func (o *ZmqObject) AddClassAnnotation(className, annotate string) bool {
	cmd := "addClassAnnotation(" + className + ", annotate=" + annotate + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "true" {
		return true
	}
	return false
}

func (o *ZmqObject) ReadSimulationResult(varNameList []string, path string) ([][]float64, bool) {
	// readSimulationResult("result_res.mat", {time,Bessel.a[1]}, 0)
	varNameStr := "time," + strings.Join(varNameList, ",")
	cmd := "readSimulationResult(\"" + path + "\",{" + varNameStr + "},0)"
	data, _ := o.SendExpressionNoParsed(cmd)
	data = bytes.ReplaceAll(data, []byte("{"), []byte("["))
	data = bytes.ReplaceAll(data, []byte("}"), []byte("]"))
	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))
	var dataList [][]float64
	err := json.Unmarshal(data, &dataList)
	if err != nil || len(dataList) == 0 {
		log.Println(err)
		return nil, false
	}
	return dataList, true

}

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

func (o *ZmqObject) LoadFile(path string) bool {
	pwd, _ := os.Getwd()
	cmd := "loadFile(\"" + pwd + "/" + path + "\",\"UTF-8\")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "true" {
		return true
	}
	return false
}

func (o *ZmqObject) ParseString(code, path string) (string, bool) {
	jsonCode, _ := json.Marshal(code)
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

func (o *ZmqObject) LoadString(code, path string) bool {
	pwd, _ := os.Getwd()
	jsonCode, _ := json.Marshal(code)
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

func (o *ZmqObject) SaveModel(fileName, className string) bool {
	cmd := "saveModel(\"" + fileName + "\"," + className + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if ok && string(result) == "true" {
		return true
	}
	return false
}
