package structs

import "github.com/dgrijalva/jwt-go"

// CustomClaims 自定义声明
type CustomClaims struct {
	UserID   uint
	Username string
	jwt.StandardClaims
}
