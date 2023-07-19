package service

import (
	"encoding/xml"
	"errors"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/bytedance/sonic"
)

type root struct {
	XMLName xml.Name `xml:"Root"`
	Tube    []tube   `xml:"tube"`
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

func CADParseParts(path string) []map[string]any {

	v := root{}
	err := parseXML(path, &v)
	if err != nil {
		log.Printf("CADxml文件解析有误 err: %#v", err)
		return nil
	}

	var parts []map[string]any

	for i := 0; i < len(v.Tube); i++ {
		t := v.Tube[i]
		var pList []map[string]any
		pipeData := make([]map[string]any, 0)
		bendsData := make([]map[string]any, 0)
		data := map[string]any{"partnumber": t.Partnumber, "type": t.Type}
		outsideDiameter := 0.0
		insulationThickness := 0.0
		diameter := 0.0
		for _, attr := range t.Solid.Attributes.Attribute {
			switch attr.AttrName {
			case "OutsideDiameter":
				outsideDiameter, _ = strconv.ParseFloat(attr.AttrValue[:len(attr.AttrValue)-2], 64)
			case "Insulation thickness":
				insulationThickness, _ = strconv.ParseFloat(attr.AttrValue[:len(attr.AttrValue)-2], 64)
			}
		}
		outsideDiameter = outsideDiameter / 1000
		insulationThickness = insulationThickness / 1000
		diameter = outsideDiameter - insulationThickness
		for _, p := range t.Run.Points.Point {
			pData := map[string]any{"x": p.X.Value, "y": p.Y.Value}
			pList = append(pList, pData)
		}
		data["points"] = pList
		parts = append(parts, data)
		sList := t.Run.Segments.Segment

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

			pipeData = append(pipeData, map[string]any{"length": length, "height_ab": height, "diameter": diameter, "bend_radius": bendRadius, "coordinate": coordinate})

		}
		for p := 0; p < len(pipeData)-1; p++ {
			pipe := pipeData[p]
			nextPipe := pipeData[p+1]
			coordinate := pipe["coordinate"].(map[string]float64)
			nextCoordinate := nextPipe["coordinate"].(map[string]float64)
			pipeModulus := coordinate["x"]*coordinate["x"] + coordinate["y"]*coordinate["y"] + coordinate["z"]*coordinate["z"]
			nextPipeModulus := nextCoordinate["x"]*nextCoordinate["x"] + nextCoordinate["y"]*nextCoordinate["y"] + nextCoordinate["z"]*nextCoordinate["z"]
			delta := math.Acos((nextCoordinate["x"]*coordinate["x"] + nextCoordinate["y"]*coordinate["y"] + nextCoordinate["z"]*coordinate["z"]) / pipeModulus / nextPipeModulus)
			R_0 := pipe["bend_radius"]
			d_hyd := pipe["diameter"]
			bendsData = append(bendsData, map[string]any{"delta": delta, "R_0": R_0, "d_hyd": d_hyd})
		}

		modelInformation := make([]map[string]any, 0)
		if t.Type == "CATTubBendableTube" {
			pointList := t.Run.Points.Point
			for index, pointData := range pointList {
				xNum := pointData.X.Value
				yNum := pointData.Y.Value
				pData := map[string]any{"rotation": 0, "yz": 1, "xz": 1}
				pzData := map[string]any{"rotation": 0, "yz": 1, "xz": 1}
				pData["originDiagram"] = []float64{xNum, yNum}

				if index == len(pointList)-1 {
					continue
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
					pzData["originDiagram"] = []float64{x1Num - (x1Num-xNum)/2, yNum}
				}
				if yNum-y_1Num > 0 && x1Num-xNum < 0 {
					pData["rotation"] = 90
					pzData["xz"] = -1
					pzData["originDiagram"] = []float64{x1Num - (x1Num-xNum)/2, yNum}
				}
				if yNum-y_1Num < 0 && x1Num-xNum > 0 {
					pData["rotation"] = -90
					pzData["originDiagram"] = []float64{x1Num - (x1Num-xNum)/2, yNum}
				}
				if yNum-y_1Num < 0 && x1Num-xNum < 0 {
					pData["rotation"] = -90
					pData["yz"] = -1
					pzData["xz"] = -1
					pzData["originDiagram"] = []float64{x1Num - (x1Num-xNum)/2, yNum}
				}
				if xNum-x_1Num > 0 && y1Num-yNum > 0 {
					pzData["rotation"] = 90
					pzData["originDiagram"] = []float64{xNum, y1Num - (y1Num-yNum)/2}
				}
				if xNum-x_1Num < 0 && y1Num-yNum > 0 {
					pData["rotation"] = 180
					pData["yz"] = -1
					pzData["rotation"] = 90
					pzData["originDiagram"] = []float64{xNum, y1Num - (y1Num-yNum)/2}
				}
				if xNum-x_1Num > 0 && y1Num-yNum < 0 {
					pData["yz"] = -1
					pzData["rotation"] = -90
					pzData["originDiagram"] = []float64{xNum, y1Num - (y1Num-yNum)/2}
				}
				if xNum-x_1Num < 0 && y1Num-yNum < 0 {
					pData["rotation"] = 180
					pzData["rotation"] = -90
					pzData["originDiagram"] = []float64{xNum, y1Num - (y1Num-yNum)/2}
				}
				modelInformation = append(modelInformation, pData)
				modelInformation = append(modelInformation, pzData)
			}
		} else {
			pointList := t.Run.Points.Point
			for index, pointData := range pointList {
				xNum := pointData.X.Value
				yNum := pointData.Y.Value
				pData := map[string]any{"rotation": 0, "yz": 1, "xz": 1}
				pData["originDiagram"] = []float64{xNum, yNum}
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
		data["model_information"] = modelInformation
		data["pipes"] = pipeData
		data["bends"] = bendsData
	}
	return parts
}

func CADMappingModel() {
	d := []map[string]any{}

	a, _ := sonic.Marshal(d)
	log.Println(string(a))
}

func parseXML(path string, obj any) error {

	data, err := os.ReadFile(path)
	if err != nil {
		return errors.New("读取消息错误错误：" + err.Error())
	}
	err = xml.Unmarshal(data, obj)
	if err != nil {
		return errors.New("反序列化错误：" + err.Error())
	}
	return nil
}
