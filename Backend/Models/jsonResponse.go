package models

import (
	"encoding/json"

	errors "github.com/dstejas19/HRManagementSoftware-SoftwareEngineering/Backend/Models/Errors"
)

type JsonResponse struct {
	Error string `json:"error"`
	Msg   string `json:"msg"`
	Data  string `json:"data"`
}

func (m JsonResponse) ToJSON() []byte {
	jsonResponse, jsonError := json.Marshal(m)

	if jsonError != nil {
		internalServerError := errors.InternalServerError

		message := JsonResponse{}

		message.Error = internalServerError.Err
		message.Data = ""
		message.Msg = ""

		res := message.ToJSON()

		return res
	}

	return jsonResponse
}
