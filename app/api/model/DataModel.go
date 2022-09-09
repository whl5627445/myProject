package API

type ResponseData struct {
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Status int         `json:"status"` // 正常是0，系统级错误是1， 用戶错误是2
	Err    string      `json:"err"`
}

type ModelGraphicsData struct {
	PackageId     string `json:"package_id" binding:"required"`
	ModelName     string `json:"model_name" binding:"required"`
	ComponentName string `json:"component_name,omitempty" binding:""`
}

type SetComponentModifierValueData struct {
	PackageId      string            `json:"package_id" binding:"required"`
	ModelName      string            `json:"model_name" binding:"required"`
	ParameterValue map[string]string `json:"parameter_value" binding:"required"`
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
	PackageId          string   `json:"package_id" binding:"required"`
	ComponentName      string   `json:"component_name" binding:"required"`
	ComponentClassName string   `json:"component_class_name" binding:"required"`
	ModelName          string   `json:"model_name" binding:"required"`
	Origin             string   `json:"origin" binding:"required"`
	Extent             []string `json:"extent" binding:"required"`
	Rotation           string   `json:"rotation" binding:"required"`
}

type UpdateConnectionAnnotationData struct {
	PackageId    string   `json:"package_id" binding:"required"`
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
