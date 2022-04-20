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
	"strconv"
	"strings"
	"testing"
)

func TestNotifications(t *testing.T) {
	router := mux.NewRouter()

	loginRouter := router.PathPrefix("/login").Subrouter()
	loginRouter.HandleFunc("", Login).Methods("POST")
	loginRouter.Use(middleware.ValidateUser)
	loginRouter.Use(middleware.Authorize)

	notificationRouter := router.PathPrefix("/notify").Subrouter()
	notificationRouter.HandleFunc("/", ReceiveNotification).Methods("POST")
	notificationRouter.HandleFunc("/teamMembers", GetTeamMembers).Methods("GET")
	notificationRouter.HandleFunc("/markRead", MarkRead).Methods("PUT")
	notificationRouter.Use(middleware.ValidateAccessToken)

	dashboardRouter := router.PathPrefix("/dashboard").Subrouter()
	dashboardRouter.HandleFunc("", GetDashboard).Methods("GET")
	dashboardRouter.Use(middleware.ValidateAccessToken)

	utils.Init()

	sender := "vkashyapdeepak1@hrtools.com"
	receiver := "vkashyapdeepak4@hrtools.com"
	message := "testing this now"
	messageId := 0

	sentMessage := models.ReceiveNotificationMessage{
		Receiver: receiver,
		Message:  message,
	}

	t.Run("If an employee sends a notification, it should be saved in the DB and read flag should be false", func(t *testing.T) {
		user := models.UserLogin{
			Email:    sender,
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

		payload, err = json.Marshal(&sentMessage)
		assert.NoError(t, err)

		createRequest, err := http.NewRequest("POST", "/notify/", bytes.NewBuffer(payload))
		createRequest.Header.Set("Content-Type", "application/json")
		createRequest.Header.Set("Authorization", "Bearer "+target.AccessToken)
		assert.NoError(t, err)

		w = httptest.NewRecorder()

		router.ServeHTTP(w, createRequest)
		assert.Equal(t, 201, w.Code)

		notification := gormModels.Notification{}
		row := utils.Db.Raw("SELECT * FROM NOTIFICATIONS WHERE LOWER(SENDER) = ? AND MESSAGE = ?", strings.ToLower(sender), message).Row()

		if row.Err() != nil {
			assert.NoError(t, row.Err())
		}

		row.Scan(&notification.Sender, &notification.Receiver, &notification.Message, &notification.Read, &notification.MessageID)

		assert.Equal(t, message, notification.Message)
		assert.Equal(t, false, notification.Read)

	})

	t.Run("If there are any unread notifications, they should be displayed to the receiver", func(t *testing.T) {
		user := models.UserLogin{
			Email:    receiver,
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

		createRequest, err := http.NewRequest("GET", "/dashboard", nil)
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
		json.Unmarshal([]byte(response.Data), &dashboardResponse)

		messageId = dashboardResponse.Messages[0].MessageID

		assert.Equal(t, 1, len(dashboardResponse.Messages))
		assert.Equal(t, message, dashboardResponse.Messages[0].Message)
		assert.Equal(t, sender, dashboardResponse.Messages[0].Sender)
	})

	t.Run("After marking the notification as read, user should not see the message again", func(t *testing.T) {
		user := models.UserLogin{
			Email:    receiver,
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

		createRequest, err := http.NewRequest("PUT", "/notify/markRead?id="+strconv.Itoa(messageId), nil)
		createRequest.Header.Set("Content-Type", "application/json")
		createRequest.Header.Set("Authorization", "Bearer "+target.AccessToken)
		assert.NoError(t, err)

		w = httptest.NewRecorder()

		router.ServeHTTP(w, createRequest)
		assert.Equal(t, 200, w.Code)

		createRequest, err = http.NewRequest("GET", "/dashboard", nil)
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
		json.Unmarshal([]byte(response.Data), &dashboardResponse)

		assert.Equal(t, 0, len(dashboardResponse.Messages))
	})

	t.Run("Only HRs should be able to get team members of particular business units", func(t *testing.T) {
		user := models.UserLogin{
			Email:    "vkashyapdeepak1@hrtools.com",
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

		createRequest, err := http.NewRequest("GET", "/notify/teamMembers?businessUnit=Management", nil)
		createRequest.Header.Set("Content-Type", "application/json")
		createRequest.Header.Set("Authorization", "Bearer "+target.AccessToken)
		assert.NoError(t, err)

		w = httptest.NewRecorder()

		router.ServeHTTP(w, createRequest)
		assert.Equal(t, 200, w.Code)

		var membersResponse struct {
			TeamMembers []models.TeamMember `json:"teamMembers"`
		}

		json.NewDecoder(w.Body).Decode(&response)
		json.Unmarshal([]byte(response.Data), &membersResponse)

		assert.Less(t, 0, len(membersResponse.TeamMembers))
	})

	t.Run("Team members should not work for other employees", func(t *testing.T) {
		user := models.UserLogin{
			Email:    "vkashyapdeepak4@hrtools.com",
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

		createRequest, err := http.NewRequest("GET", "/notify/teamMembers?businessUnit=Management", nil)
		createRequest.Header.Set("Content-Type", "application/json")
		createRequest.Header.Set("Authorization", "Bearer "+target.AccessToken)
		assert.NoError(t, err)

		w = httptest.NewRecorder()

		router.ServeHTTP(w, createRequest)
		assert.Equal(t, 403, w.Code)
	})
}
