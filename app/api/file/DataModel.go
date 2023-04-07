package API

type responseData struct {
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Status int         `json:"status"` // 正常是0，系统级错误是1， 用戶错误是2
	Err    string      `json:"err"`
}

type updateModelPackageData struct {
	PackageId  string `json:"package_id" binding:"required"`
	UpdateName string `json:"package_name" binding:"required"`
	ModelStr   string `json:"model_str" binding:"required"`
}

type createModelPackageDataVars struct {
	Encapsulated bool   `json:"encapsulated" binding:""`
	Expand       string `json:"expand" binding:""`
	InsertTo     string `json:"insert_to" binding:""`
	Partial      bool   `json:"partial" binding:""`
	State        bool   `json:"state" binding:""`
}

type createModelPackageData struct {
	PackageId string                     `json:"package_id" binding:""`
	Name      string                     `json:"package_name" binding:"required"`
	Comment   string                     `json:"comment" binding:""`
	StrType   string                     `json:"str_type" binding:"required"`
	Vars      createModelPackageDataVars `json:"vars" binding:"required"`
}

type filterResultFileData struct {
	RecordId string   `json:"record_id" binding:"required"`
	VarList  []string `json:"var_list" binding:"required"`
}

type fmuExportData struct {
	PackageId     string                 `json:"package_id" binding:"required"`
	PackageName   string                 `json:"package_name" binding:"required"`
	ModelName     string                 `json:"model_name" binding:"required"`
	FmuName       string                 `json:"fmu_name" binding:"required"`
	FmuPar        map[string]interface{} `json:"fmuPar" binding:"required"`
	DownloadLocal bool                   `json:"download_local" binding:"required"`
}

type packageFileData struct {
	PackageId string `json:"package_id" binding:"required"`
}

type resultFileData struct {
	RecordId string `json:"record_id" binding:"required"`
}

type modelCodeSaveData struct {
	PackageId string `json:"package_id" binding:"required"`
	//ModelName string `json:"model_name" binding:"required"`
}

type packageResourcesData struct {
	PackageId string `json:"package_id" binding:"required"`
	Parent    string `json:"parent" binding:""`
	Path      string `json:"path" binding:""`
}

type resourcesImagesPathData struct {
	PackageId string `json:"package_id" binding:"required"`
	KeyWord   string `json:"key_word" binding:""`
}

type getResourcesImagesData struct {
	Path string `json:"path" binding:"required"`
}

type setResourcesImagesIconData struct {
	PackageId string `json:"package_id" binding:"required"`
	ModelName string `json:"model_name" binding:"required"`
	Path      string `json:"path" binding:"required"`
}
