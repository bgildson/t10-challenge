package entity

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// UserRole defines which roles the user could have
type UserRole string

// Available user role
const (
	ExternalApp UserRole = "externalapp"
	SuperUser   UserRole = "superuser"
)

// User represents an user
type User struct {
	ID             string
	Name           string
	Email          string
	HashedPassword string
	Role           UserRole
	CreatedAt      time.Time
}

// NewUser creates a new user instance
func NewUser(name string, email string, password string, role UserRole) (*User, error) {
	id := uuid.NewV4()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("could not create a password: %v", err)
	}

	return &User{
		ID:             id.String(),
		Name:           name,
		Email:          email,
		HashedPassword: string(hashedPassword),
		Role:           role,
		CreatedAt:      time.Now(),
	}, nil
}

// PasswordMatch validates the param with the current password
func (u User) PasswordMatch(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password)) == nil
}
