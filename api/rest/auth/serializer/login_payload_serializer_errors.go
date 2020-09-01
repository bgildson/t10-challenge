package serializer

import "errors"

// LoginPayloadSerializer generic error
var (
	ErrDecodeLoginPayload = errors.New("could not decode login payload")
)
