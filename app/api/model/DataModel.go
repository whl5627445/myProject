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

type SetComponentModifierValueModel struct {
	PackageId      string            `json:"package_id" binding:"required"`
	ModelName      string            `json:"model_name" binding:"required"`
	ParameterValue map[string]string `json:"parameter_value" binding:"required"`
}

type SetComponentPropertiesModel struct {
	PackageId        string `json:"package_id" binding:"required"`
	ModelName        string `json:"model_name" binding:"required"`
	OldComponentName string `json:"old_component_name" binding:"required"`
	NewComponentName string `json:"new_component_name" binding:"required"`
	Final            bool   `json:"final" binding:""`
	Protected        bool   `json:"protected" binding:""`
	Replaceable      bool   `json:"replaceable" binding:""`
	Variability      string `json:"variability" binding:"required"`
	Inner            bool   `json:"inner" binding:""`
	Outer            bool   `json:"outer" binding:""`
	Causality        string `json:"causality" binding:"required"`
}

type CopyClassModel struct {
	PackageId       string `json:"package_id" binding:"required"`
	ParentName      string `json:"parent_name" binding:""`
	ClassName       string `json:"class_name" binding:"required"`
	CopiedClassName string `json:"copied_class_name" binding:"required"`
}

type DeleteClassModel struct {
	PackageId  string `json:"package_id" binding:"required"`
	ParentName string `json:"parent_name" binding:""`
	ClassName  string `json:"class_name" binding:"required"`
}

type AddComponentModel struct {
	PackageId        string   `json:"package_id" binding:"required"`
	NewComponentName string   `json:"new_component_name" binding:"required"`
	OldComponentName string   `json:"old_component_name" binding:"required"`
	ClassName        string   `json:"class_name" binding:"required"`
	Origin           string   `json:"origin" binding:"required"`
	Extent           []string `json:"extent" binding:"required"`
	Rotation         string   `json:"rotation" binding:"required"`
}
