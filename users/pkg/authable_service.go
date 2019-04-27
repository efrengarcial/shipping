package users

import (
	"github.com/efrengarcial/shipping/users/pkg/model"
)


//Authable Service 
type Authable interface {
	Decode(token string) (*model.CustomClaims, error)
	Encode(user *model.User) (string, error)
}
