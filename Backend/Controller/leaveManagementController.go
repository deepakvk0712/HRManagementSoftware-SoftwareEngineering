package Controller

import (
	"encoding/json"
	"fmt"
	Dao "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Database/DAO"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	errorResponses "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils/ErrorHandler/ErrorResponse"
	"net/http"
)

func ApplyLeave(w http.ResponseWriter, req *http.Request) {
	email := req.Context().Value("email").(string)

	l := models.ApplyLeave{}

	if err := json.NewDecoder(req.Body).Decode(&l); err != nil {
		fmt.Println(err)

		errorResponses.SendBadRequestResponse(w, "")

		return
	}

	err := Dao.UpdateLeave(l, email)
	if err != 1 {
		errorResponses.SendInternalServerErrorResponse(w)

		return
	}

	res := models.JsonResponse{}
	res.Error = ""
	res.Msg = "Leave applied"
	res.Data = ""

	jsonResponse, jsonError := json.Marshal(res)
	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)
}
