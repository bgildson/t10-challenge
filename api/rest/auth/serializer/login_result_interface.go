package serializer

// LoginResultSerializer defines how a login result serializer should be implemented
type LoginResultSerializer func(string) ([]byte, error)
