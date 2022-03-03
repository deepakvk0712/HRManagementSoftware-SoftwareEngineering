package Controller

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	Dao "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Database/DAO"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	errorResponses "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils/ErrorHandler/ErrorResponse"
	"net/http"
)

func ChangePassword(w http.ResponseWriter, req *http.Request) {
	loginInfo := models.ChangeLoginInfo{}
	role := req.Context().Value("role").(byte)

	fmt.Println(role)

	if err := json.NewDecoder(req.Body).Decode(&loginInfo); err != nil {
		fmt.Println(err)

		errorResponses.SendBadRequestResponse(w, "")

		return
	}

	dbUser, err := Dao.GetUserDetailsDAO(loginInfo.Email)
	if err == 0 {
		errorResponses.SendInternalServerErrorResponse(w)

		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginInfo.OldPassword)); err != nil {
		fmt.Println("Passwords do not match")

		errorResponses.SendBadRequestResponse(w, "Passwords do not match")

		return
	}

	newHashedPassword, err := utils.HashPassword(loginInfo.NewPassword)
	if err == 0 {
		errorResponses.SendInternalServerErrorResponse(w)

		return
	}

	if Dao.UpdatePasswordDAO(loginInfo.Email, newHashedPassword) == 0 {
		errorResponses.SendInternalServerErrorResponse(w)

		return
	}

	res := models.JsonResponse{}
	res.Error = ""
	res.Msg = "Password updated successfully"
	res.Data = ""

	jsonResponse, jsonError := json.Marshal(res)
	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)
}
