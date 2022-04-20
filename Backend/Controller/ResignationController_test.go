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

func TestInsertFeedBack(t *testing.T) {
	router := mux.NewRouter()

	loginRouter := router.PathPrefix("/login").Subrouter()
	loginRouter.HandleFunc("", Login).Methods("POST")
	loginRouter.Use(middleware.ValidateUser)
	loginRouter.Use(middleware.Authorize)

	resignRouter := router.PathPrefix("/resign").Subrouter()
	resignRouter.HandleFunc("/insertFeedback", InsertFeedBack).Methods("POST")
	resignRouter.Use(middleware.ValidateAccessToken)

	utils.Init()

	t.Run("Insert FeedBack", func(t *testing.T) {
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

		Feedback := gormModels.Feedback{}
		Feedback.Feedback = "leave company"
		Feedback.Q1 = 1
		Feedback.Q2 = 2
		Feedback.Q3 = 3
		Feedback.Q4 = 4
		Feedback.Q5 = 5
		Feedback.ResignDate = "2022-03-04"

		payload, err = json.Marshal(&Feedback)
		assert.NoError(t, err)

		FeedbackRequest, err := http.NewRequest("POST", "/resign/insertFeedback", bytes.NewBuffer(payload))
		FeedbackRequest.Header.Set("Content-Type", "application/json")
		FeedbackRequest.Header.Set("Authorization", "Bearer "+target.AccessToken)
		assert.NoError(t, err)

		w = httptest.NewRecorder()

		resignRouter.ServeHTTP(w, FeedbackRequest)
		assert.Equal(t, 200, w.Code)
	})
}
