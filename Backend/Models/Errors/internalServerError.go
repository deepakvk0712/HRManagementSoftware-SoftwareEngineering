package errors

var InternalServerError = Error{
	Code: 500,
	Msg:  "",
	Err:  "Internal Server Error. Please try again.",
	Data: "",
}

var InternalServerErrorJSON []byte
