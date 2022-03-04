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

func TestLogin(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/login", Login).Methods("POST")
	router.Use(middleware.ValidateUser)
	router.Use(middleware.Authorize)

	utils.Init()

	t.Run("Valid Credentials", func(t *testing.T) {
		user := models.UserLogin{
			Email:    "dstejas191@hrtools.com",
			Password: "123",
		}

		payload, err := json.Marshal(&user)
		assert.NoError(t, err)

		request, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
		request.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		w := httptest.NewRecorder()

		router.ServeHTTP(w, request)
		assert.Equal(t, 200, w.Code)
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		user := struct {
			Email string
		}{
			Email: "dstejas191@hrtools.com",
		}

		payload, err := json.Marshal(&user)
		assert.NoError(t, err)

		request, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
		request.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		w := httptest.NewRecorder()

		router.ServeHTTP(w, request)
		assert.Equal(t, 400, w.Code)
	})

	t.Run("Invalid Credentials", func(t *testing.T) {
		user := models.UserLogin{
			Email:    "dstejas191@hrtools.com",
			Password: "dfdf",
		}

		payload, err := json.Marshal(&user)
		assert.NoError(t, err)

		request, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
		request.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		w := httptest.NewRecorder()

		router.ServeHTTP(w, request)
		assert.Equal(t, 400, w.Code)
	})
}
