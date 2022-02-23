package errors

import "net/http"

var UnauthorizedError = Error{
	Code: http.StatusUnauthorized,
	Err:  "Unauthorized user. Please login.",
}
