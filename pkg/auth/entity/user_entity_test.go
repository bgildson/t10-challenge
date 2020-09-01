package entity

import (
	"fmt"
	"testing"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func TestNewUser(t *testing.T) {
	tt := []struct {
		Name     string
		Email    string
		Password string
		Role     UserRole
	}{
		{
			Name:     "newuser",
			Email:    "newuser@email.com",
			Password: "123456",
			Role:     SuperUser,
		},
		{
			Name:     "newuser2",
			Email:    "newuser2@email.com",
			Password: "654321",
			Role:     ExternalApp,
		},
	}

	for _, tc := range tt {
		user, err := NewUser(tc.Name, tc.Email, tc.Password, tc.Role)
		if err != nil {
			t.Errorf("could not create a new user: %v", err)
		}

		t.Run("popuales id correctly", func(t *testing.T) {
			id, err := uuid.FromString(user.ID)
			if err != nil {
				t.Errorf("could not load the generated id \"%s\": %v", user.ID, err)
			}

			if id.Version() != uuid.V4 {
				t.Errorf("was expecting version 4, but was used version %d", id.Version())
			}
		})

		if user.Name != tc.Name {
			t.Errorf("the name was not setted properly, was expecting \"%s\", but was setted \"%s\"", tc.Name, user.Name)
		}

		if user.Email != tc.Email {
			t.Errorf("the email was not setted properly, was expecting \"%s\", but was setted \"%s\"", tc.Email, user.Email)
		}

		if bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(tc.Password)) != nil {
			t.Errorf("the password was not setted properly")
		}

		if user.Role != tc.Role {
			t.Errorf("the role was not setted properly, was expecting \"%s\", but was setted \"%s\"", tc.Role, user.Role)
		}

		if user.CreatedAt.IsZero() {
			t.Errorf("the CreatedAt was not setted proporly")
		}
	}
}

func TestUserPasswordMatch(t *testing.T) {
	password := "654321"
	wrongPassword := fmt.Sprintf("wrong%s", password)

	user, err := NewUser("", "", password, SuperUser)
	if err != nil {
		t.Error(err)
	}

	t.Run("correct password", func(t *testing.T) {
		if !user.PasswordMatch(password) {
			t.Errorf("was expecting that the correct password match")
		}
	})

	t.Run("wrong password", func(t *testing.T) {
		if user.PasswordMatch(wrongPassword) {
			t.Errorf("was expecting that the wrong password doesnt match")
		}
	})
}
