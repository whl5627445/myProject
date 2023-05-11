package API

import "gorm.io/datatypes"

type responseData struct {
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Status int         `json:"status"` // 正常是0，系统级错误是1， 用戶错误是2
	Err    string      `json:"err"`
}

type AppModelMarkData struct {
	PackageId      string `json:"package_id" binding:"required"`
	ModelName      string `json:"model_name" binding:"required"`
	CompileType    string `json:"compile_type,omitempty" binding:""`
	MandatorySave  bool   `json:"save,omitempty" binding:""`
	GroundName     string `json:"ground_name" binding:"required"`
	DataSourceName string `json:"data_source_name" binding:"required"`
	ExperimentId   string `json:"experiment_id" binding:""`
}

type CreateAppSpaceData struct {
	SpaceName        string `json:"space_name" binding:"required"`
	SpaceDescription string `json:"space_description,omitempty" binding:""`
	Background       string `json:"space_background,omitempty" binding:""`
	Icon             string `json:"space_icon,omitempty" binding:"required"`
	IconColor        string `json:"space_icon_color,omitempty" binding:"required"`
}

type EditAppSpaceData struct {
	SpaceId          string `json:"space_id" binding:"required"`
	SpaceName        string `json:"space_name" binding:"required"`
	SpaceDescription string `json:"space_description,omitempty" binding:""`
	Background       string `json:"space_background,omitempty" binding:""`
	Icon             string `json:"space_icon,omitempty" binding:""`
	IconColor        string `json:"space_icon_color,omitempty" binding:""`
	Collect          bool   `json:"space_collect,omitempty" binding:""`
}

type AppSpaceCollectData struct {
	SpaceId []string `json:"space_id" binding:"required"`
	Collect bool     `json:"collect" binding:""`
}

type DeleteAppSpaceData struct {
	SpaceId []string `json:"space_id" binding:"required"`
}

type CreateAppPageData struct {
	//SpaceId  string `json:"space_id" binding:"required"`
	PageName string `json:"name" binding:"required"`
	Tag      string `json:"tag" binding:"required"`
}

type EditAppPageData struct {
	SpaceId  string `json:"space_id" binding:"required"`
	PageId   string `json:"page_id" binding:"required"`
	PageName string `json:"page_name" binding:"required"`
	Tag      string `json:"tag" binding:"required"`
}

type DeleteAppPageData struct {
	SpaceId string   `json:"space_id" binding:"required"`
	PageId  []string `json:"page_id" binding:"required"`
}

type CreatePageComponentData struct {
	PageId             string         `json:"page_id" binding:"required"`
	Type               string         `json:"type" binding:"required"`
	DataObject         datatypes.JSON `json:"data,omitempty" binding:""`
	Width              string         `json:"width" binding:"required"`
	Height             string         `json:"height" binding:"required"`
	PositionX          float64        `json:"position_x" binding:"required"`
	PositionY          float64        `json:"position_y" binding:"required"`
	Angle              float64        `json:"angle" binding:"required"`
	HorizontalFlip     bool           `json:"horizontal_flip" binding:""`
	VerticalFlip       bool           `json:"vertical_flip" binding:""`
	Opacity            int            `json:"opacity" binding:"required"`
	OtherConfiguration datatypes.JSON `json:"other_configuration,omitempty" binding:"required"`
	ZIndex             int            `json:"z_index,omitempty" binding:"required"`
	Styles             datatypes.JSON `json:"styles,omitempty" binding:"required"`
	Events             datatypes.JSON `json:"events,omitempty" binding:"required"`
	ChartConfig        datatypes.JSON `json:"chart_config,omitempty" binding:"required"`
	Option             datatypes.JSON `json:"option,omitempty" binding:"required"`
	ComponentPath      string         `json:"component_path,omitempty" binding:"required"`
	Hide               bool           `json:"hide,omitempty" binding:""`
	Lock               bool           `json:"lock,omitempty" binding:""`
	IsGroup            bool           `json:"is_group,omitempty" binding:""`
}

type EditPageComponentData struct {
	Id                 string         `json:"id" binding:"required"`
	PageId             string         `json:"page_id" binding:"required"`
	Type               string         `json:"type,omitempty" binding:""`
	DataObject         datatypes.JSON `json:"data,omitempty" binding:""`
	Width              string         `json:"width,omitempty" binding:""`
	Height             string         `json:"height,omitempty" binding:""`
	PositionX          float64        `json:"position_x,omitempty" binding:""`
	PositionY          float64        `json:"position_y,omitempty" binding:""`
	Angle              float64        `json:"angle,omitempty" binding:""`
	HorizontalFlip     bool           `json:"horizontal_flip,omitempty" binding:""`
	VerticalFlip       bool           `json:"vertical_flip,omitempty" binding:""`
	Opacity            int            `json:"opacity,omitempty" binding:""`
	OtherConfiguration datatypes.JSON `json:"other_configuration,omitempty" binding:""`
	ZIndex             int            `json:"z_index,omitempty" binding:""`
	Styles             datatypes.JSON `json:"styles,omitempty" binding:""`
	Events             datatypes.JSON `json:"events,omitempty" binding:""`
	ChartConfig        datatypes.JSON `json:"chart_config,omitempty" binding:""`
	Option             datatypes.JSON `json:"option,omitempty" binding:""`
	ComponentPath      string         `json:"component_path,omitempty" binding:""`
	Hide               bool           `json:"hide,omitempty" binding:""`
	Lock               bool           `json:"lock,omitempty" binding:""`
	IsGroup            bool           `json:"is_group,omitempty" binding:""`
}

type CreateBaseComponentData struct {
	Type               string         `json:"type" binding:"required"`
	DataObject         datatypes.JSON `json:"data,omitempty" binding:""`
	Width              string         `json:"width" binding:"required"`
	Height             string         `json:"height" binding:"required"`
	PositionX          float64        `json:"position_x" binding:"required"`
	PositionY          float64        `json:"position_y" binding:"required"`
	Angle              float64        `json:"angle" binding:"required"`
	HorizontalFlip     bool           `json:"horizontal_flip" binding:""`
	VerticalFlip       bool           `json:"vertical_flip" binding:""`
	Opacity            int            `json:"opacity" binding:"required"`
	OtherConfiguration datatypes.JSON `json:"other_configuration" binding:"required"`
	ZIndex             int            `json:"z_index" binding:"required"`
	Styles             datatypes.JSON `json:"styles" binding:"required"`
	Events             datatypes.JSON `json:"events" binding:"required"`
	ChartConfig        datatypes.JSON `json:"chart_config" binding:"required"`
	Option             datatypes.JSON `json:"option" binding:"required"`
	ComponentPath      string         `json:"component_path" binding:"required"`
	Hide               bool           `json:"hide" binding:""`
	Lock               bool           `json:"lock" binding:""`
	IsGroup            bool           `json:"is_group" binding:""`
}

type EditBaseComponentData struct {
	Id                 string         `json:"id" binding:"required"`
	Type               string         `json:"type" binding:"required"`
	DataObject         datatypes.JSON `json:"data,omitempty" binding:""`
	Width              string         `json:"width,omitempty" binding:"required"`
	Height             string         `json:"height,omitempty" binding:"required"`
	PositionX          float64        `json:"position_x,omitempty" binding:"required"`
	PositionY          float64        `json:"position_y,omitempty" binding:"required"`
	Angle              float64        `json:"angle,omitempty" binding:"required"`
	HorizontalFlip     bool           `json:"horizontal_flip,omitempty" binding:""`
	VerticalFlip       bool           `json:"vertical_flip,omitempty" binding:""`
	Opacity            int            `json:"opacity,omitempty" binding:"required"`
	OtherConfiguration datatypes.JSON `json:"other_configuration,omitempty" binding:"required"`
	ZIndex             int            `json:"z_index,omitempty" binding:"required"`
	Styles             datatypes.JSON `json:"styles,omitempty" binding:"required"`
	Events             datatypes.JSON `json:"events,omitempty" binding:"required"`
	ChartConfig        datatypes.JSON `json:"chart_config,omitempty" binding:"required"`
	Option             datatypes.JSON `json:"option,omitempty" binding:"required"`
	ComponentPath      string         `json:"component_path,omitempty" binding:"required"`
	Hide               bool           `json:"hide,omitempty" binding:""`
	Lock               bool           `json:"lock,omitempty" binding:""`
	IsGroup            bool           `json:"is_group,omitempty" binding:""`
}

type DeleteBaseComponentData struct {
	Id string `json:"id" binding:"required"`
}
