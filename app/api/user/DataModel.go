package API

type ResponseData struct {
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Status int         `json:"status"` // 正常是0，系统级错误是1， 用戶错误是2
	Err    string      `json:"err"`
}

type UserSpaceModel struct {
	SpaceId   string `json:"space_id"`
	SpaceName string `json:"space_name"`
}
