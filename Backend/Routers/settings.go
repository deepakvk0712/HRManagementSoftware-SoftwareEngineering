package Routers

import (
	"github.com/gorilla/mux"
	"hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Controller"
	"net/http"
)

func SettingsRouter() *mux.Router {
	router := mux.NewRouter()

	router.Path("/resetPassword").Methods(http.MethodPut).HandlerFunc(Controller.ChangePassword)

	return router
}
