package models

type UpdateSalary struct {
	EmployeeID int     `json:"employeeID"`
	NewSalary  float32 `json:"newSalary"`
}
