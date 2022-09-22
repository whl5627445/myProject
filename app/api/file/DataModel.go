package API

type ResponseData struct {
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Status int         `json:"status"` // 正常是0，系统级错误是1， 用戶错误是2
	Err    string      `json:"err"`
}

type UpdateModelPackageData struct {
	PackageId  string `json:"package_id" binding:"required"`
	UpdateName string `json:"package_name" binding:"required"`
	ModelStr   string `json:"model_str" binding:"required"`
}

type CreateModelPackageDataVars struct {
	Encapsulated bool   `json:"encapsulated" binding:""`
	Expand       string `json:"expand" binding:""`
	InsertTo     string `json:"insert_to" binding:""`
	Partial      bool   `json:"partial" binding:""`
	State        bool   `json:"state" binding:""`
}

type CreateModelPackageData struct {
	PackageId string                     `json:"package_id" binding:""`
	Name      string                     `json:"package_name" binding:"required"`
	Comment   string                     `json:"comment" binding:""`
	StrType   string                     `json:"str_type" binding:"required"`
	Vars      CreateModelPackageDataVars `json:"vars" binding:"required"`
}

type FilterResultFileData struct {
	RecordId string   `json:"record_id" binding:"required"`
	VarList  []string `json:"var_list" binding:"required"`
}

type FmuExportData struct {
	PackageId     string                 `json:"package_id" binding:"required"`
	PackageName   string                 `json:"package_name" binding:"required"`
	ModelName     string                 `json:"model_name" binding:"required"`
	FmuName       string                 `json:"fmu_name" binding:"required"`
	FmuPar        map[string]interface{} `json:"fmuPar" binding:"required"`
	DownloadLocal bool                   `json:"download_local" binding:"required"`
}

type PackageFileData struct {
	PackageId string `json:"package_id" binding:"required"`
}

type ResultFileData struct {
	RecordId string `json:"record_id" binding:"required"`
}

type ModelCodeSaveData struct {
	PackageId string `json:"package_id" binding:"required"`
	//ModelName string `json:"model_name" binding:"required"`
}
