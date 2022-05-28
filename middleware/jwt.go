package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"server/common/resp"
	"server/global"
	"server/structs"
)

var (
	TokenInvalid     = errors.New("token invalid")
	TokenExpired     = errors.New("token expired")
	TokenMalformed   = errors.New("token malformed")
	TokenNotValidYet = errors.New("token not valid yet")
)

// JWT 定义JWT对象
type JWT struct {
	// 声明签名信息
	SigningKey []byte
}

// NewJWT 初始化JWT对象
func NewJWT() *JWT {
	return &JWT{
		[]byte(global.CONFIG.Jwt.SignKey),
	}
}

// GenerateToken 生成Token
func (j *JWT) GenerateToken(claims structs.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(global.CONFIG.Jwt.SignKey))
}

// ParseToken 解析Token
func (j *JWT) ParseToken(tokenString string) (*structs.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &structs.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(global.CONFIG.Jwt.SignKey), nil
	})
	if err != nil {
		// https://gowalker.org/github.com/dgrijalva/jwt-go#ValidationError
		// jwt.ValidationError 是一个无效token的错误结构
		if v, ok := err.(*jwt.ValidationError); ok {
			if v.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if v.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if v.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		// 将 token 中的 claims 信息解析出来并断言成用户自定义的有效载荷结构
		if claims, ok := token.Claims.(*structs.CustomClaims); ok && token.Valid {
			return claims, nil
		} else {
			return nil, TokenInvalid
		}
	} else {
		return nil, TokenInvalid
	}
}

// JWTAuth 定义JWTAuth中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头中的 token
		token := c.Request.Header.Get("Token")
		if token == "" {
			resp.Result(resp.Failed, "未登录或非法请求", nil, c)
			c.Abort()
			return
		}
		// 初始化一个 JWT 对象实例，并根据结构体方法来解析 token
		j := NewJWT()
		// 解析 token 中包含的相关信息 (有效载荷)
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				resp.Result(resp.Failed, "认证过期，请重新登录", nil, c)
				c.Abort()
				return
			}
			// 其它错误
			resp.Result(resp.Failed, err.Error(), nil, c)
			c.Abort()
			return
		}
		// 将解析后的有效载荷 claims 重新写入 gin.Context 引用对象中
		c.Set("claims", claims)
		c.Next()
	}
}

// GetUserName 获取 JWT 解析出来的用户名
func GetUserName(c *gin.Context) string {
	claims, _ := c.Get("claims")
	return claims.(*structs.CustomClaims).Username
}

// GetUserID 获取 JWT 解析出来的用户 ID
func GetUserID(c *gin.Context) uint {
	claims, _ := c.Get("claims")
	return claims.(*structs.CustomClaims).UserID
}
