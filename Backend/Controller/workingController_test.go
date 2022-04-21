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

func TestSchedule(t *testing.T) {
	router := mux.NewRouter()

	workingRouter := router.PathPrefix("/working").Subrouter()
	workingRouter.HandleFunc("/getSchedule", GetSchedule).Methods("GET")
	workingRouter.HandleFunc("/getWorkingDetails", GetWorkingDetailsBetween).Methods("GET")
	workingRouter.HandleFunc("/insertWorkingRecord", SetWorkingHours).Methods("POST")

	workingRouter.Use(middleware.ValidateAccessToken)

	utils.Init()

	t.Run("Get schedule", func(t *testing.T) {

		EmployeeID := 34
		payload, err := json.Marshal(&EmployeeID)
		assert.NoError(t, err)

		scheduleRequest, err := http.NewRequest("GET", "/working/getSchedule", bytes.NewBuffer(payload))
		scheduleRequest.Header.Set("Content-Type", "application/json")
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImRzdGVqYXMxOTFAaHJ0b29scy5jb20iLCJTdGFuZGFyZENsYWltIjp7ImV4cCI6MTY1NDcwMTg2MywiaXNzIjoiTG9naW4ifSwiUm9sZSI6N30.jOH-ukfC09X0hjMSJsywpspU2GiP2_whAmLRlfRVVpM"
		scheduleRequest.Header.Set("Authorization", "Bearer "+token)
		assert.NoError(t, err)

		w := httptest.NewRecorder()

		workingRouter.ServeHTTP(w, scheduleRequest)
		assert.Equal(t, 200, w.Code)
	})

	t.Run("Get Working Details", func(t *testing.T) {

		FD := models.FrontendDate{}
		FD.StartDate = "2022-04-01"
		FD.EndDate = "2022-04-15"

		payload, err := json.Marshal(&FD)
		assert.NoError(t, err)

		weekHRequest, err := http.NewRequest("GET", "/working/getWorkingDetails", bytes.NewBuffer(payload))
		weekHRequest.Header.Set("Content-Type", "application/json")
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImRzdGVqYXMxOTFAaHJ0b29scy5jb20iLCJTdGFuZGFyZENsYWltIjp7ImV4cCI6MTY1NDcwMTg2MywiaXNzIjoiTG9naW4ifSwiUm9sZSI6N30.jOH-ukfC09X0hjMSJsywpspU2GiP2_whAmLRlfRVVpM"
		weekHRequest.Header.Set("Authorization", "Bearer "+token)
		assert.NoError(t, err)

		w := httptest.NewRecorder()

		workingRouter.ServeHTTP(w, weekHRequest)
		assert.Equal(t, 200, w.Code)
	})
	t.Run("Insert Working Record", func(t *testing.T) {

		FT := models.FrontendTime{}
		FT.Date = "2022-04-17"
		FT.StartTime = "09:00"
		FT.EndTime = "17:00"

		payload, err := json.Marshal(&FT)
		assert.NoError(t, err)

		weekHRequest, err := http.NewRequest("POST", "/working/insertWorkingRecord", bytes.NewBuffer(payload))
		weekHRequest.Header.Set("Content-Type", "application/json")
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImRzdGVqYXMxOTFAaHJ0b29scy5jb20iLCJTdGFuZGFyZENsYWltIjp7ImV4cCI6MTY1NDcwMTg2MywiaXNzIjoiTG9naW4ifSwiUm9sZSI6N30.jOH-ukfC09X0hjMSJsywpspU2GiP2_whAmLRlfRVVpM"
		weekHRequest.Header.Set("Authorization", "Bearer "+token)
		assert.NoError(t, err)

		w := httptest.NewRecorder()

		workingRouter.ServeHTTP(w, weekHRequest)
		assert.Equal(t, 201, w.Code)
	})
}
