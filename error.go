package r4r

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	StatusCode        int    `json:"statusCode"`
	StatusMessage     string `json:"statusMessage"`
	AdditionalMessage string `json:"additionalMessage"`
}

func NewError(statusCode int, statusMessage string, additionalMessage string) *Error {
	return &Error{
		StatusCode:        statusCode,
		StatusMessage:     statusMessage,
		AdditionalMessage: additionalMessage,
	}
}

func errorHandler(err *Error, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)
	json.NewEncoder(w).Encode(err)
}
