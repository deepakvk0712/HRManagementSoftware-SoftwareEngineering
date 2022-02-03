package main

import (
	Routers "github.com/AlexWeiZH/HRManagementSoftware-SoftwareEngineering/Backend/Routers"
	utils "github.com/AlexWeiZH/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	utils.Init()

	r := mux.NewRouter()

	Routers.RegisterUsers(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
