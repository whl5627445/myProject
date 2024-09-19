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
	Filename    string `form:"filename" binding:""`
	Description string `form:"description" binding:""`
}

type DownloadPipeNetInfoFileData struct {
	PipeNetInfoFileIdList []string `json:"id_list" binding:"required"`
}

type DeletePipeNetInfoFileData struct {
	PipeNetInfoFileIdList []string `json:"id_list" binding:"required"`
}

type EditMappingConfigDetailsData struct {
	ID     string  `json:"id" binding:"required"`
	Op     string  `json:"op" binding:"required"`
	System string  `json:"system" binding:""`
	Medium string  `json:"medium" binding:""`
	Parts  []*Part `json:"parts" binding:""`
}

type Part struct {
	Kind          string  `json:"kind" binding:"required"`
	Name          string  `json:"name" binding:"required"`
	ModelicaClass string  `json:"modelica_class" binding:"required"`
	ParameterList []*Pair `json:"parameter_list" binding:"required"`
	PortList      []*Pair `json:"port_list" binding:"required"`
}

type Pair struct {
	SourceName string `json:"source_name" binding:"required"`
	TargetName string `json:"target_name" binding:"required"`
}
