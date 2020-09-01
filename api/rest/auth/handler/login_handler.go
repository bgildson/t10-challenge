package handler

import (
	"io/ioutil"
	"net/http"

	authSerializer "github.com/bgildson/t10-challenge/api/rest/auth/serializer"
	"github.com/bgildson/t10-challenge/api/rest/util"
	utilSerializer "github.com/bgildson/t10-challenge/api/rest/util/serializer"
	"github.com/bgildson/t10-challenge/pkg/auth/service"
)

type loginHandler struct {
	authService        service.AuthService
	payloadSerializers map[string]authSerializer.LoginPayloadSerializer
	resultSerializers  map[string]authSerializer.LoginResultSerializer
	errorSerializers   map[string]utilSerializer.ErrorSerializer
}

// NewLoginHandler creates a login handler instance
func NewLoginHandler(
	authService service.AuthService,
	payloadSerializers map[string]authSerializer.LoginPayloadSerializer,
	resultSerializers map[string]authSerializer.LoginResultSerializer,
	errorSerializers map[string]utilSerializer.ErrorSerializer,
) util.Handler {
	return &loginHandler{
		authService,
		payloadSerializers,
		resultSerializers,
		errorSerializers,
	}
}

func (h loginHandler) Handle(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")

	errorSerializer, exists := h.errorSerializers[contentType]
	if !exists {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte("unsupported content type"))
		return
	}

	w.Header().Add("Content-Type", contentType)

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data, _ := errorSerializer("could not read the request payload")
		w.Write(data)
		return
	}
	defer r.Body.Close()

	payloadSerializer, exists := h.payloadSerializers[contentType]
	if !exists {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		data, _ := errorSerializer("unsuported payload content type")
		w.Write(data)
		return
	}

	payload, err := payloadSerializer(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data, _ := errorSerializer(err.Error())
		w.Write(data)
		return
	}

	if err = payload.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data, _ := errorSerializer(err.Error())
		w.Write(data)
		return
	}

	token, err := h.authService.Login(payload.Email, payload.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		data, _ := errorSerializer(err.Error())
		w.Write(data)
		return
	}

	resultSerializer, exists := h.resultSerializers[contentType]
	if !exists {
		w.WriteHeader(http.StatusInternalServerError)
		data, _ := errorSerializer("result serializer not found")
		w.Write(data)
		return
	}

	result, err := resultSerializer(token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		data, _ := errorSerializer("could not serialize the result")
		w.Write(data)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
