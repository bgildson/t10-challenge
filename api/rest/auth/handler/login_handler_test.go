package handler

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"

	authSerializer "github.com/bgildson/t10-challenge/api/rest/auth/serializer"
	authSerializerJson "github.com/bgildson/t10-challenge/api/rest/auth/serializer/json"
	utilSerializer "github.com/bgildson/t10-challenge/api/rest/util/serializer"
	utilSerializerJson "github.com/bgildson/t10-challenge/api/rest/util/serializer/json"
	authServiceMock "github.com/bgildson/t10-challenge/pkg/auth/service/mock"
)

func TestLoginHandlerHandle(t *testing.T) {
	ctrl := gomock.NewController(t)

	invalidEmail := "wrong@email.com"
	validEmail := "correct@email.com"
	validPassword := "123456"
	validToken := "token"
	validPayloadFormat := `{"email": "%s", "password": "%s"}`
	authService := authServiceMock.NewMockAuthService(ctrl)
	authService.
		EXPECT().
		Login(invalidEmail, validPassword).
		Return("", errors.New("user does not exists"))
	authService.
		EXPECT().
		Login(validEmail, validPassword).
		Return(validToken, nil)

	contentType := "application/json"
	handler := NewLoginHandler(
		authService,
		map[string]authSerializer.LoginPayloadSerializer{
			contentType: authSerializerJson.LoginPayloadSerializer,
		},
		map[string]authSerializer.LoginResultSerializer{
			contentType: authSerializerJson.LoginResultSerializer,
		},
		map[string]utilSerializer.ErrorSerializer{
			contentType: utilSerializerJson.ErrorSerializer,
		},
	)

	tt := []struct {
		description string
		// inBody      io.Reader
		// inHeaders   map[string]string
		in  *http.Request
		out int
	}{
		{
			description: "an unsupported content-type",
			in: &http.Request{
				Method: http.MethodPost,
				Header: http.Header{"Content-Type": []string{"application/xml"}},
				Body:   ioutil.NopCloser(bytes.NewReader([]byte(""))),
			},
			out: http.StatusUnsupportedMediaType,
		},
		{
			description: "an invalid payload syntax",
			in: &http.Request{
				Method: http.MethodPost,
				Header: http.Header{"Content-Type": []string{"application/json"}},
				Body:   ioutil.NopCloser(bytes.NewReader([]byte("invalid payload"))),
			},
			out: http.StatusBadRequest,
		},
		{
			description: "an invalid payload",
			in: &http.Request{
				Method: http.MethodPost,
				Header: http.Header{"Content-Type": []string{"application/json"}},
				Body:   ioutil.NopCloser(bytes.NewReader([]byte(`{"email": "user@email.com"}`))),
			},
			out: http.StatusBadRequest,
		},
		{
			description: "an invalid user",
			in: &http.Request{
				Method: http.MethodPost,
				Header: http.Header{"Content-Type": []string{"application/json"}},
				Body: ioutil.NopCloser(
					bytes.NewReader(
						[]byte(fmt.Sprintf(validPayloadFormat, invalidEmail, validPassword)),
					),
				),
			},
			out: http.StatusUnauthorized,
		},
		{
			description: "an invalid user",
			in: &http.Request{
				Method: http.MethodPost,
				Header: http.Header{"Content-Type": []string{"application/json"}},
				Body: ioutil.NopCloser(
					bytes.NewReader(
						[]byte(fmt.Sprintf(validPayloadFormat, validEmail, validPassword)),
					),
				),
			},
			out: http.StatusOK,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			rec := httptest.NewRecorder()
			handler.Handle(rec, tc.in)

			if rec.Code != tc.out {
				t.Errorf("was expecting %d, but returns %d", tc.out, rec.Code)
			}
		})
	}

}
