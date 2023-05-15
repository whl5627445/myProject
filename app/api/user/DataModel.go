package API

type responseData struct {
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Status int         `json:"status"` // 正常是0，系统级错误是1， 用戶错误是2
	Err    string      `json:"err"`
}

type CreateUserSpaceModel struct {
	SpaceName   string `json:"space_name" binding:"required"`
	Description string `json:"space_description" binding:""`
	Background  string `json:"space_background" binding:""`
	Icon        string `json:"space_icon" binding:"required"`
	IconColor   string `json:"space_icon_color" binding:"required"`
}

type EditUserSpaceModel struct {
	SpaceId     string `json:"space_id"`
	SpaceName   string `json:"space_name" binding:"required"`
	Description string `json:"space_description" binding:""`
	Background  string `json:"space_background" binding:""`
	Icon        string `json:"space_icon" binding:"required"`
	IconColor   string `json:"space_icon_color" binding:"required"`
	Collect     string `json:"collect"`
}

type DeleteUserSpaceModel struct {
	SpaceId []string `json:"space_id" binding:"required"`
}

type LoginUserSpaceModel struct {
	SpaceId string `json:"space_id" binding:"required"`
}

type CollectUserSpaceData struct {
	SpaceId []string `json:"space_id" binding:"required"`
	Collect bool     `json:"collect" binding:""`
}

type userSettingsModel struct {
	//UserName    string `json:"username"  binding:""`
	GridDisplay bool `json:"grid_display" binding:""`
}
