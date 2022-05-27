package service

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"yssim-go/config"
)

var omc = config.OMC

type GraphicsData struct {
	data [][]map[string]interface{}
}

func GetGraphicsData(nameList []string) [][]map[string]interface{} {
	var g = GraphicsData{}

	g.data = [][]map[string]interface{}{{}, {}}
	nameList = g.getICList(nameList)
	diagramAnnotationData := omc.GetDiagramAnnotationList(nameList)
	if len(diagramAnnotationData) >= 8 {
		dData := diagramAnnotationData[len(diagramAnnotationData)-1]
		data1 := g.data01(dData.([]interface{}))
		g.data[0] = append(g.data[0], data1...)
	}
	g.getnthconnectionData(nameList)
	componentsData := omc.GetComponentsList(nameList)
	componentannotationsData := omc.GetComponentAnnotationsList(nameList)
	data2 := g.data02(componentsData, componentannotationsData, false, "")
	g.data[1] = append(g.data[1], data2...)
	return g.data
}

func oneDimensionalProcessing(drawingData []interface{}) string {

	var data []string
	defer func() {
		if err := recover(); err != nil {
			data = []string{}
		}
	}()
	if len(drawingData) >= 1 && reflect.TypeOf(drawingData[0]).String() != "string" {
		return ""
	}
	for i := 0; i < len(drawingData); i++ {
		data = append(data, drawingData[i].(string))
	}

	return strings.Join(data, ",")
}

func twoDimensionalProcessing(drawingData []interface{}) []string {
	var data []string
	defer func() {
		if err := recover(); err != nil {
			data = []string{""}
		}
	}()
	for i := 0; i < len(drawingData); i++ {
		var piList []string
		for pi := 0; pi < len(drawingData[i].([]interface{})); pi++ {
			piList = append(piList, drawingData[i].([]interface{})[pi].(string))
		}
		data = append(data, strings.Join(piList, ","))
	}
	return data
}

func (g *GraphicsData) getICList(name []string) []string {
	dataList := name
	nameList := name
	for {
		InheritedClassesData, ok := omc.GetInheritedClassesList(nameList)
		if ok {
			dataList = append(dataList, InheritedClassesData...)
			nameList = InheritedClassesData
		} else {
			break
		}
	}
	return dataList
}

func (g *GraphicsData) data01(cData []interface{}) []map[string]interface{} {
	dataList := make([]map[string]interface{}, 0, 1)
	for i := 0; i < len(cData); i += 2 {
		data := map[string]interface{}{}
		drawingDataList := cData[i+1].([]interface{})
		if len(drawingDataList) < 9 {
			continue
		}
		dataType := cData[i]
		data["type"] = dataType
		data["visible"] = strconv.FormatBool(drawingDataList[0].(bool))
		data["originalPoint"] = oneDimensionalProcessing(drawingDataList[1].([]interface{}))
		data["rotation"] = drawingDataList[2]
		data["color"] = oneDimensionalProcessing(drawingDataList[3].([]interface{}))
		data["fillColor"] = oneDimensionalProcessing(drawingDataList[4].([]interface{}))
		data["linePattern"] = drawingDataList[5]
		data["fillPattern"] = drawingDataList[6]
		data["lineThickness"] = drawingDataList[7]
		switch dataType {
		case "Polygon":
			data["polygonPoints"] = twoDimensionalProcessing(drawingDataList[8].([]interface{}))
			data["smooth"] = drawingDataList[9]
		case "Line":
			delete(data, "fillColor")
			delete(data, "fillPattern")
			data["points"] = twoDimensionalProcessing(drawingDataList[3].([]interface{}))
			data["color"] = oneDimensionalProcessing(drawingDataList[4].([]interface{}))
			data["lineThickness"] = drawingDataList[6]
			data["arrow"] = oneDimensionalProcessing(drawingDataList[7].([]interface{}))
			data["arrowSize"] = drawingDataList[8]
			data["smooth"] = drawingDataList[9]
		case "Text":
			data["fillPattern"] = drawingDataList[6]
			data["extentsPoints"] = twoDimensionalProcessing(drawingDataList[8].([]interface{}))
			typeOriginalTextString, ok := drawingDataList[9].([]interface{})
			if ok {
				data["originalTextString"] = typeOriginalTextString[0]
			} else {
				data["originalTextString"] = drawingDataList[9]
			}
			data["fontSize"] = drawingDataList[10]
			data["textColor"] = oneDimensionalProcessing(drawingDataList[11].([]interface{}))
			data["fontName"] = drawingDataList[12]
			data["textStyles"] = drawingDataList[13]
			data["horizontalAlignment"] = drawingDataList[14]
		case "Rectangle":
			data["fillPattern"] = drawingDataList[6]
			data["borderPattern"] = drawingDataList[8]
			data["extentsPoints"] = twoDimensionalProcessing(drawingDataList[9].([]interface{}))
			data["radius"] = drawingDataList[10]
		case "Ellipse":
			data["fillPattern"] = drawingDataList[6]
			data["extentsPoints"] = twoDimensionalProcessing(drawingDataList[8].([]interface{}))
			data["startAngle"] = drawingDataList[9]
			data["endAngle"] = drawingDataList[10]
		}
		dataList = append(dataList, data)
	}
	return dataList
}

func (g *GraphicsData) data02(cData [][]interface{}, caData [][]interface{}, isIcon bool, parent string) []map[string]interface{} {

	dataList := make([]map[string]interface{}, 0, 1)
	var cDataFilter [][]interface{}
	var caDataFilter [][]interface{}
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("出错了: ", err)
	// 		fmt.Printf("cDataFilter: %s \n", cDataFilter)
	// 		fmt.Printf("caDataFilter: %s \n", caDataFilter)
	// 	}
	// }()
	if isIcon == true && cData != nil && caData != nil {
		for i := 0; i < len(cData); i++ {
			cDataSplit := strings.Split(cData[i][0].(string), ".")
			for ii := 0; ii < len(cDataSplit); ii++ {
				if "Interfaces" == cDataSplit[ii] {
					cDataFilter = append(cDataFilter, cData[i])
					caDataFilter = append(caDataFilter, caData[i])
				}
			}
		}
	} else {
		cDataFilter = cData
		caDataFilter = caData
	}

	if cDataFilter == nil || caDataFilter == nil {
		return dataList
	}
	for i := 0; i < len(cDataFilter); i++ {
		name := cDataFilter[i][0].(string)
		nameList := g.getICList([]string{name})
		placementIndex := func() int {
			for pi := 0; pi < len(caDataFilter[i]); pi++ {
				if caDataFilter[i][pi].(string) == "Placement" {
					return pi
				}
			}
			return -1
		}()
		if placementIndex != -1 {
			componentsData := omc.GetComponentsList(nameList)
			componentannotationsData := omc.GetComponentAnnotationsList(nameList)
			IconAnnotationData := omc.GetIconAnnotationList(nameList)
			caf := caDataFilter[i][placementIndex+1].([]interface{})
			rotateAngle := func() string {
				if caf[7] == "-" {
					return "0"
				} else {
					return caf[7].(string)
				}
			}()
			data := map[string]interface{}{"type": "Transformation"}
			data["graphType"] = func() string {
				for di := 0; di < len(cDataFilter); di++ {
					dList := strings.Split(cDataFilter[di][0].(string), ".")
					for dii := 0; dii < len(dList); dii++ {
						if dList[dii] == "Interfaces" {
							return "connecter"
						}
					}
				}
				return ""
			}()
			data["ID"] = strconv.Itoa(i)
			data["classname"] = cDataFilter[i][0]
			data["name"] = cDataFilter[i][1]
			data["original_name"] = cDataFilter[i][1]
			data["parent"] = parent
			data["visible"] = strconv.FormatBool(caf[0].(bool))
			data["rotateAngle"] = rotateAngle
			data["originDiagram"] = strings.Join([]string{caDataFilter[i][1].([]interface{})[1].(string), caDataFilter[i][1].([]interface{})[2].(string)}, ",")
			data["extent1Diagram"] = strings.Join([]string{caDataFilter[i][1].([]interface{})[3].(string), caDataFilter[i][1].([]interface{})[4].(string)}, ",")
			data["extent2Diagram"] = strings.Join([]string{caDataFilter[i][1].([]interface{})[5].(string), caDataFilter[i][1].([]interface{})[6].(string)}, ",")
			data["originDiagram"] = strings.Join([]string{caf[1].(string), caf[2].(string)}, ",")
			data["extent1Diagram"] = strings.Join([]string{caf[3].(string), caf[4].(string)}, ",")
			data["extent2Diagram"] = strings.Join([]string{caf[5].(string), caf[6].(string)}, ",")
			data["rotation"] = rotateAngle
			data["output_type"] = cDataFilter[i][len(cDataFilter[i])-1].([]interface{})[0]
			data["inputOutputs"] = g.data02(componentsData, componentannotationsData, true, data["name"].(string))
			data["subShapes"] = g.data01(IconAnnotationData)
			dataList = append(dataList, data)
		}
	}
	return dataList
}

func (g *GraphicsData) getnthconnectionData(nameList []string) {
	ConnectionCount := omc.GetConnectionCountList(nameList)
	for i := 0; i < len(ConnectionCount); i++ {
		for c := 0; c < ConnectionCount[i]; c++ {
			ncData := omc.GetNthConnection(nameList[i], c+1)
			ncaData := omc.GetNthConnectionAnnotation(nameList[i], c+1)
			daData := g.data01(ncaData)[0]
			daData["connectionfrom_original_name"] = ncData[0]
			daData["connectionto_original_name"] = ncData[1]
			re1, _ := regexp.Compile("[[0-9]+]$")
			re2, _ := regexp.Compile("[[0-9]+].")
			connectionfrom := re1.ReplaceAll([]byte(ncData[0]), []byte(""))
			connectionto := re1.ReplaceAll([]byte(ncData[1]), []byte(""))
			daData["connectionfrom"] = string(re2.ReplaceAll(connectionfrom, []byte(".")))
			daData["connectionto"] = string(re2.ReplaceAll(connectionto, []byte(".")))
			g.data[0] = append(g.data[0], daData)
		}
	}
}
