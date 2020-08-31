package postgres

import (
	"context"
	"errors"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bgildson/t10-challenge/pkg/auth/entity"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

func TestUserRepositoryGetByEmail(t *testing.T) {
	// prepare mocked database
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")

	repository := NewUserRepository(db)

	user := entity.User{
		ID:             uuid.NewV4().String(),
		Name:           "myuser",
		Email:          "myuser@email.com",
		HashedPassword: "123456",
		Role:           entity.SuperUser,
		CreatedAt:      time.Now(),
	}

	// with success
	rows := sqlmock.NewRows([]string{"id", "name", "email", "hashed_password", "role", "created_at"}).
		AddRow(user.ID, user.Name, user.Email, user.HashedPassword, user.Role, user.CreatedAt)
	expectedSQL := regexp.QuoteMeta(QueryGetByEmail)
	mock.ExpectQuery(expectedSQL).WillReturnRows(rows)

	result, err := repository.GetByEmail(context.Background(), user.Email)
	if err != nil {
		t.Errorf("an unexpected error has occurred: %v", err)
	}

	if !reflect.DeepEqual(*result, user) {
		t.Errorf("was expecting:\n%#v\nbut returns:\n%#v", user, result)
	}

	// with failure
	mock.ExpectQuery(expectedSQL).WillReturnError(errors.New("occur an error"))

	_, err = repository.GetByEmail(context.Background(), user.Email)
	if err == nil {
		t.Errorf("an unexpected error has occurred: %v", err)
	}
}
