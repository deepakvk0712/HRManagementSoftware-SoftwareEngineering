package errors

import (
	"encoding/json"
	"fmt"

	errorHandler "../../ErrorHandler"
)

type Error struct {
	Code int
	Msg  string
	Err  string
	Data string
}

func (e Error) ToJSON() []byte {
	jsonResponse, jsonError := json.Marshal(e)

	errorHandler.JsonErrorHandler(jsonError, fmt.Sprintf("Could not create json for %d error code. Exiting.\n", e.Code))

	return jsonResponse
}
