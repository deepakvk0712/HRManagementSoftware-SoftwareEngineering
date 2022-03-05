package Controller

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	gormModels "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/GormModels"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateEmployeeInfo(t *testing.T) {
	router := mux.NewRouter()

	updateRouter := router.PathPrefix("/users").Subrouter()
	updateRouter.HandleFunc("", UpdateEmployeeInfo).Methods("POST")

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

		updateRequest, err := http.NewRequest("POST", "/UpdateEmployeeInfo1", bytes.NewBuffer(payload))
		updateRequest.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		w := httptest.NewRecorder()

		updateRouter.ServeHTTP(w, updateRequest)
		assert.Equal(t, 200, w.Code)
	})

}

func TestUpdateEmployeeBankingInfo(t *testing.T) {
	router := mux.NewRouter()

	updateRouter := router.PathPrefix("/users").Subrouter()
	updateRouter.HandleFunc("", UpdateEmployeeBankingInfo).Methods("POST")

	utils.Init()

	t.Run("Update employee Banking info", func(t *testing.T) {

		employeeInfo := gormModels.User{
			EmployeeID:    34,
			RoutingNumber: "123",
			AccountNumber: "456",
			Bank:          "Wells Fargo"}

		payload, err := json.Marshal(&employeeInfo)
		assert.NoError(t, err)

		updateRequest, err := http.NewRequest("POST", "/UpdateEmployeeInfo2", bytes.NewBuffer(payload))
		updateRequest.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		w := httptest.NewRecorder()

		updateRouter.ServeHTTP(w, updateRequest)
		assert.Equal(t, 200, w.Code)
	})

}
