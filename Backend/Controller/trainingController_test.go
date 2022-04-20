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
	"testing"
)

func TestTraining(t *testing.T) {
	router := mux.NewRouter()

	loginRouter := router.PathPrefix("/login").Subrouter()
	loginRouter.HandleFunc("", Login).Methods("POST")
	loginRouter.Use(middleware.ValidateUser)
	loginRouter.Use(middleware.Authorize)

	trainingRouter := router.PathPrefix("/training").Subrouter()
	trainingRouter.HandleFunc("/", GetTrainingStatus).Methods("GET")
	trainingRouter.HandleFunc("/send", PostTrainingQuestions).Methods("POST")
	trainingRouter.Use(middleware.ValidateAccessToken)

	utils.Init()

	t.Run("When an employee hasn't completed training, he should get training completed as false", func(t *testing.T) {
		user := models.UserLogin{
			Email:    "vkashyapdeepak5@hrtools.com",
			Password: "abcd1234",
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

		//payload, err = json.Marshal(&sentMessage)
		//assert.NoError(t, err)

		createRequest, err := http.NewRequest("GET", "/training/", nil)
		createRequest.Header.Set("Content-Type", "application/json")
		createRequest.Header.Set("Authorization", "Bearer "+target.AccessToken)
		assert.NoError(t, err)

		w = httptest.NewRecorder()

		router.ServeHTTP(w, createRequest)
		assert.Equal(t, 200, w.Code)

		var trainingResponse struct {
			TrainingCompleted bool `json:"trainingCompleted"`
			Score             int  `json:"score"`
		}

		json.NewDecoder(w.Body).Decode(&response)
		json.Unmarshal([]byte(response.Data), &trainingResponse)

		assert.Equal(t, false, trainingResponse.TrainingCompleted)
		assert.Equal(t, -1, trainingResponse.Score)
	})

	t.Run("When an employee submits his answers, the response should contain the score", func(t *testing.T) {
		user := models.UserLogin{
			Email:    "vkashyapdeepak5@hrtools.com",
			Password: "abcd1234",
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

		trainingSolutions := models.TrainingQuestions{
			QuestionOne:   3,
			QuestionTwo:   3,
			QuestionThree: 1,
			QuestionFour:  1,
			QuestionFive:  2,
			QuestionSix:   4,
			QuestionSeven: 3,
			QuestionEight: 4,
			QuestionNine:  4,
			QuestionTen:   4,
		}

		payload, err = json.Marshal(&trainingSolutions)
		assert.NoError(t, err)

		createRequest, err := http.NewRequest("POST", "/training/send", bytes.NewBuffer(payload))
		createRequest.Header.Set("Content-Type", "application/json")
		createRequest.Header.Set("Authorization", "Bearer "+target.AccessToken)
		assert.NoError(t, err)

		w = httptest.NewRecorder()

		router.ServeHTTP(w, createRequest)
		assert.Equal(t, 201, w.Code)

		var trainingResponse struct {
			Score int `json:"score"`
		}

		json.NewDecoder(w.Body).Decode(&response)
		json.Unmarshal([]byte(response.Data), &trainingResponse)

		fmt.Println(trainingResponse.Score)

		assert.Less(t, -1, trainingResponse.Score)
	})

	t.Run("When an employee has completed training, he should get training completed as true and his score", func(t *testing.T) {
		user := models.UserLogin{
			Email:    "vkashyapdeepak5@hrtools.com",
			Password: "abcd1234",
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

		//payload, err = json.Marshal(&sentMessage)
		//assert.NoError(t, err)

		createRequest, err := http.NewRequest("GET", "/training/", nil)
		createRequest.Header.Set("Content-Type", "application/json")
		createRequest.Header.Set("Authorization", "Bearer "+target.AccessToken)
		assert.NoError(t, err)

		w = httptest.NewRecorder()

		router.ServeHTTP(w, createRequest)
		assert.Equal(t, 200, w.Code)

		var trainingResponse struct {
			TrainingCompleted bool `json:"trainingCompleted"`
			Score             int  `json:"score"`
		}

		json.NewDecoder(w.Body).Decode(&response)
		json.Unmarshal([]byte(response.Data), &trainingResponse)

		assert.Equal(t, true, trainingResponse.TrainingCompleted)
		assert.Less(t, -1, trainingResponse.Score)
	})

	utils.Db.Exec("DELETE FROM TRAININGS WHERE EMAIL = ?", "vkashyapdeepak5@hrtools.com")
}
