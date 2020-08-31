package service

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/bgildson/t10-challenge/pkg/auth/entity"
	authRepositoryMock "github.com/bgildson/t10-challenge/pkg/auth/repository/mock"
)

func TestAuthServiceLogin(t *testing.T) {
	token := "newtoken"
	userPassword := "123456"
	user, err := entity.NewUser("myuser", "myuser@email.com", userPassword, entity.SuperUser)
	if err != nil {
		t.Errorf("an unexpected error has occurred: %v", err)
	}

	ctrl := gomock.NewController(t)

	tokenRepository := authRepositoryMock.NewMockTokenRepository(ctrl)
	userRepository := authRepositoryMock.NewMockUserRepository(ctrl)
	service := NewAuthService(tokenRepository, userRepository)

	ctx := context.Background()

	t.Run("when occur an error while getting the user by email", func(t *testing.T) {
		userRepository.
			EXPECT().
			GetByEmail(ctx, user.Email).
			Return(nil, errors.New("occur an error"))

		result, err := service.Login(user.Email, userPassword)
		if result != "" {
			t.Errorf("the result should be empty")
		}
		if err == nil {
			t.Errorf("the error, should not be nil")
		}
	})

	t.Run("when user not found", func(t *testing.T) {
		userRepository.
			EXPECT().
			GetByEmail(ctx, user.Email).
			Return(nil, nil)

		result, err := service.Login(user.Email, userPassword)
		if result != "" {
			t.Errorf("the result should be empty")
		}
		if err == nil {
			t.Errorf("the error should not be nil")
		}
	})

	t.Run("when get the user but password does not match", func(t *testing.T) {
		userRepository.
			EXPECT().
			GetByEmail(ctx, user.Email).
			Return(user, nil)

		result, err := service.Login(user.Email, fmt.Sprint(userPassword, userPassword))
		if result != "" {
			t.Errorf("the result should be empty")
		}
		if err == nil {
			t.Errorf("the error should not be nil")
		}
	})

	t.Run("when get an error while generating the token", func(t *testing.T) {
		userRepository.
			EXPECT().
			GetByEmail(ctx, user.Email).
			Return(user, nil)
		tokenRepository.
			EXPECT().
			Generate(user).
			Return("", errors.New("occur an error"))

		result, err := service.Login(user.Email, userPassword)
		if result != "" {
			t.Errorf("the result should be empty")
		}
		if err == nil {
			t.Errorf("the error should not be nil")
		}
	})

	t.Run("when everything is ok", func(t *testing.T) {
		userRepository.
			EXPECT().
			GetByEmail(ctx, user.Email).
			Return(user, nil)
		tokenRepository.
			EXPECT().
			Generate(user).
			Return(token, nil)

		newToken, err := service.Login(user.Email, userPassword)
		if err != nil {
			t.Errorf("an unexpected error has occurred: %v", err)
		}

		if newToken != token {
			t.Errorf("was expecting\n%v\nbut returns\n%v", token, newToken)
		}
	})
}
