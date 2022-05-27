package omc

import (
	"strconv"

	"github.com/pebbe/zmq4"
)

type omcZMQ struct {
	*zmq4.Socket
}

func (o *omcZMQ) SendExpression(cmd string) ([]interface{}, bool) {
	// s := time.Now().UnixNano()
	_, _ = o.Send(cmd, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, false
	// }
	msg, _ := o.Recv(0)
	// fmt.Println("接收数据： ", msg)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, false
	// }
	// e := time.Now().UnixNano()/1e6 - s/1e6
	// if e > 2 {
	// 	fmt.Println("通信数据： ", msg)
	// 	fmt.Println("通信时长： ", e)
	// 	fmt.Println("通信命令： ", cmd)
	// 	fmt.Println("通信数据大小： ", len(msg))
	// 	fmt.Printf("数据占的字节数是 %d \n", unsafe.Sizeof(msg))
	// }
	data, _ := DataToGo(msg)
	// if err != nil || len(data) == 0 {
	if len(data) == 0 {
		return nil, false
	}
	return data, true
}

func (o *omcZMQ) GetInheritedClassesList(classNameList []string) ([]string, bool) {
	var dataList []string
	for i := 0; i < len(classNameList); i++ {
		cmd := "getInheritedClasses(" + classNameList[i] + ")"
		InheritedclassesData, ok := o.SendExpression(cmd)
		if ok {
			for p := 0; p < len(InheritedclassesData); p++ {
				dataList = append(dataList, InheritedclassesData[p].(string))
			}

		}
	}
	if len(dataList) == 0 {
		return dataList, false
	}
	return dataList, true
}

func (o *omcZMQ) GetDiagramAnnotationList(classNameList []string) []interface{} {
	var dataList []interface{}
	for i := 0; i < len(classNameList); i++ {
		cmd := "getDiagramAnnotation(" + classNameList[i] + ")"
		diagramannotationData, ok := o.SendExpression(cmd)
		if ok {
			for di := 0; di < len(diagramannotationData); di++ {
				dataList = append(dataList, diagramannotationData[di])
			}

		}
	}
	return dataList
}

func (o *omcZMQ) GetConnectionCountList(classNameList []string) []int {
	var dataList []int
	for i := 0; i < len(classNameList); i++ {
		cmd := "getConnectionCount(" + classNameList[i] + ")"
		ConnectionCountNum, ok := o.SendExpression(cmd)
		if ok {
			for p := 0; p < len(ConnectionCountNum); p++ {
				dataList = append(dataList, int(ConnectionCountNum[p].(float64)))
			}
		}
	}
	return dataList
}

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

func (o *omcZMQ) GetNthConnectionAnnotation(className string, num int) []interface{} {
	var data []interface{}
	cmd := "getNthConnectionAnnotation(" + className + "," + strconv.Itoa(num) + ")"
	NthConnectionAnnotationData, _ := o.SendExpression(cmd)
	data = append(data, NthConnectionAnnotationData...)
	return data
}

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
	}
	return dataList
}

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
	}
	return dataList
}

func (o *omcZMQ) GetIconAnnotationList(classNameList []string) []interface{} {
	var dataList []interface{}
	for i := 0; i < len(classNameList); i++ {
		cmd := "getIconAnnotation(" + classNameList[i] + ")"
		iconAnnotationData, ok := o.SendExpression(cmd)
		if ok && len(iconAnnotationData) > 8 {
			data := iconAnnotationData[8]
			// for p := 0; p < len(data); p++ {
			dataList = append(dataList, data.([]interface{})...)
			// }
		}
	}
	return dataList
}
