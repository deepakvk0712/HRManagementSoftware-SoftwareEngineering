package Controller

import (
	"encoding/json"
	"fmt"
	Dao "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Database/DAO"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	errorResponses "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils/ErrorHandler/ErrorResponse"
	"net/http"
	"os"
)

func RegisterHR(w http.ResponseWriter, req *http.Request) {
	u := models.User{}
	var email string

	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		fmt.Println(err)

		errorResponses.SendBadRequestResponse(w, "")

		return
	}

	tempPassword := utils.GenerateRandomString(len(u.PersonalEmail))

	hashedPassword, err := utils.HashPassword(tempPassword)
	if err == 0 {
		errorResponses.SendInternalServerErrorResponse(w)

		return
	}

	if email, err = Dao.CreateUserDAO(u, u.PersonalEmail); err == 0 {
		errorResponses.SendInternalServerErrorResponse(w)

		return
	}

	if Dao.CreateLoginDAO(email, hashedPassword) == 0 {
		errorResponses.SendInternalServerErrorResponse(w)

		return
	}

	fmt.Println("Password ", tempPassword)
	fmt.Println("Hashed Password", hashedPassword)

	welcomeMail := models.MailTemplate{
		From:        "HR Admin",
		To:          u.FirstName + " " + u.LastName,
		FromEmail:   os.Getenv("SENDGRID_SENDER_MAIL"),
		ToEmail:     u.PersonalEmail,
		Subject:     "Welcome to the company",
		TextContent: "Dear " + u.FirstName + " " + u.LastName + ",\n\n Welcome onboard. Your official email id is " + email + ". Your temporary password is " + tempPassword + "\n\nRegards,\nHR Team",
		HTMLContent: "",
	}

	if utils.SendMail(welcomeMail) == 0 {
		Dao.DeleteUserDAO(email)
		Dao.DeleteLoginDAO(email)

		errorResponses.SendInternalServerErrorResponse(w)

		return
	}

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
