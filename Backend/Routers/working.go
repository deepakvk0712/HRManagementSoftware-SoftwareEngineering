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
	router.Path("/setSchedule").Methods(http.MethodPost).HandlerFunc(Controller.SetSchedule)
	router.Path("/getSchedule").Methods(http.MethodGet).HandlerFunc(Controller.GetSchedule)
	router.Path("/delSchedule").Methods(http.MethodDelete).HandlerFunc(Controller.DelSchedule)
	router.Path("/getWeekWork").Methods(http.MethodGet).HandlerFunc(Controller.GetWeekWorkingByID)
	router.Path("/getTodayWork").Methods(http.MethodGet).HandlerFunc(Controller.GetTodayWorkingHoursByID)
	router.Path("/getWorkingDetails").Methods(http.MethodGet).HandlerFunc(Controller.GetWorkingDetailsBetween)

	router.Use(middleware.ValidateAccessToken)

	return router
}
