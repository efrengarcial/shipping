package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/efrengarcial/shipping/users/pkg"
	"github.com/efrengarcial/shipping/users/pkg/model"
	"golang.org/x/crypto/bcrypt"
	"log"
)

const topic = "user.created"

type userService struct {
	repo         users.Repository
	tokenService users.Authable
}

// NewUserService will create new an userService object representation of users.AuthService interface
func NewUserService(r users.Repository,t users.Authable) users.AuthService {
	return &userService{
		repo: r,
		tokenService: t,
	}
}

func (srv *userService) Get(ctx context.Context, id int64) (*model.User, error) {
	user, err := srv.repo.Get(id)
	if err != nil {
		return nil, err
	}

	return  user, nil
}

func (srv *userService) GetAll(ctx context.Context) ([]*model.User, error) {
	users, err := srv.repo.GetAll()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (srv *userService) Auth(ctx context.Context, req *model.User) (*model.Token, error) {
	log.Println("Logging in with:", req.Email, req.Password)
	user, err := srv.repo.GetByEmail(req.Email)
	log.Println(user, err)
	if err != nil {
		return nil, err
	}

	// Compares our given password against the hashed password
	// stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, err
	}

	tokenString, err := srv.tokenService.Encode(user)
	if err != nil {
		return nil, err
	}
	token := &model.Token {
		Token: tokenString,
	}
	return token, nil
}

func (srv *userService) Create(ctx context.Context, req *model.User) (*model.User ,error) {

	log.Println("Creating user: ", req)

	// Generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error hashing password: %v", err))
	}

	req.Password = string(hashedPass)
	if err := srv.repo.Create(req); err != nil {
		return nil, errors.New(fmt.Sprintf("error creating user: %v", err))
	}

	token, err := srv.tokenService.Encode(req)
	if err != nil {
		return nil, err
	}

	req.Token = token
	/*
	if err := srv.Publisher.Publish(ctx, req); err != nil {
		return errors.New(fmt.Sprintf("error publishing event: %v", err))
	} */

	return req, nil
}

func (srv *userService) ValidateToken(ctx context.Context, req *model.Token)  (*model.Token, error) {

	// Decode token
	claims, err := srv.tokenService.Decode(req.Token)

	if err != nil {
		return nil, err
	}

	if claims.User.ID == 0 {
		return nil, errors.New("invalid user")
	}

	token := &model.Token{Valid: true}

	return token, nil
}
