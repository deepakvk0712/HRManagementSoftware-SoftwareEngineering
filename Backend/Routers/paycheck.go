package Routers

import (
	"github.com/gorilla/mux"
	"hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Controller"
	middleware "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Middleware"
	"net/http"
)

func PaycheckRouter() *mux.Router {
	router := mux.NewRouter()

	router.Path("/").Methods(http.MethodGet).HandlerFunc(Controller.GetPaycheck)
	router.Path("/team").Methods(http.MethodGet).HandlerFunc(Controller.GetAllSalaries)
	router.Use(middleware.ValidateAccessToken)

	return router
}
