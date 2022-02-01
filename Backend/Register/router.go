package register

import (
	"encoding/json"
	"fmt"
	"net/http"

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

		utils.SendBadRequestResponse()

		w = utils.ResponseHandler(w, "Error decoding response object, incorrect input", "", "", http.StatusBadRequest)

		return
	}

	fmt.Println(u)

	res := models.JsonResponse{}

	res.Error = ""
	res.Msg = "Employee record created"
	res.Data = ""

	jsonResponse, jsonError := json.Marshal(res)

	if jsonError != nil {
		fmt.Println(jsonError)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonResponse)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
