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
	Public              *bool       `json:"public,omitempty"`
	Final               bool        `json:"final,omitempty"`
	Inner               bool        `json:"inner,omitempty"`
	Outer               bool        `json:"outer,omitempty"`
	ReplaceableOriginal any         `json:"replaceable,omitempty"` // 值为bool的true，或replaceableObject的结构体类型
	Replaceable         replaceable `json:"template,omitempty"`    // replaceableObject的结构体类型
	Redeclare           bool        `json:"redeclare,omitempty"`
	Partial             bool        `json:"partial,omitempty"`
	Encapsulated        bool        `json:"encapsulated,omitempty"`
	Connector           string      `json:"connector,omitempty"`   // ["flow", "stream"]
	Variability         string      `json:"variability,omitempty"` // ["constant", "parameter", "discrete"]
	Direction           string      `json:"direction,omitempty"`   // ["input", "output"]
}
type replaceable struct {
	IsReplaceable bool       `json:"replaceable,omitempty"`
	Constrainedby string     `json:"constrainedby,omitempty"`
	Comment       string     `json:"comment,omitempty"`
	Annotation    annotation `json:"annotation,omitempty"`
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
	DefaultComponentName string                       `json:"defaultComponentName,omitempty"`
	Experiment           map[string]any               `json:"experiment,omitempty"`
	Diagram              Diagram                      `json:"Diagram,omitempty"`
	Icon                 Icon                         `json:"Icon,omitempty"`
	Uses                 map[string]map[string]string `json:"uses,omitempty"`
	Placement            *placement                   `json:"Placement,omitempty"`
	Evaluate             bool                         `json:"Evaluate,omitempty"`
	HideResult           bool                         `json:"HideResult,omitempty"`
	Choices              choices                      `json:"choices,omitempty"`
	Dialog               dialog                       `json:"Dialog,omitempty"`
	ChoicesAllMatching   bool                         `json:"choicesAllMatching,omitempty"`
}
type dialog struct {
	Tab                string `json:"tab,omitempty"`
	Group              string `json:"group,omitempty"`
	ShowStartAttribute bool   `json:"showStartAttribute,omitempty"`
	GroupImage         string `json:"groupImage,omitempty"`
	ConnectorSizing    bool   `json:"connectorSizing,omitempty"`
}
type choices struct {
	CheckBox       bool                `json:"checkBox,omitempty"`
	ChoiceOriginal []any               `json:"choice,omitempty"`
	Choice         []map[string]string `json:"choicePreprocessing,omitempty"`
}
type placement struct {
	Transformation     *transformation `json:"transformation,omitempty"`
	IconTransformation *transformation `json:"iconTransformation,omitempty"`
}
type transformation struct {
	Extents  [][]float64 `json:"extent,omitempty"`
	Origin   []float64   `json:"origin,omitempty"`
	Rotation *float64    `json:"rotation,omitempty"`
}
type Diagram struct {
	CoordinateSystem *coordinateSystem `json:"coordinateSystem,omitempty"`
	GraphicsOriginal any               `json:"graphics,omitempty"`
	Graphics         []*graphics
}
type Icon struct {
	CoordinateSystem *coordinateSystem `json:"coordinateSystem,omitempty"`
	GraphicsOriginal any               `json:"graphics,omitempty"`
	Graphics         []*graphics
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

type TypeConnector struct {
	ClassName   string           `json:"classname,"`
	Name        string           `json:"name"`
	Restriction string           `json:"restriction"`
	Direction   string           `json:"direction"`
	Elements    []*TypeConnector `json:"elements"`
	Extends     []*TypeConnector `json:"extends"`
	Type        string           `json:"type"`
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
			// if m.Elements[i].Prefixes.Public != nil && *m.Elements[i].Prefixes.Public == false {
			// 	m.Elements = append(m.Elements[:i], m.Elements[i+1:]...)
			// 	i -= 1
			// 	continue
			// }
			if t, ok := m.Elements[i].TypeOriginal.(string); ok {
				m.Elements[i].Type = &ModelInstance{Name: t, BasicType: true}
			} else {
				tInstance := &ModelInstance{}
				convert.S2S(m.Elements[i].TypeOriginal, &tInstance)
				m.Elements[i].Type = tInstance
				tInstance.DataPreprocessing()
			}
			m.Elements[i].TypeOriginal = nil
			for c := 0; c < len(m.Elements[i].Annotation.Choices.ChoiceOriginal); c++ {
				co := m.Elements[i].Annotation.Choices.ChoiceOriginal[c]
				coMap, ok := co.(map[string]any)
				cStr := ""
				if ok {
					pre := coMap["prefixes"].(map[string]any)
					restriction := coMap["restriction"].(string)
					name := coMap["name"].(string)
					baseClass := coMap["baseClass"].(string)
					comment := coMap["comment"].(string)
					if _, reOk := pre["redeclare"].(bool); reOk {
						cStr += "redeclare "
					}
					cStr += restriction
					cStr += " " + name
					cStr += " =  " + baseClass
					m.Elements[i].Annotation.Choices.Choice = append(m.Elements[i].Annotation.Choices.Choice, map[string]string{"value": cStr, "comment": comment})
				}
			}
		case m.Elements[i].TypeOriginal == nil:
			if b, ok := m.Elements[i].BaseClassOriginal.(string); ok {
				m.Elements[i].Type = &ModelInstance{Name: b, BasicType: false}
			} else {
				m.Elements[i].Type = &ModelInstance{}
			}
		}
		m.Elements[i].Modifiers = m.Elements[i].getElementModifiers()
		m.Elements[i].ModifiersOriginal = nil
		m.Elements[i].getPrefixesReplaceable()
	}
}

// getElementModifiers 预处理组件modifier数据， 被设置过的参数与参数属性
func (e *elements) getElementModifiers() map[string]any {
	modifiers := make(map[string]any, 0)
	if modifier, ok := e.ModifiersOriginal.(map[string]any); ok {
		for k1, v1 := range modifier { // map可能存在多层结构，这里表示第一层的k与v的值
			if vMap, vOk := v1.(map[string]string); vOk && e.Kind == "extends" {
				for k2, v2 := range vMap {
					modifiers[k1+"."+k2] = v2 // map可能存在多层结构，这里表示第二层的k与v的值， 相当于某组件的某个参数
				}
			} else {
				modifiers[k1] = v1
			}
		}
	}
	if modifier, ok := e.ModifiersOriginal.(string); ok {
		modifiers["value"] = modifier
	}
	return modifiers
}

// getExtendsParameter 获取设置的继承模型参数
func (e *elements) getExtendsModifiers(extendModelParameterMap map[string]map[string]*Parameter, n int) {
	for elementName, v := range e.Modifiers {
		if nvMap, ok := v.(map[string]any); ok {
			for pName, pValue := range nvMap {
				if parameterMap, ok := extendModelParameterMap[elementName]; ok {
					if p, ok := parameterMap[pName]; ok {
						p.DefaultValue = pValue
						continue
					} else {
						extendModelParameterMap[elementName][pName] = &Parameter{Name: e.Name, IsExtend: n > 1, Type: "Normal", Visible: !(e.Prefixes.Public != nil && *e.Prefixes.Public == false)}
					}
				} else {
					extendModelParameterMap[elementName] = map[string]*Parameter{pName: {Name: pName, IsExtend: n > 1, Type: "Normal", Visible: !(e.Prefixes.Public != nil && *e.Prefixes.Public == false)}}
				}
				if n > 1 {
					extendModelParameterMap[elementName][pName].DefaultValue = pValue
				} else {
					extendModelParameterMap[elementName][pName].Value = pValue
				}
			}
		}
	}
}

// getPrefixesReplaceable 预处理组件Prefixes中的Replaceable数据
func (e *elements) getPrefixesReplaceable() {
	if e.Prefixes.ReplaceableOriginal == nil {
		return
	}
	switch e.Prefixes.ReplaceableOriginal.(type) {
	case bool:
		e.Prefixes.Replaceable.IsReplaceable = true
	case map[string]any:
		e.Prefixes.Replaceable.IsReplaceable = true
		convert.S2S(e.Prefixes.ReplaceableOriginal, &e.Prefixes.Replaceable)
	}
}

// GetModelParameterValue 获取模型的参数数据，extendModelParameterMap是当模型继承了其他模型，又设置了继承模型是参数时会有用
func (m *ModelInstance) GetModelParameterValue(modelParameterMap map[string]map[string]*Parameter, isExtend bool, n int) []map[string]any {
	eList := make([]map[string]any, 0)
	for _, e := range m.Elements {
		if e.Kind == "extends" {
			e.getExtendsModifiers(modelParameterMap, n)
			eList = append(eList, e.BaseClass.GetModelParameterValue(modelParameterMap, true, n+1)...)
		}
		if e.Kind == "component" {
			e.ElementsParameter = map[string]*Parameter{}
			e.ParameterList = []*Parameter{}
			if extend, ok := modelParameterMap[e.Name]; ok {
				e.ElementsParameter = extend
			}
			e.GetElementsParameters(e.ElementsParameter, &e.ParameterList, isExtend, e.Type.Name, n)
			p := map[string]any{"name": e.Name, "parameter": e.ParameterList, "type": "component"}
			if e.Prefixes.Variability == "parameter" {
				p["type"] = "model"
			} else {
				p["properties"] = e.getProperties()
			}
			eList = append(eList, p)
		}
	}
	return eList
}

// 获取模型组件属性
func (e *elements) getProperties() map[string]any {
	p := map[string]any{
		"variability": "unspecified",
		"causality":   "unspecified",
		"dimension":   e.Dims.Typed,
		"inner/outer": "none",
		"comment":     e.Comment,
		"path":        e.Type.Name,
	}
	if e.Prefixes.Variability != "" {
		p["variability"] = e.Prefixes.Variability
	}
	if e.Prefixes.Direction != "" {
		p["causality"] = e.Prefixes.Direction
	}
	if e.Prefixes.Inner {
		p["inner/outer"] = "inner"
	} else if e.Prefixes.Outer {
		p["inner/outer"] = "outer"
	}
	properties := make([]any, 3)
	properties[0] = e.Prefixes.Final
	properties[1] = "public"
	properties[2] = false
	if _, ok := e.Prefixes.ReplaceableOriginal.(bool); ok {
		properties[2] = e.Prefixes.Replaceable.IsReplaceable
	}
	p["properties"] = properties
	return p
}

// GetConnectionsList 将给定连接信息处理成结构化信息
func (m *ModelInstance) GetConnectionsList() []map[string]any {
	data := make([]map[string]any, 0)
	for _, c := range m.Connections {
		line := map[string]any{
			"points":      make([][]float64, 0),                                                    // 连线经过的拐点，第一个与最后一个表示起始位置
			"color":       []int{0, 0, 127},                                                        // 连线颜色
			"arrow":       []string{"Arrow.None", "Arrow.None"},                                    // 连线的开始和结束点箭头样式
			"arrowSize":   3,                                                                       // 箭头大小
			"linePattern": map[string]any{"name": "LinePattern.Solid", "index": 2, "kind": "enum"}, // 连线样式
			"thickness":   0.25,                                                                    // 连线粗细
			"smooth":      map[string]any{"index": 1, "kind": "enum", "name": "Smooth.None"},       // 平滑样式
			"rotation":    0,                                                                       // 旋转角度
			"type":        "Line",                                                                  // 数据类型
			"offset":      []float64{0, 0},                                                         // 偏移量
			"lhs":         c.Lhs.Parts,                                                             // 起点数据 ，包括组件名字和接口名字， 如果有下标会有下标数组
			"rhs":         c.Rhs.Parts,                                                             // 结束点数据，包括组件名字和接口名字， 如果有下标会有下标数组
		}
		if c.Annotation != nil {
			for k, v := range c.Annotation.Line {
				line[k] = v
			}
		}
		data = append(data, line)
	}
	return data
}

// GetAnnotationDiagram 获取Diagram中的图形以及坐标系信息
func (m *Diagram) GetAnnotationDiagram() []map[string]any {
	diagram := make(map[string]any, 0)
	diagramData := m.GetDiagramList(nil)
	diagram["diagram"] = diagramData
	if len(diagramData) > 0 {
		diagram["coordinateSystem"] = m.GetCoordinateSystem()
		return []map[string]any{diagram}
	}
	return nil
}

// GetCoordinateSystem 获取Diagram的坐标系数据
func (m *Diagram) GetCoordinateSystem() map[string]any {
	return getCoordinateSystem(m.CoordinateSystem)
}

// GetDiagramList 将给定Diagram数据处理成结构化信息
func (m *Diagram) GetDiagramList(modelElements *elements) []map[string]any {
	graphicsList := make([]map[string]any, 0)
	for _, c := range m.Graphics {
		graphicsData := getGraphicsData(c, modelElements)
		graphicsList = append(graphicsList, graphicsData)
	}
	return graphicsList
}

// GetCoordinateSystem 获取Icon的坐标系数据
func (m *Icon) GetCoordinateSystem() map[string]any {
	return getCoordinateSystem(m.CoordinateSystem)
}

// GetIconList 将给定Icon数据处理成结构化信息，elements是为了传入参数信息
func (m *Icon) GetIconList(modelElements *elements) []map[string]any {
	graphicsList := make([]map[string]any, 0)
	for _, c := range m.Graphics {
		graphicsData := getGraphicsData(c, modelElements)
		graphicsList = append(graphicsList, graphicsData)
	}
	return graphicsList
}

// GetIconListALL 获取给定ModelInstance的全部图标数据，递归查找，elements是为了传入参数信息
func (m *ModelInstance) GetIconListALL(modelElements *elements, isElement bool) []map[string]any {

	graphicsList := make([]map[string]any, 0)
	if (m.Restriction == "connector" || m.Restriction == "expandable connector") && isElement {
		graphicsList = append(graphicsList, m.Annotation.Diagram.GetDiagramList(modelElements)...)
	} else {
		graphicsList = append(graphicsList, m.Annotation.Icon.GetIconList(modelElements)...)
	}
	for _, element := range m.Elements {
		if element.BaseClass != nil && element.Kind == "extends" {
			graphicsList = append(element.BaseClass.GetIconListALL(modelElements, isElement), graphicsList...)
		}
	}
	return graphicsList
}

// GetElementsParameters 获取组件参数与值信息，写入给定的map当中
func (e *elements) GetElementsParameters(parameterMap map[string]*Parameter, parameterList *[]*Parameter, end bool, parentName string, n int) {
	if e.Prefixes.Final || e.Annotation.Dialog.ConnectorSizing {
		return
	}
	if e.Kind == "component" && e.Prefixes.Variability != "parameter" && e.ElementsParameter != nil {
		getElementsParameterModifier(e, n)
	}
	p := getElementsParameter(parameterMap, e, parentName, n)
	if p != nil {
		*parameterList = append(*parameterList, p)
		return
	}
	if e.Type != nil && !end {
		for _, element := range e.Type.Elements {
			element.GetElementsParameters(parameterMap, parameterList, true, e.Type.Name, n+1)
		}
	}
	if e.BaseClass != nil {
		getElementsParameterBaseClass(parameterMap, parameterList, e, n)
	}
}

func getElementsParameter(parameterMap map[string]*Parameter, e *elements, parentName string, n int) *Parameter {
	p, _ := parameterMap[e.Name]
	eStartValue, eStartOk, eFixedValue, eFixedOk := getParameterFixedAndStart(e, p)
	if e.Prefixes.Variability == "parameter" || e.Annotation.Dialog.ShowStartAttribute || (e.Prefixes.Variability != "parameter" && (eStartOk || eFixedOk)) {
		value := getElementsParameterValue(e)
		if eStartOk || eFixedOk {
			value = map[string]any{"start": eStartValue, "fixed": eFixedValue}
		}
		p = getElementsParameterInstance(p, e, value, n)
		if p == nil {
			return nil
		}
		p.Comment = e.Comment
		p.ExtendName = parentName
		e.Annotation.Dialog.getParameterDialog(p, eStartOk)
		e.Annotation.getParameterChoices(p)
		if e.Type != nil && (!e.Type.BasicType || e.Type.Restriction == "type") {
			getElementsParameterEnumeration(p, e)
		}
		if e.Type != nil && e.Type.BasicType && e.Type.Name == "Boolean" {
			p.Type = "CheckBox"
		}
		parameterMap[e.Name] = p
		return p
	}
	return nil
}

func getElementsParameterModifier(e *elements, n int) {
	for k, v := range e.Modifiers {
		if p, ok := e.ElementsParameter[k]; ok {
			if p.DefaultValue == nil {
				p.DefaultValue = v
			}
		} else {
			e.ElementsParameter[k] = &Parameter{Name: k, Type: "Normal", Visible: !(e.Prefixes.Public != nil && *e.Prefixes.Public == false)}
			switch true {
			case n > 0 && e.ElementsParameter[k].DefaultValue == nil:
				e.ElementsParameter[k].DefaultValue = v
			case n < 1:
				e.ElementsParameter[k].Value = v
			}
		}
	}
}

func getElementsParameterValue(e *elements) any {
	value := any(nil)
	if _, modifierOk := e.Modifiers["value"]; modifierOk {
		value = e.Modifiers["value"].(string)
	}
	return value
}

func getElementsParameterInstance(p *Parameter, e *elements, value any, n int) *Parameter {
	if p != nil {
		if final, finalOk := p.ParameterAttributes["final"]; finalOk && final.(bool) {
			return nil
		}
		p.IsExtend = n > 1
		if p.DefaultValue == nil {
			p.DefaultValue = value
		}
	} else {
		p = &Parameter{ParameterAttributesData: e.Modifiers, Name: e.Name, IsExtend: n > 1, Type: "Normal", Visible: !(e.Prefixes.Public != nil && *e.Prefixes.Public == false)}
		switch true {
		case n > 0:
			p.DefaultValue = value
		default:
			p.Value = value
		}
	}
	return p
}

func getElementsParameterEnumeration(p *Parameter, e *elements) {
	if len(e.Type.Elements) > 0 && e.Type.Elements[0].BaseClass != nil && e.Type.Elements[0].BaseClass.Name == "enumeration" {
		p.Type = "Enumeration"
		options := []map[string]string{}
		for i := 1; i < len(e.Type.Elements); i++ {
			options = append(options, map[string]string{"value": e.Type.Elements[i].Name, "comment": e.Type.Elements[i].Comment})
		}
		p.Options = options
	}
	p.ParameterUnit = e.getParameterUnit()
}

func getElementsParameterBaseClass(parameterMap map[string]*Parameter, parameterList *[]*Parameter, e *elements, n int) {
	if e.Kind == "extends" {
		for modifierName, modifierValue := range e.Modifiers {
			parameterMap[modifierName] = &Parameter{Name: modifierName, ParameterAttributesData: modifierValue, IsExtend: n > 1, Type: "Normal", Visible: !(e.Prefixes.Public != nil && *e.Prefixes.Public == false)}
			p, ok := modifierValue.(map[string]any)
			if ok {
				parameterMap[modifierName].ParameterAttributes = p
			}
		}
	}
	for _, element := range e.BaseClass.Elements {
		element.GetElementsParameters(parameterMap, parameterList, true, e.BaseClass.Name, n+1)
	}
}

func getParameterFixedAndStart(e *elements, p *Parameter) (any, bool, any, bool) {
	if e.Prefixes.Variability == "parameter" {
		return "", false, "", false
	}
	startValue, startOk := e.Modifiers["start"]
	fixedValue, fixedOk := e.Modifiers["fixed"]
	if p != nil {
		if !startOk {
			startValue, startOk = p.ParameterAttributes["start"]
		}
		if !startOk {
			startValue = nil
		}
		if !fixedOk {
			fixedValue, fixedOk = p.ParameterAttributes["fixed"]
		}
		if !fixedOk {
			fixedValue = nil
		}
	}
	switch startValue.(type) {
	case map[string]any:
		startValue = startValue.(map[string]any)["value"]
	}
	switch fixedValue.(type) {
	case map[string]any:
		fixedValue = fixedValue.(map[string]any)["value"]
	}
	return startValue, startOk, fixedValue, fixedOk
}

// GetParameterAttributes 获取该类型的属性数据，max、min、start等等
func (m *ModelInstance) GetParameterAttributes() map[string]any {
	attr := make(map[string]any, 0)
	for i := 0; i < len(m.Elements); i++ {
		if m.Elements[i].Kind == "extends" && m.Elements[i].BaseClass.BasicType {
			attr = m.Elements[i].Modifiers
			break
		}
	}
	return attr
}

// Parameter 定义的参数结构体
type Parameter struct {
	Comment                 string              `json:"comment"`
	DefaultValue            any                 `json:"defaultValue"`
	Value                   any                 `json:"value"`
	ExtendName              string              `json:"extendName"`
	Group                   string              `json:"group"`
	Name                    string              `json:"name"`
	Tab                     string              `json:"tab"`
	Type                    string              `json:"type"`
	Options                 []map[string]string `json:"options,omitempty"`
	ParameterAttributes     map[string]any      `json:"-"`
	ParameterAttributesData any                 `json:"attributes,omitempty"`
	ParameterUnit           map[string]any      `json:"parameterUnit"`
	IsExtend                bool                `json:"isExtend"`
	Visible                 bool                `json:"visible"`
}

// 获取单位数据，是一个map，包含源码当中定义的该类型的单位属性
func (e *elements) getParameterUnit() map[string]any {
	if e.BaseClass != nil && e.BaseClass.BasicType {
		return e.Modifiers
	}
	if e.Type != nil {
		for _, element := range e.Type.Elements {
			if element.Kind == "extends" {
				if element.BaseClass != nil && element.BaseClass.BasicType {
					return element.Modifiers
				}
				for _, bElement := range element.BaseClass.Elements {
					return bElement.getParameterUnit()
				}
			}
		}
	}
	return map[string]any{}
}

// 获取参数的dialog数据， 包括分组，tab页，是否显示开始属性以及部分专属处理
func (d *dialog) getParameterDialog(parameter *Parameter, startOk bool) {
	parameter.Tab = "General"
	parameter.Group = "Parameters"
	if d.Tab != "" {
		parameter.Tab = d.Tab
	}
	if d.Group != "" {
		parameter.Group = d.Group
	}
	if d.ShowStartAttribute || startOk {
		parameter.Group = "Initialization"
		parameter.Type = "checkWrite"
		parameter.Name = parameter.Name + ".start"
	}
}

func (a *annotation) getParameterChoices(parameter *Parameter) {
	switch true {
	case len(a.Choices.Choice) > 0:
		options := []map[string]string{}
		for _, value := range a.Choices.Choice {
			choiceMap := map[string]string{}
			choiceMap["comment"] = value["comment"]
			choiceMap["value"] = value["value"]
			options = append(options, choiceMap)
		}
		parameter.Options = options
		parameter.Type = "Enumeration"
		return
	case a.Choices.CheckBox:
		parameter.Type = "CheckBox"
	}
}

// GetElementsIconExtents 获取元素的Extents数据，从IconTransformation中
func (p *placement) GetElementsIconExtents() [][]float64 {
	if p.IconTransformation != nil {
		return p.IconTransformation.Extents
	}

	if p.Transformation != nil {
		return p.Transformation.Extents
	}

	return nil
}

// GetElementsIconOrigin 获取元素的Origin数据，从IconTransformation中
func (p *placement) GetElementsIconOrigin() []float64 {
	if p.IconTransformation != nil {
		return p.IconTransformation.Origin
	}

	if p.Transformation != nil {
		return p.Transformation.Origin
	}

	return nil
}

// GetElementsIconRotation 获取元素的Rotation数据，从IconTransformation中
func (p *placement) GetElementsIconRotation() float64 {
	if p.IconTransformation != nil {
		if p.IconTransformation.Rotation != nil {
			return *(p.IconTransformation.Rotation)
		}
		return 0
	}

	if p.Transformation != nil {
		if p.Transformation.Rotation != nil {
			return *(p.Transformation.Rotation)
		}
		return 0
	}

	return 0
}

// GetElementsExtents 获取元素的Extents数据，从Transformation中
func (p *placement) GetElementsExtents() [][]float64 {
	if p.Transformation != nil {
		return p.Transformation.Extents
	}

	return nil
}

// GetElementsOrigin 获取元素的Origin数据，从Transformation中
func (p *placement) GetElementsOrigin() []float64 {
	if p.Transformation != nil {
		return p.Transformation.Origin
	}

	return nil
}

// GetElementsRotation 获取元素的Rotation数据，从Transformation中
func (p *placement) GetElementsRotation() float64 {
	if p.Transformation != nil {
		if p.Transformation.Rotation != nil {
			return *(p.Transformation.Rotation)
		}
		return 0
	}

	return 0
}

// hasReplaceablePlacement 判断Prefixes中是否包含放置点范围数据，并且返回范围数据
func (p *prefixes) HasReplaceableExtent() ([][]float64, bool) {
	if p.Replaceable.Annotation.Placement != nil {
		return p.Replaceable.Annotation.Placement.Transformation.Extents, true
	}
	return nil, false
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

// 将给定coordinateSystem处理成结构化的坐标系数据，具有默认值， 当前默认值是不具备修改能力的，如果是配置项，则需要传入对应的配置数据
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

// setElementModifierValue(BatteryDischargeCharge, battery1, $Code((redeclare Modelica.Electrical.Batteries.ParameterRecords.CellData cellData(Qnom = 4432428010))))
