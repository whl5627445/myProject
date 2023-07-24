package DataType

type CreateUserSpaceModel struct {
	SpaceName   string `json:"space_name" binding:"required"`
	Description string `json:"space_description" binding:""`
	Background  string `json:"space_background" binding:""`
	Icon        string `json:"space_icon" binding:""`
	IconColor   string `json:"space_icon_color" binding:""`
}

type EditUserSpaceModel struct {
	SpaceId     string `json:"space_id"`
	SpaceName   string `json:"space_name" binding:"required"`
	Description string `json:"space_description" binding:""`
	Background  string `json:"space_background" binding:""`
	Icon        string `json:"space_icon" binding:""`
	IconColor   string `json:"space_icon_color" binding:""`
	Collect     string `json:"collect"`
}

type DeleteUserSpaceModel struct {
	SpaceId []string `json:"space_id" binding:"required"`
}

type CollectUserSpaceData struct {
	SpaceId []string `json:"space_id" binding:"required"`
	Collect bool     `json:"collect" binding:""`
}

type UserSettingsModel struct {
	//UserName    string `json:"username"  binding:""`
	GridDisplay bool `json:"grid_display" binding:""`
}
