package DataBaseModel

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type YssimSimulateRecord struct {
	ID                      string         `gorm:"primaryKey"`
	PackageId               string         `gorm:"column:package_id"`
	UserspaceId             string         `gorm:"column:userspace_id"`
	Username                string         `gorm:"column:username"`
	SimulateModelName       string         `gorm:"column:simulate_model_name"`
	SimulateModelResultPath string         `gorm:"column:simulate_model_result_path"`
	SimulateResultStr       string         `gorm:"column:simulate_result_str"`
	SimulateStatus          string         `gorm:"column:simulate_status"`
	SimulateStartTime       string         `gorm:"column:simulate_start_time"`
	SimulateEndTime         string         `gorm:"column:simulate_end_time"`
	FmiVersion              string         `gorm:"column:fmi_version"`
	Description             string         `gorm:"column:description"`
	StartTime               string         `gorm:"column:start_time"`
	StopTime                string         `gorm:"column:stop_time"`
	Method                  string         `gorm:"column:method"`
	SimulateType            string         `gorm:"column:simulate_type"`
	NumberOfIntervals       string         `gorm:"column:number_intervals"`
	StepSize                string         `gorm:"column:step_size"`
	Tolerance               string         `gorm:"column:tolerance"`
	Solver                  string         `gorm:"column:solver"`
	OutputFormat            string         `gorm:"column:output_format"`
	VariableFilter          string         `gorm:"column:variable_filter"`
	SimulateStart           bool           `gorm:"column:simulate_start"`
	CreatedAt               *time.Time     `gorm:"column:create_time;autoCreateTime"`
	Deleted                 gorm.DeletedAt `gorm:"column:deleted_at"`
}

type YssimExperimentRecord struct {
	ID                string         `gorm:"primaryKey"`
	PackageId         string         `gorm:"column:package_id"`
	UserspaceId       string         `gorm:"column:userspace_id"`
	Username          string         `gorm:"column:username"`
	ExperimentName    string         `gorm:"column:experiment_name"`
	ModelName         string         `gorm:"column:model_name"`
	ModelVarData      datatypes.JSON `gorm:"column:model_var_data;TYPE:json"`
	StartTime         string         `gorm:"column:start_time"`
	StopTime          string         `gorm:"column:stop_time"`
	Method            string         `gorm:"column:method"`
	SimulateType      string         `gorm:"column:simulate_type"`
	NumberOfIntervals string         `gorm:"column:number_intervals"`
	Tolerance         string         `gorm:"column:tolerance"`
	Interval          string         `gorm:"column:interval"`
	CreatedAt         *time.Time     `gorm:"column:create_time;autoCreateTime"`
	Deleted           gorm.DeletedAt `gorm:"column:deleted_at"`
}
