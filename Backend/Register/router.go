package register

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	db "github.com/dstejas19/HRManagementSoftware-SoftwareEngineering/Backend/Database"
	errorResponses "github.com/dstejas19/HRManagementSoftware-SoftwareEngineering/Backend/ErrorHandler/ErrorResponse"
	models "github.com/dstejas19/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	utils "github.com/dstejas19/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	fmt.Println("Hey")
	router := mux.NewRouter()

	router.Path("/HR").Methods(http.MethodPost).HandlerFunc(registerHR)

	return router
}

func registerHR(w http.ResponseWriter, req *http.Request) {
	u := models.User{}

	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		fmt.Println(err)

		errorResponses.SendBadRequestResponse(w)

		return
	}

	emailSlice := strings.Split(u.PersonalEmail, "@")
	email := emailSlice[0] + utils.Constants.EmailDomain

	queryString := utils.Queries.RegisterHR + "\"" + u.FirstName + "\", " + "\"" + u.LastName + "\", " + "\"" + u.BusinessUnit + "\", " + string(rune(u.ManagerId)) + "\"" + u.Grade + "\", " + "\"" + u.Location + "\", " + "\"" + u.Country + "\", " + "\"" + u.Title + "\", " + "\"" + u.Type + "\", " + "\"" + u.PersonalEmail + "\", " + "\"" + email + "\", " + "CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"

	_, err := db.Db.Exec(queryString)

	if err != nil {
		fmt.Println(err)

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
