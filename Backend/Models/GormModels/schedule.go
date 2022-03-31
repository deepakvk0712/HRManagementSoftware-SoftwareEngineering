package gormModels

type Schedule struct {
	ID              int    `json:"ID";gorm:"primaryKey"`
	EmployeeID      int    `json:"EmployeeID";gorm:"type:integer"`
	ScheduleTitle   string `json:"ScheduleTitle";gorm:"type:varchar(32)"`
	ScheduleContent string `json:"ScheduleContent";gorm:"type:varchar(256)"`
}
