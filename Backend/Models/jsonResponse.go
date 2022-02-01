package models

import (
	"encoding/json"
	"fmt"

	errorHandler "github.com/dstejas19/HRManagementSoftware-SoftwareEngineering/Backend/ErrorHandler"
)

type JsonResponse struct {
	Error string `json:"error"`
	Msg   string `json:"msg"`
	Data  string `json:"data"`
}

func (m JsonResponse) ToJSON() []byte {
	jsonResponse, jsonError := json.Marshal(m)

	errorHandler.JsonErrorHandler(jsonError, fmt.Sprintf("Could not create json for %v. Exiting.\n", m))

	return jsonResponse
}
