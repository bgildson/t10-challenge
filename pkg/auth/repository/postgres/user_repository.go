package postgres

import (
	"context"
	"fmt"

	"github.com/bgildson/t10-challenge/pkg/auth/entity"
	"github.com/bgildson/t10-challenge/pkg/auth/repository"
	"github.com/jmoiron/sqlx"
)

// QueryGetByEmail selects the user by email
var QueryGetByEmail = `
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
	rows, err := r.db.Query(QueryGetByEmail, email)
	if err != nil {
		return nil, fmt.Errorf("could not get an user by email: %v", err)
	}
	defer rows.Close()

	user := &entity.User{}
	if rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.HashedPassword,
			&user.Role,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}
