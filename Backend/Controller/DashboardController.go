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

func GetDashboard(w http.ResponseWriter, req *http.Request) {
	role := req.Context().Value("role").(byte)
	email := req.Context().Value("email").(string)

	err, isOnboard, isFinance := Dao.GetFormStatus(email)
	if err == 0 {
		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	teamDetails, err := Dao.GetTeamDetails(email)
	if err == 0 {
		errorResponses.SendInternalServerErrorResponse(w)

		return
	}

	businessUnits, err := Dao.GetBusinessUnits()
	if err != 1 {
		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	Msg := struct {
		AccountType   byte                `json:"accountType"`
		Username      string              `json:"username"`
		IsManager     bool                `json:"isManager"`
		IsHR          bool                `json:"isHR"`
		IsOnboard     bool                `json:"isOnboard"`
		IsFinance     bool                `json:"isFinance"`
		TeamMembers   []models.TeamMember `json:"teamMembers"`
		BusinessUnits []string            `json:"businessUnits"`
	}{
		AccountType:   role,
		Username:      utils.GetUsername(email),
		IsManager:     role&utils.IsManager == utils.IsManager,
		IsHR:          role&utils.IsHR == utils.IsHR,
		IsOnboard:     isOnboard,
		IsFinance:     isFinance,
		TeamMembers:   teamDetails.TeamMembers,
		BusinessUnits: businessUnits,
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

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)
}
