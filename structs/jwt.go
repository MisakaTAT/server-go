package structs

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	UserID   uint
	Username string
	jwt.StandardClaims
}
