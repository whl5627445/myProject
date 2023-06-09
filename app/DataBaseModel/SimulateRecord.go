package DataBaseModel

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type YssimSimulateRecord struct {
	ID                      string         `gorm:"primaryKey;type:varchar(128)"`
	PackageId               string         `gorm:"column:package_id;type:varchar(128)"`
	UserspaceId             string         `gorm:"column:userspace_id;type:varchar(128)"`
	ExperimentId            string         `gorm:"column:experiment_id;type:varchar(128)"`
	UserName                string         `gorm:"column:username;type:varchar(32)"`
	SimulateModelName       string         `gorm:"column:simulate_model_name;type:varchar(32)"`
	SimulateModelResultPath string         `gorm:"column:simulate_model_result_path;type:varchar(32)"`
	SimulateResultStr       string         `gorm:"column:simulate_result_str;type:varchar(256)"`
	SimulateStatus          string         `gorm:"column:simulate_status;type:varchar(32)"`
	SimulateStartTime       int64          `gorm:"column:simulate_start_time;type:int"`
	SimulateEndTime         int64          `gorm:"column:simulate_end_time;type:int"`
	FmiVersion              string         `gorm:"column:fmi_version;default:\"\";type:varchar(16)"`
	Description             string         `gorm:"column:description;default:\"\";type:varchar(128)"`
	StartTime               string         `gorm:"column:start_time;default:\"\";type:varchar(32)"`
	StopTime                string         `gorm:"column:stop_time;default:\"\";type:varchar(32)"`
	Method                  string         `gorm:"column:method;default:\"\";type:varchar(32)"`
	SimulateType            string         `gorm:"column:simulate_type;default:\"\";type:varchar(32)"`
	NumberOfIntervals       string         `gorm:"column:number_intervals;default:\"\";type:varchar(32)"`
	Intervals               string         `gorm:"column:intervals;default:\"\";type:varchar(32)"`
	StepSize                string         `gorm:"column:step_size;default:\"\";type:varchar(32)"`
	Tolerance               string         `gorm:"column:tolerance;default:\"\";type:varchar(32)"`
	Solver                  string         `gorm:"column:solver;default:\"\";type:varchar(32)"`
	OutputFormat            string         `gorm:"column:output_format;default:\"\";type:varchar(32)"`
	VariableFilter          string         `gorm:"column:variable_filter;default:\"\";type:varchar(128)"`
	SimulateStart           bool           `gorm:"column:simulate_start;type:int"`
	AnotherName             string         `gorm:"column:another_name;type:varchar(32)"`
	EnvModelData            datatypes.JSON `gorm:"column:env_model_data;TYPE:json;type:json"`
	CreatedAt               *time.Time     `gorm:"column:create_time;autoCreateTime"`
	Deleted                 gorm.DeletedAt `gorm:"column:deleted_at"`
}

type YssimExperimentRecord struct {
	ID                string         `gorm:"primaryKey;type:varchar(128)"`
	PackageId         string         `gorm:"column:package_id;type:varchar(128)"`
	UserspaceId       string         `gorm:"column:userspace_id;type:varchar(128)"`
	UserName          string         `gorm:"column:username;type:varchar(32)"`
	ExperimentName    string         `gorm:"column:experiment_name;type:varchar(32)"`
	ModelName         string         `gorm:"column:model_name;type:varchar(32)"`
	ModelVarData      datatypes.JSON `gorm:"column:model_var_data;type:json"`
	StartTime         string         `gorm:"column:start_time;type:varchar(32)"`
	StopTime          string         `gorm:"column:stop_time;type:varchar(32)"`
	Method            string         `gorm:"column:method;default:\"\";type:varchar(32)"`
	SimulateType      string         `gorm:"column:simulate_type;default:\"\";type:varchar(32)"`
	NumberOfIntervals string         `gorm:"column:number_intervals;default:\"\";type:varchar(32)"`
	Tolerance         string         `gorm:"column:tolerance;default:\"\";type:varchar(32)"`
	Interval          string         `gorm:"column:interval;default:\"\";type:varchar(32)"`
	CreatedAt         *time.Time     `gorm:"column:create_time;autoCreateTime"`
	Deleted           gorm.DeletedAt `gorm:"column:deleted_at"`
}
