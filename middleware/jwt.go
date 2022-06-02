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
	TokenInvalid     = errors.New("令牌无效")
	TokenExpired     = errors.New("令牌过期")
	TokenMalformed   = errors.New("令牌格式非法")
	TokenNotValidYet = errors.New("令牌未生效")
)

// JWT 定义 JWT 对象
type JWT struct {
	// 声明签名信息
	SigningKey []byte
}

// NewJWT 初始化 JWT 对象
func NewJWT() *JWT {
	return &JWT{
		[]byte(global.CONFIG.Jwt.SignKey),
	}
}

// GenerateToken 生成 Token
func (j *JWT) GenerateToken(claims structs.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(global.CONFIG.Jwt.SignKey))
}

// ParseToken 解析 Token
func (j *JWT) ParseToken(tokenString string) (*structs.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &structs.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(global.CONFIG.Jwt.SignKey), nil
	})
	if err != nil {
		// https://gowalker.org/github.com/dgrijalva/jwt-go#ValidationError
		// jwt.ValidationError 是一个无效 Token 的错误信息
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
		// 将 Token 中的 claims 信息解析出来并断言成用户自定义的有效载荷结构
		if claims, ok := token.Claims.(*structs.CustomClaims); ok && token.Valid {
			return claims, nil
		} else {
			return nil, TokenInvalid
		}
	} else {
		return nil, TokenInvalid
	}
}

// JWTAuth 定义 JWTAuth 中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头中的 Token
		token := c.Request.Header.Get("X-Token")
		if token == "" {
			resp.Result(resp.Failed, "未登录或非法请求", nil, c)
			c.Abort()
			return
		}
		// 初始化一个 JWT 对象实例，并根据结构体方法来解析 Token
		j := NewJWT()
		// 解析 Token 中包含的相关信息 (有效载荷)
		claims, err := j.ParseToken(token)
		if err != nil {
			resp.Result(resp.Unauthorized, err.Error(), nil, c)
			c.Abort()
			return
		}
		// 将解析后的有效载荷 claims 重新写入 gin.Context 引用对象中
		c.Set("claims", claims)
		c.Next()
	}
}

// GetUserName 从 Token 中解析 Username
func GetUserName(c *gin.Context) string {
	claims, _ := c.Get("claims")
	return claims.(*structs.CustomClaims).Username
}

// GetUserID 从 Token 中解析 UserID
func GetUserID(c *gin.Context) uint {
	claims, _ := c.Get("claims")
	return claims.(*structs.CustomClaims).UserID
}
