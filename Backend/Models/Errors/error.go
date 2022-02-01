package errors

type Error struct {
	Code int
	Err  string
}

// func (e Error) ToJSON() []byte {
// 	jsonResponse, jsonError := json.Marshal(e)

// 	errorHandler.JsonErrorHandler(jsonError, fmt.Sprintf("Could not create json for %d error code. Exiting.\n", e.Code))

// 	return jsonResponse
// }
