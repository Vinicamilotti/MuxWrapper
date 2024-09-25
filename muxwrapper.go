package server

import (
	"net/http"
)

type ApiError struct {
	StatusCode int
	Message    string
}

type ApiHandler func(w http.ResponseWriter, r *http.Request) *ApiError

func CreateApiHandler(handler ApiHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			w.WriteHeader(err.StatusCode)
			w.Write([]byte(err.Message))
		}
	}
}
