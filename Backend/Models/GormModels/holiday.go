package gormModels

import "time"

type Holiday struct {
	ID      int       `gorm:"primaryKey"`
	Date    time.Time `gorm:"type:timestamp"`
	Comment string    `gorm:"type:varchar(64)"`
}
