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
	"os"
)

func RegisterHR(w http.ResponseWriter, req *http.Request) {
	userRole := req.Context().Value("role").(byte)
	if userRole&utils.IsHR != utils.IsHR {
		errorResponses.SendForbiddenRequestResponse(w)

		return
	}

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

	var role byte = utils.IsEmployee
	if u.IsHR {
		role = role | utils.IsHR
	}
	if u.IsManager {
		role = role | utils.IsManager
	}

	if Dao.CreateLoginDAO(email, hashedPassword, role) == 0 {
		errorResponses.SendInternalServerErrorResponse(w)

		return
	}

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

	resData := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    email,
		Password: tempPassword,
	}

	data, jsonError := json.Marshal(resData)

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

func GetProfile(w http.ResponseWriter, req *http.Request) {
	email := req.Context().Value("email").(string)

	profileDetails, err := Dao.GetProfileDetails(email)
	if err == 0 {
		errorResponses.SendInternalServerErrorResponse(w)

		return
	}

	Msg := struct {
		FirstName         string `json:"firstName"`
		LastName          string `json:"lastName"`
		Title             string `json:"title"`
		AboutMe           string `json:"aboutMe"`
		ProductivityScore int    `json:"productivityScore"`
	}{
		FirstName:         profileDetails.FirstName,
		LastName:          profileDetails.LastName,
		Title:             profileDetails.Title,
		AboutMe:           profileDetails.AboutMe,
		ProductivityScore: 70,
	}

	data, jsonError := json.Marshal(Msg)
	if jsonError != nil {
		fmt.Println(err)

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

func UpdateProfile(w http.ResponseWriter, req *http.Request) {
	email := req.Context().Value("email").(string)

	userProfile := models.ProfileDetails{}
	if err := json.NewDecoder(req.Body).Decode(&userProfile); err != nil {
		fmt.Println(err)

		errorResponses.SendBadRequestResponse(w, "")

		return
	}

	err := Dao.UpdateProfileDetails(userProfile, email)
	if err == 0 {
		errorResponses.SendInternalServerErrorResponse(w)

		return
	}

	res := models.JsonResponse{}
	res.Error = ""
	res.Msg = "Details updated successfully"
	res.Data = ""

	jsonResponse, jsonError := json.Marshal(res)
	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)
}

func UpdateEmployeeInfo(w http.ResponseWriter, r *http.Request) {

	var user gormModels.User
	//decode json from the front end and convert to object employee
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println(err)
		errorResponses.SendBadRequestResponse(w, "bad request!")
		return
	}

	//Database operation

	//Get EmployeeID from token
	email := r.Context().Value("email").(string)

	user.EmployeeID = Dao.GetEmployeeIDByEmail(email)

	if Dao.UpdateEmployeeDao(user) == 0 {
		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	res := models.JsonResponse{}

	res.Error = ""
	res.Msg = "Employee record updated"
	res.Data = ""

	jsonResponse, jsonError := json.Marshal(res)

	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)
}
func UpdateEmployeeBankingInfo(w http.ResponseWriter, r *http.Request) {
	var user gormModels.User
	//decode json from the front end and convert to object employee
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println(err)
		errorResponses.SendBadRequestResponse(w, "bad request!")
		return
	}

	//Database operation

	//Get EmployeeID from token
	email := r.Context().Value("email").(string)

	user.EmployeeID = Dao.GetEmployeeIDByEmail(email)

	if Dao.UpdateEmployeeBankingDao(user) == 0 {
		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	res := models.JsonResponse{}

	res.Error = ""
	res.Msg = "Employee banking info updated"
	res.Data = ""

	jsonResponse, jsonError := json.Marshal(res)

	if jsonError != nil {
		fmt.Println(jsonError)

		errorResponses.SendInternalServerErrorResponse(w)
		return
	}

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)
}
