package users

import (
	"github.com/efrengarcial/shipping/users/pkg/model"
)

import (
	"context"
)

// AuthService has the logic authentication
type AuthService interface {
	Create(ctx context.Context, req *model.User) (*model.User, error)
	Get(ctx context.Context,  id int64) (*model.User, error)
	GetAll(ctx context.Context) ([]*model.User, error)
	Auth(ctx context.Context, in *model.User) (*model.Token, error)
	ValidateToken(ctx context.Context, in *model.Token) (*model.Token, error)
}
