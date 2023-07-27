package DataBaseModel

import (
	"time"

	"gorm.io/datatypes"

	"gorm.io/gorm"
)

type AppSpace struct {
	ID            string         `gorm:"primaryKey;type:varchar(128);comment:应用空间唯一识别标识"`
	SpaceName     string         `gorm:"column:space_name;type:varchar(32);comment:应用空间名称"`
	UserName      string         `gorm:"column:username;type:varchar(32);comment:用户名"`
	Description   string         `gorm:"column:description;type:varchar(128);comment:应用空间描述简介"`
	Background    string         `gorm:"column:background;default:\"static/UserFiles/Images/default_back_ground.png\";type:varchar(128);comment:应用空间背景图片"`
	Icon          string         `gorm:"column:icon;type:varchar(32);comment:应用空间图标"`
	IconColor     string         `gorm:"column:icon_color;type:varchar(64);comment:应用空间图标颜色"`
	Collect       bool           `gorm:"column:collect;type:bool;comment:应用空间收藏"`
	Release       bool           `gorm:"column:is_release;type:bool;comment:应用空间是否发布"`
	CreatedAt     *time.Time     `gorm:"column:create_time;autoCreateTime;comment:应用空间创建时间"`
	UpdatedAt     *time.Time     `gorm:"column:update_time;comment:应用空间更新时间"`
	LastLoginTime int64          `gorm:"column:last_login_time;type:int;comment:应用空间最近打开时间"`
	Deleted       gorm.DeletedAt `gorm:"column:deleted_at;comment:应用空间删除时间"`
}

type AppDataSource struct {
	ID                string         `gorm:"primaryKey;type:varchar(128);comment:应用数据源唯一识别标识"`
	UserName          string         `gorm:"column:username;type:varchar(32);comment:用户名"`
	UserSpaceId       string         `gorm:"column:user_space_id;type:varchar(128);comment:用户空间唯一识别标识"`
	PackageId         string         `gorm:"column:package_id;type:varchar(128);comment:导出数据源模型所在package的唯一识别标识"`
	ModelName         string         `gorm:"column:model_name;type:varchar(128);comment:导出数据源的模型名称"`
	CompileType       string         `gorm:"column:compile_type;type:varchar(128);default:\"OM\";comment:导出数据源所用的编译器类型，有OM与dymola"`
	CompilePath       string         `gorm:"column:compile_path;type:varchar(128);comment:编译好的文件存放路径"`
	ZipMoPath         string         `gorm:"column:zip_mo_path;type:varchar(128);comment:数据源模型所在的mo文件路径"`
	CompileStatus     int64          `gorm:"column:compile_status;type:int;comment:编译结果状态码，1(初始化)、2(运行中)、3(失败)、4(成功)"`
	CompileStartTime  int64          `gorm:"column:compile_start_time;type:int;comment:编译开始时间"`
	CompileStopTime   int64          `gorm:"column:compile_stop_time;type:int;comment:编译结束时间"`
	GroupName         string         `gorm:"column:group_name;type:varchar(32);comment:分组名称"`
	DataSourceName    string         `gorm:"column:data_source_name;type:varchar(128);comment:数据源别名"`
	ExperimentId      string         `gorm:"column:experiment_id;type:varchar(32);comment:仿真实验唯一标识"`
	EnvModelData      datatypes.JSON `gorm:"column:env_model_data;type:json;comment:数据源依赖相关"`
	StartTime         string         `gorm:"column:start_time;type:varchar(32);comment:仿真设置当中的开始时间"`
	StopTime          string         `gorm:"column:stop_time;type:varchar(32);comment:仿真设置当中的结束时间"`
	Method            string         `gorm:"column:method;type:varchar(32);comment:仿真求解方法"`
	NumberOfIntervals string         `gorm:"column:number_intervals;type:varchar(32);comment:仿真间隔数与步长"`
	Tolerance         string         `gorm:"column:tolerance;type:varchar(32);comment:仿真间隔数与步长"`
	CreatedAt         *time.Time     `gorm:"column:create_time;autoCreateTime;comment:记录创建时间"`
	Deleted           gorm.DeletedAt `gorm:"column:deleted_at;comment:记录删除时间"`
}

type AppPage struct {
	ID                     string         `gorm:"primaryKey;type:varchar(128);comment:应用页面唯一标识"`
	UserName               string         `gorm:"column:username;type:varchar(32);comment:用户名"`
	PageType               string         `gorm:"column:page_type;type:varchar(32);comment:页面类型"`
	AppSpaceId             string         `gorm:"column:app_space_id;type:varchar(128);comment:页面所属应用的唯一标识"`
	Input                  datatypes.JSON `gorm:"column:input;type:json;comment:页面的输入相关数据，在webApp类型页面中使用"`
	Output                 datatypes.JSON `gorm:"column:output;type:json;comment:页面的输入相关数据，在webApp类型页面中使用"`
	PageName               string         `gorm:"column:page_name;type:varchar(32);comment:页面的输入相关数据，在webApp类型页面中使用"`
	PagePath               string         `gorm:"column:page_path;type:varchar(32);comment:页面路径"`
	DataSourceId           string         `gorm:"column:data_source_id;type:varchar(128);comment:页面使用到的数据源唯一识别id，在webApp类型页面中使用"`
	PageWidth              int            `gorm:"column:page_width;type:int;comment:页面宽度"`
	PageHeight             int            `gorm:"column:page_height;type:int;comment:页面高度"`
	Background             string         `gorm:"column:background;type:varchar(64);comment:页面背景图片"`
	Color                  string         `gorm:"column:color;type:varchar(32);comment:页面颜色"`
	BackgroundColor        string         `gorm:"column:background_color;type:varchar(32);comment:页面背景颜色"`
	Release                bool           `gorm:"column:is_release;type:bool;comment:页面是否发布"`
	IsMulSimulate          bool           `gorm:"column:is_mul_simulate;type:bool;comment:是否进行过多轮仿真"`
	IsPreview              bool           `gorm:"column:is_preview;type:bool"`
	MulResultPath          string         `gorm:"column:mul_result_path;type:varchar(128)"`
	MulSimulateState       int            `gorm:"column:mul_sim_state;type:int;default:0"`
	MulSimulateMessageRead bool           `gorm:"column:mul_sim_message_read;type:bool;default:true"`
	MulSimulateErr         string         `gorm:"column:mul_sim_err;type:longtext"`
	MulSimulateTime        int            `gorm:"column:mul_sim_time;type:int"`
	ReleaseMessageRead     bool           `gorm:"column:release_message_read;type:bool;default:true"`
	ReleaseErr             string         `gorm:"column:release_err;type:longtext"`
	ReleaseState           int            `gorm:"column:release_state;type:int;default:0"`
	ReleaseTime            int            `gorm:"column:release_time;type:int"`
	NamingOrder            datatypes.JSON `gorm:"column:naming_order;type:json"`
	AlignmentLine          datatypes.JSON `gorm:"column:alignment_line;type:json"`
	CreatedAt              *time.Time     `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt              *time.Time     `gorm:"column:update_time"`
	Deleted                gorm.DeletedAt `gorm:"column:deleted_at"`
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

type AppPageComponentsPreview struct {
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

type AppPageComponentsRelease struct {
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
