package omc

import (
	"fmt"
)

var omcInit, _ = Connect("127.0.0.1", "23456")

func init() {
	fmt.Println("初始化标准库...")
	//omcInit.SendExpression("GC_gcollect_and_unmap()")
	omcInit.SendExpression("clearCommandLineOptions()")
	omcInit.SendExpression("clear()")
	omcInit.SendExpression("clearVariables()")
	omcInit.SendExpression("clearProgram()")
	omcInit.SendExpression("setCommandLineOptions(\"+ignoreSimulationFlagsAnnotation=false\")")
	omcInit.SendExpression("setCommandLineOptions(\"-d=nogen,noevalfunc,newInst,nfAPI\")")
	omcInit.SendExpression("loadModel(Modelica, {\"3.2.3\"},true,\"\",false)")
	omcInit.SendExpression("loadModel(Complex, {\"3.2.3\"},true,\"\",false)")
	omcInit.SendExpression("loadModel(ModelicaServices, {\"3.2.3\"},true,\"\",false)")
	omcInit.SendExpression("loadModel(ModelicaReference, {\"\"},true,\"\",false)")
	omcInit.SendExpression("loadModel(Applications, {\"\"},true,\"\",false)")
	omcInit.SendExpression("loadModel(Buildings, {\"7.0.0\"},true,\"\",false)")
	fmt.Println("标准库初始化完成")
}

var OMC = *omcInit
