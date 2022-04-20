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

func TestDashboard(t *testing.T) {
	router := mux.NewRouter()

	loginRouter := router.PathPrefix("/login").Subrouter()
	loginRouter.HandleFunc("", Login).Methods("POST")
	loginRouter.Use(middleware.ValidateUser)
	loginRouter.Use(middleware.Authorize)

	trainingRouter := router.PathPrefix("/dashboard").Subrouter()
	trainingRouter.HandleFunc("/", GetDashboard).Methods("GET")
	trainingRouter.Use(middleware.ValidateAccessToken)

	utils.Init()

	t.Run("When a call to get dashboard has been made, he should get all the necessary fields", func(t *testing.T) {
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

		//payload, err = json.Marshal(&sentMessage)
		//assert.NoError(t, err)

		createRequest, err := http.NewRequest("GET", "/dashboard/", nil)
		createRequest.Header.Set("Content-Type", "application/json")
		createRequest.Header.Set("Authorization", "Bearer "+target.AccessToken)
		assert.NoError(t, err)

		w = httptest.NewRecorder()

		router.ServeHTTP(w, createRequest)
		assert.Equal(t, 200, w.Code)

		var dashboardResponse struct {
			AccountType   byte                             `json:"accountType"`
			Username      string                           `json:"username"`
			IsManager     bool                             `json:"isManager"`
			IsHR          bool                             `json:"isHR"`
			IsOnboard     bool                             `json:"isOnboard"`
			IsFinance     bool                             `json:"isFinance"`
			TeamMembers   []models.TeamMember              `json:"teamMembers"`
			BusinessUnits []string                         `json:"businessUnits"`
			Messages      []models.SendNotificationMessage `json:"messages"`
			PaidLeaves    int                              `json:"paidLeaves"`
			UnpaidLeaves  int                              `json:"unpaidLeaves"`
		}

		json.NewDecoder(w.Body).Decode(&response)
		err = json.Unmarshal([]byte(response.Data), &dashboardResponse)

		assert.NoError(t, err)
	})
}
