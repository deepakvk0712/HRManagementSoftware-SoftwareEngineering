package gormModels

import "time"

type Feedback struct {
	Q1         int       `json:"q1";gorm:"type:integer"`
	Q2         int       `json:"q2";gorm:"type:integer"`
	Q3         int       `json:"q3";gorm:"type:integer"`
	Q4         int       `json:"q4";gorm:"type:integer"`
	Q5         int       `json:"q5";gorm:"type:integer"`
	Feedback   string    `json:"feedback";gorm:"type:varchar(256)"`
	ResignDate string    `json:"resigndate";gorm:"type:varchar(32)"`
	EmployeeID int       `gorm:"type:integer"`
	Name       string    `gorm:"type:varchar(32)"`
	Created_TS time.Time `gorm:"type:timestamp"`
}
