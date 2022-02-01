package errors

import "net/http"

var BadServerError = Error{
	Code: http.StatusBadRequest,
	Err:  "Bad Request.",
}

var BadServerErrorJSON []byte
