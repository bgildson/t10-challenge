package jwt

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/bgildson/t10-challenge/pkg/auth/entity"
	jwt "github.com/dgrijalva/jwt-go"
)

func TestNewTokenRepository(t *testing.T) {
	tt := []struct {
		secretKey   string
		expiresWith int64
	}{
		{
			secretKey:   "mysecretkey",
			expiresWith: 2520,
		},
	}

	for _, tc := range tt {
		repository := NewTokenRepository(tc.secretKey, tc.expiresWith)

		r := reflect.ValueOf(repository)

		secretKey := r.Elem().FieldByName("secretKey").String()
		if secretKey != tc.secretKey {
			t.Errorf("was expecting %s, but returns %s", tc.secretKey, secretKey)
		}

		expiresWith := r.Elem().FieldByName("expiresWith").Int()
		if secretKey != tc.secretKey {
			t.Errorf("was expecting %d, but returns %d", tc.expiresWith, expiresWith)
		}
	}
}

func TestTokenRepositoryGenerate(t *testing.T) {
	secretKey := "mysecretkey"
	expiresWith := int64(2520)

	repository := NewTokenRepository(secretKey, expiresWith)

	user, err := entity.NewUser("myuser", "myuser@email.com", "pass", entity.SuperUser)
	if err != nil {
		t.Errorf("an unexpected error has occurred: %v", err)
	}

	token, err := repository.Generate(user)
	if err != nil {
		t.Errorf("an unexpected error has occurred: %v", err)
	}

	if token == "" {
		t.Error("the token when success should not be empty")
	}
}

func TestTokenRepositoryValidate(t *testing.T) {
	secretKey := "mysecretkey"
	expiresWith := int64(2520)
	user, err := entity.NewUser("myuser", "myuser@email.com", "pass", entity.SuperUser)
	if err != nil {
		t.Errorf("an unexpected error has occurred: %v", err)
	}

	repository := NewTokenRepository(secretKey, expiresWith)

	now := time.Now()

	tt := []struct {
		description  string
		inIssuedAt   int64
		inExpiresAt  int64
		inSignString string
		outValid     bool
	}{
		{
			description:  "valid token",
			inIssuedAt:   now.Unix(),
			inExpiresAt:  now.Add(time.Second * time.Duration(expiresWith)).Unix(),
			inSignString: secretKey,
			outValid:     true,
		},
		{
			description:  "invalid token",
			inIssuedAt:   now.Add(time.Second * time.Duration(-expiresWith)).Unix(),
			inExpiresAt:  now.Add(time.Second * time.Duration(-1)).Unix(),
			inSignString: secretKey,
			outValid:     false,
		},
		{
			description:  "signed with a different secret key",
			inIssuedAt:   now.Unix(),
			inExpiresAt:  now.Add(time.Second * time.Duration(expiresWith)).Unix(),
			inSignString: fmt.Sprint(secretKey, secretKey),
			outValid:     false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			claims := TokenClaims{
				StandardClaims: jwt.StandardClaims{
					Subject:   user.ID,
					IssuedAt:  tc.inIssuedAt,
					ExpiresAt: tc.inExpiresAt,
				},
				Role: entity.ExternalApp,
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenStr, err := token.SignedString([]byte(tc.inSignString))
			if err != nil {
				t.Errorf("could not sign the token: %v", err)
			}

			if valid, _ := repository.Validate(tokenStr); valid != tc.outValid {
				t.Error("could not validate the token")
			}
		})
	}
}
