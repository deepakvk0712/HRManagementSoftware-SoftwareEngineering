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

func TestPaychecks(t *testing.T) {
	router := mux.NewRouter()

	loginRouter := router.PathPrefix("/login").Subrouter()
	loginRouter.HandleFunc("", Login).Methods("POST")
	loginRouter.Use(middleware.ValidateUser)
	loginRouter.Use(middleware.Authorize)

	paycheckRouter := router.PathPrefix("/paycheck").Subrouter()
	paycheckRouter.HandleFunc("/", GetPaycheck).Methods("GET")
	paycheckRouter.HandleFunc("/updateSalary", UpdateEmployeeSalary).Methods("PUT")
	paycheckRouter.Use(middleware.ValidateAccessToken)

	utils.Init()

	type targetPaycheck struct {
		Paychecks    []gormModels.Paycheck `json:"paychecks"`
		IsManager    bool                  `json:"isManager"`
		TeamSalaries []models.TeamSalary   `json:"teamSalaries"`
	}

	t.Run("If employee is a manager, he should get his paycheck along with paychecks of his employees", func(t *testing.T) {
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

		//payload, err = json.Marshal(&{})
		//assert.NoError(t, err)

		createRequest, err := http.NewRequest("GET", "/paycheck/", nil)
		createRequest.Header.Set("Content-Type", "application/json")
		createRequest.Header.Set("Authorization", "Bearer "+target.AccessToken)
		assert.NoError(t, err)

		w = httptest.NewRecorder()

		router.ServeHTTP(w, createRequest)
		assert.Equal(t, 200, w.Code)

		paychecks := targetPaycheck{}

		json.NewDecoder(w.Body).Decode(&response)
		json.Unmarshal([]byte(response.Data), &paychecks)

		paychecksLength := len(paychecks.Paychecks)
		isManager := paychecks.IsManager
		teamSalariesLength := len(paychecks.TeamSalaries)

		assert.GreaterOrEqual(t, paychecksLength, 1)
		assert.GreaterOrEqual(t, teamSalariesLength, 1)
		assert.Equal(t, true, isManager)
	})

	t.Run("If employee is not a manager, he should get only his paycheck", func(t *testing.T) {
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

		createRequest, err := http.NewRequest("GET", "/paycheck/", nil)
		createRequest.Header.Set("Content-Type", "application/json")
		createRequest.Header.Set("Authorization", "Bearer "+target.AccessToken)
		assert.NoError(t, err)

		w = httptest.NewRecorder()

		router.ServeHTTP(w, createRequest)
		assert.Equal(t, 200, w.Code)

		paychecks := targetPaycheck{}

		json.NewDecoder(w.Body).Decode(&response)
		json.Unmarshal([]byte(response.Data), &paychecks)

		paychecksLength := len(paychecks.Paychecks)
		isManager := paychecks.IsManager
		teamSalariesLength := len(paychecks.TeamSalaries)

		assert.GreaterOrEqual(t, paychecksLength, 1)
		assert.Equal(t, 0, teamSalariesLength)
		assert.Equal(t, false, isManager)
	})

	t.Run("Only a manager must be able to update the salary of his employees", func(t *testing.T) {
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

		updatedSalary := models.UpdateSalary{
			EmployeeID: 58,
			NewSalary:  70000,
		}

		payload, err = json.Marshal(&updatedSalary)
		assert.NoError(t, err)

		createRequest, err := http.NewRequest("PUT", "/paycheck/updateSalary", bytes.NewBuffer(payload))
		createRequest.Header.Set("Content-Type", "application/json")
		createRequest.Header.Set("Authorization", "Bearer "+target.AccessToken)
		assert.NoError(t, err)

		w = httptest.NewRecorder()

		router.ServeHTTP(w, createRequest)
		assert.Equal(t, 200, w.Code)

		row := utils.Db.Raw("SELECT SALARY FROM USERS WHERE EMPLOYEE_ID = 58").Row()

		salary := 0.0

		row.Scan(&salary)

		assert.Equal(t, 70000.0, salary)
	})

	t.Run("If an employee is not a manager, he should not be able to access update salary", func(t *testing.T) {
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

		updatedSalary := models.UpdateSalary{
			EmployeeID: 58,
			NewSalary:  70000,
		}

		payload, err = json.Marshal(&updatedSalary)
		assert.NoError(t, err)

		createRequest, err := http.NewRequest("PUT", "/paycheck/updateSalary", bytes.NewBuffer(payload))
		createRequest.Header.Set("Content-Type", "application/json")
		createRequest.Header.Set("Authorization", "Bearer "+target.AccessToken)
		assert.NoError(t, err)

		w = httptest.NewRecorder()

		router.ServeHTTP(w, createRequest)
		assert.Equal(t, 403, w.Code)
	})
}
