package Routers

import (
	"fmt"
	"github.com/gorilla/mux"
	"hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Controller"
	"net/http"
)

func Router() *mux.Router {
	fmt.Println("Hey")
	router := mux.NewRouter()

	router.Path("/HR").Methods(http.MethodPost).HandlerFunc(Controller.RegisterHR)

	return router
}
