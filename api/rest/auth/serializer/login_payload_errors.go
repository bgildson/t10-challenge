package serializer

import "errors"

// LoginPayload validation error
var (
	ErrRequiredEmail    = errors.New(`the email is required`)
	ErrRequiredPassword = errors.New(`the password is required`)
)
