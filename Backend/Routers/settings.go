package Routers

import (
	"github.com/gorilla/mux"
	"hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Controller"
	middleware "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Middleware"
	"net/http"
)

func SettingsRouter() *mux.Router {
	router := mux.NewRouter()

	router.Path("/resetPassword").Methods(http.MethodPut).HandlerFunc(Controller.ChangePassword)
	router.Use(middleware.ValidateAccessToken)

	return router
}
