package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"server/common/resp"
	"server/common/utils"
	"server/global"
	"server/middleware"
	"server/models"
	"server/service"
	"server/structs"
	"time"
)

// Login 登陆
func Login(c *gin.Context) {
	loginRequest := structs.LoginRequest{}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		resp.Result(resp.Success, utils.Translate(err), nil, c)
		return
	}
	// 账号密码检查
	if user, ok := service.LoginCheck(loginRequest.Username, loginRequest.Password); ok {
		createToken(c, user)
		return
	}
	resp.Result(resp.Unauthorized, "Login failed, please check username or password.", nil, c)
}

// createToken 创建Token
func createToken(c *gin.Context, user models.User) {
	// 构造SignKey,签名和解签名需要使用同一个值
	j := middleware.NewJWT()
	// 创建claims
	claims := structs.CustomClaims{
		UserID:   user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),                         // 生效时间
			ExpiresAt: time.Now().Unix() + global.CONFIG.Jwt.Exp, // 过期时间
			Issuer:    global.CONFIG.Jwt.Iss,                     // 签发人
		},
	}
	// 生成Token
	token, err := j.GenerateToken(claims)
	if err != nil {
		resp.Result(resp.Error, err.Error(), nil, c)
		return
	}
	// 封装一个响应数据,返回用户名与Token
	data := structs.LoginResult{
		Username: user.Username,
		Token:    token,
	}
	resp.Result(resp.Success, "Login success", data, c)
}
