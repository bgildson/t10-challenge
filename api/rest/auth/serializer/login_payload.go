package serializer

// LoginPayload represents a valid login payload
type LoginPayload struct {
	Email    string
	Password string
}

// Validate applies the login payload validations
func (p LoginPayload) Validate() error {
	if p.Email == "" {
		return ErrRequiredEmail
	}
	if p.Password == "" {
		return ErrRequiredPassword
	}
	return nil
}
