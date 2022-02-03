package Controller

import (
	"encoding/json"
	"fmt"
	models "github.com/AlexWeiZH/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	"net/http"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var employee models.Employee
	_ = json.NewDecoder(r.Body).Decode(&employee) //decode json from the front end and convert to employee
	fmt.Println(employee)
	/*
		Database operation
	*/
	json.NewEncoder(w).Encode(employee)
}
