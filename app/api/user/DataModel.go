package API

type responseData struct {
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Status int         `json:"status"` // 正常是0，系统级错误是1， 用戶错误是2
	Err    string      `json:"err"`
}

type userSpaceModel struct {
	SpaceId   string `json:"space_id"`
	SpaceName string `json:"space_name"`
}

type userSettingsModel struct {
	UserName    string `json:"username"  binding:"required"`
	GridDisplay string `json:"grid_display" binding:""`
}
