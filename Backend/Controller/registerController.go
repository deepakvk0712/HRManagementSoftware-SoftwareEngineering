package Controller

import (
	"encoding/json"
	"fmt"
	errorHandler "github.com/AlexWeiZH/HRManagementSoftware-SoftwareEngineering/Backend/ErrorHandler"
	errorResponses "github.com/AlexWeiZH/HRManagementSoftware-SoftwareEngineering/Backend/ErrorHandler/ErrorResponse"
	models "github.com/AlexWeiZH/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	utils "github.com/AlexWeiZH/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	"net/http"
)

func Home(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Home"))
}
func RegisterHR(w http.ResponseWriter, req *http.Request) {
	u := models.User{}

	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		fmt.Println(err)

		errorResponses.SendBadRequestResponse(w)

		return
	}

	fmt.Println(u)

	res := models.JsonResponse{}

	res.Error = ""
	res.Msg = "Employee record created"
	res.Data = ""

	jsonResponse, jsonError := json.Marshal(res)

	errorHandler.JsonErrorHandler(jsonError, fmt.Sprintf("Could not create json for %v. Exiting.\n", res))

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)

}
