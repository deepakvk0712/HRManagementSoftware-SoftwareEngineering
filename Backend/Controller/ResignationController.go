package Controller

import (
	"encoding/json"
	"fmt"
	Dao "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Database/DAO"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	gormModels "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/GormModels"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	errorResponses "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils/ErrorHandler/ErrorResponse"
	"net/http"
	"strconv"
	"time"
)

func InsertFeedBack(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value("email").(string)
	employee_id := Dao.GetEmployeeIDByEmail(email)
	//FirstName + LastName
	employeeName := Dao.GetEmployeeNameByEmail(email)

	Feedback := gormModels.Feedback{}
	if err := json.NewDecoder(r.Body).Decode(&Feedback); err != nil {
		fmt.Println(err)
		errorResponses.SendBadRequestResponse(w, "bad request!")
		return
	}

	if employeeName == "0" {
		errorResponses.SendInternalServerErrorResponse(w)
		return
	}
	Feedback.EmployeeID = employee_id
	Feedback.Name = employeeName
	Feedback.Created_TS = time.Now()
	insertFeed := Dao.InsertFeedback(Feedback)

	if insertFeed == 0 {
		errorResponses.SendInternalServerErrorResponse(w)
		return
	}
	res := models.JsonResponseObject{}

	res.Error = ""
	res.Msg = "Feedback added!"
	res.Data = "ID: " + strconv.Itoa(employee_id) + " Name: " + employeeName

	jsonResponse, jsonError := json.Marshal(res)

	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)
}
