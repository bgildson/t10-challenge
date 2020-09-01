package serializer

// ErrorSerializer defines how an error serializer should be implemented
type ErrorSerializer func(string) ([]byte, error)
