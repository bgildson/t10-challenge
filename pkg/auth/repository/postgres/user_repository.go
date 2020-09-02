package postgres

import (
	"context"

	"github.com/bgildson/t10-challenge/pkg/auth/entity"
	"github.com/bgildson/t10-challenge/pkg/auth/repository"
	"github.com/jmoiron/sqlx"
)

// QueryGetByEmail selects the user by email
const QueryGetByEmail = `
SELECT
	id,
	name,
	email,
	hashed_password,
	role,
	created_at
FROM users
WHERE email = $1;
`

type userRepository struct {
	db *sqlx.DB
}

// NewUserRepository creates a postgres user repository instance
func NewUserRepository(db *sqlx.DB) repository.UserRepository {
	return &userRepository{db}
}

func (r userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	row := r.db.QueryRowContext(ctx, QueryGetByEmail, email)

	var user entity.User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.HashedPassword,
		&user.Role,
		&user.CreatedAt,
	)

	return &user, err
}
