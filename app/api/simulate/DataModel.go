package API

type ResponseData struct {
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Status int         `json:"status"` // 正常是0，系统级错误是1， 用戶错误是2
	Err    string      `json:"err"`
}

type SetSimulationOptionsData struct {
	PackageId         string `json:"package_id" binding:"required"`
	ModelName         string `json:"model_name" binding:"required"`
	StartTime         string `json:"startTime" binding:"required"`
	StopTime          string `json:"stopTime" binding:"required"`
	Tolerance         string `json:"tolerance" binding:"required"`
	NumberOfIntervals string `json:"numberOfIntervals" binding:"required"`
	Interval          string `json:"interval" binding:"required"`
}

type ModelSimulateData struct {
	PackageId         string `json:"package_id" binding:"required"`
	ModelName         string `json:"model_name" binding:"required"`
	SimulateType      string `json:"simulate_type" binding:"required"`
	StartTime         string `json:"startTime" binding:"required"`
	StopTime          string `json:"stopTime" binding:"required"`
	Tolerance         string `json:"tolerance" binding:"required"`
	NumberOfIntervals string `json:"numberOfIntervals" binding:"required"`
	//Interval          string `json:"interval" binding:"required"`
	Method string `json:"method" binding:""`
}

type ModelSimulateResultData struct {
	RecordId string `json:"id" binding:"required"`
	Variable string `json:"variable" binding:"required"`
	S1       string `json:"s1" binding:""`
	S2       string `json:"s2" binding:""`
}

type ExperimentCreateData struct {
	PackageId       string            `json:"package_id" binding:"required"`
	ModelName       string            `json:"model_name" binding:"required"`
	ExperimentName  string            `json:"experiment_name" binding:"required"`
	SimulateVarData map[string]string `json:"simulate_var_data" binding:"required"`
}

type ExperimentDeleteData struct {
	ExperimentId string `json:"experiment_id" binding:"required"`
}

type ExperimentEditData struct {
	ExperimentId    string            `json:"experiment_id" binding:"required"`
	ExperimentName  string            `json:"experiment_name" binding:""`
	SimulateVarData map[string]string `json:"simulate_var_data" binding:"required"`
}

type ModelCodeSaveData struct {
	PackageId string `json:"package_id" binding:"required"`
	//PackageName string `json:"package_name" binding:"required"`
}
