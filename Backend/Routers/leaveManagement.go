package Routers

import (
	"github.com/gorilla/mux"
	"hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Controller"
	middleware "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Middleware"
	"net/http"
)

func LeaveManagementRouter() *mux.Router {
	router := mux.NewRouter()

	router.Path("/apply").Methods(http.MethodPost).HandlerFunc(Controller.ApplyLeave)
	router.Use(middleware.ValidateAccessToken)

	return router
}
