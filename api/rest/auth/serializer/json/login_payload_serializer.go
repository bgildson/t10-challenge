package json

import (
	"encoding/json"

	"github.com/bgildson/t10-challenge/api/rest/auth/serializer"
)

// LoginPayloadSerializer is a json login payload serializer implementation
func LoginPayloadSerializer(data []byte) (*serializer.LoginPayload, error) {
	result := make(map[string]interface{})
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, serializer.ErrDecodeLoginPayload
	}

	email, _ := result["email"].(string)
	password, _ := result["password"].(string)

	return &serializer.LoginPayload{
		Email:    email,
		Password: password,
	}, nil
}
