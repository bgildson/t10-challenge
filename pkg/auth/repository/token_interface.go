package repository

import "github.com/bgildson/t10-challenge/pkg/auth/entity"

// TokenRepository defines how a token repository should be implemented
type TokenRepository interface {
	Generate(*entity.User) (string, error)
	Validate(string) (bool, error)
}
