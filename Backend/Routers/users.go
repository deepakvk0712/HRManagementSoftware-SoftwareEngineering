package Routers

import (
	"fmt"
	"github.com/gorilla/mux"
	"hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Controller"
	middleware "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Middleware"
	"net/http"
)

func Router() *mux.Router {
	fmt.Println("Hey")
	router := mux.NewRouter()

	router.Path("/registerHR").Methods(http.MethodPost).HandlerFunc(Controller.RegisterHR)
	router.Path("/profile").Methods(http.MethodGet).HandlerFunc(Controller.GetProfile)
	router.Path("/profile").Methods(http.MethodPut).HandlerFunc(Controller.UpdateProfile)
	router.Path("/UpdateEmployeeInfo1").Methods(http.MethodPost).HandlerFunc(Controller.UpdateEmployeeInfo)
	router.Path("/UpdateEmployeeInfo2").Methods(http.MethodPost).HandlerFunc(Controller.UpdateEmployeeBankingInfo)
	router.Use(middleware.ValidateAccessToken)

	return router
}
