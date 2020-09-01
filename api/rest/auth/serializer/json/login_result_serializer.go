package json

import (
	"encoding/json"
)

// LoginResultSerializer is a json login result serializer implementation
func LoginResultSerializer(token string) ([]byte, error) {
	return json.Marshal(map[string]string{"token": token})
}
