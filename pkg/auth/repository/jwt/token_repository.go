package jwt

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/bgildson/t10-challenge/pkg/auth/entity"
	"github.com/bgildson/t10-challenge/pkg/auth/repository"
)

// TokenClaims represents the custom claims with role
type TokenClaims struct {
	jwt.StandardClaims
	Role entity.UserRole `json:"role"`
}

type tokenRepository struct {
	secretKey   string
	expiresWith int64
}

// NewTokenRepository creates a jwt token repository instance
func NewTokenRepository(secretKey string, expiresWith int64) repository.TokenRepository {
	return &tokenRepository{secretKey, expiresWith}
}

func (r tokenRepository) Generate(user *entity.User) (string, error) {
	issuedAt := time.Now()
	expiresAt := issuedAt.Add(time.Second * time.Duration(r.expiresWith))
	claims := TokenClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   user.ID,
			IssuedAt:  issuedAt.Unix(),
			ExpiresAt: expiresAt.Unix(),
		},
		Role: user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(r.secretKey))
	if err != nil {
		return "", repository.ErrTokenGenerationProblem
	}

	return tokenStr, nil
}

func (r tokenRepository) Validate(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(r.secretKey), nil
	})
	if err != nil {
		return false, err
	}

	_, ok := token.Claims.(jwt.MapClaims)

	return ok && token.Valid, nil
}
