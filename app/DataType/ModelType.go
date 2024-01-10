package DataType

type ResponseData struct {
	Data   any    `json:"data"`
	Msg    string `json:"msg"`
	Status int    `json:"status"` // 正常是0，系统级错误是1， 用戶错误是2, omc未启动则是3， 资源已被删除是4
	Err    string `json:"err"`
}

type ModelGraphicsData struct {
	PackageId     string `json:"package_id" binding:"required"`
	ModelName     string `json:"model_name" binding:"required"`
	ComponentName string `json:"component_name,omitempty" binding:""`
}

type SetComponentModifierValueData struct {
	PackageId string          `json:"package_id" binding:"required"`
	ModelName string          `json:"model_name" binding:"required"`
	Parameter []ParameterData `json:"parameter" binding:"required"`
}

type ParameterData struct {
	ExtendName     string `json:"extend_name" binding:"required"`
	IsExtend       bool   `json:"is_extend" binding:"required"`
	ParameterName  string `json:"parameter_name" binding:"required"`
	ParameterValue string `json:"parameter_value" binding:"required"`
}

type SetComponentUintData struct {
	PackageId      string           `json:"package_id" binding:"required"`
	ModelName      string           `json:"model_name" binding:"required"`
	UnitEditorData []UnitEditorData `json:"unit_editor" binding:"required"`
}

type UnitEditorData struct {
	ExtendName    string `json:"extend_name" binding:"required"`
	IsExtend      bool   `json:"is_extend" binding:"required"`
	ParameterName string `json:"parameter_name" binding:"required"`
	ParameterUnit string `json:"parameter_unit" binding:"required"`
}

type AddComponentParametersData struct {
	PackageId     string `json:"package_id" binding:"required"`
	ModelName     string `json:"model_name" binding:"required"`
	ParameterName string `json:"parameter_name" binding:"required"`
	VarType       string `json:"var_type" binding:"required"`
}

type DeleteComponentParametersData struct {
	PackageId     string `json:"package_id" binding:"required"`
	ModelName     string `json:"model_name" binding:"required"`
	ParameterName string `json:"parameter_name" binding:"required"`
}

type SetComponentPropertiesData struct {
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
	Comment          string `json:"comment" binding:""`
	Dimensions       string `json:"dimensions" binding:""`
}

type CopyClassData struct {
	FromPackageId   string `json:"from_package_id" binding:"required"`
	ToPackageId     string `json:"to_package_id" binding:""`
	ParentName      string `json:"parent_name" binding:""`
	ModelName       string `json:"model_name" binding:"required"`
	CopiedClassName string `json:"copied_class_name" binding:"required"`
}

type DeleteClassData struct {
	PackageId  string `json:"package_id" binding:"required"`
	ParentName string `json:"parent_name" binding:""`
	ModelName  string `json:"model_name" binding:"required"`
}

type AddComponentData struct {
	PackageId        string `json:"package_id" binding:"required"`
	NewComponentName string `json:"new_component_name" binding:"required"`
	OldComponentName string `json:"old_component_name" binding:"required"`
	ModelName        string `json:"model_name" binding:"required"`
	Origin           string `json:"origin" binding:"required"`
	// Extent           []string `json:"extent" binding:"required"`
	Rotation int `json:"rotation" binding:""`
}

type DeleteComponentMap struct {
	DeleteType    string `json:"delete_type" binding:"required"`
	ComponentName string `json:"component_name" binding:"required"`
	ModelName     string `json:"model_name" binding:"required"`
	ConnectStart  string `json:"connect_start" binding:"required"`
	ConnectEnd    string `json:"connect_end" binding:"required"`
}

type DeleteComponentData struct {
	PackageId     string               `json:"package_id" binding:"required"`
	ComponentList []DeleteComponentMap `json:"component_list" binding:"required"`
}

type BatchUpdateComponentData struct {
	PackageId            string                `json:"package_id" binding:"required"`
	ModelName            string                `json:"model_name" binding:"required"`
	UpdateComponentDatas []UpdateComponentData `json:"batch_update_data" binding:"required"`
}

type UpdateComponentData struct {
	PackageId          string                           `json:"package_id" binding:"required"`
	ConnectionList     []UpdateConnectionAnnotationData `json:"connection_list" binding:""`
	ComponentName      string                           `json:"component_name" binding:"required"`
	ComponentClassName string                           `json:"component_class_name" binding:"required"`
	ModelName          string                           `json:"model_name" binding:"required"`
	Origin             string                           `json:"origin" binding:"required"`
	Extent             []string                         `json:"extent" binding:"required"`
	Rotation           string                           `json:"rotation" binding:"required"`
}

type UpdateConnectionAnnotationData struct {
	PackageId    string   `json:"package_id" binding:""`
	ModelName    string   `json:"model_name" binding:"required"`
	ConnectStart string   `json:"connect_start" binding:"required"`
	ConnectEnd   string   `json:"connect_end" binding:"required"`
	Color        string   `json:"color" binding:"required"`
	LinePoints   []string `json:"line_points" binding:"required"`
}

type UpdateConnectionNamesData struct {
	PackageId   string `json:"package_id" binding:"required"`
	ModelName   string `json:"model_name" binding:"required"`
	FromName    string `json:"from_name" binding:"required"`
	ToName      string `json:"to_name" binding:"required"`
	FromNameNew string `json:"from_name_new" binding:"required"`
	ToNameNew   string `json:"to_name_new" binding:"required"`
}

type DeleteConnectionData struct {
	PackageId    string `json:"package_id" binding:"required"`
	ModelName    string `json:"model_name" binding:"required"`
	ConnectStart string `json:"connect_start" binding:"required"`
	ConnectEnd   string `json:"connect_end" binding:"required"`
}

type SetModelDocumentData struct {
	PackageId string `json:"package_id" binding:"required"`
	ModelName string `json:"model_name" binding:"required"`
	Document  string `json:"document" binding:"required"`
	Revisions string `json:"revisions" binding:""`
}

type ConvertUnitsData struct {
	S1 string `json:"s1" binding:"required"`
	S2 string `json:"s2" binding:"required"`
}

type ModelCollectionData struct {
	PackageId string `json:"package_id" binding:"required"`
	ModelName string `json:"model_name" binding:"required"`
}

type LoadPackageData struct {
	PackageId           string                `json:"package_id" binding:"required"`
	LoadPackageConflict []LoadPackageConflict `json:"conflict" binding:""`
}

type UnLoadPackageData struct {
	PackageId string `json:"package_id" binding:"required"`
}

type LoadPackageConflict struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	// LoadOrUnload string `json:"type,omitempty"`
}

type LoginUserSpaceModel struct {
	SpaceId string `json:"space_id" binding:"required"`
}

type AppModelMarkData struct {
	PackageId      string `json:"package_id" binding:"required"`
	ModelName      string `json:"model_name" binding:"required"`
	CompileType    string `json:"compiler_type,omitempty" binding:""`
	MandatorySave  bool   `json:"save,omitempty" binding:""`
	GroupName      string `json:"group_name" binding:"required"`
	DataSourceName string `json:"data_source_name" binding:"required"`
	ExperimentId   string `json:"experiment_id" binding:""`
}

type CADMappingModelData struct {
	PackageId    string                       `json:"package_id" binding:"required"`
	ModelName    string                       `json:"model_name" binding:"required"`
	Information  []CadMappingModelInformation `json:"information" binding:"required"`
	ModelMapping []CadModelMapping            `json:"model_mapping" binding:"required"`
}

type CadModelMapping struct {
	Id        int      `json:"id" binding:"required"`
	ModelName []string `json:"model_name" binding:"required"`
}

type CadMappingModelInformation struct {
	ModelInformation  []CadInformation     `json:"model_information" binding:"required"`
	PartNumber        string               `json:"partnumber" binding:"required"`
	Type              string               `json:"type" binding:"required"`
	ConnectedRelation []map[string]any     `json:"connected_relation" binding:""`
	Points            []map[string]float64 `json:"points" binding:""`
	Name              string               `json:"name" binding:""`
}

type CadData struct {
	Msg  string `json:"msg" binding:""`
	Code int    `json:"code" binding:""`
	Data string `json:"data" binding:""`
}

type CadInformation struct {
	GeometryData  cadMappingGeometry `json:"geometry_data" binding:"required"`
	OriginDiagram []float64          `json:"origin" binding:"required"`
	Rotation      float64            `json:"rotation" binding:"required"`
	Xz            float64            `json:"xz" binding:"required"`
	Yz            float64            `json:"yz" binding:"required"`
}
type cadMappingGeometry struct {
	BendRadius float64            `json:"bend_radius" binding:""`
	Coordinate map[string]float64 `json:"coordinate" binding:""`
	Diameter   float64            `json:"diameter" binding:""`
	HeightAb   float64            `json:"height_ab" binding:""`
	Length     float64            `json:"length" binding:""`
	R0         float64            `json:"R_0" binding:""`
	DHyd       float64            `json:"d_hyd" binding:""`
	Delta      float64            `json:"delta" binding:""`
}

type CreateDependencyLibraryData struct {
	ID       string `json:"id" binding:"required"`
	UserName string `json:"user_name" binding:"required"`
	SpaceId  string `json:"space_id" binding:"required"`
}

type DeleteDependencyLibraryData struct {
	Ids []string `json:"id" binding:"required"`
}

type DeleteIsHaveVersionLibraryData struct {
	Id          string `json:"id" binding:"required"`
	SysUser     string `json:"sys_or_user" binding:"required"`
	UserSpaceId string `json:"userspace_id" binding:"required"`
}

type GetVersionLibraryData struct {
	Id             string `json:"id" binding:"required"`
	PackageName    string `json:"package_name" binding:"required"`
	LibraryId      string `json:"library_id" binding:"required"`
	Version        string `json:"version" binding:""`
	VersionControl bool   `json:"version_control" binding:""`
	VersionBranch  string `json:"version_branch" binding:"required"`
}

type DeleteVersionLibraryData struct {
	Id             string `json:"id" binding:"required"`
	LibraryId      string `json:"library_id" binding:""`
	SpaceId        string `json:"space_id" binding:"required"`
	VersionControl bool   `json:"version_control" binding:""`
}

type CreateVersionLibraryData struct {
	Id      string `json:"id" binding:"required"`
	SpaceId string `json:"space_id" binding:"required"`
}

type GetUMLData struct {
	ClassName        string             `json:"class_name" binding:""`
	ModelType        string             `json:"model_type" binding:""`
	Description      string             `json:"description" binding:""`
	Library          []string           `json:"library" binding:""`
	Level            int                `json:"level" binding:""`
	ExtendsModelData []ExtendsModelData `json:"extends_model_data" binding:""`
}

type ExtendsModelData struct {
	ClassName    string   `json:"class_name" binding:""`
	Count        int      `json:"count" binding:""`
	RelationShip []string `json:"relation_ship" binding:""`
}

type RepositoryCloneData struct {
	RepositoryAddress string `json:"repository_address" binding:"required"`
	Name              string `json:"name" binding:"required"`
	Branch            string `json:"branch" binding:""`
}
type InitVersionControlData struct {
	NoVersionPackageId string `json:"no_version_package_id" binding:"required"`
	RepositoryAddress  string `json:"repository_address" binding:"required"`
	// AnotherName        string `json:"another_name" binding:"required"`
	UserName string `json:"user_name" binding:"required"`
	PassWord string `json:"pass_word" binding:"required"`
}

type RepositoryDeleteData struct {
	ID string `json:"id" binding:"required"`
}

type FilePathData struct {
	FilePath []string `json:"file_path" binding:"required"`
}

type GetPackageUMLData struct {
	ClassName    string              `json:"class_name" binding:""`
	Description  string              `json:"description"`
	Library      []string            `json:"library"`
	ModelType    string              `json:"model_type" binding:""`
	ParentNode   []GetPackageUMLData `json:"parent_node"`
	ChildNode    []GetPackageUMLData `json:"child_node"`
	RelationShip []string            `json:"relation_ship" binding:""`
}

type SetRatedConditionData struct {
	ID                 string           `json:"id" binding:"required"`
	PackageId          string           `json:"package_id" binding:"required"`
	ModelName          string           `json:"model_name" binding:"required"`
	RatedConditionList []ratedCondition `json:"rated_condition_list" binding:"required"`
	// RatedConditionList datatypes.JSON `json:"rated_condition_list" binding:"required"`
}
type ratedCondition struct {
	Name         string `json:"name" binding:"required"`
	Value        string `json:"value" binding:"required"`
	DefaultValue string `json:"default_value" binding:"required"`
	Unit         string `json:"unit" binding:"required"`
}

type SetActualData struct {
	ID             string       `json:"id" binding:"required"`
	PackageId      string       `json:"package_id" binding:"required"`
	ModelName      string       `json:"model_name" binding:"required"`
	ActualDataList []actualData `json:"actual_data_list" binding:"required"`
}
type actualData struct {
	Name  string   `json:"name" binding:"required"`
	Value []string `json:"value" binding:"required"`
}

type SetConditionParametersData struct {
	ID                      string                `json:"id" binding:"required"`
	PackageId               string                `json:"package_id" binding:"required"`
	ModelName               string                `json:"model_name" binding:"required"`
	ConditionParametersList []conditionParameters `json:"condition_parameters_list" binding:"required"`
}

type conditionParameters struct {
	Name  string   `json:"name" binding:"required"`
	Value []string `json:"value" binding:"required"`
}

type SetResultParametersData struct {
	ID                   string             `json:"id" binding:"required"`
	PackageId            string             `json:"package_id" binding:"required"`
	ModelName            string             `json:"model_name" binding:"required"`
	ResultParametersList []resultParameters `json:"result_parameters_list" binding:"required"`
}

type resultParameters struct {
	ResultName string `json:"result_name" binding:"required"`
	ActualName string `json:"actual_name" binding:"required"`
}

type FormulaParserData struct {
	ID         string `json:"id" binding:"required"`
	PackageId  string `json:"package_id" binding:"required"`
	ModelName  string `json:"model_name" binding:"required"`
	FormulaStr string `json:"formula" binding:"required"`
}

type AssociatedParametersData struct {
	ID         string       `json:"id" binding:"required"`
	PackageId  string       `json:"package_id" binding:"required"`
	ModelName  string       `json:"model_name" binding:"required"`
	Parameters []associated `json:"parameters" binding:"required"`
}
type associated struct {
	FormulaVariable  string `json:"formula_variable" binding:"required"`
	MeasuredVariable string `json:"measured_variable" binding:"required"`
}

type SimulationOptionsData struct {
	ID        string   `json:"id" binding:"required"`
	PackageId string   `json:"package_id" binding:"required"`
	ModelName string   `json:"model_name" binding:"required"`
	Options   sOptions `json:"options" binding:"required"`
}
type sOptions struct {
	StartTime         string `json:"start_time" binding:"required"`
	StopTime          string `json:"stop_time" binding:"required"`
	Tolerance         string `json:"tolerance" binding:"required"`
	NumberOfIntervals string `json:"number_of_intervals" binding:"required"`
	Interval          string `json:"interval" binding:"required"`
	Method            string `json:"method" binding:"required"`
}

type CreateCalibrationTemplateData struct {
	ID           string `json:"id" binding:"required"`
	TemplateName string `json:"template_name" binding:"required"`
}

type FittingCalculationData struct {
	ID string `json:"id" binding:"required"`
}

type FittingCoefficientSetData struct {
	ID    string `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type DeleteCalibrationTemplateData struct {
	ID string `json:"id" binding:"required"`
}
