package Routers

import (
	"github.com/gorilla/mux"
	"hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Controller"
	middleware "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Middleware"
	"net/http"
)

func DashboardRouter() *mux.Router {
	router := mux.NewRouter()

	router.Path("/").Methods(http.MethodGet).HandlerFunc(Controller.GetDashboard)
	router.Use(middleware.ValidateAccessToken)

	return router
}
