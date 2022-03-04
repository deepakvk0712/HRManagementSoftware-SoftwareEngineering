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

//func TestSignUp(t *testing.T) {
//	var actualResult models.User
//
//	user := models.User{
//		Name:     "Test User",
//		Email:    "jwt@email.com",
//		Password: "secret",
//	}
//
//	payload, err := json.Marshal(&user)
//	assert.NoError(t, err)
//
//	request, err := http.NewRequest("POST", "/api/public/signup", bytes.NewBuffer(payload))
//	assert.NoError(t, err)
//
//	w := httptest.NewRecorder()
//
//	c, _ := gin.CreateTestContext(w)
//	c.Request = request
//
//	err = database.InitDatabase()
//	assert.NoError(t, err)
//
//	database.GlobalDB.AutoMigrate(&models.User{})
//
//	Signup(c)
//
//	assert.Equal(t, 200, w.Code)
//
//	err = json.Unmarshal(w.Body.Bytes(), &actualResult)
//	assert.NoError(t, err)
//
//	assert.Equal(t, user.Name, actualResult.Name)
//	assert.Equal(t, user.Email, actualResult.Email)
//}
//
//func TestSignUpInvalidJSON(t *testing.T) {
//	user := "test"
//
//	payload, err := json.Marshal(&user)
//	assert.NoError(t, err)
//
//	request, err := http.NewRequest("POST", "/api/public/signup", bytes.NewBuffer(payload))
//	assert.NoError(t, err)
//
//	w := httptest.NewRecorder()
//
//	c, _ := gin.CreateTestContext(w)
//	c.Request = request
//
//	Signup(c)
//
//	assert.Equal(t, 400, w.Code)
//}

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

		fmt.Println(response.Data)

		json.Unmarshal([]byte(response.Data), &target)

		fmt.Println("target   ", target)

		updatedUser := models.ChangeLoginInfo{
			Email:       "dstejas191@hrtools.com",
			OldPassword: "123",
			NewPassword: "123",
		}

		payload, err = json.Marshal(&updatedUser)
		assert.NoError(t, err)

		fmt.Println(target.AccessToken)

		passwordRequest, err := http.NewRequest("PUT", "/resetPassword", bytes.NewBuffer(payload))
		passwordRequest.Header.Set("Content-Type", "application/json")
		passwordRequest.Header.Set("Authorization", "Bearer "+target.AccessToken)
		assert.NoError(t, err)

		w = httptest.NewRecorder()

		router.ServeHTTP(w, passwordRequest)
		assert.Equal(t, 401, w.Code)
	})
}
