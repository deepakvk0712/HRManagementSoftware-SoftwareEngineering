package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/dstejas19/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	errors "github.com/dstejas19/HRManagementSoftware-SoftwareEngineering/Backend/Models/Errors"
)

func ResponseHandler(w http.ResponseWriter, err string, msg string, data string, code http.ConnState) http.ResponseWriter {
	res := models.JsonResponse{}

	res.Error = err
	res.Msg = msg
	res.Data = data

	jsonResponse, jsonError := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")

	if jsonError != nil {
		fmt.Println(jsonError)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errors.InternalServerErrorJSON)

		return w
	}

	w.WriteHeader(int(code))
	w.Write(jsonResponse)

	return w
}
