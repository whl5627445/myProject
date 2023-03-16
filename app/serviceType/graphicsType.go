package serviceType

type DiagramGraphics struct {
	BorderPattern       string      `json:"border_pattern,omitempty"`
	Color               string      `json:"color,omitempty"`
	Diagram             bool        `json:"diagram,omitempty"`
	ExtentsPoints       []string    `json:"extents_points,omitempty"`
	FillColor           string      `json:"fill_color,omitempty"`
	Points              []string    `json:"points,omitempty"`
	ImageBase64         string      `json:"image_base_64,omitempty"`
	PolygonPoints       []string    `json:"polygon_points,omitempty"`
	Arrow               string      `json:"arrow,omitempty"`
	ArrowSize           string      `json:"arrow_size,omitempty"`
	Smooth              string      `json:"smooth,omitempty"`
	StartAngle          string      `json:"start_angle,omitempty"`
	EndAngle            string      `json:"end_angle,omitempty"`
	HorizontalAlignment string      `json:"horizontal_alignment,omitempty"`
	TextStyles          interface{} `json:"text_styles,omitempty"`
	FontName            string      `json:"font_name,omitempty"`
	TextColor           string      `json:"text_color,omitempty"`
	FontSize            string      `json:"font_size,omitempty"`
	TextType            string      `json:"text_type,omitempty"`
	OriginalTextString  string      `json:"original_text_string,omitempty"`
	FillPattern         string      `json:"fill_pattern,omitempty"`
	LinePattern         string      `json:"line_pattern,omitempty"`
	LineThickness       string      `json:"line_thickness,omitempty"`
	Mobility            bool        `json:"mobility,omitempty"`
	OriginalPoint       string      `json:"original_point,omitempty"`
	Radius              string      `json:"radius,omitempty"`
	Rotation            string      `json:"rotation,omitempty"`
	Type                string      `json:"type,omitempty"`
	Visible             string      `json:"visible,omitempty"`
	ModelLineGraphics
	InputOutputsGraphics
}

type ModelLineGraphics struct {
	Arrow                      string   `json:"arrow,omitempty"`
	ArrowSize                  string   `json:"arrow_size,omitempty"`
	Color                      string   `json:"color,omitempty"`
	ConnectionFrom             string   `json:"connectionfrom,omitempty"`
	ConnectionFromOriginalName string   `json:"connectionfrom_original_name,omitempty"`
	ConnectionTo               string   `json:"connectionto,omitempty"`
	ConnectionToOriginalName   string   `json:"connectionto_original_name,omitempty"`
	LinePattern                string   `json:"line_pattern,omitempty"`
	LineThickness              string   `json:"line_thickness,omitempty"`
	Mobility                   bool     `json:"mobility,omitempty"`
	OriginalPoint              string   `json:"original_point,omitempty"`
	Points                     []string `json:"points,omitempty"`
	Rotation                   string   `json:"rotation,omitempty"`
	Smooth                     string   `json:"smooth,omitempty"`
	Type                       string   `json:"type,omitempty"`
	Visible                    string   `json:"visible,omitempty"`
}

type ComponentIconGraphics struct {
	ID             string                   `json:"ID,omitempty"`
	ClassName      string                   `json:"classname,omitempty"`
	Extent1Diagram string                   `json:"extent_1_diagram,omitempty"`
	Extent2Diagram string                   `json:"extent_2_diagram,omitempty"`
	GraphType      string                   `json:"graph_type,omitempty"`
	InputOutputs   []*ComponentIconGraphics `json:"input_outputs"`
	Mobility       bool                     `json:"mobility,omitempty"`
	Name           string                   `json:"name,omitempty"`
	OriginDiagram  string                   `json:"origin_diagram,omitempty"`
	OriginalName   string                   `json:"original_name,omitempty"`
	OutputType     string                   `json:"output_type,omitempty"`
	Parent         string                   `json:"parent,omitempty"`
	RotateAngle    string                   `json:"rotate_angle,omitempty"`
	Rotation       string                   `json:"rotation,omitempty"`
	SubShapes      []*DiagramGraphics       `json:"sub_shapes"`
	Type           string                   `json:"type,omitempty"`
	Visible        string                   `json:"visible,omitempty"`
}
type InputOutputsGraphics struct {
	ClassName      string             `json:"classname,omitempty"`
	Extent1Diagram string             `json:"extent_1_diagram,omitempty"`
	Extent2Diagram string             `json:"extent_2_diagram,omitempty"`
	GraphType      string             `json:"graph_type,omitempty"`
	InputOutputs   []string           `json:"input_outputs,omitempty"`
	Mobility       bool               `json:"mobility,omitempty"`
	Name           string             `json:"name,omitempty"`
	OriginDiagram  string             `json:"origin_diagram,omitempty"`
	OriginalName   string             `json:"original_name,omitempty"`
	OutputType     string             `json:"output_type,omitempty"`
	Parent         string             `json:"parent,omitempty"`
	RotateAngle    string             `json:"rotate_angle,omitempty"`
	Rotation       string             `json:"rotation,omitempty"`
	SubShapes      []*DiagramGraphics `json:"sub_shapes"`
	Type           string             `json:"type,omitempty"`
	Visible        string             `json:"visible,omitempty"`
}

//type SubShapesGraphics struct {
//	BorderPattern       string   `json:"border_pattern,omitempty"`
//	Color               string   `json:"color,omitempty"`
//	ExtentsPoints       []string `json:"extents_points,omitempty"`
//	FillColor           string   `json:"fill_color,omitempty"`
//	FillPattern         string   `json:"fill_pattern,omitempty"`
//	LinePattern         string   `json:"line_pattern,omitempty"`
//	LineThickness       string   `json:"line_thickness,omitempty"`
//	OriginalPoint       string   `json:"original_point,omitempty"`
//	Radius              string   `json:"radius,omitempty"`
//	Rotation            string   `json:"rotation,omitempty"`
//	ImageBase64         string   `json:"image_base_64,omitempty"`
//	PolygonPoints       string   `json:"polygon_points,omitempty"`
//	Arrow               string   `json:"arrow,omitempty"`
//	ArrowSize           string   `json:"arrow_size,omitempty"`
//	Smooth              string   `json:"smooth,omitempty"`
//	StartAngle          string   `json:"start_angle,omitempty"`
//	EndAngle            string   `json:"end_angle,omitempty"`
//	HorizontalAlignment string   `json:"horizontal_alignment,omitempty"`
//	TextStyles          string   `json:"text_styles,omitempty"`
//	FontName            string   `json:"font_name,omitempty"`
//	TextColor           string   `json:"text_color,omitempty"`
//	FontSize            string   `json:"font_size,omitempty"`
//	TextType            string   `json:"text_type,omitempty"`
//	OriginalTextString  string   `json:"original_text_string,omitempty"`
//	Type                string   `json:"type,omitempty"`
//	Visible             string   `json:"visible,omitempty"`
//}
