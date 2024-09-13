package DataType

type UploadMappingConfigData struct {
	Filename    string `form:"filename" binding:"required"`
	Description string `form:"description" binding:"required"`
}
