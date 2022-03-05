package Controller

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSchedule(t *testing.T) {
	router := mux.NewRouter()

	workingRouter := router.PathPrefix("/working").Subrouter()
	workingRouter.HandleFunc("", GetSchedule).Methods("GET")

	utils.Init()

	t.Run("Get schedule", func(t *testing.T) {

		EmployeeID := 34

		payload, err := json.Marshal(&EmployeeID)
		assert.NoError(t, err)

		scheduleRequest, err := http.NewRequest("GET", "/getSchedule", bytes.NewBuffer(payload))
		scheduleRequest.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		w := httptest.NewRecorder()

		workingRouter.ServeHTTP(w, scheduleRequest)
		assert.Equal(t, 200, w.Code)
	})

}
func TestGetTodayWorkingHoursByID(t *testing.T) {
	router := mux.NewRouter()

	workingRouter := router.PathPrefix("/working").Subrouter()
	workingRouter.HandleFunc("", GetTodayWorkingHoursByID).Methods("GET")

	utils.Init()

	t.Run("Get Today Working Hours", func(t *testing.T) {

		EmployeeID := 34

		payload, err := json.Marshal(&EmployeeID)
		assert.NoError(t, err)

		todayHRequest, err := http.NewRequest("GET", "/getTodayWork", bytes.NewBuffer(payload))
		todayHRequest.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		w := httptest.NewRecorder()

		workingRouter.ServeHTTP(w, todayHRequest)
		assert.Equal(t, 200, w.Code)
	})

}
func TestGetWeekWorkingByID(t *testing.T) {
	router := mux.NewRouter()

	workingRouter := router.PathPrefix("/working").Subrouter()
	workingRouter.HandleFunc("", GetWeekWorkingByID).Methods("GET")

	utils.Init()

	t.Run("Get This Week's Working Hours", func(t *testing.T) {

		EmployeeID := 34

		payload, err := json.Marshal(&EmployeeID)
		assert.NoError(t, err)

		weekHRequest, err := http.NewRequest("GET", "/getWeekWork", bytes.NewBuffer(payload))
		weekHRequest.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		w := httptest.NewRecorder()

		workingRouter.ServeHTTP(w, weekHRequest)
		assert.Equal(t, 200, w.Code)
	})

}
