package API

import (
	"gorm.io/datatypes"
)

type responseData struct {
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Status int         `json:"status"` // 正常是0，系统级错误是1， 用戶错误是2
	Err    string      `json:"err"`
}

type setSimulationOptionsData struct {
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

type modelSimulateData struct {
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

type modelSimulateResultData struct {
	RecordId []string `json:"id" binding:"required"`
	Variable string   `json:"variable" binding:"required"`
	S1       string   `json:"s1" binding:""`
	S2       string   `json:"s2" binding:""`
}

type modelSimulateResultSingularData struct {
	RecordId string `json:"id" binding:"required"`
	Variable string `json:"variable" binding:"required"`
	S1       string `json:"s1" binding:""`
	S2       string `json:"s2" binding:""`
}

type experimentCreateData struct {
	PackageId       string            `json:"package_id" binding:"required"`
	ModelName       string            `json:"model_name" binding:"required"`
	ExperimentName  string            `json:"experiment_name" binding:"required"`
	SimulateVarData map[string]string `json:"simulate_var_data" binding:"required"`
	ModelVarData    datatypes.JSON    `json:"model_var_data" binding:""`
}

type experimentDeleteData struct {
	ExperimentId string `json:"experiment_id" binding:"required"`
}

type experimentEditData struct {
	ExperimentId    string            `json:"experiment_id" binding:"required"`
	ExperimentName  string            `json:"experiment_name" binding:""`
	SimulateVarData map[string]string `json:"simulate_var_data" binding:"required"`
	ModelVarData    datatypes.JSON    `json:"model_var_data" binding:""`
}

type modelCodeSaveData struct {
	PackageId string `json:"package_id" binding:"required"`
	// PackageName string `json:"package_name" binding:"required"`
}

type snapshotCreatData struct {
	SnapshotName      string         `json:"snapshot_name" binding:"required"`
	ModelName         string         `json:"model_name" binding:"required"`
	ComponentName     string         `json:"component_name" binding:""`
	PackageId         string         `json:"package_id" binding:"required"`
	ModelVarData      datatypes.JSON `json:"model_var_data" binding:""`
	ExperimentId      string         `json:"experiment_id" binding:""`
	SimulateVarData   datatypes.JSON `json:"simulate_var_data" binding:""`
	SimulateResultId  string         `json:"simulate_result_id" binding:""`
	SimulateResultObj datatypes.JSON `json:"simulate_result_obj" binding:""`
}

type snapshotDeleteData struct {
	SnapshotId string `json:"snapshot_id" binding:"required"`
}

// snapshotEditData的字段数和名称必须和数据库模型YssimSnapshots一致
type snapshotEditData struct {
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

type recordRenameData struct {
	RecordId       string `json:"record_id" binding:"required"`
	NewAnotherName string `json:"new_another_name" binding:"required"`
}
