package Routers

import (
	"fmt"
	Controller "github.com/AlexWeiZH/HRManagementSoftware-SoftwareEngineering/Backend/Controller"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func RegisterUsers(router *mux.Router) {
	router.HandleFunc("/", Controller.Home)
	mount(router, "/register ", Router())

}

func Router() *mux.Router {
	fmt.Println("Hey")
	router := mux.NewRouter()
	router.Path("/HR").Methods(http.MethodPost).HandlerFunc(Controller.RegisterHR)
	//collecting employee's personal info
	router.Path("/CreateEmployee").Methods(http.MethodPost).HandlerFunc(Controller.CreateEmployee)

	return router
}
func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
