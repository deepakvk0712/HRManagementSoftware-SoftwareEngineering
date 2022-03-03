package errors

import "net/http"

var ForbiddenRequestError = Error{
	Code: http.StatusForbidden,
	Err:  "Forbidden Request.",
}
