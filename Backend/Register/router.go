package register

import (
	"encoding/json"
	"fmt"
	"net/http"

	errorHandler "github.com/dstejas19/HRManagementSoftware-SoftwareEngineering/Backend/ErrorHandler"
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

	fmt.Println(u)

	res := models.JsonResponse{}

	res.Error = ""
	res.Msg = "Employee record created"
	res.Data = ""

	jsonResponse, jsonError := json.Marshal(res)

	errorHandler.JsonErrorHandler(jsonError, fmt.Sprintf("Could not create json for %v. Exiting.\n", res))

	utils.MessageHandler(w, jsonResponse, http.StatusCreated)
}
