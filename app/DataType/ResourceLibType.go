package DataType

type EditResourceInfoData struct {
	ParentId    string `json:"parent_id" binding:""`
	ID          string `json:"id" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:""`
}

type DeleteResourceData struct {
	ID string `json:"id" binding:"required"`
}

type CreateResourceFolderData struct {
	ParentId string `json:"parent_id" binding:""`
	Name     string `json:"name" binding:"required"`
}

type UploadResourceFileData struct {
	ParentId    string `form:"parent_id" binding:""`
	Filename    string `form:"filename" binding:"required"`
	Description string `form:"description" binding:"required"`
}

type CopyLibFileToResourcesData struct {
	ID        string `json:"id" binding:"required"`
	PackageId string `json:"package_id" binding:"required"`
	Parent    string `json:"parent" binding:""`
}
