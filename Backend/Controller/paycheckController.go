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
)

func GetPaycheck(w http.ResponseWriter, req *http.Request) {
	email := req.Context().Value("email").(string)
	role := req.Context().Value("role").(byte)

	var paychecks []gormModels.Paycheck
	var teamSalaries []models.TeamSalary
	isManager := false
	v := req.URL.Query()

	startDate := v.Get("startDate")
	endDate := v.Get("endDate")

	employeeID, err := Dao.GetEmployeeID(email)
	if err == 0 {
		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	paychecks = Dao.GetPaycheck(employeeID, startDate, endDate)

	if role&utils.IsManager == utils.IsManager {
		isManager = true

		teamSalaries = Dao.GetSalaries(employeeID, startDate, endDate)

	}

	Msg := struct {
		Paychecks    []gormModels.Paycheck `json:"paychecks"`
		IsManager    bool                  `json:"isManager"`
		TeamSalaries []models.TeamSalary   `json:"teamSalaries"`
	}{
		Paychecks:    paychecks,
		IsManager:    isManager,
		TeamSalaries: teamSalaries,
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

//func GetAllSalaries(w http.ResponseWriter, req *http.Request) {
//	email := req.Context().Value("email").(string)
//	role := req.Context().Value("role").(byte)
//
//	if role&utils.IsManager != utils.IsManager {
//		errorResponses.SendForbiddenRequestResponse(w)
//
//		return
//	}
//
//	v := req.URL.Query()
//
//	startDate := v.Get("startDate")
//	endDate := v.Get("endDate")
//
//	managerID, err := Dao.GetEmployeeID(email)
//	if err == 0 {
//		errorResponses.SendInternalServerErrorResponse(w)
//		return
//	}
//
//	teamSalaries := Dao.GetSalaries(managerID, startDate, endDate)
//
//	Msg := struct {
//		TeamSalaries []models.TeamSalary `json:"paychecks"`
//	}{
//		TeamSalaries: teamSalaries,
//	}
//
//	data, jsonError := json.Marshal(Msg)
//	if jsonError != nil {
//		fmt.Println(jsonError)
//
//		errorResponses.SendInternalServerErrorResponse(w)
//		return
//	}
//
//	res := models.JsonResponse{}
//
//	res.Error = ""
//	res.Msg = ""
//	res.Data = string(data)
//
//	jsonResponse, jsonError := json.Marshal(res)
//	if jsonError != nil {
//		fmt.Println(jsonError)
//
//		errorResponses.SendInternalServerErrorResponse(w)
//		return
//	}
//
//	utils.MessageHandler(w, jsonResponse, http.StatusOK)
//}

func UpdateEmployeeSalary(w http.ResponseWriter, req *http.Request) {
	role := req.Context().Value("role").(byte)

	if role&utils.IsManager != utils.IsManager {
		errorResponses.SendForbiddenRequestResponse(w)

		return
	}

	u := models.UpdateSalary{}
	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		fmt.Println(err)

		errorResponses.SendBadRequestResponse(w, "")
		return
	}

	err := Dao.UpdateSalary(u.EmployeeID, u.NewSalary)
	if err == 0 {
		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	res := models.JsonResponse{}

	res.Error = ""
	res.Msg = "Salary updated successfully"
	res.Data = ""

	jsonResponse, jsonError := json.Marshal(res)
	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusOK)
}
