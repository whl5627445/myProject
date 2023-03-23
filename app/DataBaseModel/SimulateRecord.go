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
	ExperimentId            string         `gorm:"column:experiment_id"`
	UserName                string         `gorm:"column:username"`
	SimulateModelName       string         `gorm:"column:simulate_model_name"`
	SimulateModelResultPath string         `gorm:"column:simulate_model_result_path"`
	SimulateResultStr       string         `gorm:"column:simulate_result_str"`
	SimulateStatus          string         `gorm:"column:simulate_status"`
	SimulateStartTime       int64          `gorm:"column:simulate_start_time"`
	SimulateEndTime         int64          `gorm:"column:simulate_end_time"`
	FmiVersion              string         `gorm:"column:fmi_version;default:\"\""`
	Description             string         `gorm:"column:description;default:\"\""`
	StartTime               string         `gorm:"column:start_time;default:\"\""`
	StopTime                string         `gorm:"column:stop_time;default:\"\""`
	Method                  string         `gorm:"column:method;default:\"\""`
	SimulateType            string         `gorm:"column:simulate_type;default:\"\""`
	NumberOfIntervals       string         `gorm:"column:number_intervals;default:\"\""`
	Intervals               string         `gorm:"column:intervals;default:\"\""`
	StepSize                string         `gorm:"column:step_size;default:\"\""`
	Tolerance               string         `gorm:"column:tolerance;default:\"\""`
	Solver                  string         `gorm:"column:solver;default:\"\""`
	OutputFormat            string         `gorm:"column:output_format;default:\"\""`
	VariableFilter          string         `gorm:"column:variable_filter;default:\"\""`
	SimulateStart           bool           `gorm:"column:simulate_start"`
	AnotherName             string         `gorm:"column:another_name"`
	CreatedAt               *time.Time     `gorm:"column:create_time;autoCreateTime"`
	Deleted                 gorm.DeletedAt `gorm:"column:deleted_at"`
}

type YssimExperimentRecord struct {
	ID                string         `gorm:"primaryKey"`
	PackageId         string         `gorm:"column:package_id"`
	UserspaceId       string         `gorm:"column:userspace_id"`
	UserName          string         `gorm:"column:username"`
	ExperimentName    string         `gorm:"column:experiment_name"`
	ModelName         string         `gorm:"column:model_name"`
	ModelVarData      datatypes.JSON `gorm:"column:model_var_data;TYPE:json"`
	StartTime         string         `gorm:"column:start_time"`
	StopTime          string         `gorm:"column:stop_time"`
	Method            string         `gorm:"column:method;default:\"\""`
	SimulateType      string         `gorm:"column:simulate_type;default:\"\""`
	NumberOfIntervals string         `gorm:"column:number_intervals;default:\"\""`
	Tolerance         string         `gorm:"column:tolerance;default:\"\""`
	Interval          string         `gorm:"column:interval;default:\"\""`
	CreatedAt         *time.Time     `gorm:"column:create_time;autoCreateTime"`
	Deleted           gorm.DeletedAt `gorm:"column:deleted_at"`
}
