package service_test

import (
	"context"
	"github.com/efrengarcial/shipping/users/pkg/mocks"
	"github.com/efrengarcial/shipping/users/pkg/model"
	"github.com/efrengarcial/shipping/users/pkg/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGet(t *testing.T) {
	mockRepository :=  new(mocks.Repository)

	mockUser := &model.User {
		ID: 1,
		Name: "I am Efren",
		Email: "efren.gl@gmail.com",

	}

	t.Run("success", func(t *testing.T) {
		mockRepository.On("Get", mock.AnythingOfType("int64")).Return(mockUser, nil).Once()
		tokenService := service.NewTokenService()
		u := service.NewUserService(mockRepository ,tokenService)

		user, err := u.Get(context.TODO(), mockUser.ID)

		assert.NoError(t, err)
		assert.NotNil(t, user)

		mockRepository.AssertExpectations(t)
	})
}