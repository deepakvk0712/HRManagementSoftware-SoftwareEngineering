package Controller

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	middleware "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Middleware"
	gormModels "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/GormModels"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateEmployeeInfo(t *testing.T) {
	router := mux.NewRouter()

	loginRouter := router.PathPrefix("/login").Subrouter()
	loginRouter.HandleFunc("", Login).Methods("POST")
	loginRouter.Use(middleware.ValidateUser)
	loginRouter.Use(middleware.Authorize)

	updateRouter := router.PathPrefix("/users").Subrouter()
	updateRouter.HandleFunc("/UpdateEmployeeInfo1", UpdateEmployeeInfo).Methods("POST")
	updateRouter.HandleFunc("/UpdateEmployeeInfo2", UpdateEmployeeBankingInfo).Methods("POST")
	updateRouter.Use(middleware.ValidateAccessToken)

	utils.Init()

	t.Run("Update employee info", func(t *testing.T) {

		employeeInfo := gormModels.User{
			EmployeeID:      34,
			DriversLicense:  "abcdefg",
			SSN:             "45678942",
			StateID:         "cn123456798",
			Address:         "4915 SW 14th Pl",
			Phone:           "3528883848",
			AlternateEmails: "weizihan@ufl.edu",
			FirstName:       "Zihan",
			LastName:        "Wei",
			Gender:          "Male",
			DateOfBirth:     "05/03/1998",
			Race:            "Asian",
			Ethnicity:       "Han",
			Citizenship:     "Chinese",
			Nationality:     "China",
			Pronouns:        "Alex"}

		payload, err := json.Marshal(&employeeInfo)
		assert.NoError(t, err)

		updateRequest, err := http.NewRequest("POST", "/users/UpdateEmployeeInfo1", bytes.NewBuffer(payload))
		updateRequest.Header.Set("Content-Type", "application/json")
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImRzdGVqYXMxOTFAaHJ0b29scy5jb20iLCJTdGFuZGFyZENsYWltIjp7ImV4cCI6MTY1NDcwMTg2MywiaXNzIjoiTG9naW4ifSwiUm9sZSI6N30.jOH-ukfC09X0hjMSJsywpspU2GiP2_whAmLRlfRVVpM"
		updateRequest.Header.Set("Authorization", "Bearer "+token)
		assert.NoError(t, err)

		w := httptest.NewRecorder()

		updateRouter.ServeHTTP(w, updateRequest)
		assert.Equal(t, 201, w.Code)
	})
	t.Run("Update employee Banking info", func(t *testing.T) {

		employeeInfo := gormModels.User{
			EmployeeID:    34,
			RoutingNumber: "123",
			AccountNumber: "456",
			Bank:          "Wells Fargo"}

		payload, err := json.Marshal(&employeeInfo)
		assert.NoError(t, err)

		updateRequest, err := http.NewRequest("POST", "/users/UpdateEmployeeInfo2", bytes.NewBuffer(payload))
		updateRequest.Header.Set("Content-Type", "application/json")
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImRzdGVqYXMxOTFAaHJ0b29scy5jb20iLCJTdGFuZGFyZENsYWltIjp7ImV4cCI6MTY1NDcwMTg2MywiaXNzIjoiTG9naW4ifSwiUm9sZSI6N30.jOH-ukfC09X0hjMSJsywpspU2GiP2_whAmLRlfRVVpM"
		updateRequest.Header.Set("Authorization", "Bearer "+token)
		assert.NoError(t, err)

		w := httptest.NewRecorder()

		updateRouter.ServeHTTP(w, updateRequest)
		assert.Equal(t, 201, w.Code)
	})
}
