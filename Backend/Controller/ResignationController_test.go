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

func TestInsertFeedBack(t *testing.T) {
	router := mux.NewRouter()

	resignRouter := router.PathPrefix("/resign").Subrouter()
	resignRouter.HandleFunc("/insertFeedback", InsertFeedBack).Methods("POST")
	resignRouter.Use(middleware.ValidateAccessToken)

	utils.Init()

	t.Run("Insert FeedBack", func(t *testing.T) {

		Feedback := gormModels.Feedback{}
		Feedback.Q1 = 1
		Feedback.Q2 = 2
		Feedback.Q3 = 3
		Feedback.Q4 = 4
		Feedback.Q5 = 5
		Feedback.Feedback = "leave company"
		Feedback.ResignDate = "2022-03-04"

		payload, err := json.Marshal(&Feedback)
		assert.NoError(t, err)

		FeedbackRequest, err := http.NewRequest("POST", "/resign/insertFeedback", bytes.NewBuffer(payload))
		FeedbackRequest.Header.Set("Content-Type", "application/json")
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImRzdGVqYXMxOTFAaHJ0b29scy5jb20iLCJTdGFuZGFyZENsYWltIjp7ImV4cCI6MTY1NDcwMTg2MywiaXNzIjoiTG9naW4ifSwiUm9sZSI6N30.jOH-ukfC09X0hjMSJsywpspU2GiP2_whAmLRlfRVVpM"
		FeedbackRequest.Header.Set("Authorization", "Bearer "+token)
		assert.NoError(t, err)

		w := httptest.NewRecorder()

		resignRouter.ServeHTTP(w, FeedbackRequest)
		assert.Equal(t, 201, w.Code)
	})
}

