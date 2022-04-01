package models

import "time"

type TeamSalary struct {
	EmployeeName string    `json:"employeeName"`
	EmployeeID   string    `json:"employeeID"`
	CheckDate    time.Time `json:"checkDate"`
	PayBeginDate time.Time `json:"payBeginDate"`
	PayEndDate   time.Time `json:"payEndDate"`
	AmountPaid   float32   `json:"amountPaid"`
}
