package serviceType

type ModelInstance struct {
	Name        string                     `json:"name,omitempty"`
	Restriction string                     `json:"restriction,omitempty"`
	Extends     []map[string]ModelInstance `json:"extends,omitempty"`
	Comment     string                     `json:"comment,omitempty"`
	Annotation  Annotation                 `json:"annotation,omitempty"`
	Components  []Component                `json:"components,omitempty"`
	Connections []Connections              `json:"connections,omitempty"`
	//Source      map[string]interface{}   `json:"source,omitempty"`
}

//type BaseClass struct {
//	Name        string     `json:"name,omitempty"`
//	Restriction string     `json:"restriction,omitempty"`
//	Prefixes    Prefixes   `json:"prefixes,omitempty"`
//	Comment     string     `json:"comment,omitempty"`
//	Annotation  Annotation `json:"annotation,omitempty"`
//}

type Annotation struct {
	Icon                 Icon      `json:"Icon,omitempty"`
	Diagram              Diagram   `json:"Diagram,omitempty"`
	Placement            Placement `json:"Placement,omitempty"`
	Dialog               Dialog    `json:"Dialog,omitempty"`
	Choices              Choices   `json:"choices,omitempty"`
	DefaultComponentName string    `json:"defaultComponentName,omitempty"`
}

type Icon struct {
	CoordinateSystem CoordinateSystem `json:"coordinateSystem,omitempty"`
	Graphics         []Graphics       `json:"graphics,omitempty"`
}
type Diagram struct {
	CoordinateSystem CoordinateSystem `json:"coordinateSystem,omitempty"`
	Graphics         []Graphics       `json:"graphics,omitempty"`
}

type CoordinateSystem struct {
	PreserveAspectRatio bool    `json:"preserveAspectRatio,omitempty"`
	InitialScale        float32 `json:"initialScale,omitempty"`
	Extent              [][]int `json:"extent,omitempty"`
}

type Graphics struct {
	Name     string        `json:"name,omitempty"`
	Elements []interface{} `json:"elements,omitempty"`
}

type Component struct {
	Name       string                 `json:"name,omitempty"`
	Type       ComponentType          `json:"type,omitempty"`
	Extends    []ModelInstance        `json:"extends,omitempty"`
	Modifiers  map[string]interface{} `json:"modifiers,omitempty"`
	Condition  interface{}            `json:"condition,omitempty"`
	Dims       Dims                   `json:"dims,omitempty"`
	Value      map[string]interface{} `json:"value,omitempty"`
	Prefixes   Prefixes               `json:"prefixes,omitempty"`
	Comment    string                 `json:"comment,omitempty"`
	Annotation Annotation             `json:"annotation,omitempty"`
}

type Dims struct {
	Absyn []string `json:"absyn,omitempty"`
	Typed []string `json:"typed,omitempty"`
}

type ComponentType struct {
	Name        string                 `json:"name,omitempty"`
	Restriction string                 `json:"restriction,omitempty"`
	Modifiers   map[string]interface{} `json:"modifiers,omitempty"`
}
type Modifiers struct {
	Quantity    Quantity `json:"quantity,omitempty"`
	Unit        Unit     `json:"unit,omitempty"`
	DisplayUnit string   `json:"displayUnit,omitempty"`
}
type Quantity struct {
	Final bool   `json:"final,omitempty"`
	Value string `json:"$value,omitempty"`
}

type Unit struct {
	Final bool   `json:"final,omitempty"`
	Value string `json:"$value,omitempty"`
}

type Prefixes struct {
	Public       bool   `json:"public,omitempty"`
	Final        bool   `json:"final,omitempty"`
	Inner        bool   `json:"inner,omitempty"`
	Outer        bool   `json:"outer,omitempty"`
	Redeclare    bool   `json:"redeclare,omitempty"`
	Partial      bool   `json:"partial,omitempty"`
	Encapsulated bool   `json:"encapsulated,omitempty"`
	Connector    string `json:"connector,omitempty"`
	Variability  string `json:"variability,omitempty"`
	Direction    string `json:"direction,omitempty"`
}

type Placement struct {
	Transformation Transformation `json:"transformation,omitempty"`
}

type Transformation struct {
	Extent   [][]int `json:"extent,omitempty"`
	Rotation int     `json:"rotation,omitempty"`
	Origin   []int   `json:"origin,omitempty"`
}

type Dialog struct {
	ShowStartAttribute string `json:"showStartAttribute,omitempty"`
	Group              string `json:"group,omitempty"`
	Tab                string `json:"tab,omitempty"`
}

type PlacementExtent struct {
	Extent [][]int `json:"extent,omitempty"`
}

type Connections struct {
	Lhs        Lhs                   `json:"lhs,omitempty"`
	Rhs        Rhs                   `json:"rhs,omitempty"`
	Annotation ConnectionsAnnotation `json:"annotation,omitempty"`
}

type Lhs struct {
	Parts []PartsName `json:"parts,omitempty"`
	Kind  string      `json:"$kind,omitempty"`
}

type Rhs struct {
	Parts []PartsName `json:"parts,omitempty"`
	Kind  string      `json:"$kind,omitempty"`
}

type PartsName struct {
	Name string `json:"name,omitempty"`
}

type ConnectionsAnnotation struct {
	Line ConnectionsAnnotationLine `json:"line,omitempty"`
}

type ConnectionsAnnotationLine struct {
	Points [][]int `json:"points,omitempty"`
	Color  []int   `json:"color,omitempty"`
}

type Choices struct {
	CheckBox bool
}
