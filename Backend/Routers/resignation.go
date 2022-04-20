package Routers

import (
	"github.com/gorilla/mux"
	"hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Controller"
	middleware "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Middleware"
	"net/http"
)

func ResignationRouter() *mux.Router {
	router := mux.NewRouter()

	router.Path("/insertFeedback").Methods(http.MethodPost).HandlerFunc(Controller.InsertFeedBack)
	router.Use(middleware.ValidateAccessToken)

	return router
}
