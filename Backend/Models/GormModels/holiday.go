package gormModels

import "time"

type Holiday struct {
	ID      int       `gorm:"primaryKey"`
	Date    time.Time `json:"date";gorm:"type:timestamp"`
	Comment string    `json:"comment";gorm:"type:varchar(64)"`
}
