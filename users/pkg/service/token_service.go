package service

import (
	"github.com/efrengarcial/shipping/users/pkg"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/efrengarcial/shipping/users/pkg/model"
)

var (

	// Define a secure key string used
	// as a salt when hashing our tokens.
	// Please make your own way more secure than this,
	// use a randomly generated md5 hash or something.
	key = []byte("mySuperSecretKeyLol")
)

type tokenService struct {

}

// NewTokenService will create new an tokenService object representation of users.Authable interface
func NewTokenService() users.Authable {
	return &tokenService{}
}

// Decode a token string into a token object
func (srv *tokenService) Decode(tokenString string) (*model.CustomClaims, error) {

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	// Validate the token and return the custom claims
	if claims, ok := token.Claims.(*model.CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

// Encode a claim into a JWT
func (srv *tokenService) Encode(user *model.User) (string, error) {

	expireToken := time.Now().Add(time.Hour * 72).Unix()

	// Create the Claims
	claims := model.CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "shipping.user",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString(key)
}
