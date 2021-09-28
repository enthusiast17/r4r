package r4r

import (
	"net/http"
)

type Handler = func(w http.ResponseWriter, r *http.Request)

type Middleware = func(w http.ResponseWriter, r *http.Request) *Error
