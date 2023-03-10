package API

import "gorm.io/datatypes"

type responseData struct {
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Status int         `json:"status"` // 正常是0，系统级错误是1， 用戶错误是2
	Err    string      `json:"err"`
}

type modelGraphicsData struct {
	PackageId     string `json:"package_id" binding:"required"`
	ModelName     string `json:"model_name" binding:"required"`
	ComponentName string `json:"component_name,omitempty" binding:""`
}

type setComponentModifierValueData struct {
	PackageId      string            `json:"package_id" binding:"required"`
	ModelName      string            `json:"model_name" binding:"required"`
	ParameterValue map[string]string `json:"parameter_value" binding:"required"`
}

type addComponentParametersData struct {
	PackageId     string `json:"package_id" binding:"required"`
	ModelName     string `json:"model_name" binding:"required"`
	ParameterName string `json:"parameter_name" binding:"required"`
	VarType       string `json:"var_type" binding:"required"`
}

type deleteComponentParametersData struct {
	PackageId     string `json:"package_id" binding:"required"`
	ModelName     string `json:"model_name" binding:"required"`
	ParameterName string `json:"parameter_name" binding:"required"`
}

type setComponentPropertiesData struct {
	PackageId        string `json:"package_id" binding:"required"`
	ModelName        string `json:"model_name" binding:"required"`
	OldComponentName string `json:"old_component_name" binding:"required"`
	NewComponentName string `json:"new_component_name" binding:"required"`
	Final            bool   `json:"final" binding:""`
	Protected        bool   `json:"protected" binding:""`
	Replaceable      bool   `json:"replaceable" binding:""`
	Variability      string `json:"variability" binding:""`
	Inner            bool   `json:"inner" binding:""`
	Outer            bool   `json:"outer" binding:""`
	Causality        string `json:"causality" binding:""`
}

type copyClassData struct {
	PackageId       string `json:"package_id" binding:""`
	ParentName      string `json:"parent_name" binding:""`
	ModelName       string `json:"model_name" binding:"required"`
	CopiedClassName string `json:"copied_class_name" binding:"required"`
}

type deleteClassData struct {
	PackageId  string `json:"package_id" binding:"required"`
	ParentName string `json:"parent_name" binding:""`
	ModelName  string `json:"model_name" binding:"required"`
}

type addComponentData struct {
	PackageId        string   `json:"package_id" binding:"required"`
	NewComponentName string   `json:"new_component_name" binding:"required"`
	OldComponentName string   `json:"old_component_name" binding:"required"`
	ModelName        string   `json:"model_name" binding:"required"`
	Origin           string   `json:"origin" binding:"required"`
	Extent           []string `json:"extent" binding:"required"`
	Rotation         int      `json:"rotation" binding:""`
}

type deleteComponentMap struct {
	DeleteType    string `json:"delete_type" binding:"required"`
	ComponentName string `json:"component_name" binding:"required"`
	ModelName     string `json:"model_name" binding:"required"`
	ConnectStart  string `json:"connect_start" binding:"required"`
	ConnectEnd    string `json:"connect_end" binding:"required"`
}

type deleteComponentData struct {
	PackageId     string               `json:"package_id" binding:"required"`
	ComponentList []deleteComponentMap `json:"component_list" binding:"required"`
}

type updateComponentData struct {
	PackageId          string                           `json:"package_id" binding:"required"`
	ConnectionList     []updateConnectionAnnotationData `json:"connection_list" binding:""`
	ComponentName      string                           `json:"component_name" binding:"required"`
	ComponentClassName string                           `json:"component_class_name" binding:"required"`
	ModelName          string                           `json:"model_name" binding:"required"`
	Origin             string                           `json:"origin" binding:"required"`
	Extent             []string                         `json:"extent" binding:"required"`
	Rotation           string                           `json:"rotation" binding:"required"`
}

type updateConnectionAnnotationData struct {
	PackageId    string   `json:"package_id" binding:""`
	ModelName    string   `json:"model_name" binding:"required"`
	ConnectStart string   `json:"connect_start" binding:"required"`
	ConnectEnd   string   `json:"connect_end" binding:"required"`
	Color        string   `json:"color" binding:"required"`
	LinePoints   []string `json:"line_points" binding:"required"`
}

type updateConnectionNamesData struct {
	PackageId   string `json:"package_id" binding:"required"`
	ModelName   string `json:"model_name" binding:"required"`
	FromName    string `json:"from_name" binding:"required"`
	ToName      string `json:"to_name" binding:"required"`
	FromNameNew string `json:"from_name_new" binding:"required"`
	ToNameNew   string `json:"to_name_new" binding:"required"`
}

type deleteConnectionData struct {
	PackageId    string `json:"package_id" binding:"required"`
	ModelName    string `json:"model_name" binding:"required"`
	ConnectStart string `json:"connect_start" binding:"required"`
	ConnectEnd   string `json:"connect_end" binding:"required"`
}

type setModelDocumentData struct {
	PackageId string `json:"package_id" binding:"required"`
	ModelName string `json:"model_name" binding:"required"`
	Document  string `json:"document" binding:"required"`
	Revisions string `json:"revisions" binding:"required"`
}

type convertUnitsData struct {
	S1 string `json:"s1" binding:"required"`
	S2 string `json:"s2" binding:"required"`
}

type modelCollectionData struct {
	PackageId string `json:"package_id" binding:"required"`
	ModelName string `json:"model_name" binding:"required"`
}

type loadPackageData struct {
	PackageId           string                `json:"package_id" binding:"required"`
	LoadPackageConflict []loadPackageConflict `json:"conflict" binding:""`
}

type loadPackageConflict struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	//LoadOrUnload string `json:"type,omitempty"`
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

type snapshotEditData struct {
	SnapshotId       string `json:"snapshot_id" binding:"required"`
	SnapshotName     string `json:"snapshot_name" binding:"required"`
	ModelName        string `json:"model_name"  binding:"required"`
	ComponentName    string `json:"component_name" binding:""`
	PackageId        string `json:"package_id" binding:"required"`
	ExperimentId     string `json:"experiment_id" binding:""`
	SimulateResultId string `json:"simulate_result_id" binding:""`

	SimulateVarData   datatypes.JSON `json:"simulate_var_data" binding:""`
	SimulateResultObj datatypes.JSON `json:"simulate_result_obj" binding:""`
	ModelVarData      datatypes.JSON `json:"model_var_data" binding:""`
}
