package main

import (
	"log"
	"net/http"
	"strings"

	db "github.com/dstejas19/HRManagementSoftware-SoftwareEngineering/Backend/Database"
	register "github.com/dstejas19/HRManagementSoftware-SoftwareEngineering/Backend/Register"

	"github.com/gorilla/mux"
)

func main() {
	db.ConnectToDB()

	router := mux.NewRouter()

	router.HandleFunc("/", home)

	mount(router, "/register", register.Router())

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
