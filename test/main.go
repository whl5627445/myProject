package main

func main() {

	//pprof.StartCPUProfile(os.Stdout)
	//defer pprof.StopCPUProfile()
	//for i := 0; i < 30; i++ {
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

}
