package serializer

// LoginPayloadSerializer defines how a login payload serializer should be implemented
type LoginPayloadSerializer func([]byte) (*LoginPayload, error)
