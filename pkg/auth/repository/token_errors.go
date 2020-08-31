package repository

import "errors"

// TokenRepository context generic error
var (
	ErrTokenGenerationProblem = errors.New("could not generate a token")
)
