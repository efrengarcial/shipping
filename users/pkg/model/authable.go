package model


import (
	"github.com/dgrijalva/jwt-go"
)


// CustomClaims is our custom metadata, which will be hashed
// and sent as the second segment in our JWT
type CustomClaims struct {
	User *User
	jwt.StandardClaims
}
