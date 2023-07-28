package service

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
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
	"yssim-go/library/omc"

	"github.com/shopspring/decimal"
)

type root struct {
	XMLName xml.Name `xml:"Root"`
	Tube    []tube   `xml:"tube"`
	Result  string   `xml:"result,attr"`
	Message string   `xml:"message,attr"`
}

type tube struct {
	XMLName    xml.Name `xml:"tube"`
	Partnumber string   `xml:"partnumber,attr"`
	Type       string   `xml:"type,attr"`
	Run        run      `xml:"run"`
	Solid      solid    `xml:"solid"`
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
	XMLName    xml.Name   `xml:"point-start"`
	X          x          `xml:"x"`
	Y          y          `xml:"y"`
	Z          z          `xml:"z"`
	Bendradius bendradius `xml:"bendradius"`
}

type pointEnd struct {
	XMLName    xml.Name   `xml:"point-end"`
	X          x          `xml:"x"`
	Y          y          `xml:"y"`
	Z          z          `xml:"z"`
	Bendradius bendradius `xml:"bendradius"`
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
	Value   float64  `xml:"value,attr"`
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

func GetXmlData(form *multipart.Form, header string) string {

	files := form.File["files"] // 获取名为 "files" 的文件数组
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	for _, file := range files {
		currentDir, _ := os.Getwd()
		filePath := currentDir + "/static/UserFiles/CAD/"
		fileName := file.Filename
		fileWriter, _ := bodyWriter.CreateFormFile("file", filePath+fileName)

		// 打开文件并将内容复制到fileWriter
		fileCopy, _ := file.Open()
		_, _ = io.Copy(fileWriter, fileCopy)
		_ = fileCopy.Close()
		_ = os.RemoveAll(filePath)
	}
	_ = bodyWriter.WriteField("url", header+"/xml")
	// 完成multipart/form-data表单
	contentType := bodyWriter.FormDataContentType()
	err := bodyWriter.Close()

	// 创建一个POST请求，并设置请求头和请求体
	req, err := http.NewRequest("POST", "http://192.168.1.200:8081/file/batch", bodyBuf)
	if err != nil {
		fmt.Println("error creating request")
		return ""
	}

	// 设置请求头中的Content-Type字段
	req.Header.Set("Content-Type", contentType)

	// 发送HTTP请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error sending request")
		return ""
	}

	// 读取响应的内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return ""
	}

	// 关闭响应的 Body
	_ = resp.Body.Close()
	var data DataType.CadData
	_ = sonic.Unmarshal(body, &data)
	return data.Data
}

func CADParseParts(path string) []map[string]any {

	v := root{}
	err := parseXML(path, &v)
	if err != nil {
		log.Printf("CADxml文件解析有误 err: %#v", err)
		return nil
	}
	if v.Result != "success" {
		log.Printf("CADxml文件解析有误 err: 解析失败")
		return nil
	}
	var parts []map[string]any
	for i := 0; i < len(v.Tube); i++ {
		t := v.Tube[i]
		if t.Type == "CATTubBendableTube" {
			if len(t.Run.Points.Point) != len(t.Run.Segments.Segment)+1 {
				continue
			}
		}
		var pList []map[string]any
		data := map[string]any{"partnumber": t.Partnumber, "type": t.Type}
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
		sx := sList[s].PointStart.X.Value
		sy := sList[s].PointStart.Y.Value
		sz := sList[s].PointStart.Z.Value
		ex := sList[s].PointEnd.X.Value
		ey := sList[s].PointEnd.Y.Value
		ez := sList[s].PointEnd.Z.Value
		bendRadius := sList[s].PointEnd.Bendradius.Value
		coordinate := map[string]float64{"x": ex - sx, "y": ey - sy, "z": ez - sz}
		length := math.Sqrt((ex-sx)*(ex-sx)+(ey-sy)*(ey-sy)+(ez-sz)*(ez-sz)) - bendRadius
		height := ez - sz
		pipeData = append(pipeData, map[string]any{"length": length / 1000, "height_ab": height / 1000, "diameter": diameter / 1000, "bend_radius": bendRadius, "coordinate": coordinate})
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
				pData["yz"] = -1
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
			pzData["geometry_data"] = pipeData[0]
			pipeData = pipeData[1:]
			pData["geometry_data"] = bendsData[0]
			bendsData = bendsData[1:]
			modelInformation = append(modelInformation, pData)
			modelInformation = append(modelInformation, pzData)
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

func CADMappingModel(modelName string, classNameList []string, modelInformationList DataType.CadMappingModelInformation) {
	componentNames := []map[string]string{}
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
		component["name"] = componentName
		component["origin"] = origin
		componentNames = append(componentNames, component)
	}
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
