package models

import "time"

type WorkingDetails struct {
	StartDate         time.Time `json:"startDate"`
	EndDate           time.Time `json:"endDate"`
	Totaldays         int       `json:"totaldays"`
	Holidays          int       `json:"holidays"`
	PresentDays       int       `json:"presentDays"`
	AbsentDays        int       `json:"absentDays"`
	AverageWorkHour   float64   `json:"averageWorkHour"`
	TotalWorkHour     float64   `json:"totalWorkHour"`
	TotalRegularHour  float64   `json:"totalRegularHour"`
	TotalOvertimeHour float64   `json:"totalOvertimeHour"`
}
