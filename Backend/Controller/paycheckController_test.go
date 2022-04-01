package Controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	middleware "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Middleware"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPaychecks(t *testing.T) {
	router := mux.NewRouter()

	loginRouter := router.PathPrefix("/login").Subrouter()
	loginRouter.HandleFunc("", Login).Methods("POST")
	loginRouter.Use(middleware.ValidateUser)
	loginRouter.Use(middleware.Authorize)

	changePasswordRouter := router.PathPrefix("/registerHR").Subrouter()
	changePasswordRouter.HandleFunc("", RegisterHR).Methods("POST")
	changePasswordRouter.Use(middleware.ValidateAccessToken)

	utils.Init()

	newHR := models.User{
		FirstName:     "Test",
		LastName:      "One",
		BusinessUnit:  "test",
		ManagerId:     1213,
		Grade:         "6A",
		Location:      "Bangalore",
		Country:       "India",
		Title:         "SWE",
		Type:          "",
		PersonalEmail: "tejasdharmastala@ufl.edu",
		IsHR:          true,
		IsManager:     false,
	}

	newEmployee := models.User{
		FirstName:     "Test",
		LastName:      "One",
		BusinessUnit:  "test",
		ManagerId:     1213,
		Grade:         "6A",
		Location:      "Bangalore",
		Country:       "India",
		Title:         "SWE",
		Type:          "",
		PersonalEmail: "tejasdharmastala@ufl.edu",
		IsHR:          false,
		IsManager:     false,
	}

	newHREmail := ""
	newEmployeeEmail := ""

	newHRPassword := ""
	newEmployeePassword := ""

	type targetEmployee struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	t.Run("Register employee successfully", func(t *testing.T) {
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

		payload, err = json.Marshal(&newHR)
		assert.NoError(t, err)

		createRequest, err := http.NewRequest("POST", "/registerHR", bytes.NewBuffer(payload))
		createRequest.Header.Set("Content-Type", "application/json")
		createRequest.Header.Set("Authorization", "Bearer "+target.AccessToken)
		assert.NoError(t, err)

		w = httptest.NewRecorder()

		router.ServeHTTP(w, createRequest)
		assert.Equal(t, 201, w.Code)

		employee := targetEmployee{}

		json.NewDecoder(w.Body).Decode(&response)
		json.Unmarshal([]byte(response.Data), &employee)
		newHRPassword = employee.Password
		newHREmail = employee.Email
	})

	t.Run("If the new user is HR, he should be able to create other users", func(t *testing.T) {
		fmt.Println(newHREmail, newHRPassword)
		user := models.UserLogin{
			Email:    newHREmail,
			Password: newHRPassword,
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

		payload, err = json.Marshal(&newEmployee)
		assert.NoError(t, err)

		createRequest, err := http.NewRequest("POST", "/registerHR", bytes.NewBuffer(payload))
		createRequest.Header.Set("Content-Type", "application/json")
		createRequest.Header.Set("Authorization", "Bearer "+target.AccessToken)
		assert.NoError(t, err)

		w = httptest.NewRecorder()

		router.ServeHTTP(w, createRequest)
		assert.Equal(t, 201, w.Code)

		employee := targetEmployee{}

		json.NewDecoder(w.Body).Decode(&response)
		json.Unmarshal([]byte(response.Data), &employee)
		newEmployeePassword = employee.Password
		newEmployeeEmail = employee.Email
	})

	t.Run("If the user is not a HR, he cannot register an employee", func(t *testing.T) {
		user := models.UserLogin{
			Email:    newEmployeeEmail,
			Password: newEmployeePassword,
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

		payload, err = json.Marshal(&newEmployee)
		assert.NoError(t, err)

		createRequest, err := http.NewRequest("POST", "/registerHR", bytes.NewBuffer(payload))
		createRequest.Header.Set("Content-Type", "application/json")
		createRequest.Header.Set("Authorization", "Bearer "+target.AccessToken)
		assert.NoError(t, err)
		w = httptest.NewRecorder()

		router.ServeHTTP(w, createRequest)
		assert.Equal(t, 403, w.Code)

		employee := targetEmployee{}

		json.NewDecoder(w.Body).Decode(&response)
		json.Unmarshal([]byte(response.Data), &employee)
	})

	utils.Db.Exec("DELETE FROM USERS WHERE LOWER(OFFICIAL_EMAIL) = ? OR LOWER(OFFICIAL_EMAIL) = ? ", strings.ToLower(newHREmail), strings.ToLower(newEmployeeEmail))
	utils.Db.Exec("DELETE FROM LOGIN_USERS WHERE LOWER(EMAIL) = ? OR LOWER(EMAIL) = ? ", strings.ToLower(newHREmail), strings.ToLower(newEmployeeEmail))
}
