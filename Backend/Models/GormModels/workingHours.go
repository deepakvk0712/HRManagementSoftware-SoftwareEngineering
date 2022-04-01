package gormModels

import (
	"time"
)

type WorkingHours struct {
	EmployeeID int       `json:"EmployeeID"`
	StartTime  time.Time `json:"StartTime";gorm:"type:timestamp"`
	EndTime    time.Time `json:"EndTime";gorm:"type:timestamp"`
	TimeWorked float64   `json:"TimeWorked"`
}
