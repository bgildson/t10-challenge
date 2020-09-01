package util

import "net/http"

// Handler defines how to implements a handler
type Handler interface {
	Handle(http.ResponseWriter, *http.Request)
}
