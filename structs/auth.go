package structs

// LoginReq 登录请求
type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResp 登录成功响应
type LoginResp struct {
	Token string `json:"token"`
}
