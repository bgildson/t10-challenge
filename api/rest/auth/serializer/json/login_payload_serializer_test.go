package json

import (
	"reflect"
	"testing"

	"github.com/bgildson/t10-challenge/api/rest/auth/serializer"
)

func TestLoginPayloadSerializer(t *testing.T) {
	tt := []struct {
		description     string
		in              []byte
		outLoginPayload *serializer.LoginPayload
		outErr          error
	}{
		{
			description:     "an invalid input",
			in:              []byte("invalid"),
			outLoginPayload: nil,
			outErr:          serializer.ErrDecodeLoginPayload,
		},
		{
			description:     "an empty and valid input",
			in:              []byte("{}"),
			outLoginPayload: &serializer.LoginPayload{},
			outErr:          nil,
		},
		{
			description:     "a valid input with just email",
			in:              []byte(`{"email": "user@email.com"}`),
			outLoginPayload: &serializer.LoginPayload{Email: "user@email.com"},
			outErr:          nil,
		},
		{
			description:     "a valid input with just password",
			in:              []byte(`{"password": "123456"}`),
			outLoginPayload: &serializer.LoginPayload{Password: "123456"},
			outErr:          nil,
		},
		{
			description:     "a valid input",
			in:              []byte(`{"email": "user@email.com", "password": "123456"}`),
			outLoginPayload: &serializer.LoginPayload{Email: "user@email.com", Password: "123456"},
			outErr:          nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			result, err := LoginPayloadSerializer(tc.in)
			if !reflect.DeepEqual(result, tc.outLoginPayload) {
				t.Errorf("was expecting\n%#v\nbut returns\n%#v", tc.outLoginPayload, result)
			}

			if err != tc.outErr {
				t.Errorf("was expecting\n%v\nbut returns\n%v", tc.outErr, err)
			}
		})
	}
}
