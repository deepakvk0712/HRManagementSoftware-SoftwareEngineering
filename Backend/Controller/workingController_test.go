package Controller

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	middleware "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Middleware"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSchedule(t *testing.T) {
	router := mux.NewRouter()

	loginRouter := router.PathPrefix("/login").Subrouter()
	loginRouter.HandleFunc("", Login).Methods("POST")
	loginRouter.Use(middleware.ValidateUser)
	loginRouter.Use(middleware.Authorize)

	workingRouter := router.PathPrefix("/working").Subrouter()
	workingRouter.HandleFunc("/getSchedule", GetSchedule).Methods("GET")
	workingRouter.HandleFunc("/getWorkingDetails", GetWorkingDetailsBetween).Methods("GET")
	workingRouter.Use(middleware.ValidateAccessToken)

	utils.Init()

	t.Run("Get schedule", func(t *testing.T) {
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

		EmployeeID := 34
		payload, err = json.Marshal(&EmployeeID)
		assert.NoError(t, err)

		scheduleRequest, err := http.NewRequest("GET", "/working/getSchedule", bytes.NewBuffer(payload))
		scheduleRequest.Header.Set("Content-Type", "application/json")
		scheduleRequest.Header.Set("Authorization", "Bearer "+target.AccessToken)
		assert.NoError(t, err)

		w = httptest.NewRecorder()

		workingRouter.ServeHTTP(w, scheduleRequest)
		assert.Equal(t, 200, w.Code)
	})

	t.Run("Get Working Details", func(t *testing.T) {
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

		FD := models.FrontendDate{}
		FD.StartDate = "2022-03-31"
		FD.EndDate = "2022-04-05"

		payload, err = json.Marshal(&FD)
		assert.NoError(t, err)

		weekHRequest, err := http.NewRequest("GET", "/working/getWorkingDetails", bytes.NewBuffer(payload))
		weekHRequest.Header.Set("Content-Type", "application/json")
		weekHRequest.Header.Set("Authorization", "Bearer "+target.AccessToken)
		assert.NoError(t, err)

		w = httptest.NewRecorder()

		workingRouter.ServeHTTP(w, weekHRequest)
		assert.Equal(t, 200, w.Code)
	})

}
