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

type CopyMappingConfigData struct {
	MappingConfigIdList []string `json:"mapping_config_id_list" binding:"required"`
}

type EditMappingConfigData struct {
	ID          string `json:"id" binding:"required"`
	Name        string `json:"name,omitempty" binding:"required"`
	Description string `json:"description,omitempty" binding:""`
}

type UploadPipeNetInfoFileData struct {
	Filename    string `form:"filename" binding:"required"`
	Description string `form:"description" binding:"required"`
}

type DownloadPipeNetInfoFileData struct {
	PipeNetInfoFileIdList []string `json:"id_list" binding:"required"`
}
