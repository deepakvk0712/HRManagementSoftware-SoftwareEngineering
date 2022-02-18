package gormModels

import "time"

type LoginUser struct {
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedTS time.Time `json:"createdTS"`
	UpdatedTS time.Time `json:"updatedTS"`
}
