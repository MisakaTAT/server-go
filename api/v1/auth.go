package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/middleware"
	"server/models"
	"server/service"
	"server/structs"
	"server/utils"
	"time"
)

// Login 登陆
func Login(c *gin.Context) {
	loginRequest := structs.LoginReq{}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		utils.Result(utils.Succeed, utils.Translate(err), nil, c)
		return
	}
	// 账号密码检查
	user, ok, err := service.LoginCheck(loginRequest.Username, loginRequest.Password)
	if err != nil {
		global.Zap.Errorf("数据库查询失败：%v", err)
		utils.Result(utils.Failed, "内部错误", nil, c)
		return
	}
	if ok {
		createToken(c, user)
		return
	}
	utils.Result(utils.Failed, "用户名或密码无效", nil, c)
}

// createToken 创建 Token
func createToken(c *gin.Context, user models.User) {
	// 构造 SignKey 签名和解签名需要使用同一个值
	j := middleware.NewJWT()
	// 创建 claims
	claims := structs.CustomClaims{
		UUID: user.UUID,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),                         // 生效时间
			ExpiresAt: time.Now().Unix() + global.CONFIG.Jwt.Exp, // 过期时间
			Issuer:    global.CONFIG.Jwt.Iss,                     // 签发人
		},
	}
	// 生成 Token
	token, err := j.GenerateToken(claims)
	if err != nil {
		utils.Result(utils.Failed, err.Error(), nil, c)
		return
	}
	// 封装一个响应数据，返回 Username 与 Token
	data := structs.LoginResp{
		Token: token,
	}
	utils.Result(utils.Succeed, "登录成功", data, c)
}
