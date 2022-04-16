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

func ReceiveNotification(w http.ResponseWriter, req *http.Request) {
	//role := req.Context().Value("role").(byte)
	email := req.Context().Value("email").(string)

	var m = models.ReceiveNotificationMessage{}

	if err := json.NewDecoder(req.Body).Decode(&m); err != nil {
		fmt.Println(err)

		errorResponses.SendBadRequestResponse(w, "")

		return
	}

	err := Dao.StoreMessage(m, email)
	if err == 0 {
		errorResponses.SendInternalServerErrorResponse(w)

		return
	}

	res := models.JsonResponse{}

	res.Error = ""
	res.Msg = "Notified successfully"
	res.Data = ""

	jsonResponse, jsonError := json.Marshal(res)
	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)
}

func GetTeamMembers(w http.ResponseWriter, req *http.Request) {
	role := req.Context().Value("role").(byte)
	email := req.Context().Value("email").(string)

	if role&utils.IsHR != utils.IsHR {
		errorResponses.SendForbiddenRequestResponse(w)

		return
	}

	v := req.URL.Query()

	businessUnit := v.Get("businessUnit")

	teamMembers, err := Dao.GetTeamDetailsByBU(businessUnit, email)
	if err != 1 {
		errorResponses.SendInternalServerErrorResponse(w)

		return
	}

	Msg := struct {
		TeamMembers []models.TeamMember `json:"teamMembers"`
	}{
		TeamMembers: teamMembers,
	}

	data, jsonError := json.Marshal(Msg)
	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	res := models.JsonResponse{}

	res.Error = ""
	res.Msg = ""
	res.Data = string(data)

	jsonResponse, jsonError := json.Marshal(res)
	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusOK)
}

func MarkRead(w http.ResponseWriter, req *http.Request) {
	res := models.JsonResponse{}

	v := req.URL.Query()

	messageID := v.Get("id")

	err := Dao.MarkAsRead(messageID)
	if err != 1 {
		errorResponses.SendInternalServerErrorResponse(w)

		return
	}

	res.Error = ""
	res.Msg = "Mark as read"
	res.Data = ""

	jsonResponse, jsonError := json.Marshal(res)
	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusOK)
}