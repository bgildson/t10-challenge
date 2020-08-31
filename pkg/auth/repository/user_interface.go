package repository

import (
	"context"

	"github.com/bgildson/t10-challenge/pkg/auth/entity"
)

// UserRepository defines how a user repository should be implemented
type UserRepository interface {
	GetByEmail(context.Context, string) (*entity.User, error)
}
