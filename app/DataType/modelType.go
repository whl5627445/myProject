package DataType

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
