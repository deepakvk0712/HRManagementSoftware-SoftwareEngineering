package Controller

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	middleware "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Middleware"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
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
	updateRouter.Use(middleware.ValidateAccessToken)

	utils.Init()

	t.Run("Update employee info", func(t *testing.T) {
		user := models.UserLogin{
			Email:    "dstejas191@hrtools.com",
			Password: "123",
		}

		payload, err := json.Marshal(&user)
		assert.NoError(t, err)

		loginRequest, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
		loginRequest.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		w := httptest.NewRecorder()

		router.ServeHTTP(w, loginRequest)
		assert.Equal(t, 200, w.Code)

		type targetType struct {
			AccessToken string `json:"accessToken"`
			FirstLogin  bool   `json:"firstLogin"`
		}

		response := models.JsonResponse{}
		var target targetType

		json.NewDecoder(w.Body).Decode(&response)
		json.Unmarshal([]byte(response.Data), &target)

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

		payload, err = json.Marshal(&employeeInfo)
		assert.NoError(t, err)

		updateRequest, err := http.NewRequest("POST", "/users/UpdateEmployeeInfo1", bytes.NewBuffer(payload))
		updateRequest.Header.Set("Content-Type", "application/json")
		updateRequest.Header.Set("Authorization", "Bearer "+target.AccessToken)
		assert.NoError(t, err)

		w = httptest.NewRecorder()

		updateRouter.ServeHTTP(w, updateRequest)
		assert.Equal(t, 200, w.Code)
	})

}

func TestUpdateEmployeeBankingInfo(t *testing.T) {
	router := mux.NewRouter()

	loginRouter := router.PathPrefix("/login").Subrouter()
	loginRouter.HandleFunc("", Login).Methods("POST")
	loginRouter.Use(middleware.ValidateUser)
	loginRouter.Use(middleware.Authorize)

	updateRouter := router.PathPrefix("/users").Subrouter()
	updateRouter.HandleFunc("/UpdateEmployeeInfo2", UpdateEmployeeBankingInfo).Methods("POST")
	updateRouter.Use(middleware.ValidateAccessToken)

	utils.Init()

	t.Run("Update employee Banking info", func(t *testing.T) {

		user := models.UserLogin{
			Email:    "dstejas191@hrtools.com",
			Password: "123",
		}

		payload, err := json.Marshal(&user)
		assert.NoError(t, err)

		loginRequest, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
		loginRequest.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		w := httptest.NewRecorder()

		router.ServeHTTP(w, loginRequest)
		assert.Equal(t, 200, w.Code)

		type targetType struct {
			AccessToken string `json:"accessToken"`
			FirstLogin  bool   `json:"firstLogin"`
		}

		response := models.JsonResponse{}
		var target targetType

		json.NewDecoder(w.Body).Decode(&response)
		json.Unmarshal([]byte(response.Data), &target)

		employeeInfo := gormModels.User{
			EmployeeID:    34,
			RoutingNumber: "123",
			AccountNumber: "456",
			Bank:          "Wells Fargo"}

		payload, err = json.Marshal(&employeeInfo)
		assert.NoError(t, err)

		updateRequest, err := http.NewRequest("POST", "/users/UpdateEmployeeInfo2", bytes.NewBuffer(payload))
		updateRequest.Header.Set("Content-Type", "application/json")
		updateRequest.Header.Set("Authorization", "Bearer "+target.AccessToken)
		assert.NoError(t, err)

		w = httptest.NewRecorder()

		updateRouter.ServeHTTP(w, updateRequest)
		assert.Equal(t, 200, w.Code)
	})

}
