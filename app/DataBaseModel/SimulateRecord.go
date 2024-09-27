package DataBaseModel

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type YssimSimulateRecord struct {
	ID                      string         `gorm:"index;primaryKey;type:varchar(128);comment:仿真记录唯一识别标识"`
	PackageId               string         `gorm:"index;column:package_id;type:varchar(128);comment:package唯一识别标识"`
	UserspaceId             string         `gorm:"index;column:userspace_id;type:varchar(128);comment:用户空间唯一识别标识"`
	ExperimentId            string         `gorm:"index;column:experiment_id;type:varchar(128);comment:仿真实验唯一识别标识"`
	UserName                string         `gorm:"index;column:username;type:varchar(32);comment:用户名"`
	SimulateModelName       string         `gorm:"column:simulate_model_name;type:varchar(256);comment:仿真模型名称"`
	SimulateModelResultPath string         `gorm:"column:simulate_model_result_path;type:varchar(256);comment:仿真结果存储路径"`
	SimulateResultStr       string         `gorm:"column:simulate_result_str;comment:仿真结果输出字符串"`
	SimulateStatus          string         `gorm:"column:simulate_status;type:varchar(32);comment:仿真状态"`
	SimulateStartTime       int64          `gorm:"column:simulate_start_time;type:int;comment:仿真动作开始执行时间"`
	SimulateEndTime         int64          `gorm:"column:simulate_end_time;type:int;comment:仿真动作结束执行时间"`
	Percentage              int64          `gorm:"column:percentage;type:int;comment:仿真进度(0-100)"`
	StartTime               string         `gorm:"column:start_time;default:\"\";type:varchar(32);comment:仿真配置参数中的开始时间"`
	StopTime                string         `gorm:"column:stop_time;default:\"\";type:varchar(32);comment:仿真配置参数中的结束时间"`
	Method                  string         `gorm:"column:method;default:\"\";type:varchar(32);comment:仿真方法"`
	SimulateType            string         `gorm:"column:simulate_type;default:\"\";type:varchar(32);comment:仿真编译求解器类型，OM和DM"`
	NumberOfIntervals       string         `gorm:"column:number_intervals;default:\"\";type:varchar(32);comment:仿真间隔数和间隔"`
	Intervals               string         `gorm:"column:intervals;default:\"\";type:varchar(32);comment:仿真间隔"`
	Tolerance               string         `gorm:"column:tolerance;default:\"\";type:varchar(32);comment:仿真时的容差"`
	SimulateStart           bool           `gorm:"column:simulate_start;type:int;comment:仿真任务是否开始"`
	AnotherName             string         `gorm:"column:another_name;type:varchar(32);comment:仿真记录别名"`
	TaskId                  string         `gorm:"column:task_id;type:varchar(128);comment:仿真任务唯一标识"`
	EnvModelData            datatypes.JSON `gorm:"column:env_model_data;TYPE:json;type:json;comment:仿真模型的依赖"`
	CreatedAt               *time.Time     `gorm:"column:create_time;autoCreateTime;comment:创建时间"`
	Deleted                 gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间"`
	//StepSize          string `gorm:"column:step_size;default:\"\";type:varchar(32)"`
	//Solver            string `gorm:"column:solver;default:\"\";type:varchar(32)"`
	//OutputFormat      string `gorm:"column:output_format;default:\"\";type:varchar(32)"`
	//VariableFilter          string         `gorm:"column:variable_filter;default:\"\";type:varchar(128)"`
	//FmiVersion              string         `gorm:"column:fmi_version;default:\"\";type:varchar(16)"`
	//Description             string         `gorm:"column:description;default:\"\";type:varchar(128)"`
}

type YssimExperimentRecord struct {
	ID                string         `gorm:"index;primaryKey;type:varchar(128);comment:仿真实验记录唯一标识"`
	PackageId         string         `gorm:"index;column:package_id;type:varchar(128);comment:package唯一标识"`
	UserspaceId       string         `gorm:"index;column:userspace_id;type:varchar(128);comment:用户空间唯一标识"`
	UserName          string         `gorm:"index;column:username;type:varchar(32);comment:用户名"`
	ExperimentName    string         `gorm:"column:experiment_name;type:varchar(32);comment:仿真实验名称"`
	ModelName         string         `gorm:"column:model_name;type:varchar(256);comment:模型名称"`
	ModelVarData      datatypes.JSON `gorm:"column:model_var_data;type:json;comment:模型参数数据"`
	IsFullModelVar    bool           `gorm:"column:is_full_model_var;type:bool;comment:是否是全量组件参数"`
	StartTime         string         `gorm:"column:start_time;type:varchar(32);comment:仿真配置当中的开始时间"`
	StopTime          string         `gorm:"column:stop_time;type:varchar(32);comment:仿真配置当中的结束时间"`
	Method            string         `gorm:"column:method;default:\"\";type:varchar(32);comment:仿真方法"`
	SimulateType      string         `gorm:"column:simulate_type;default:\"\";type:varchar(32);comment:仿真编译器类型，OM和DM"`
	NumberOfIntervals string         `gorm:"column:number_intervals;default:\"\";type:varchar(32);comment:仿真间隔数和间隔"`
	Tolerance         string         `gorm:"column:tolerance;default:\"\";type:varchar(32);comment:仿真时的容差"`
	Interval          string         `gorm:"column:interval;default:\"\";type:varchar(32);comment:仿真间隔"`
	CreatedAt         *time.Time     `gorm:"column:create_time;autoCreateTime;comment:创建时间"`
	Deleted           gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间"`
}
