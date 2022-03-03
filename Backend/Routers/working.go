package Routers

import (
	"github.com/gorilla/mux"
	"hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Controller"
	middleware "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Middleware"
	"net/http"
)

func WorkingRouter() *mux.Router {
	router := mux.NewRouter()

	router.Path("/startWorking").Methods(http.MethodGet).HandlerFunc(Controller.StartWorking)
	router.Path("/stopWorking").Methods(http.MethodPost).HandlerFunc(Controller.StopWorking)
	router.Use(middleware.ValidateAccessToken)

	return router
}
