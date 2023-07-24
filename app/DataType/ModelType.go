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
}

type CopyClassData struct {
	PackageId       string `json:"package_id" binding:""`
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
	PackageId        string   `json:"package_id" binding:"required"`
	NewComponentName string   `json:"new_component_name" binding:"required"`
	OldComponentName string   `json:"old_component_name" binding:"required"`
	ModelName        string   `json:"model_name" binding:"required"`
	Origin           string   `json:"origin" binding:"required"`
	Extent           []string `json:"extent" binding:"required"`
	Rotation         int      `json:"rotation" binding:""`
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
	Revisions string `json:"revisions" binding:"required"`
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
	//LoadOrUnload string `json:"type,omitempty"`
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
	ModelInformation []CadInformation `json:"model_information" binding:"required"`
	PartNumber       string           `json:"partnumber" binding:"required"`
	Type             string           `json:"type" binding:"required"`
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
