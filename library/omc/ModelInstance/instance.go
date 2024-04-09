package instance

import (
	"yssim-go/library/convert"
)

type ModelInstance struct {
	Name        string        `json:"name"`
	Dims        dimensions    `json:"dims,omitempty"`
	Restriction string        `json:"restriction,omitempty"`
	Prefixes    prefixes      `json:"prefixes,omitempty"`
	Comment     string        `json:"comment,omitempty"`
	Elements    []*elements   `json:"elements,omitempty"`
	Connections []*connection `json:"connections,omitempty"`
	Annotation  annotation    `json:"annotation,omitempty"`
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
	Modifiers         any            `json:"modifiers,omitempty"`
	Annotation        annotation     `json:"annotation,omitempty"`
	BaseClassOriginal any            `json:"baseClass,omitempty"` // 字符串或baseClass
	BaseClass         *ModelInstance `json:"baseClassPreprocessing,omitempty"`
	// Source            *source        `json:"source,omitempty"`
	Value     any        `json:"value,omitempty"`
	Condition any        `json:"condition,omitempty"`
	Dims      dimensions `json:"dims,omitempty"`
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
	Choices    *choices                     `json:"choices,omitempty"`
	Dialog     *dialog                      `json:"Dialog,omitempty"`
}
type dialog struct {
	Tab                string `json:"tab,omitempty"`
	Group              string `json:"group,omitempty"`
	ShowStartAttribute bool   `json:"showStartAttribute,omitempty"`
	GroupImage         string `json:"groupImage,omitempty"`
	// ConnectorSizing    bool   `json:"connectorSizing,omitempty"`
}
type choices struct {
	CheckBox bool `json:"checkBox,omitempty"`
}
type placement struct {
	Transformation     transformation `json:"transformation,omitempty"`
	IconTransformation transformation `json:"iconTransformation,omitempty"`
}
type transformation struct {
	Extent   [][]float64 `json:"extent,omitempty"`
	Origin   []float64   `json:"origin,omitempty"`
	Rotation float64     `json:"rotation,omitempty"`
}
type Diagram struct {
	CoordinateSystem coordinateSystem `json:"coordinateSystem,omitempty"`
	GraphicsOriginal any              `json:"graphics,omitempty"`
	Graphics         []*graphics
	// Graphics []*graphics `json:"graphics,omitempty"`
	// TypeOriginal      any            `json:"type,omitempty"`
	// Type              *ModelInstance `json:"typePreprocessing,omitempty"`
}
type Icon struct {
	CoordinateSystem coordinateSystem `json:"coordinateSystem,omitempty"`
	GraphicsOriginal any              `json:"graphics,omitempty"`
	Graphics         []*graphics
	// Graphics []*graphics `json:"graphics,omitempty"`
	// Graphics []*graphics `json:"graphics,omitempty"`
	// Graphics         []*graphics      `json:"graphicsPreprocessing,omitempty"`
}
type coordinateSystem struct {
	PreserveAspectRatio bool        `json:"preserveAspectRatio,omitempty"`
	Extent              [][]float64 `json:"extent,omitempty"`
	InitialScale        float64     `json:"initialScale,omitempty"`
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
		element := m.Elements[i]
		switch true {
		case element.Kind == "extends" && element.BaseClassOriginal != nil:
			if b, ok := element.BaseClassOriginal.(string); ok {
				m.Elements[i].BaseClass = &ModelInstance{Name: b, BasicType: true}
			} else {
				bInstance := &ModelInstance{}
				convert.S2S(m.Elements[i].BaseClassOriginal, &bInstance)
				m.Elements[i].BaseClass = bInstance
				bInstance.DataPreprocessing()
			}
			element.BaseClassOriginal = nil
		case element.Kind == "component" && element.TypeOriginal != nil:
			if element.Prefixes.Public != nil && *element.Prefixes.Public == false {
				m.Elements = append(m.Elements[:i], m.Elements[i+1:]...)
				continue
			}
			if t, ok := element.TypeOriginal.(string); ok {
				m.Elements[i].Type = &ModelInstance{Name: t, BasicType: true}
			} else {
				tInstance := &ModelInstance{}
				convert.S2S(m.Elements[i].TypeOriginal, &tInstance)
				m.Elements[i].Type = tInstance
				tInstance.DataPreprocessing()
			}
			element.TypeOriginal = nil
		}
	}
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
	diagram["coordinateSystem"] = m.CoordinateSystem
	diagram["diagram"] = m.GetDiagramList()
	return diagram
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

// GetIconList 将给定Icon数据处理成结构化信息
func (m *Icon) GetIconList(modelElements *elements) []map[string]any {
	graphicsList := make([]map[string]any, 0)
	for _, c := range m.Graphics {
		graphicsData := getGraphicsData(c, modelElements)
		graphicsList = append(graphicsList, graphicsData)
	}
	return graphicsList
}
func getGraphicsData(c *graphics, modelElements *elements) map[string]any {
	graphicsData := map[string]any{}
	switch c.Name {
	case "Rectangle":
		graphicsData = getRectangle(c.Elements, graphicsData)
	case "Text":
		graphicsData = getText(c.Elements, graphicsData, modelElements)
	case "Polygon":
		graphicsData = getPolygon(c.Elements, graphicsData)
	case "Line":
		graphicsData = getLine(c.Elements, graphicsData)
	case "Ellipse":
		graphicsData = getEllipse(c.Elements, graphicsData)
	case "Bitmap":
		graphicsData = getBitmap(c.Elements, graphicsData)
	}
	return graphicsData
}
