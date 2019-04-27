package users

import (
	"github.com/efrengarcial/shipping/users/pkg/model"
)


// Repository interface
type Repository interface {
	GetAll() ([]*model.User, error)
	Get(id int64) (*model.User, error)
	Create(user *model.User) error
	GetByEmail(email string) (*model.User, error)
}
