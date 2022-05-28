package structs

// LoginRequest 登录请求结构体
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResult 登录成功返回结构体
type LoginResult struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
