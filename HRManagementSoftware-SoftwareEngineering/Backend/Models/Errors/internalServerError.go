package errors

import "net/http"

var InternalServerError = Error{
	Code: http.StatusInternalServerError,
	Err:  "Internal Server Error. Please try again.",
}
