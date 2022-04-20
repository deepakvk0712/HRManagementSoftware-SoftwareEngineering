package Routers

import (
	"github.com/gorilla/mux"
	"hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Controller"
	middleware "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Middleware"
	"net/http"
)

func NotificationRouter() *mux.Router {
	router := mux.NewRouter()

	router.Path("/").Methods(http.MethodPost).HandlerFunc(Controller.ReceiveNotification)
	router.Path("/teamMembers").Methods(http.MethodGet).HandlerFunc(Controller.GetTeamMembers)
	router.Path("/markRead").Methods(http.MethodPut).HandlerFunc(Controller.MarkRead)
	//router.Path("/teamMembers").Methods(http.MethodGet).HandlerFunc(Controller.GetTeamMembers)
	router.Use(middleware.ValidateAccessToken)

	return router
}
