package DataBaseModel

import (
	"time"

	"gorm.io/datatypes"

	"gorm.io/gorm"
)

type AppSpace struct {
	ID            string         `gorm:"primaryKey;type:varchar(128)"`
	SpaceName     string         `gorm:"column:space_name;type:varchar(32)"`
	UserName      string         `gorm:"column:username;type:varchar(32)"`
	Description   string         `gorm:"column:description;type:varchar(128)"`
	Background    string         `gorm:"column:background;default:\"static/UserFiles/Images/default_back_ground.png\";type:varchar(128)"`
	Icon          string         `gorm:"column:icon;type:varchar(32)"`
	IconColor     string         `gorm:"column:icon_color;type:varchar(64)"`
	Collect       bool           `gorm:"column:collect;type:bool"`
	Release       bool           `gorm:"column:is_release;type:bool"`
	CreatedAt     *time.Time     `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt     *time.Time     `gorm:"column:update_time"`
	LastLoginTime int64          `gorm:"column:last_login_time;type:int"`
	Deleted       gorm.DeletedAt `gorm:"column:deleted_at"`
}

type AppDataSource struct {
	ID                string         `gorm:"primaryKey;type:varchar(128)"`
	UserName          string         `gorm:"column:username;type:varchar(32)"`
	UserSpaceId       string         `gorm:"column:user_space_id;type:varchar(128)"`
	PackageId         string         `gorm:"column:package_id;type:varchar(128)"`
	ModelName         string         `gorm:"column:model_name;type:varchar(128)"`
	CompileType       string         `gorm:"column:compile_type;type:varchar(128)"`
	CompilePath       string         `gorm:"column:compile_path;type:varchar(128)"`
	CompileStatus     int64          `gorm:"column:compile_status;type:int"`
	CompileStartTime  int64          `gorm:"column:compile_start_time;type:int"`
	CompileStopTime   int64          `gorm:"column:compile_stop_time;type:int"`
	GroupName         string         `gorm:"column:group_name;type:varchar(32)"`
	DataSourceName    string         `gorm:"column:data_source_name;type:varchar(128)"`
	ExperimentId      string         `gorm:"column:experiment_id;type:varchar(32)"`
	EnvModelData      datatypes.JSON `gorm:"column:env_model_data;type:json"`
	StartTime         string         `gorm:"column:start_time;type:varchar(32)"`
	StopTime          string         `gorm:"column:stop_time;type:varchar(32)"`
	Method            string         `gorm:"column:method;type:varchar(32)"`
	NumberOfIntervals string         `gorm:"column:number_intervals;type:varchar(32)"`
	Tolerance         string         `gorm:"column:tolerance;type:varchar(32)"`
	CreatedAt         *time.Time     `gorm:"column:create_time;autoCreateTime"`
	Deleted           gorm.DeletedAt `gorm:"column:deleted_at"`
}

type AppPage struct {
	ID                  string         `gorm:"primaryKey;type:varchar(128)"`
	UserName            string         `gorm:"column:username;type:varchar(32)"`
	AppSpaceId          string         `gorm:"column:app_space_id;type:varchar(128)"`
	Input               datatypes.JSON `gorm:"column:input;type:json"`
	Output              datatypes.JSON `gorm:"column:output;type:json"`
	PageName            string         `gorm:"column:page_name;type:varchar(32)"`
	PagePath            string         `gorm:"column:page_path;type:varchar(32)"`
	DataSourceId        string         `gorm:"column:data_source_id;type:varchar(128)"`
	PageWidth           int            `gorm:"column:page_width;type:int"`
	PageHeight          int            `gorm:"column:page_height;type:int"`
	Background          string         `gorm:"column:background;type:varchar(64)"`
	Color               string         `gorm:"column:color;type:varchar(32)"`
	BackgroundColor     string         `gorm:"column:background_color;type:varchar(32)"`
	Release             bool           `gorm:"column:is_release;type:bool"`
	MulResultPath       string         `gorm:"column:mul_result_path;type:varchar(128)"`
	SimulateState       int            `gorm:"column:simulate_state;type:int"`
	SimulateMessageRead bool           `gorm:"column:simulate_message_read;type:bool;default:true"`
	SimulateTime        int            `gorm:"column:simulate_time;type:int"`
	ReleaseMessageRead  bool           `gorm:"column:release_message_read;type:bool;default:true"`
	ReleaseState        int            `gorm:"column:release_state;type:int"`
	ReleaseTime         int            `gorm:"column:release_time;type:int"`
	NamingOrder         datatypes.JSON `gorm:"column:naming_order;type:json"`
	AlignmentLine       datatypes.JSON `gorm:"column:alignment_line;type:json"`
	CreatedAt           *time.Time     `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt           *time.Time     `gorm:"column:update_time"`
	Deleted             gorm.DeletedAt `gorm:"column:deleted_at"`
}

type AppComponentBases struct {
	ID                 string         `gorm:"primaryKey;type:varchar(128)"`
	TopLevelName       string         `gorm:"column:top_level _name;type:varchar(32)"`
	SecondLevelName    string         `gorm:"column:second_level_name;type:varchar(32)"`
	Type               string         `gorm:"column:type;type:varchar(32)"`
	Width              int            `gorm:"column:width;type:int"`
	Height             int            `gorm:"column:height;type:int"`
	Margin             int            `gorm:"column:margin;type:int"`
	PositionX          int            `gorm:"column:position_x;type:int"`
	PositionY          int            `gorm:"column:position_y;type:int"`
	Angle              int            `gorm:"column:angle;type:int"`
	HorizontalFlip     bool           `gorm:"column:horizontal_flip;type:bool"`
	VerticalFlip       bool           `gorm:"column:vertical_flip;type:bool"`
	Opacity            int            `gorm:"column:opacity;type:int"`
	OtherConfiguration datatypes.JSON `gorm:"column:other_configuration;type:json"`
	CreatedAt          *time.Time     `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt          *time.Time     `gorm:"column:update_time"`
	Deleted            gorm.DeletedAt `gorm:"column:deleted_at"`
}

type AppPageComponent struct {
	ID                 string         `gorm:"primaryKey;type:varchar(128)"`
	PageId             string         `gorm:"column:page_id;type:varchar(128)"`
	InputName          string         `gorm:"column:input_name;type:varchar(32)"`
	Output             datatypes.JSON `gorm:"column:output;type:json"`
	Max                float64        `gorm:"column:max;type:float"`
	Min                float64        `gorm:"column:min;type:float"`
	Interval           float64        `gorm:"column:interval;type:float"`
	Type               string         `gorm:"column:type;type:varchar(32)"`
	Width              int            `gorm:"column:width;type:int"`
	Height             int            `gorm:"column:height;type:int"`
	PositionX          int            `gorm:"column:position_x;type:int"`
	PositionY          int            `gorm:"column:position_y;type:int"`
	Angle              int            `gorm:"column:angle;type:int"`
	HorizontalFlip     bool           `gorm:"column:horizontal_flip;type:bool"`
	VerticalFlip       bool           `gorm:"column:vertical_flip;type:bool"`
	Opacity            int            `gorm:"column:opacity;type:int"`
	OtherConfiguration datatypes.JSON `gorm:"column:other_configuration;type:json"`
	CreatedAt          *time.Time     `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt          *time.Time     `gorm:"column:update_time"`
	Deleted            gorm.DeletedAt `gorm:"column:deleted_at"`
	ZIndex             int            `gorm:"column:z_index;type:int"`
	Styles             datatypes.JSON `gorm:"column:styles;type:json"`
	Events             datatypes.JSON `gorm:"column:events;type:json"`
	ChartConfig        datatypes.JSON `gorm:"column:chart_config;type:json"`
	Option             datatypes.JSON `gorm:"column:option;type:json"`
	ComponentPath      string         `gorm:"column:component_path;type:varchar(32)"`
	Hide               bool           `gorm:"column:hide;type:bool"`
	Lock               bool           `gorm:"column:lock;type:bool"`
	IsGroup            bool           `gorm:"column:is_group;type:bool"`
}
