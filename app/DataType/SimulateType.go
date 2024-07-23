package DataType

import (
	"gorm.io/datatypes"
)

type SetSimulationOptionsData struct {
	PackageId         string `json:"package_id" binding:"required"`
	ModelName         string `json:"model_name" binding:"required"`
	StartTime         string `json:"startTime" binding:"required"`
	StopTime          string `json:"stopTime" binding:"required"`
	Tolerance         string `json:"tolerance" binding:""`
	NumberOfIntervals string `json:"numberOfIntervals" binding:"required"`
	Interval          string `json:"interval" binding:"required"`
	SimulationFlags   string `json:"method" binding:"required"`
	SimulateType      string `json:"simulate_type" binding:"required"`
}

type ModelSimulateData struct {
	PackageId         string `json:"package_id" binding:"required"`
	ModelName         string `json:"model_name" binding:"required"`
	SimulateType      string `json:"simulate_type" binding:"required"`
	StartTime         string `json:"startTime" binding:"required"`
	StopTime          string `json:"stopTime" binding:"required"`
	Tolerance         string `json:"tolerance" binding:"required"`
	NumberOfIntervals string `json:"numberOfIntervals" binding:"required"`
	Interval          string `json:"interval" binding:"required"`
	Method            string `json:"method" binding:""`
	ExperimentId      string `json:"experiment_id" binding:""`
}

type ModelSimulateResultData struct {
	RecordId []string `json:"id" binding:"required"`
	Variable string   `json:"variable" binding:"required"`
	S1       string   `json:"s1" binding:""`
	S2       string   `json:"s2" binding:""`
}

type ModelSimulateResultSingularData struct {
	RecordId string `json:"id" binding:"required"`
	Variable string `json:"variable" binding:"required"`
	S1       string `json:"s1" binding:""`
	S2       string `json:"s2" binding:""`
}

type SimulateTerminateData struct {
	RecordId string `json:"record_id" binding:"required"`
}

type ExperimentExistsData struct {
	PackageId       string            `json:"package_id" binding:"required"`
	SimulateVarData map[string]string `json:"simulate_var_data" binding:"required"`
	ModelVarData    datatypes.JSON    `json:"model_var_data" binding:""`
}

type ExperimentCreateData struct {
	PackageId       string            `json:"package_id" binding:"required"`
	ModelName       string            `json:"model_name" binding:"required"`
	ExperimentName  string            `json:"experiment_name" binding:"required"`
	SimulateVarData map[string]string `json:"simulate_var_data" binding:"required"`
	ModelVarData    datatypes.JSON    `json:"model_var_data" binding:""`
}

type ExperimentDeleteData struct {
	ExperimentId string `json:"experiment_id" binding:"required"`
}

type ExperimentEditData struct {
	PackageId       string            `json:"package_id" binding:"required"`
	ModelName       string            `json:"model_name" binding:"required"`
	ExperimentId    string            `json:"experiment_id" binding:"required"`
	ExperimentName  string            `json:"experiment_name" binding:""`
	SimulateVarData map[string]string `json:"simulate_var_data" binding:"required"`
	ModelVarData    datatypes.JSON    `json:"model_var_data" binding:""`
}

type ExperimentNameEditData struct {
	PackageId         string `json:"package_id" binding:"required"`
	ModelName         string `json:"model_name" binding:"required"`
	ExperimentId      string `json:"experiment_id" binding:"required"`
	NewExperimentName string `json:"new_experiment_name" binding:"required"`
}

type ExperimentCompareData struct {
	PackageId        string   `json:"package_id" binding:"required"`
	ExperimentIdList []string `json:"experiment_id_list" binding:"required"`
}

type SimulateRecordCompareData struct {
	PackageId    string   `json:"package_id" binding:"required"`
	RecordIdList []string `json:"record_id_list" binding:"required"`
}

type SnapshotCreatData struct {
	SnapshotName      string         `json:"snapshot_name" binding:"required"`
	ModelName         string         `json:"model_name" binding:"required"`
	ComponentName     string         `json:"component_name" binding:""`
	PackageId         string         `json:"package_id" binding:"required"`
	ModelVarData      datatypes.JSON `json:"model_var_data" binding:""`
	ExperimentId      string         `json:"experiment_id" binding:""`
	SimulateVarData   datatypes.JSON `json:"simulate_var_data" binding:""`
	SimulateResultId  string         `json:"simulate_result_id" binding:"required"`
	SimulateResultObj datatypes.JSON `json:"simulate_result_obj" binding:""`
}

type SnapshotDeleteData struct {
	SnapshotId string `json:"snapshot_id" binding:"required"`
}

// SnapshotEditData 的字段数和名称必须和数据库模型YssimSnapshots一致
type SnapshotEditData struct {
	ID                string         `json:"snapshot_id"  binding:"required"`
	SpaceId           string         `json:"space_id"  binding:""`
	SnapshotName      string         `json:"snapshot_name" binding:"required"`
	UserName          string         `json:"user_name"  binding:""`
	ModelName         string         `json:"model_name" binding:"required"`
	PackageId         string         `json:"package_id" binding:"required"`
	ComponentName     string         `json:"component_name" binding:""`
	ExperimentId      string         `json:"experiment_id" binding:""`
	SimulateResultId  string         `json:"simulate_result_id" binding:""`
	SimulateResultObj datatypes.JSON `json:"simulate_result_obj" binding:""`
	SimulateVarData   datatypes.JSON `json:"simulate_var_data" binding:""`
	ModelVarData      datatypes.JSON `json:"model_var_data" binding:""`
}

type RecordRenameData struct {
	RecordId       string `json:"record_id" binding:"required"`
	NewAnotherName string `json:"new_another_name" binding:"required"`
}

type CalibrationCompileData struct {
	ID          string `json:"id" binding:"required"`
	UserSpaceId string `json:"user_space_id" binding:"required"`
	PackageId   string `json:"package_id" binding:"required"`
	ModelName   string `json:"model_name" binding:"required"`
}

type CalibrationSimulateData struct {
	ID string `json:"id" binding:"required"`
}
