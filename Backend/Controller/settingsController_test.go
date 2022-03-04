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

func TestChangePassword(t *testing.T) {
	router := mux.NewRouter()

	loginRouter := router.PathPrefix("/login").Subrouter()
	loginRouter.HandleFunc("", Login).Methods("POST")
	loginRouter.Use(middleware.ValidateUser)
	loginRouter.Use(middleware.Authorize)

	changePasswordRouter := router.PathPrefix("/resetPassword").Subrouter()
	changePasswordRouter.HandleFunc("", ChangePassword).Methods("PUT")
	changePasswordRouter.Use(middleware.ValidateAccessToken)

	utils.Init()

	t.Run("Change Password with correct old password", func(t *testing.T) {
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

		updatedUser := models.ChangeLoginInfo{
			Email:       "dstejas191@hrtools.com",
			OldPassword: "123",
			NewPassword: "123",
		}

		payload, err = json.Marshal(&updatedUser)
		assert.NoError(t, err)

		passwordRequest, err := http.NewRequest("PUT", "/resetPassword", bytes.NewBuffer(payload))
		passwordRequest.Header.Set("Content-Type", "application/json")
		passwordRequest.Header.Set("Authorization", "Bearer "+target.AccessToken)
		assert.NoError(t, err)

		w = httptest.NewRecorder()

		router.ServeHTTP(w, passwordRequest)
		assert.Equal(t, 201, w.Code)
	})

	t.Run("Should fail if incorrect credentials are provided", func(t *testing.T) {
		user := models.UserLogin{
			Email:    "dstejas19@hrtools.com",
			Password: "123",
		}

		payload, err := json.Marshal(&user)
		assert.NoError(t, err)

		loginRequest, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
		loginRequest.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		w := httptest.NewRecorder()

		router.ServeHTTP(w, loginRequest)
		assert.Equal(t, 400, w.Code)

		type targetType struct {
			AccessToken string `json:"accessToken"`
			FirstLogin  bool   `json:"firstLogin"`
		}

		response := models.JsonResponse{}
		var target targetType

		json.NewDecoder(w.Body).Decode(&response)
		json.Unmarshal([]byte(response.Data), &target)

		updatedUser := models.ChangeLoginInfo{
			Email:       "dstejas191@hrtools.com",
			OldPassword: "123",
			NewPassword: "123",
		}

		payload, err = json.Marshal(&updatedUser)
		assert.NoError(t, err)

		passwordRequest, err := http.NewRequest("PUT", "/resetPassword", bytes.NewBuffer(payload))
		passwordRequest.Header.Set("Content-Type", "application/json")
		passwordRequest.Header.Set("Authorization", "Bearer "+target.AccessToken)
		assert.NoError(t, err)

		w = httptest.NewRecorder()

		router.ServeHTTP(w, passwordRequest)
		assert.Equal(t, 401, w.Code)
	})

	t.Run("Cannot change password if user is not authenticated", func(t *testing.T) {
		updatedUser := models.ChangeLoginInfo{
			Email:       "dstejas191@hrtools.com",
			OldPassword: "123",
			NewPassword: "123",
		}

		payload, err := json.Marshal(&updatedUser)
		assert.NoError(t, err)

		passwordRequest, err := http.NewRequest("PUT", "/resetPassword", bytes.NewBuffer(payload))
		passwordRequest.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		w := httptest.NewRecorder()

		router.ServeHTTP(w, passwordRequest)
		assert.Equal(t, 401, w.Code)
	})
}
