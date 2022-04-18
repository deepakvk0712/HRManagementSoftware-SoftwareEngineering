package gormModels

import "time"

type Feedback struct {
	EmployeeID int       `gorm:"type:int"`
	Name       string    `gorm:"type:varchar(32)"`
	Feedback   string    `json:"feedback";gorm:"type:varchar(256)"`
	Created_TS time.Time `gorm:"type:timestamp"`
}
