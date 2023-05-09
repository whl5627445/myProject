package API

type responseData struct {
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Status int         `json:"status"` // 正常是0，系统级错误是1， 用戶错误是2
	Err    string      `json:"err"`
}

type AppModelMarkData struct {
	PackageId     string `json:"package_id" binding:"required"`
	ModelName     string `json:"model_name" binding:"required"`
	CompilerType  string `json:"compiler_type,omitempty" binding:""`
	MandatorySave bool   `json:"save,omitempty" binding:""`
}

type CreateAppSpaceData struct {
	SpaceName        string `json:"space_name" binding:"required"`
	SpaceDescription string `json:"space_description,omitempty" binding:""`
	Background       string `json:"space_background,omitempty" binding:""`
	Icon             string `json:"space_icon,omitempty" binding:"required"`
	IconColor        string `json:"space_icon_color,omitempty" binding:"required"`
}

type EditAppSpaceData struct {
	SpaceId          string `json:"space_id" binding:"required"`
	SpaceName        string `json:"space_name" binding:"required"`
	SpaceDescription string `json:"space_description,omitempty" binding:""`
	Background       string `json:"space_background,omitempty" binding:""`
	Icon             string `json:"space_icon,omitempty" binding:""`
	IconColor        string `json:"space_icon_color,omitempty" binding:""`
	Collect          string `json:"space_collect,omitempty" binding:""`
}

type DeleteAppSpaceData struct {
	SpaceId string `json:"space_id" binding:"required"`
}

type CreateAppPageData struct {
	//SpaceId  string `json:"space_id" binding:"required"`
	PageName string `json:"name" binding:"required"`
	Tag      string `json:"tag" binding:"required"`
}

type EditAppPageData struct {
	SpaceId  string `json:"space_id" binding:"required"`
	PageId   string `json:"page_id" binding:"required"`
	PageName string `json:"page_name" binding:"required"`
	Tag      string `json:"tag" binding:"required"`
}

type DeleteAppPageData struct {
	SpaceId string `json:"space_id" binding:"required"`
	PageId  string `json:"page_id" binding:"required"`
}
