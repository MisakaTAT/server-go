package structs

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// CustomClaims 自定义声明
type CustomClaims struct {
	UUID uuid.UUID
	jwt.StandardClaims
}
