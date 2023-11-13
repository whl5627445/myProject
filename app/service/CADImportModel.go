package service

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"yssim-go/app/DataType"
	"yssim-go/config"
	"yssim-go/library/fileOperation"
	"yssim-go/library/omc"

	"github.com/bytedance/sonic"

	"github.com/shopspring/decimal"
)

type root struct {
	XMLName xml.Name `xml:"Root"`
	Tube    []tube   `xml:"tube"`
	Lines   []lines  `xml:"lines"`
	Result  string   `xml:"result,attr"`
	Message string   `xml:"message,attr"`
}

type tube struct {
	XMLName    xml.Name `xml:"tube"`
	Partnumber string   `xml:"partnumber,attr"`
	Type       string   `xml:"type,attr"`
	Name       string   `xml:"name,attr"`
	Run        run      `xml:"run"`
	Solid      solid    `xml:"solid"`
}

type lines struct {
	XMLName xml.Name `xml:"lines"`
	Line    []line   `xml:"line"`
}

type line struct {
	XMLName xml.Name `xml:"line"`
	Start   start    `xml:"start"`
	End     end      `xml:"end"`
}

type start struct {
	XMLName xml.Name `xml:"start"`
	Name    string   `xml:"name,attr"`
	X       x        `xml:"x"`
	Y       y        `xml:"y"`
	Z       z        `xml:"z"`
}

type end struct {
	XMLName xml.Name `xml:"end"`
	Name    string   `xml:"name,attr"`
	X       x        `xml:"x"`
	Y       y        `xml:"y"`
	Z       z        `xml:"z"`
}

type solid struct {
	XMLName    xml.Name   `xml:"solid"`
	Attributes attributes `xml:"attributes"`
}

type run struct {
	XMLName  xml.Name `xml:"run"`
	Points   points   `xml:"points"`
	Segments segments `xml:"segments"`
}

type segments struct {
	XMLName xml.Name  `xml:"segments"`
	Segment []segment `xml:"segment"`
}

type segment struct {
	XMLName    xml.Name   `xml:"segment"`
	PointStart pointStart `xml:"point-start"`
	PointEnd   pointEnd   `xml:"point-end"`
}

type pointStart struct {
	XMLName xml.Name `xml:"point-start"`
	Point   point    `xml:"point"`
}

type pointEnd struct {
	XMLName xml.Name `xml:"point-end"`
	Point   point    `xml:"point"`
}

type points struct {
	XMLName xml.Name `xml:"points"`
	Point   []point  `xml:"point"`
}

type point struct {
	XMLName    xml.Name   `xml:"point"`
	X          x          `xml:"x"`
	Y          y          `xml:"y"`
	Z          z          `xml:"z"`
	Bendradius bendradius `xml:"bendradius"`
}

type x struct {
	XMLName xml.Name `xml:"x"`
	Value   float64  `xml:"value,attr"`
}

type y struct {
	XMLName xml.Name `xml:"y"`
	Value   float64  `xml:"value,attr"`
}

type z struct {
	XMLName xml.Name `xml:"z"`
	Value   float64  `xml:"value,attr"`
}

type bendradius struct {
	XMLName xml.Name `xml:"bendradius"`
	Value   string   `xml:"value,attr"`
}

type attributes struct {
	XMLName   xml.Name    `xml:"attributes"`
	Attribute []attribute `xml:"attribute"`
}

type attribute struct {
	XMLName   xml.Name `xml:"attribute"`
	AttrName  string   `xml:"attr-name,attr"`
	AttrValue string   `xml:"attr-value,attr"`
}

func CadFilesUpload(form *multipart.Form, userName string) []string {
	files := form.File["files"] // 获取名为 "files" 的文件数组
	var filePath []string
	for _, file := range files {
		currentDir, _ := os.Getwd()
		relativePath := "/static/UserFiles/CAD/" + userName + "/"
		src, err := file.Open()
		fileData, _ := io.ReadAll(src)
		//fileSavePath := resourcesDir(relativePath, userName)

		if err != nil {
			log.Printf("文件打开失败")
		}
		os.MkdirAll(currentDir+relativePath, os.ModePerm)
		dst, err := os.Create(currentDir + relativePath + file.Filename)
		if err != nil {
			// 处理错误
			log.Printf("文件创建失败")
		}
		_ = dst.Close()

		// 将上传的文件内容复制到目标文件
		_, err = io.Copy(dst, src)
		if err != nil {
			// 处理错误
			log.Printf("拷贝文件失败")
		}
		fileOperation.WriteFileByte(currentDir+relativePath+file.Filename, fileData)
		filePath = append(filePath, relativePath+file.Filename)
	}
	return filePath
}

func GetXmlData(files []string, userName string) (string, int) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	currentDir, _ := os.Getwd()
	for _, filePath := range files {
		fileWriter, err := bodyWriter.CreateFormFile("file", currentDir+filePath)
		if err != nil {
			log.Println("error writing to buffer", err)
		}
		// 打开文件并将内容复制到fileWriter
		file, err := os.Open(currentDir + filePath)
		if err != nil {
			log.Println("error opening file", err)
		}
		defer file.Close()
		_, err = io.Copy(fileWriter, file)
		if err != nil {
			log.Println("error copying file", err)
		}
	}

	// 完成multipart/form-data表单
	_ = bodyWriter.WriteField("url", userName+"/xml")
	_ = bodyWriter.Close()

	// 创建一个POST请求，并设置请求头和请求体
	req, err := http.NewRequest("POST", config.CADConnect+"/file/batch", bodyBuf)
	//req, err := http.NewRequest("POST", "http://121.37.183.103:8081/file/batch", bodyBuf)
	if err != nil {
		log.Println("error creating request", err)
	}

	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())

	// 发送HTTP请求
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error sending request", err)
	}
	defer resp.Body.Close()
	// 读取响应的内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取响应失败:", err)
		return "", 0
	}

	// 关闭响应的 Body
	defer resp.Body.Close()
	var data DataType.CadData
	_ = sonic.Unmarshal(body, &data)
	if data.Code != 200 {
		log.Println(data)
		return data.Msg, data.Code
	}
	return data.Data, data.Code
}

func HandleXMLUpload(from *multipart.FileHeader) string {
	open, err2 := from.Open()
	if err2 != nil {

	}
	defer open.Close()

	// 解析XML文件
	data, err := ioutil.ReadAll(open)
	if err != nil {
		fmt.Println("读取文件失败")
		return ""
	}
	var xmlData string
	if err := xml.Unmarshal(data, &xmlData); err != nil {
		fmt.Println("文件转换失败")
		return ""
	}

	return string(data)
}

func CADParseParts(path string) []map[string]any {

	v := root{}
	err := parseXML(path, &v)
	if err != nil {
		log.Printf("CADxml文件解析有误 err: %#v", err)
		return nil
	}
	//if v.Result != "success" {
	//	log.Printf("CADxml文件解析有误 err: 解析失败")
	//	return nil
	//}
	var parts []map[string]any
	for i := 0; i < len(v.Tube); i++ {
		t := v.Tube[i]
		var pList []map[string]any
		name := t.Name
		data := map[string]any{"partnumber": t.Partnumber, "type": t.Type, "name": t.Name}
		if t.Type == "CATTubBendableTube" {
			if len(t.Run.Points.Point) != len(t.Run.Segments.Segment)+1 {
				continue
			}
		}
		var lineList []map[string]any
		for _, ls := range v.Lines {
			for _, l := range ls.Line {
				startName := l.Start.Name
				endName := l.End.Name
				var lineMap map[string]any
				if strings.Contains(startName, name) || strings.Contains(endName, name) {
					startLine := map[string]any{"x": l.Start.X.Value, "y": l.Start.Y.Value}
					endLine := map[string]any{"x": l.End.X.Value, "y": l.End.Y.Value}
					lineMap = map[string]any{
						"start_line": startLine,
						"end_line":   endLine,
						"start_name": l.Start.Name,
						"end_name":   l.End.Name,
					}
					lineList = append(lineList, lineMap)
				}
			}
		}
		if t.Type == "CATTubTeeJunction" {

			data["connected_relation"] = lineList
			modelInformation := getModelInformationData(t, nil, nil)
			for _, p := range t.Run.Points.Point {
				pData := map[string]any{"x": p.X.Value, "y": p.Y.Value}
				pList = append(pList, pData)
			}
			data["points"] = pList
			data["model_information"] = modelInformation
		} else {
			diameter := getAttributes(t.Solid.Attributes)
			for _, p := range t.Run.Points.Point {
				pData := map[string]any{"x": p.X.Value, "y": p.Y.Value}
				pList = append(pList, pData)
			}
			data["points"] = pList
			sList := t.Run.Segments.Segment
			pipeData := getPipeData(diameter, sList)
			bendsData := getBendsData(pipeData)
			modelInformation := getModelInformationData(t, pipeData, bendsData)
			data["model_information"] = modelInformation
		}
		parts = append(parts, data)
	}
	return parts
}

func getAttributes(Attribute attributes) float64 {
	outsideDiameter := 0.0
	insulationThickness := 0.0

	for _, attr := range Attribute.Attribute {
		switch attr.AttrName {
		case "OutsideDiameter":
			outsideDiameterStr := strings.TrimSuffix(attr.AttrValue, "mm")
			outsideDiameter, _ = strconv.ParseFloat(outsideDiameterStr, 64)
		case "Insulation thickness":
			insulationThicknessStr := strings.TrimSuffix(attr.AttrValue, "mm")
			insulationThickness, _ = strconv.ParseFloat(insulationThicknessStr, 64)
		}
	}
	outsideDiameterDecimal := decimal.NewFromFloat(outsideDiameter)
	insulationThicknessDecimal := decimal.NewFromFloat(insulationThickness)
	diameter, _ := outsideDiameterDecimal.Sub(insulationThicknessDecimal).Float64()
	return diameter
}

func getPipeData(diameter float64, sList []segment) []map[string]any {
	pipeData := []map[string]any{}
	for s := 0; s < len(sList); s++ {
		sx := sList[s].PointStart.Point.X.Value
		sy := sList[s].PointStart.Point.Y.Value
		sz := sList[s].PointStart.Point.Z.Value
		ex := sList[s].PointEnd.Point.X.Value
		ey := sList[s].PointEnd.Point.Y.Value
		ez := sList[s].PointEnd.Point.Z.Value
		bendRadius := sList[s].PointEnd.Point.Bendradius.Value
		trim := strings.Trim(bendRadius, "mm")
		float, _ := strconv.ParseFloat(trim, 64)
		coordinate := map[string]float64{"x": ex - sx, "y": ey - sy, "z": ez - sz}
		length := math.Sqrt((ex-sx)*(ex-sx)+(ey-sy)*(ey-sy)+(ez-sz)*(ez-sz)) - float
		height := ez - sz
		pipeData = append(pipeData, map[string]any{"length": length / 1000, "height_ab": height / 1000, "diameter": diameter / 1000, "bend_radius": float, "coordinate": coordinate})
	}
	return pipeData
}

func getBendsData(pipeData []map[string]any) []map[string]any {
	bendsData := []map[string]any{}
	for p := 0; p < len(pipeData)-1; p++ {
		pipe := pipeData[p]
		nextPipe := pipeData[p+1]
		coordinate := pipe["coordinate"].(map[string]float64)
		nextCoordinate := nextPipe["coordinate"].(map[string]float64)
		pipeModulus := coordinate["x"]*coordinate["x"] + coordinate["y"]*coordinate["y"] + coordinate["z"]*coordinate["z"]
		nextPipeModulus := nextCoordinate["x"]*nextCoordinate["x"] + nextCoordinate["y"]*nextCoordinate["y"] + nextCoordinate["z"]*nextCoordinate["z"]
		delta := math.Acos((nextCoordinate["x"]*coordinate["x"] + nextCoordinate["y"]*coordinate["y"] + nextCoordinate["z"]*coordinate["z"]) / pipeModulus / nextPipeModulus)
		if math.IsNaN(delta) {
			delta = 0
		}
		R0 := pipe["bend_radius"]
		dHyd := pipe["diameter"]
		bendsData = append(bendsData, map[string]any{"delta": delta, "R_0": R0, "d_hyd": dHyd})
	}
	return bendsData
}

func getModelInformationData(tube tube, pipeData []map[string]any, bendsData []map[string]any) []map[string]any {
	modelInformation := make([]map[string]any, 0)
	if tube.Type == "CATTubBendableTube" {
		pointList := tube.Run.Points.Point
		for index, pointData := range pointList {
			xNum := pointData.X.Value
			yNum := pointData.Y.Value
			pData := map[string]any{"rotation": 0, "yz": 1, "xz": 1}
			pzData := map[string]any{"rotation": 0, "yz": 1, "xz": 1}
			pData["origin"] = []float64{xNum, yNum}

			if index == len(pointList)-1 {
				break
			}
			if index == 0 {
				if pointList[index+1].X.Value-xNum < 0 {
					pData["xz"] = -1
				}
				if pointList[index+1].Y.Value-yNum > 0 {
					pData["rotation"] = 90
				}
				if pointList[index+1].Y.Value-yNum < 0 {
					pData["rotation"] = -90
				}
				pData["geometry_data"] = pipeData[0]
				pipeData = pipeData[1:]
				modelInformation = append(modelInformation, pData)
				continue
			}
			x_1Num := pointList[index-1].X.Value
			y_1Num := pointList[index-1].Y.Value

			x1Num := pointList[index+1].X.Value
			y1Num := pointList[index+1].Y.Value

			if yNum-y_1Num > 0 && x1Num-xNum > 0 {
				pData["rotation"] = -90
				pData["xz"] = -1
				pzData["origin"] = []float64{x1Num - (x1Num-xNum)/2, yNum}
				pzData["port_a"] = []float64{x1Num - (x1Num-xNum)/2, yNum}
			}
			if yNum-y_1Num > 0 && x1Num-xNum < 0 {
				pData["rotation"] = 90
				pzData["xz"] = -1
				pzData["origin"] = []float64{x1Num - (x1Num-xNum)/2, yNum}
			}
			if yNum-y_1Num < 0 && x1Num-xNum > 0 {
				pData["rotation"] = -90
				pzData["origin"] = []float64{x1Num - (x1Num-xNum)/2, yNum}
			}
			if yNum-y_1Num < 0 && x1Num-xNum < 0 {
				pData["rotation"] = -90
				pData["yz"] = -1
				pzData["xz"] = -1
				pzData["origin"] = []float64{x1Num - (x1Num-xNum)/2, yNum}
			}
			if xNum-x_1Num > 0 && y1Num-yNum > 0 {
				pzData["rotation"] = 90
				pzData["origin"] = []float64{xNum, y1Num - (y1Num-yNum)/2}
			}
			if xNum-x_1Num < 0 && y1Num-yNum > 0 {
				pData["rotation"] = 180
				pData["yz"] = -1
				pzData["rotation"] = 90
				pzData["origin"] = []float64{xNum, y1Num - (y1Num-yNum)/2}
			}
			if xNum-x_1Num > 0 && y1Num-yNum < 0 {
				pData["yz"] = -1
				pzData["rotation"] = -90
				pzData["origin"] = []float64{xNum, y1Num - (y1Num-yNum)/2}
			}
			if xNum-x_1Num < 0 && y1Num-yNum < 0 {
				pData["rotation"] = 180
				pzData["rotation"] = -90
				pzData["origin"] = []float64{xNum, y1Num - (y1Num-yNum)/2}
			}

			if index == len(pointList)-2 {

				xLast := pointList[index+1].X.Value
				yLast := pointList[index+1].Y.Value
				pzData["origin"] = []float64{xLast, yLast}
			}
			pzData["geometry_data"] = pipeData[0]
			pipeData = pipeData[1:]
			pData["geometry_data"] = bendsData[0]
			bendsData = bendsData[1:]
			modelInformation = append(modelInformation, pData)
			modelInformation = append(modelInformation, pzData)
		}
	} else if tube.Type == "CATTubTeeJunction" {
		pointList := tube.Run.Points.Point
		for index, pointData := range pointList {
			xNum := pointData.X.Value
			yNum := pointData.Y.Value
			pData := map[string]any{"rotation": 0, "yz": 1, "xz": 1}
			pData["origin"] = []float64{xNum, yNum}
			if index == 0 {
				if pointList[index+1].X.Value-xNum < 0 {
					pData["xz"] = -1
				}
				if pointList[index+1].Y.Value-yNum > 0 {
					pData["rotation"] = 90
				}
				if pointList[index+1].Y.Value-yNum < 0 {
					pData["rotation"] = -90
				}
				modelInformation = append(modelInformation, pData)
				break
			}
		}
	} else {
		pointList := tube.Run.Points.Point
		for index, pointData := range pointList {
			xNum := pointData.X.Value
			yNum := pointData.Y.Value
			pData := map[string]any{"rotation": 0, "yz": 1, "xz": 1}
			pData["origin"] = []float64{xNum, yNum}
			pData["geometry_data"] = pipeData[0]
			if index == 0 {
				if pointList[index+1].X.Value-xNum < 0 {
					pData["xz"] = -1
				}
				if pointList[index+1].Y.Value-yNum > 0 {
					pData["rotation"] = 90
				}
				if pointList[index+1].Y.Value-yNum < 0 {
					pData["rotation"] = -90
				}
				modelInformation = append(modelInformation, pData)
				break
			}
		}
	}
	return modelInformation
}

func CADMappingModel(modelName string, classNameList []string, modelInformationList DataType.CadMappingModelInformation, componentsName map[string][]map[string]string) {
	var componentNames []map[string]string
	for index, className := range classNameList {
		component := map[string]string{}
		componentName := GetComponentName(modelName, className)
		modelInformation := modelInformationList.ModelInformation[index]
		originX := strconv.FormatFloat(modelInformation.OriginDiagram[0], 'f', -1, 64)
		originY := strconv.FormatFloat(modelInformation.OriginDiagram[1], 'f', -1, 64)
		origin := strings.Join([]string{originX, originY}, ",")
		rotation := strconv.FormatFloat(modelInformation.Rotation, 'f', -1, 64)
		extent := getExtents(className, modelInformation.Xz, modelInformation.Yz)
		AddComponent(componentName, className, modelName, origin, rotation, extent)
		if modelInformationList.Type == "CATTubBendableTube" {
			if (index+1)%2 == 1 {
				length := strconv.FormatFloat(modelInformation.GeometryData.Length, 'f', -1, 64)
				heightAb := strconv.FormatFloat(modelInformation.GeometryData.HeightAb, 'f', -1, 64)
				diameter := strconv.FormatFloat(modelInformation.GeometryData.Diameter, 'f', -1, 64)
				omc.OMC.SetElementModifierValue(modelName, componentName+".length", length)
				omc.OMC.SetElementModifierValue(modelName, componentName+".height_ab", heightAb)
				omc.OMC.SetElementModifierValue(modelName, componentName+".diameter", diameter)
			} else {
				DHyd := strconv.FormatFloat(modelInformation.GeometryData.DHyd, 'f', -1, 64)
				R0 := strconv.FormatFloat(modelInformation.GeometryData.R0, 'f', -1, 64)
				Delta := strconv.FormatFloat(modelInformation.GeometryData.Delta, 'f', -1, 64)
				parameters := "geometry(d_hyd=" + DHyd + ", R_0=" + R0 + ", delta=" + Delta + ")"
				omc.OMC.SetElementModifierValue(modelName, componentName+".geometry", parameters)
			}
		}
		component["name"] = componentName
		component["origin"] = origin
		componentNames = append(componentNames, component)
	}
	componentsName[modelInformationList.Name] = componentNames
	if len(classNameList) > 1 {
		for i := 0; i < len(componentNames)-1; i++ {
			startName := componentNames[i]["name"] + ".port_b"
			endName := componentNames[i+1]["name"] + ".port_a"
			startCoordinate := componentNames[i]["origin"]
			endCoordinate := componentNames[i+1]["origin"]
			AddConnection(modelName, startName, endName, "0,127,255", []string{startCoordinate, endCoordinate})
		}
	}
}

func ThreeWayManage(modelName string, componentsNames map[string][]map[string]string, connectedRelation []map[string]any) {
	for _, con := range connectedRelation {
		endDot := con["start_line"]
		startDot := con["end_line"]
		endMap := endDot.(map[string]interface{})
		startMap := startDot.(map[string]interface{})
		endX := endMap["x"]
		endY := endMap["y"]
		startX := startMap["x"]
		startY := startMap["y"]
		endOriginX := strconv.FormatFloat(endX.(float64), 'f', -1, 64)
		endOriginY := strconv.FormatFloat(endY.(float64), 'f', -1, 64)
		startOriginX := strconv.FormatFloat(startX.(float64), 'f', -1, 64)
		startOriginY := strconv.FormatFloat(startY.(float64), 'f', -1, 64)
		startLine := strings.Join([]string{endOriginX, endOriginY}, ",")
		endLine := strings.Join([]string{startOriginX, startOriginY}, ",")

		startXml := con["start_name"]
		endXml := con["end_name"]
		startSplit := strings.Split(startXml.(string), ".")
		endSplit := strings.Split(endXml.(string), ".")
		startSuffix := startSplit[len(startSplit)-1]
		endSuffix := endSplit[len(endSplit)-1]
		startList := componentsNames[startSplit[0]]
		endList := componentsNames[endSplit[0]]
		startCom := make(map[string]string)
		endCom := make(map[string]string)
		for _, m := range startList {
			startCom[m["name"]] = m["origin"]
		}

		for i, m := range endList {
			endCom[m["name"]+"_"+strconv.Itoa(i+1)] = m["origin"]
		}
		if len(startCom) == 1 && len(endCom) == 1 {
			var startComponentName string
			var endComponentName string
			for k, _ := range startCom {
				startComponentName = k
			}
			for k, _ := range endCom {
				endComponentName = k
			}
			split := strings.Split(endComponentName, "_")
			if strings.Contains(endComponentName, "pipe") && endSuffix == "port_1" {
				AddConnection(modelName, startComponentName+"."+startSuffix, split[0]+"."+"port_b", "0,127,255", []string{startLine, endLine})
			} else {
				AddConnection(modelName, startComponentName+"."+startSuffix, split[0]+"."+"port_a", "0,127,255", []string{startLine, endLine})
			}
		} else if len(endCom) != 1 {
			var startComponentName string
			for k, _ := range startCom {
				startComponentName = k
			}
			endComponentName := getKeyByValue(endCom, endLine)
			split := strings.Split(endComponentName, "_")
			index := split[len(split)-1]
			if index == strconv.Itoa(len(endCom)) {
				AddConnection(modelName, startComponentName+"."+startSuffix, split[0]+"."+"port_b", "0,127,255", []string{startLine, endLine})
			} else {
				AddConnection(modelName, startComponentName+"."+startSuffix, split[0]+"."+"port_a", "0,127,255", []string{startLine, endLine})
			}
		}
	}
}

func getKeyByValue(m map[string]string, value string) string {
	for key, val := range m {
		if val == value {
			return key
		}
	}
	return "" // 如果没有找到对应的键，则返回空字符串或者其他你认为合适的默认值
}

func getExtents(className string, Xz, Yz float64) []string {
	classICList := GetICList(className)
	coordinateSystem := getCoordinateSystemRecursion(classICList, false)
	extent1Diagram := coordinateSystem.Extent1Diagram
	extent2Diagram := coordinateSystem.Extent2Diagram
	//initialScale, _ := strconv.ParseFloat(coordinateSystem["initialScale"], 64)
	extent1 := parseFloatListAndCalculate(extent1Diagram, []float64{Xz, Yz})
	extent2 := parseFloatListAndCalculate(extent2Diagram, []float64{Xz, Yz})

	return []string{strings.Join(extent1, ","), strings.Join(extent2, ",")}
}

func parseFloatListAndCalculate(strList []float64, flip []float64) []string {
	data := []string{}
	for index, s := range strList {
		//f, _ := strconv.ParseFloat(s, 64)
		f := s * flip[index]
		fStr := strconv.FormatFloat(f, 'f', -1, 64)
		data = append(data, fStr)
	}
	return data
}

func parseXML(path string, obj any) error {

	if !filepath.IsAbs(path) {
		err := xml.Unmarshal([]byte(path), obj)
		if err != nil {
			return errors.New("反序列化错误：" + err.Error())
		}
	} else {
		data, err := os.ReadFile(path)
		if err != nil {
			return errors.New("读取消息错误错误：" + err.Error())
		}
		err = xml.Unmarshal(data, obj)
		if err != nil {
			return errors.New("反序列化错误：" + err.Error())
		}
	}

	return nil
}
