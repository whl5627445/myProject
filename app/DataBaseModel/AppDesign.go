package DataBaseModel

import (
	"time"

	"gorm.io/datatypes"

	"gorm.io/gorm"
)

type AppSpace struct {
	ID            string         `gorm:"primaryKey"`
	SpaceName     string         `gorm:"column:space_name"`
	UserName      string         `gorm:"column:username"`
	Description   string         `gorm:"column:description"`
	Background    string         `gorm:"column:background;default:\"static/UserFiles/Images/default_back_ground.png\""`
	Icon          string         `gorm:"column:icon"`
	IconColor     string         `gorm:"column:icon_color"`
	Collect       bool           `gorm:"column:collect"`
	Release       bool           `gorm:"column:is_release"`
	CreatedAt     *time.Time     `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt     *time.Time     `gorm:"column:update_time"`
	LastLoginTime int64          `gorm:"column:last_login_time"`
	Deleted       gorm.DeletedAt `gorm:"column:deleted_at"`
}

type AppDataSource struct {
	ID                string         `gorm:"primaryKey"`
	UserName          string         `gorm:"column:username"`
	UserSpaceId       string         `gorm:"column:user_space_id"`
	PackageId         string         `gorm:"column:package_id"`
	ModelName         string         `gorm:"column:model_name"`
	CompileType       string         `gorm:"column:compile_type"`
	CompilePath       string         `gorm:"column:compile_path"`
	CompileStatus     int64          `gorm:"column:compile_status"`
	CompileStartTime  int64          `gorm:"column:compile_start_time"`
	CompileStopTime   int64          `gorm:"column:compile_stop_time"`
	GroundName        string         `gorm:"column:ground_name"`
	DataSourceName    string         `gorm:"column:data_source_name"`
	ExperimentId      string         `gorm:"column:experiment_id"`
	EnvModelData      datatypes.JSON `gorm:"column:env_model_data"`
	StartTime         string         `gorm:"column:start_time"`
	StopTime          string         `gorm:"column:stop_time"`
	Method            string         `gorm:"column:method"`
	NumberOfIntervals string         `gorm:"column:number_intervals"`
	Tolerance         string         `gorm:"column:tolerance"`
	CreatedAt         *time.Time     `gorm:"column:create_time;autoCreateTime"`
	Deleted           gorm.DeletedAt `gorm:"column:deleted_at"`
}

type AppPage struct {
	ID              string         `gorm:"primaryKey"`
	UserName        string         `gorm:"column:username"`
	AppSpaceId      string         `gorm:"column:app_space_id"`
	Input           datatypes.JSON `gorm:"column:input"`
	Output          datatypes.JSON `gorm:"column:output"`
	PageName        string         `gorm:"column:page_name"`
	PagePath        string         `gorm:"column:page_path"`
	DataSourceId    string         `gorm:"column:data_source_id"`
	PageWidth       int            `gorm:"column:page_width"`
	PageHeight      int            `gorm:"column:page_height"`
	Background      string         `gorm:"column:background"`
	Color           string         `gorm:"column:color"`
	BackgroundColor string         `gorm:"column:background_color"`
	Release         bool           `gorm:"column:is_release"`
	MulResultPath   string         `gorm:"column:mul_result_path"`
	SimulateState   int            `gorm:"column:simulate_state"`
	ReleaseState    int            `gorm:"column:release_state"`
	CreatedAt       *time.Time     `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt       *time.Time     `gorm:"column:update_time"`
	Deleted         gorm.DeletedAt `gorm:"column:deleted_at"`
}

type AppComponentBases struct {
	ID                 string         `gorm:"primaryKey"`
	PageId             string         `gorm:"column:page_id"`
	InputOutput        datatypes.JSON `gorm:"column:input_output"`
	Type               string         `gorm:"column:type"`
	Width              int            `gorm:"column:width"`
	Height             int            `gorm:"column:height"`
	PositionX          float64        `gorm:"column:position_x"`
	PositionY          float64        `gorm:"column:position_y"`
	Angle              float64        `gorm:"column:angle"`
	HorizontalFlip     bool           `gorm:"column:horizontal_flip"`
	VerticalFlip       bool           `gorm:"column:vertical_flip"`
	Opacity            int            `gorm:"column:opacity"`
	OtherConfiguration datatypes.JSON `gorm:"column:other_configuration"`
	CreatedAt          *time.Time     `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt          *time.Time     `gorm:"column:update_time"`
	Deleted            gorm.DeletedAt `gorm:"column:deleted_at"`
}

type AppPageComponent struct {
	ID                 string         `gorm:"primaryKey"`
	PageId             string         `gorm:"column:page_id"`
	InputName          string         `gorm:"column:input_name"`
	Output             datatypes.JSON `gorm:"column:output"`
	Max                float64        `gorm:"column:max"`
	Min                float64        `gorm:"column:min"`
	Interval           float64        `gorm:"column:interval"`
	Type               string         `gorm:"column:type"`
	Width              int            `gorm:"column:width"`
	Height             int            `gorm:"column:height"`
	PositionX          int            `gorm:"column:position_x"`
	PositionY          int            `gorm:"column:position_y"`
	Angle              int            `gorm:"column:angle"`
	HorizontalFlip     bool           `gorm:"column:horizontal_flip"`
	VerticalFlip       bool           `gorm:"column:vertical_flip"`
	Opacity            int            `gorm:"column:opacity"`
	OtherConfiguration datatypes.JSON `gorm:"column:other_configuration"`
	CreatedAt          *time.Time     `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt          *time.Time     `gorm:"column:update_time"`
	Deleted            gorm.DeletedAt `gorm:"column:deleted_at"`
	ZIndex             int            `gorm:"column:z_index"`
	Styles             datatypes.JSON `gorm:"column:styles"`
	Events             datatypes.JSON `gorm:"column:events"`
	ChartConfig        datatypes.JSON `gorm:"column:chart_config"`
	Option             datatypes.JSON `gorm:"column:option"`
	ComponentPath      string         `gorm:"column:component_path"`
	Hide               bool           `gorm:"column:hide"`
	Lock               bool           `gorm:"column:lock"`
	IsGroup            bool           `gorm:"column:is_group"`
}
