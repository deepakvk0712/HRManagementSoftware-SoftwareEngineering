package Routers

import (
	"github.com/gorilla/mux"
	"hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Controller"
	middleware "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Middleware"
	"net/http"
)

func TrainingRouter() *mux.Router {
	router := mux.NewRouter()

	router.Path("/").Methods(http.MethodGet).HandlerFunc(Controller.GetTrainingStatus)
	router.Path("/send").Methods(http.MethodPost).HandlerFunc(Controller.PostTrainingQuestions)
	
	router.Use(middleware.ValidateAccessToken)

	return router
}
