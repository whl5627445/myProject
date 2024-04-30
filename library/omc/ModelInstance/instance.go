package instance

import (
	"yssim-go/library/convert"
)

type ModelInstance struct {
	Name              string         `json:"name"`
	Dims              dimensions     `json:"dims,omitempty"`
	Restriction       string         `json:"restriction,omitempty"`
	Prefixes          prefixes       `json:"prefixes,omitempty"`
	Comment           string         `json:"comment,omitempty"`
	ModifiersOriginal any            `json:"modifiers,omitempty"`
	Modifiers         map[string]any `json:"modifiersObject,omitempty"`
	Elements          []*elements    `json:"elements,omitempty"`
	Connections       []*connection  `json:"connections,omitempty"`
	Annotation        annotation     `json:"annotation,omitempty"`
	// Source      source        `json:"source,omitempty"`
	BasicType bool
}
type dimensions struct {
	Absyn []string `json:"absyn,omitempty"`
	Typed []string `json:"typed,omitempty"`
}
type prefixes struct {
	Public       *bool  `json:"public,omitempty"`
	Final        *bool  `json:"final,omitempty"`
	Inner        *bool  `json:"inner,omitempty"`
	Outer        *bool  `json:"outer,omitempty"`
	Replaceable  any    `json:"replaceable,omitempty"` // 值为bool的true，或replaceableObject的结构体类型
	Redeclare    *bool  `json:"redeclare,omitempty"`
	Partial      *bool  `json:"partial,omitempty"`
	Encapsulated *bool  `json:"encapsulated,omitempty"`
	Connector    string `json:"connector,omitempty"`   // ["flow", "stream"]
	Variability  string `json:"variability,omitempty"` // ["constant", "parameter", "discrete"]
	Direction    string `json:"direction,omitempty"`   // ["input", "output"]
}
type elements struct {
	Kind              string         `json:"kind,omitempty"`
	Name              string         `json:"name,omitempty"`
	TypeOriginal      any            `json:"type,omitempty"`
	Type              *ModelInstance `json:"typePreprocessing,omitempty"`
	Restriction       string         `json:"restriction,omitempty"`
	Prefixes          prefixes       `json:"prefixes,omitempty"`
	Comment           string         `json:"comment,omitempty"`
	ModifiersOriginal any            `json:"modifiers,omitempty"`
	Modifiers         map[string]any `json:"modifiersObject,omitempty"`
	Annotation        annotation     `json:"annotation,omitempty"`
	BaseClassOriginal any            `json:"baseClass,omitempty"` // 字符串或baseClass
	BaseClass         *ModelInstance `json:"baseClassPreprocessing,omitempty"`
	Condition         any            `json:"condition,omitempty"`
	Dims              dimensions     `json:"dims,omitempty"`
	ParameterList     []*Parameter   `json:"parameter,omitempty"`
	ElementsParameter map[string]*Parameter
	// Source            *source        `json:"source,omitempty"`
	// Value             any               `json:"value,omitempty"`

}
type source struct {
	Filename    string `json:"filename,omitempty"`
	LineStart   int    `json:"lineStart,omitempty"`
	ColumnStart int    `json:"columnStart,omitempty"`
	LineEnd     int    `json:"lineEnd,omitempty"`
	ColumnEnd   int    `json:"columnEnd,omitempty"`
	Readonly    bool   `json:"readonly,omitempty"`
}
type annotation struct {
	DefaultComponentName string         `json:"defaultComponentName,omitempty"`
	Experiment           map[string]any `json:"experiment,omitempty"`
	// Documentation        map[string]any               `json:"Documentation,omitempty"`
	Diagram    Diagram                      `json:"Diagram,omitempty"`
	Icon       Icon                         `json:"Icon,omitempty"`
	Uses       map[string]map[string]string `json:"uses,omitempty"`
	Placement  *placement                   `json:"Placement,omitempty"`
	Evaluate   bool                         `json:"Evaluate,omitempty"`
	HideResult bool                         `json:"HideResult,omitempty"`
	Choices    choices                      `json:"choices,omitempty"`
	Dialog     dialog                       `json:"Dialog,omitempty"`
}
type dialog struct {
	Tab                string `json:"tab,omitempty"`
	Group              string `json:"group,omitempty"`
	ShowStartAttribute bool   `json:"showStartAttribute,omitempty"`
	GroupImage         string `json:"groupImage,omitempty"`
	ConnectorSizing    bool   `json:"connectorSizing,omitempty"`
}
type choices struct {
	CheckBox bool     `json:"checkBox,omitempty"`
	Choice   []string `json:"choice,omitempty"`
}
type placement struct {
	Transformation     transformation `json:"transformation,omitempty"`
	IconTransformation transformation `json:"iconTransformation,omitempty"`
}
type transformation struct {
	Extents  [][]float64 `json:"extent,omitempty"`
	Origin   []float64   `json:"origin,omitempty"`
	Rotation float64     `json:"rotation,omitempty"`
}
type Diagram struct {
	CoordinateSystem *coordinateSystem `json:"coordinateSystem,omitempty"`
	GraphicsOriginal any               `json:"graphics,omitempty"`
	Graphics         []*graphics
	// Graphics []*graphics `json:"graphics,omitempty"`
	// TypeOriginal      any            `json:"type,omitempty"`
	// Type              *ModelInstance `json:"typePreprocessing,omitempty"`
}
type Icon struct {
	CoordinateSystem *coordinateSystem `json:"coordinateSystem,omitempty"`
	GraphicsOriginal any               `json:"graphics,omitempty"`
	Graphics         []*graphics
	// Graphics []*graphics `json:"graphics,omitempty"`
	// Graphics []*graphics `json:"graphics,omitempty"`
	// Graphics         []*graphics      `json:"graphicsPreprocessing,omitempty"`
}
type coordinateSystem struct {
	PreserveAspectRatio bool        `json:"preserveAspectRatio"`
	Extents             [][]float64 `json:"extent"`
	InitialScale        float64     `json:"initialScale"`
}
type graphics struct {
	Kind     string `json:"kind,omitempty"`
	Name     string `json:"name,omitempty"`
	Elements []any  `json:"elements,omitempty"`
}
type connection struct {
	Lhs        lhs                   `json:"lhs,omitempty"`
	Rhs        rhs                   `json:"rhs,omitempty"`
	Annotation *connectionAnnotation `json:"annotation,omitempty"`
}
type lhs struct {
	Kind  string  `json:"kind,omitempty"`
	Parts []parts `json:"parts,omitempty"`
}
type rhs struct {
	Kind  string   `json:"kind,omitempty"`
	Parts []*parts `json:"parts,omitempty"`
}
type parts struct {
	Name       string `json:"name,omitempty"`
	Subscripts []any  `json:"subscripts,omitempty"`
}
type connectionAnnotation struct {
	Line map[string]any `json:"line,omitempty"`
}

// DataPreprocessing 模型实例数据预处理
func (m *ModelInstance) DataPreprocessing() {
	if _, ok := m.Annotation.Icon.GraphicsOriginal.([]any); ok {
		g := []*graphics{}
		convert.S2S(m.Annotation.Icon.GraphicsOriginal, &g)
		m.Annotation.Icon.Graphics = g
	}
	if _, ok := m.Annotation.Diagram.GraphicsOriginal.([]any); ok {
		g := []*graphics{}
		convert.S2S(m.Annotation.Diagram.GraphicsOriginal, &g)
		m.Annotation.Diagram.Graphics = g
	}
	m.Annotation.Icon.GraphicsOriginal = nil
	m.Annotation.Diagram.GraphicsOriginal = nil
	for i := 0; i < len(m.Elements); i++ {
		switch true {
		case m.Elements[i].Kind == "extends" && m.Elements[i].BaseClassOriginal != nil:
			if b, ok := m.Elements[i].BaseClassOriginal.(string); ok {
				m.Elements[i].BaseClass = &ModelInstance{Name: b, BasicType: true}
			} else {
				bInstance := &ModelInstance{}
				convert.S2S(m.Elements[i].BaseClassOriginal, &bInstance)
				m.Elements[i].BaseClass = bInstance
				bInstance.DataPreprocessing()
			}
			m.Elements[i].BaseClassOriginal = nil
		case m.Elements[i].Kind == "component" && m.Elements[i].TypeOriginal != nil:
			if m.Elements[i].Prefixes.Public != nil && *m.Elements[i].Prefixes.Public == false {
				m.Elements = append(m.Elements[:i], m.Elements[i+1:]...)
				continue
			}
			if t, ok := m.Elements[i].TypeOriginal.(string); ok {
				m.Elements[i].Type = &ModelInstance{Name: t, BasicType: true}
			} else {
				tInstance := &ModelInstance{}
				convert.S2S(m.Elements[i].TypeOriginal, &tInstance)
				m.Elements[i].Type = tInstance
				tInstance.DataPreprocessing()
			}
			m.Elements[i].TypeOriginal = nil
		}
		m.Elements[i].Modifiers = getElementModifiers(m.Elements[i].ModifiersOriginal)
		m.Elements[i].ModifiersOriginal = nil
	}
}

// 预处理组件modifier数据， 被设置过的参数与参数属性
func getElementModifiers(modifiersOriginal any) map[string]string {
	modifiers := make(map[string]string, 0)
	if modifier, ok := modifiersOriginal.(map[string]any); ok {
		for k1, v1 := range modifier { // map可能存在多层结构，这里表示第一层的k与v的值
			if vString, vOk := v1.(string); vOk {
				modifiers[k1] = vString
			}
			if vMap, vOk := v1.(map[string]string); vOk {
				for k2, v2 := range vMap {
					modifiers[k1+"."+k2] = v2 // map可能存在多层结构，这里表示第二层的k与v的值， 相当于某组件的某个参数
				}
			}
		}
	}
	return modifiers
}

// GetConnectionsList 将给定连接信息处理成结构化信息
func (m *ModelInstance) GetConnectionsList() []map[string]any {
	data := make([]map[string]any, 0)
	for _, c := range m.Connections {
		line := map[string]any{
			"points":        make([][]float64, 0),                 // 连线经过的拐点，第一个与最后一个表示起始位置
			"color":         []int{0, 0, 127},                     // 连线颜色
			"arrow":         []string{"Arrow.None", "Arrow.None"}, // 连线的开始和结束点箭头样式
			"arrowSize":     3,                                    // 箭头大小
			"linePattern":   "LinePattern.Solid",                  // 连线样式
			"lineThickness": 0.25,                                 // 连线粗细
			"smooth":        "Smooth.None",                        // 平滑样式
			"rotation":      0,                                    // 旋转角度
			"type":          "Line",                               // 数据类型
			"offset":        []float64{0, 0},                      // 偏移量
			"lhs":           c.Lhs.Parts,                          // 起点数据 ，包括组件名字和接口名字， 如果有下标会有下标数组
			"rhs":           c.Rhs.Parts,                          // 结束点数据，包括组件名字和接口名字， 如果有下标会有下标数组
		}
		for k, v := range c.Annotation.Line {
			line[k] = v
		}
		data = append(data, line)
	}
	return data
}

// GetAnnotationDiagram 获取Diagram中的图形以及坐标系信息
func (m *Diagram) GetAnnotationDiagram() map[string]any {
	diagram := make(map[string]any, 0)
	diagramData := m.GetDiagramList()
	diagram["diagram"] = diagramData
	if len(diagramData) > 0 {
		diagram["coordinateSystem"] = m.GetCoordinateSystem()
	}
	return diagram
}

func (m *Diagram) GetCoordinateSystem() map[string]any {
	return getCoordinateSystem(m.CoordinateSystem)
}

// GetDiagramList 将给定Diagram数据处理成结构化信息
func (m *Diagram) GetDiagramList() []map[string]any {
	graphicsList := make([]map[string]any, 0)
	for _, c := range m.Graphics {
		graphicsData := getGraphicsData(c, nil)
		graphicsList = append(graphicsList, graphicsData)
	}
	return graphicsList
}

func (m *Icon) GetCoordinateSystem() map[string]any {
	return getCoordinateSystem(m.CoordinateSystem)
}

// GetIconList 将给定Icon数据处理成结构化信息
func (m *Icon) GetIconList(modelElements *elements) []map[string]any {
	graphicsList := make([]map[string]any, 0)
	for _, c := range m.Graphics {
		graphicsData := getGraphicsData(c, modelElements)
		graphicsList = append(graphicsList, graphicsData)
	}
	return graphicsList
}

func (m *ModelInstance) GetIconListALL(modelElements *elements) []map[string]any {

	graphicsList := make([]map[string]any, 0)
	graphicsList = append(graphicsList, m.Annotation.Icon.GetIconList(modelElements)...)
	for _, element := range m.Elements {
		if element.Kind == "extends" && element.BaseClass != nil {
			graphicsList = append(element.BaseClass.GetIconListALL(modelElements), graphicsList...)
		}
	}
	return graphicsList
}

// 处理图形数据
func getGraphicsData(g *graphics, modelElements *elements) map[string]any {
	graphicsData := map[string]any{}
	switch g.Name {
	case "Rectangle":
		graphicsData = getRectangle(g.Elements, graphicsData)
	case "Text":
		graphicsData = getText(g.Elements, graphicsData, modelElements)
	case "Polygon":
		graphicsData = getPolygon(g.Elements, graphicsData)
	case "Line":
		graphicsData = getLine(g.Elements, graphicsData)
	case "Ellipse":
		graphicsData = getEllipse(g.Elements, graphicsData)
	case "Bitmap":
		graphicsData = getBitmap(g.Elements, graphicsData)
	}
	return graphicsData
}

func getCoordinateSystem(m *coordinateSystem) map[string]any {
	c := make(map[string]any, 0)
	c["preserveAspectRatio"] = false
	c["extent"] = [][]float64{{-100.0, -100.0}, {100.0, 100.0}}
	c["initialScale"] = 0.1
	if m == nil {
		return c
	}
	if m.Extents != nil {
		c["extent"] = m.Extents
	}
	if m.InitialScale != 0 {
		c["initialScale"] = m.InitialScale
	}
	c["preserveAspectRatio"] = m.PreserveAspectRatio
	return c
}

// getParameterValue 获取Text类型图形数据中组件参数的值的核心逻辑，返回值内容和值类型
// func getParameterValue(value any) (any, string) {
// 	switch value.(type) {
// 	case map[string]any:
// 		if v, ok := value.(map[string]any)["value"]; ok {
// 			return v, "Normal"
// 		}
// 		if v, ok := value.(map[string]any)["binding"]; ok {
// 			switch v.(type) {
// 			case map[string]any:
// 				vMap := v.(map[string]any)
// 				if vMap["kind"] == "enum" {
// 					return vMap["name"], "Enumeration"
// 				}
// 			case bool:
// 				return v, "CheckBox"
// 			}
//
// 			return v, "Normal"
// 		}
// 	}
// 	return "", "Normal"
// }

// setElementModifierValue(BatteryDischargeCharge, battery1, $Code((redeclare Modelica.Electrical.Batteries.ParameterRecords.CellData cellData(Qnom = 4432428010))))
