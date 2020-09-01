package json

import (
	"encoding/json"
)

// ErrorSerializer is a json login payload serializer implementation
func ErrorSerializer(message string) ([]byte, error) {
	return json.Marshal(map[string]string{"error": message})
}
