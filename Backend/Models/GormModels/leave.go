package gormModels

import "time"

type Leave struct {
	LeaveID   int       `json:"leaveID"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Subject   string    `json:"subject"`
	Reason    string    `json:"reason"`
	Email     string    `json:"email"`
	LeaveType string    `json:"leaveType"`
}
