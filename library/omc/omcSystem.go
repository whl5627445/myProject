package omc

import (
	"bytes"
	"github.com/go-zeromq/zmq4"
	"strconv"
	"strings"
)

type omcZMQ struct {
	zmq4.Socket
}

var CacheRefresh = false
var AllModelCache = map[string]any{}

// SendExpression 发送指令，获取数据
func (o *omcZMQ) SendExpression(cmd string) ([]interface{}, bool) {
	//msg, ok := AllModelCache[cmd].(zmq4.Msg)
	//if !ok || CacheRefresh == true {
	//	_ = o.Send(zmq4.NewMsgString(cmd))
	//	msg, _ = o.Recv()
	//	AllModelCache[cmd] = msg
	//}
	_ = o.Send(zmq4.NewMsgString(cmd))
	msg, _ := o.Recv()
	data, _ := DataToGo(msg.Bytes())
	if len(data) == 0 {
		return nil, false
	}
	return data, true
}

// SendExpressionNoParsed 发送指令，获取数据，但是不进行数据转换
func (o *omcZMQ) SendExpressionNoParsed(cmd string) ([]byte, bool) {
	_ = o.Send(zmq4.NewMsgString(cmd))
	msg, _ := o.Recv()
	data := msg.Bytes()
	data = bytes.ReplaceAll(data, []byte("\"\"\n"), []byte(""))
	data = bytes.ReplaceAll(data, []byte("\"false\""), []byte("false"))
	data = bytes.ReplaceAll(data, []byte("\"true\""), []byte("true"))
	if len(data) == 0 {
		return nil, false
	}
	return data, true
}

// 改变缓存策略
func (o *omcZMQ) CacheRefreshSet(cache bool) {
	CacheRefresh = cache
}

// 获取给定切片当中所有模型的继承项，包含原始数据
func (o *omcZMQ) GetInheritedClassesList(classNameList []string) []string {
	var dataList []string
	for i := 0; i < len(classNameList); i++ {
		cmd := "getInheritedClasses(" + classNameList[i] + ")"
		InheritedclassesData, ok := o.SendExpression(cmd)
		if ok {
			for p := 0; p < len(InheritedclassesData); p++ {
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

func (o *omcZMQ) GetInheritedClassesListAll(classNameList []string) []string {
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
func (o *omcZMQ) GetDiagramAnnotationList(classNameList []string) []interface{} {
	var dataList []interface{}
	for i := 0; i < len(classNameList); i++ {
		cmd := "getDiagramAnnotation(" + classNameList[i] + ")"
		diagramannotationData, ok := AllModelCache[cmd].([]interface{})
		if !ok {
			diagramannotationData, ok = o.SendExpression(cmd)
		}
		if ok && len(diagramannotationData) > 8 {
			for di := 0; di < len(diagramannotationData); di++ {
				dataList = append(dataList, diagramannotationData[di])
				AllModelCache[cmd] = diagramannotationData
			}
		}
	}
	return dataList
}

// 获取切片给定模型当中有多少个连接线，一个数字
func (o *omcZMQ) GetConnectionCountList(classNameList []string) []int {
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
func (o *omcZMQ) GetNthConnection(className string, num int) []string {
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
func (o *omcZMQ) GetNthConnectionAnnotation(className string, num int) []interface{} {
	var data []interface{}
	cmd := "getNthConnectionAnnotation(" + className + "," + strconv.Itoa(num) + ")"
	NthConnectionAnnotationData, _ := o.SendExpression(cmd)
	data = append(data, NthConnectionAnnotationData...)
	return data
}

// 获取给定模型的组成部分，包含组件信息
func (o *omcZMQ) GetComponents(className string) []interface{} {
	cmd := "getComponents(" + className + ")"
	components, _ := o.SendExpression(cmd)
	return components
}

// 获取切片给定模型的组成部分，包含组件信息
func (o *omcZMQ) GetComponentsList(classNameList []string) [][]interface{} {
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
func (o *omcZMQ) GetElements(className string) []interface{} {
	cmd := "getElements(" + className + ", useQuotes = true)"
	components, _ := o.SendExpression(cmd)
	return components
}

// 获取切片给定模型的组成部分，包含组件信息,新API
func (o *omcZMQ) GetElementsList(classNameList []string) [][]interface{} {
	var dataList [][]interface{}
	for i := 0; i < len(classNameList); i++ {
		cmd := "getElements(" + classNameList[i] + ", useQuotes = true)"
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
func (o *omcZMQ) GetComponentAnnotationsList(classNameList []string) [][]interface{} {
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
func (o *omcZMQ) GetElementAnnotationsList(classNameList []string) [][]interface{} {
	var dataList [][]interface{}
	for i := 0; i < len(classNameList); i++ {
		cmd := "getElementAnnotations(" + classNameList[i] + ", useQuotes = true)"
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
func (o *omcZMQ) GetIconAnnotation(className string) []interface{} {
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
func (o *omcZMQ) GetIconAnnotationList(classNameList []string) []interface{} {
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

// 获取给定模型包含的子节点， all参数为true时，递归查询所有子节点，返回切片形式
func (o *omcZMQ) GetClassNames(className string, all bool) []string {
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
func (o *omcZMQ) List(className string) string {
	code := ""
	cmd := "list(" + className + ")"
	codeData, ok := o.SendExpressionNoParsed(cmd)
	if ok {
		code = string(codeData)
	}
	return code
}

func (o *omcZMQ) GetComponentModifierValue(className string, modifierName string) string {
	var data []byte
	cmd := "getComponentModifierValue(" + className + "," + modifierName + ")"
	data, _ = o.SendExpressionNoParsed(cmd)
	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))
	data = bytes.ReplaceAll(data, []byte("\""), []byte(""))
	return string(data)
}

func (o *omcZMQ) IsEnumeration(className string) string {
	var data []byte
	cmd := "isEnumeration(" + className + ")"
	data, _ = o.SendExpressionNoParsed(cmd)
	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))
	data = bytes.ReplaceAll(data, []byte("\""), []byte(""))
	return string(data)
}

func (o *omcZMQ) GetEnumerationLiterals(parameterName string) []string {
	var dataList []string
	cmd := "getEnumerationLiterals(" + parameterName + ")"
	data, _ := o.SendExpression(cmd)
	for i := 0; i < len(data); i++ {
		dataList = append(dataList, data[i].(string))
	}
	return dataList
}

func (o *omcZMQ) GetParameterValue(className string, modifierName string) string {
	var data []byte
	cmd := "getParameterValue(" + className + ",\"" + modifierName + "\")"
	data, _ = o.SendExpressionNoParsed(cmd)
	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))
	data = bytes.ReplaceAll(data, []byte("\""), []byte(""))
	data = bytes.ReplaceAll(data, []byte("\\"), []byte(""))
	return string(data)
}

func (o *omcZMQ) GetComponentModifierNamesList(classNameList []string, componentName string) []string {
	var dataList []string
	for i := 0; i < len(classNameList); i++ {
		cmd := "getComponentModifierNames(" + classNameList[i] + ",\"" + componentName + "\")"
		Data, ok := o.SendExpression(cmd)
		if ok {
			for p := 0; p < len(Data); p++ {
				dataList = append(dataList, Data[p].(string))
			}
		}
	}
	return dataList
}

func (o *omcZMQ) GetDerivedClassModifierNames(className string) []interface{} {
	cmd := "getDerivedClassModifierNames(" + className + ")"
	Data, _ := o.SendExpression(cmd)
	return Data
}

func (o *omcZMQ) GetDerivedClassModifierValue(className string, modifierName string) string {
	cmd := "getDerivedClassModifierValue(" + className + "," + modifierName + ")"
	Data, _ := o.SendExpressionNoParsed(cmd)
	Data = bytes.ReplaceAll(Data, []byte("\n"), []byte(""))
	Data = bytes.ReplaceAll(Data, []byte("\""), []byte(""))
	Data = bytes.ReplaceAll(Data, []byte("\\"), []byte(""))
	return string(Data)
}

func (o *omcZMQ) GetExtendsModifierNames(classNameOne string, classNameTwo string) []string {
	var dataList []string
	cmd := "getExtendsModifierNames(" + classNameOne + "," + classNameTwo + ", useQuotes = true)"
	Data, ok := o.SendExpression(cmd)
	if ok {
		for p := 0; p < len(Data); p++ {
			dataList = append(dataList, Data[p].(string))
		}
	}
	return dataList
}

func (o *omcZMQ) GetExtendsModifierValue(classNameOne string, classNameTwo string, name string) string {
	cmd := "getExtendsModifierValue(" + classNameOne + "," + classNameTwo + "," + name + ")"
	Data, _ := o.SendExpressionNoParsed(cmd)
	Data = bytes.ReplaceAll(Data, []byte("\n"), []byte(""))
	Data = bytes.ReplaceAll(Data, []byte("\""), []byte(""))
	return string(Data)
}

func (o *omcZMQ) IsExtendsModifierFinal(classNameOne string, classNameTwo string, name string) string {
	cmd := "isExtendsModifierFinal(" + classNameOne + "," + classNameTwo + "," + name + ")"
	Data, _ := o.SendExpressionNoParsed(cmd)
	Data = bytes.ReplaceAll(Data, []byte("\n"), []byte(""))
	Data = bytes.ReplaceAll(Data, []byte("\""), []byte(""))
	return string(Data)
}

func (o *omcZMQ) SetComponentModifierValue(className string, parameter string, value string) string {
	code := "=" + value + ""
	if value == "" {
		code = "()"
	}
	cmd := "setComponentModifierValue(" + className + ", " + parameter + ", $Code(" + code + "))"
	Data, _ := o.SendExpressionNoParsed(cmd)
	Data = bytes.ReplaceAll(Data, []byte("\n"), []byte(""))
	return string(Data)
}

func (o *omcZMQ) RenameComponentInClass(className string, oldComponentName string, newComponentName string) bool {
	cmd := "renameComponentInClass(" + className + "," + oldComponentName + ", " + newComponentName + ")"
	_, ok := o.SendExpression(cmd)
	if ok {
		return true
	}
	return false
}

func (o *omcZMQ) SetComponentProperties(className string, newComponentName string, final string, protected string, replaceable string, variability string, inner string, outer string, causality string) bool {
	// setComponentProperties(PID_Controller,PI,{true,false,true,false}, {""}, {false,false}, {""})
	cmdParameterList := []string{className, ",", newComponentName, ",{", final, ",false,", protected, ",", replaceable,
		"},{\"", variability, "\"}", ",{", inner, ",", outer, "},{\"", causality, "\"}"}
	cmd := "setComponentProperties(" + strings.Join(cmdParameterList, "") + ")"
	data, ok := o.SendExpressionNoParsed(cmd)
	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))
	if ok && string(data) == "Ok" {
		return true
	}
	return false
}

func (o *omcZMQ) ExistClass(className string) bool {
	cmd := "existClass(" + className + ")"
	result, _ := o.SendExpressionNoParsed(cmd)
	result = bytes.ReplaceAll(result, []byte("\n"), []byte(""))
	if string(result) == "false" {
		return false
	}
	return true
}

func (o *omcZMQ) GetClassInformation(className string) []interface{} {
	cmd := "getClassInformation(" + className + ")"
	result, ok := o.SendExpression(cmd)
	if ok {
		return result
	}
	return nil
}

func (o *omcZMQ) CopyClass(className string, copiedClassName string, parentName string) bool {
	cmd := "copyClass(" + className + ",\"" + copiedClassName + "\"," + parentName + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	if ok && string(result) == "false" {
		return false
	}
	return true
}

func (o *omcZMQ) DeleteClass(className string) bool {
	cmd := "deleteClass(" + className + ")"
	result, ok := o.SendExpressionNoParsed(cmd)
	if ok && string(result) == "false" {
		return false
	}
	return true
}
