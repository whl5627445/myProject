package main

func main() {

	//pprof.StartCPUProfile(os.Stdout)
	//defer pprof.StopCPUProfile()
	//for i := 0; i < 3000; i++ {
	//	inputReader := bufio.NewReader(os.Stdin)
	//	input, _ := inputReader.ReadString('\n') // 读取用户输入
	//	inputInfo := strings.Trim(input, "\n")
	//	if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
	//		return
	//	}
	//	//nameList := []string{"Modelica.Blocks.Examples.PID_Controller"}
	//	s := time.Now().UnixNano()
	//	//dataList := service.GetGraphicsData(nameList)
	//	dataList := service.GetGraphicsData(inputInfo)
	//	_, _ = json.Marshal(dataList)
	//	fmt.Println("总耗时： ", time.Now().UnixNano()/1e6-s/1e6)
	//
	//}

	//s := time.Now().UnixNano()
	//n := 0
	//for i := 0; i < 100000000; i++ {
	//	n += 1
	//}
	//fmt.Println(time.Now().UnixNano()/1e6 - s/1e6)
	//fmt.Println(n)
	// Applications.Examples.ElectricGrid
	// Modelica.Blocks.Examples.Filter
	// Modelica.Blocks.Examples.PID_Controller
	// Buildings.Applications.DataCenters.ChillerCooled.Examples.IntegratedPrimaryLoadSideEconomizer
	// b := "({\"BFSB\",\"DFSB\",\"MC21A\",\"PF\",\"PFPlus\",\"HK\",\"HKDW\",\"ABMP\",\"PR\",\"DFSBExt\",\"BFSBExt\",\"MC21AExt\",\"PFExt\",\"PFPlusExt\",\"HKExt\",\"HKDWExt\",\"ABMPExt\",\"PRExt\",\"BB\"},{\"Breadth First Search based algorithm.\",\"Depth First Search based algorithm.\",\"Depth First Search based algorithm with look ahead feature.\",\"Depth First Search based algorithm with look ahead feature.\",\"Depth First Search based algorithm with look ahead feature and fair row traversal.\",\"Combined BFS and DFS algorithm.\",\"Combined BFS and DFS algorithm.\",\"Combined BFS and DFS algorithm.\",\"Matching algorithm using push relabel mechanism.\",\"Depth First Search based Algorithm external c implementation.\",\"Breadth First Search based Algorithm external c implementation.\",\"Depth First Search based Algorithm with look ahead feature external c implementation.\",\"Depth First Search based Algorithm with look ahead feature external c implementation.\",\"Depth First Search based Algorithm with look ahead feature and fair row traversal external c implementation.\",\"Combined BFS and DFS algorithm external c implementation.\",\"Combined BFS and DFS algorithm external c implementation.\",\"Combined BFS and DFS algorithm external c implementation.\",\"Matching algorithm using push relabel mechanism external c implementation.\",\"BBs try.\"})"
	// b := "({\"none\",\"uode\",\"dynamicStateSelection\",\"dummyDerivatives\"},{\"Skip index reduction\",\"Use the underlying ODE without the constraints.\",\"Simple index reduction method, select (dynamic) dummy states based on analysis of the system.\",\"Simple index reduction method, select (static) dummy states based on heuristic.\"}) "
	// b := "{record OpenModelica.Scripting.ErrorMessage\n    info = record OpenModelica.Scripting.SourceInfo\n    filename = \"\",\n    readonly = false,\n    lineStart = 0,\n    columnStart = 0,\n    lineEnd = 0,\n    columnEnd = 0\nend OpenModelica.Scripting.SourceInfo;,\n    message = \"Automatically loaded package ModelicaServices 3.2.3 due to uses annotation.\",\n    kind = .OpenModelica.Scripting.ErrorKind.scripting,\n    level = .OpenModelica.Scripting.ErrorLevel.notification,\n    id = 223\nend OpenModelica.Scripting.ErrorMessage;,record OpenModelica.Scripting.ErrorMessage\n    info = record OpenModelica.Scripting.SourceInfo\n    filename = \"\",\n    readonly = false,\n    lineStart = 0,\n    columnStart = 0,\n    lineEnd = 0,\n    columnEnd = 0\nend OpenModelica.Scripting.SourceInfo;,\n    message = \"Automatically loaded package Complex 3.2.3 due to uses annotation.\",\n    kind = .OpenModelica.Scripting.ErrorKind.scripting,\n    level = .OpenModelica.Scripting.ErrorLevel.notification,\n    id = 223\nend OpenModelica.Scripting.ErrorMessage;}"
	// b := "{ModelicaReference,ModelicaServices,Complex,Modelica}"
	// b := "(\"package\",\"OpenModelica internal definitions and scripting functions\",false,false,true,\"D:/OpenModelica/lib/omc/NFModelicaBuiltin.mo\",false,961,1,5399,17,{},false,false,\"\",\"text\",false,\"\")"
	// b := "{Line(true, {0.0, 0.0}, 0, {{-39, 50}, {-22, 50}}, {0, 0, 127}, LinePattern.Solid, 0.25, {Arrow.None, Arrow.None}, 3, Smooth.None)}"
	// b := "{-,-,-,-,false,-,-,}"
	// b := "{{unassignedMessage=\"An electrical potential cannot be uniquely calculated.\\nThe reason could be that\\n- a ground object is missing (Modelica.Electrical.Analog.Basic.Ground)\\n  to define the zero potential of the electrical circuit,or\\n- a connector of an electrical component is not connected.\"},{unassignedMessage=\"An electrical current cannot be uniquely calculated.\\nThe reason could be that\\n- a ground object is missing (Modelica.Electrical.Analog.Basic.Ground)\\n  to define the zero potential of the electrical circuit,or\\n- a connector of an electrical component is not connected.\"}}"
	// b := "{{unassignedMessage=\"An electrical potential cannot be uniquely calculated.\nThe reason could be that\n- a ground object is missing (Modelica.Electrical.Analog.Basic.Ground)\n  to define the zero potential of the electrical circuit, or\n- a connector of an electrical component is not connected.\"},{unassignedMessage=\"An electrical current cannot be uniquely calculated.\nThe reason could be that\n- a ground object is missing (Modelica.Electrical.Analog.Basic.Ground)\n  to define the zero potential of the electrical circuit, or\n- a connector of an electrical component is not connected.\"}}"
	// fmt.Println(omc.DataToGo(b))

	//htmlStr := "<html>\n\n<p>\nThis is a simple drive train controlled by a PID controller:\n</p>\n\n<ul>\n<li> The two blocks \\\"kinematic_PTP\\\" and \\\"integrator\\\" are used to generate\n   the reference speed (= constant acceleration phase, constant speed phase,\n   constant deceleration phase until inertia is at rest). To check\n   whether the system starts in steady state, the reference speed is\n   zero until time = 0.5 s and then follows the sketched trajectory.</li>\n\n<li> The block \\\"PI\\\" is an instance of \\\"Blocks.Continuous.LimPID\\\" which is\n   a PID controller where several practical important aspects, such as\n   anti-windup-compensation has been added. In this case, the control block\n   is used as PI controller.</li>\n\n<li> The output of the controller is a torque that drives a motor inertia\n   \\\"inertia1\\\". Via a compliant spring/damper component, the load\n   inertia \\\"inertia2\\\" is attached. A constant external torque of 10 Nm\n   is acting on the load inertia.</li>\n</ul>\n\n<p>\nThe PI controller settings included \\\"limitAtInit=false\\\", in order that\nthe controller output limits of 12 Nm are removed from the initialization\nproblem.\n</p>\n\n<p>\nThe PI controller is initialized in steady state (initType=SteadyState)\nand the drive shall also be initialized in steady state.\nHowever, it is not possible to initialize \\\"inertia1\\\" in SteadyState, because\n\\\"der(inertia1.phi)=inertia1.w=0\\\" is an input to the PI controller that\ndefines that the derivative of the integrator state is zero (= the same\ncondition that was already defined by option SteadyState of the PI controller).\nFurthermore, one initial condition is missing, because the absolute position\nof inertia1 or inertia2 is not defined. The solution shown in this examples is\nto initialize the angle and the angular acceleration of \\\"inertia1\\\".\n</p>\n\n<p>\nIn the following figure, results of a typical simulation are shown:\n</p>\n\n<img src=\\\"modelica://Modelica/Resources/Images/Blocks/PID_controller.png\\\"\n   alt=\\\"PID_controller.png\\\"><br>\n\n<img src=\\\"modelica://Modelica/Resources/Images/Blocks/PID_controller2.png\\\"\n   alt=\\\"PID_controller2.png\\\">\n\n<p>\nIn the upper figure the reference speed (= integrator.y) and\nthe actual speed (= inertia1.w) are shown. As can be seen,\nthe system initializes in steady state, since no transients\nare present. The inertia follows the reference speed quite good\nuntil the end of the constant speed phase. Then there is a deviation.\nIn the lower figure the reason can be seen: The output of the\ncontroller (PI.y) is in its limits. The anti-windup compensation\nworks reasonably, since the input to the limiter (PI.limiter.u)\nis forced back to its limit after a transient phase.\n</p>\n\n</html>\""
	//htmlIo := strings.NewReader(htmlStr)
	//doc, err := goquery.NewDocumentFromReader(htmlIo)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//doc.Find("img").Each(func(i int, selection *goquery.Selection) {
	//	fmt.Println(selection.Attr("src"))
	//	//for _, node := range selection.Nodes {
	//	//	//fmt.Println(node.Attr)
	//	//	for _, attribute := range node.Attr {
	//	//		//fmt.Println(attribute.Val)
	//	//		attribute.Val = "test"
	//	//	}
	//	//	fmt.Println(node.Attr)
	//	//}
	//	selection.SetAttr("src", "test")
	//	fmt.Println(selection.Attr("src"))
	//})

}
