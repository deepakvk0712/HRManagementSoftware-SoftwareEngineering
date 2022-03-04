package main

import (
	"hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Controller"
	middleware "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Middleware"
	"hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Routers"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	utils.Init()

	router := mux.NewRouter()

	postRequest := router.Methods(http.MethodPost).Subrouter()
	postRequest.HandleFunc("/login", Controller.Login)
	postRequest.HandleFunc("/logout", Controller.Logout)
	postRequest.Use(middleware.ValidateUser)
	postRequest.Use(middleware.Authorize)

	/*
		//  Test router

		getRequest := router.Methods(http.MethodGet).Subrouter()
		getRequest.HandleFunc("/test", test)
		getRequest.Use(middleware.ValidateAccessToken)

	*/

	mount(router, "/users", Routers.Router())
	mount(router, "/settings", Routers.SettingsRouter())
	mount(router, "/dashboard", Routers.DashboardRouter())

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin", "Referer", "Authorization"},
		AllowedMethods:   []string{"GET", "POST", "PUT"},
	})

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
	//log.Fatal(http.ListenAndServe(":8080", router))
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

/*
Test router functionality
//func test(w http.ResponseWriter, req *http.Request) {
//	w.Write([]byte(req.Context().Value("email").(string)))
//}


*/
