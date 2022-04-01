package models

type TeamMember struct {
	Name       string `json:"name"`
	EmailID    string `json:"emailID"`
	EmployeeID int    `json:"employeeID"`
}
