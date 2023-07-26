package DataType

import "gorm.io/datatypes"

type AppMultipleSimulateData struct {
	AppPageId                 string             `json:"app_page_id" binding:"required"`
	SingleSimulationInputData map[string]float64 `json:"single_simulation_input_data" binding:""`
}

type GetSimResData struct {
	AppPageId string   `json:"app_page_id" binding:"required"`
	Variable  []string `json:"variable" binding:"required"`
}

type GetReleaseResData struct {
	AppPageId string `json:"app_page_id" binding:"required"`
}

type ReleaseData struct {
	AppPageId string `json:"app_page_id" binding:"required"`
}

type CreateAppSpaceData struct {
	SpaceName        string `json:"space_name" binding:"required"`
	SpaceDescription string `json:"space_description,omitempty" binding:""`
	Background       string `json:"space_background,omitempty" binding:""`
	Icon             string `json:"space_icon,omitempty" binding:"required"`
	IconColor        string `json:"space_icon_color,omitempty" binding:"required"`
}

type EditAppSpaceData struct {
	SpaceId     string `json:"space_id" binding:"required"`
	SpaceName   string `json:"space_name" binding:"required"`
	Description string `json:"space_description" binding:"required"`
	Background  string `json:"space_background" binding:"required"`
	Icon        string `json:"space_icon" binding:"required"`
	IconColor   string `json:"space_icon_color" binding:"required"`
	Collect     bool   `json:"space_collect" binding:""`
}

type AppSpaceCollectData struct {
	SpaceId []string `json:"space_id" binding:"required"`
	Collect bool     `json:"collect" binding:""`
}

type DeleteAppSpaceData struct {
	SpaceId []string `json:"space_id" binding:"required"`
}

type CreateAppPageData struct {
	SpaceId  string `json:"space_id" binding:"required"`
	PageName string `json:"name" binding:"required,max=16"`
	Tag      string `json:"tag" binding:"required,max=10"`
	PageType string `json:"page_type" binding:"required"`
}

type EditAppPageData struct {
	SpaceId  string `json:"space_id" binding:"required"`
	PageId   string `json:"page_id" binding:"required"`
	PageName string `json:"page_name,omitempty" binding:""`
	Tag      string `json:"tag,omitempty" binding:""`
}

type EditAppPageDesignData struct {
	Id              string `json:"page_id" binding:"required"`
	AppSpaceId      string `json:"space_id" binding:"required"`
	PageWidth       int    `json:"page_width,omitempty" binding:""`
	PageHeight      int    `json:"page_height,omitempty" binding:""`
	Background      string `json:"background,omitempty" binding:""`
	BackgroundColor string `json:"background_color,omitempty" binding:""`
}

type DeleteAppPageData struct {
	SpaceId string   `json:"space_id" binding:"required"`
	PageId  []string `json:"page_id" binding:"required"`
}

type CreatePageComponentData struct {
	SpaceId            string         `json:"space_id" binding:"required"`
	PageId             string         `json:"page_id" binding:"required"`
	Type               string         `json:"type" binding:"required"`
	InputOutput        datatypes.JSON `json:"input_output,omitempty" binding:""`
	Width              int            `json:"width" binding:""`
	Height             int            `json:"height" binding:""`
	PositionX          int            `json:"position_x" binding:""`
	PositionY          int            `json:"position_y" binding:""`
	Angle              int            `json:"angle" binding:""`
	HorizontalFlip     bool           `json:"horizontal_flip" binding:""`
	VerticalFlip       bool           `json:"vertical_flip" binding:""`
	Opacity            int            `json:"opacity" binding:""`
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
	Max                float64        `json:"max" binding:""`
	Min                float64        `json:"min" binding:""`
	Interval           float64        `json:"interval" binding:""`
}

type EditPageComponent struct {
	Id                 string         `json:"id" binding:"required"`
	SpaceId            string         `json:"space_id" binding:"required"`
	PageId             string         `json:"page_id" binding:"required"`
	Type               string         `json:"type,omitempty" binding:""`
	InputName          string         `json:"input_name,omitempty" binding:""`
	Width              int            `json:"width,omitempty" binding:""`
	Height             int            `json:"height,omitempty" binding:""`
	PositionX          int            `json:"position_x,omitempty" binding:""`
	PositionY          int            `json:"position_y,omitempty" binding:""`
	Angle              int            `json:"angle,omitempty" binding:""`
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
	Max                float64        `json:"max" binding:""`
	Min                float64        `json:"min" binding:""`
	Interval           float64        `json:"interval" binding:""`
}

type ConfigEditPageComponentData struct {
	Id                 string         `json:"id" binding:"required"`
	SpaceId            string         `json:"space_id" binding:"required"`
	PageId             string         `json:"page_id" binding:"required"`
	Type               string         `json:"type,omitempty" binding:""`
	Width              int            `json:"width,omitempty" binding:""`
	Height             int            `json:"height,omitempty" binding:""`
	PositionX          int            `json:"position_x,omitempty" binding:""`
	PositionY          int            `json:"position_y,omitempty" binding:""`
	Angle              int            `json:"angle,omitempty" binding:""`
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

type EditPageComponentData struct {
	Id        string         `json:"id" binding:"required"`
	SpaceId   string         `json:"space_id" binding:"required"`
	PageId    string         `json:"page_id" binding:"required"`
	InputName string         `json:"input_name,omitempty" binding:""`
	Max       float64        `json:"max" binding:""`
	Min       float64        `json:"min" binding:""`
	Interval  float64        `json:"interval" binding:""`
	Option    datatypes.JSON `json:"option,omitempty" binding:""`
}

type DeletePageComponentData struct {
	PageId         string   `json:"page_id" binding:"required"`
	ComponentsList []string `json:"component_id" binding:"required"`
}

type SetPageInputOutputData struct {
	SpaceId      string         `json:"space_id" binding:"required"`
	PageId       string         `json:"page_id" binding:"required"`
	DataSourceId string         `json:"data_source_id" binding:"required"`
	Input        datatypes.JSON `json:"input,omitempty" binding:""`
	Output       datatypes.JSON `json:"output,omitempty" binding:""`
}

type SetPageComponentsInputOutputData struct {
	Id        string         `json:"id" binding:"required"`
	PageId    string         `json:"page_id" binding:"required"`
	InputName string         `json:"input_name" binding:""`
	Output    datatypes.JSON `json:"output" binding:""`
	Max       float64        `json:"max" binding:""`
	Min       float64        `json:"min" binding:""`
	Interval  float64        `json:"interval" binding:""`
}

type DataSourceRenameData struct {
	DataSourceID string `json:"data_source_id" binding:"required"`
	NewName      string `json:"new_name" binding:"required"`
}

type DeleteDatasourceData struct {
	DataSourceID []string `json:"data_source_id" binding:"required"`
}

type CreateComponentBasesData struct {
	TopLevelName       string         `json:"top_level_name" binding:""`
	SecondLevelName    string         `json:"second_level_name" binding:""`
	Type               string         `json:"type" binding:""`
	Width              int            `json:"width" binding:""`
	Height             int            `json:"height" binding:""`
	Margin             int            `json:"margin" binding:""`
	PositionX          int            `json:"position_x" binding:""`
	PositionY          int            `json:"position_y" binding:""`
	Angle              int            `json:"angle" binding:""`
	HorizontalFlip     bool           `json:"horizontal_flip" binding:""`
	VerticalFlip       bool           `json:"vertical_flip" binding:""`
	Opacity            int            `json:"opacity" binding:""`
	OtherConfiguration datatypes.JSON `json:"other_configuration" binding:""`
}

type SetPageAlignmentLineData struct {
	PageId           string         `json:"page_id" binding:"required"`
	AlignmentLineMap datatypes.JSON `json:"alignment_line" binding:"required"`
}

type ModelStateMessageReadData struct {
	AppPageId   string `json:"app_page_id" binding:"required"`
	MessageType string `json:"message_type" binding:"required"`
}
