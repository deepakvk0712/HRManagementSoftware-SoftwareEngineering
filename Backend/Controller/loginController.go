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
	firstLogin := req.Context().Value("firstLogin").(bool)

	//if firstLogin {
	//	if err := Dao.UpdateFirstLoginDAO(userLogin.Email); err == 0 {
	//		errorResponses.SendInternalServerErrorResponse(w)
	//
	//		return
	//	}
	//}

	token, err := utils.GenerateAccessToken(userLogin.Email, role)
	if err != nil {
		fmt.Println(err)

		errorResponses.SendInternalServerErrorResponse(w)

		return
	}

	res := models.JsonResponse{}

	Msg := struct {
		AccessToken string `json:"accessToken"`
		FirstLogin  bool   `json:"firstLogin"`
	}{
		AccessToken: token,
		FirstLogin:  firstLogin,
	}

	data, jsonError := json.Marshal(Msg)

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

	utils.MessageHandler(w, jsonResponse, http.StatusOK)
}

func Logout(w http.ResponseWriter, req *http.Request) {
	res := models.JsonResponse{}
	res.Error = ""
	res.Msg = "Logged out successfully"
	res.Data = ""

	jsonResponse, jsonError := json.Marshal(res)
	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusOK)
}
