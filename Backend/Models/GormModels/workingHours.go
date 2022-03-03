package gormModels

import "time"

type WorkingHours struct {
	EmployeeID int64
	StartTime  time.Time
	EndTime    time.Time
	TimeWorked float64
}
