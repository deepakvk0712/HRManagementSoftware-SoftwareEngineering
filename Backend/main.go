package main

import (
	"hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Routers"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	utils.Init()

	router := mux.NewRouter()

	router.HandleFunc("/", home)

	mount(router, "/register", Routers.Router())

	log.Fatal(http.ListenAndServe(":8080", router))
}

func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}

func home(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Home"))
}
