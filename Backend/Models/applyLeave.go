package models

type ApplyLeave struct {
	Reason    string `json:"reason"`
	Subject   string `json:"subject"`
	LeaveType string `json:"leaveType"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}
