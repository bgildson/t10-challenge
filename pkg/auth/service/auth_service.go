package service

import (
	"context"
	"errors"

	"github.com/bgildson/t10-challenge/pkg/auth/repository"
)

// AuthService context generic error
var (
	ErrDatabaseConnectionProblem = errors.New("occur a problem with database connection")
	ErrInvalidEmailOrPassword    = errors.New("invalid email or password")
)

// AuthService defines how an auth service should be implemented
type AuthService interface {
	Login(email, password string) (string, error)
}

type authService struct {
	tokenRepository repository.TokenRepository
	userRepository  repository.UserRepository
}

// NewAuthService creates an auth service instance
func NewAuthService(
	tokenRepository repository.TokenRepository,
	userRepository repository.UserRepository,
) AuthService {
	return &authService{tokenRepository, userRepository}
}

func (s authService) Login(email, password string) (string, error) {
	ctx := context.Background()

	user, err := s.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", ErrInvalidEmailOrPassword
	}

	valid := user.PasswordMatch(password)
	if !valid {
		return "", ErrInvalidEmailOrPassword
	}

	token, err := s.tokenRepository.Generate(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
