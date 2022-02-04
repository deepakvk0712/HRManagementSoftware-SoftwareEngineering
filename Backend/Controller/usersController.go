package Controller

import (
	"encoding/json"
	"fmt"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	errorResponses "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils/ErrorHandler/ErrorResponse"
	errorResponses2 "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils/ErrorHandler/ErrorResponse"
	"net/http"
	"os"
	"strings"
)

func RegisterHR(w http.ResponseWriter, req *http.Request) {
	u := models.User{}

	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		fmt.Println(err)

		errorResponses2.SendBadRequestResponse(w)

		return
	}

	emailSlice := strings.Split(u.PersonalEmail, "@")
	email := emailSlice[0] + os.Getenv("emailDomain")

	//queryString := utils.Queries.RegisterHR + "\"" + u.FirstName + "\", " + "\"" + u.LastName + "\", " + "\"" + u.BusinessUnit + "\", " + string(rune(u.ManagerId)) + "\"" + u.Grade + "\", " + "\"" + u.Location + "\", " + "\"" + u.Country + "\", " + "\"" + u.Title + "\", " + "\"" + u.Type + "\", " + "\"" + u.PersonalEmail + "\", " + "\"" + email + "\", " + "CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"

	//_, err := db.Db.Exec(queryString)

	//if err != nil {
	//	fmt.Println(err)
	//
	//	errorResponses.SendInternalServerErrorResponse(w)
	//
	//	return
	//}

	res := models.JsonResponse{}

	data, jsonError := json.Marshal(u)

	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	res.Error = ""
	res.Msg = "Employee record created"
	res.Data = string(data)

	jsonResponse, jsonError := json.Marshal(res)

	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")

	var employee models.Employee
	_ = json.NewDecoder(r.Body).Decode(&employee) //decode json from the front end and convert to employee
	fmt.Println(employee)
	/*
		Database operation
	*/
	//json.NewEncoder(w).Encode(employee)
	res := models.JsonResponse{}

	res.Error = ""
	res.Msg = "Employee record created"
	res.Data = ""

	jsonResponse, jsonError := json.Marshal(res)

	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)
}
