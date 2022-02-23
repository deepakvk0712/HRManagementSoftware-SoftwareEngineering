package Controller

import (
	"encoding/json"
	"fmt"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	errorResponses "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils/ErrorHandler/ErrorResponse"
	"net/http"
)

func Login(w http.ResponseWriter, req *http.Request) {
	userLogin := req.Context().Value("user").(models.UserLogin)
	role := req.Context().Value("role").(byte)

	token, err := utils.GenerateAccessToken(userLogin.Email, role)
	if err != nil {
		fmt.Println(err)

		errorResponses.SendInternalServerErrorResponse(w)

		return
	}

	res := models.JsonResponse{}

	msg := struct {
		AccessToken string
	}{
		AccessToken: token,
	}

	data, jsonError := json.Marshal(msg)

	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	res.Error = ""
	res.Msg = "Logged in successfully"
	res.Data = string(data)

	jsonResponse, jsonError := json.Marshal(res)

	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)
}
