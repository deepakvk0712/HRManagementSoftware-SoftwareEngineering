package gormModels

import "time"

type LoginUser struct {
	Email     string    `json:"email";gorm:"type:varchar(32)"`
	Password  string    `json:"password";gorm:"type:varchar(32)"`
	CreatedTS time.Time `json:"createdTS";gorm:"type:timestamp"`
	UpdatedTS time.Time `json:"updatedTS";gorm:"type:timestamp"`
	Role      byte      `json:"role";gorm:"type:blob"`
}
