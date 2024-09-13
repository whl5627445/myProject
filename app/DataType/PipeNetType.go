package DataType

type UploadMappingConfigData struct {
	Filename    string `form:"filename" binding:"required"`
	Description string `form:"description" binding:"required"`
}

type DownloadMappingConfigData struct {
	MappingConfigIdList []string `json:"mapping_config_id_list" binding:"required"`
}

type DeleteMappingConfigData struct {
	MappingConfigIdList []string `json:"mapping_config_id_list" binding:"required"`
}
