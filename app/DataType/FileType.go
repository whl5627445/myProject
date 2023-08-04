package DataType

type UpdateModelPackageData struct {
	PackageId string `json:"package_id" binding:"required"`
	ModelName string `json:"model_name" binding:"required"`
	ModelStr  string `json:"model_str" binding:"required"`
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
	PackageId     string         `json:"package_id" binding:"required"`
	PackageName   string         `json:"package_name" binding:"required"`
	ModelName     string         `json:"model_name" binding:"required"`
	FmuName       string         `json:"fmu_name" binding:"required"`
	FmuPar        map[string]any `json:"fmuPar" binding:"required"`
	DownloadLocal bool           `json:"download_local" binding:"required"`
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

type PackageResourcesData struct {
	PackageId string `json:"package_id" binding:""`
	Parent    string `json:"parent" binding:""`
	Path      string `json:"path" binding:""`
}

type ResourcesImagesPathData struct {
	PackageId string `json:"package_id" binding:"required"`
	KeyWord   string `json:"key_word" binding:""`
}

type GetResourcesImagesData struct {
	Path string `json:"path" binding:"required"`
}

type SetResourcesImagesIconData struct {
	PackageId string `json:"package_id" binding:"required"`
	ModelName string `json:"model_name" binding:"required"`
	Path      string `json:"path" binding:"required"`
}
