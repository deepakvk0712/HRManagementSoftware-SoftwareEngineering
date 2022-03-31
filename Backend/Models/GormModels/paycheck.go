package gormModels

import "time"

type Paycheck struct {
	EmployeeID   int       `json:"employeeID";gorm:"type:integer"`
	CheckNumber  int       `json:"checkNumber";gorm:"type:integer"`
	CheckDate    time.Time `json:"checkDate";gorm:"type:timestamp"`
	PayBeginDate time.Time `json:"payBeginDate";gorm:"type:timestamp"`
	PayEndDate   time.Time `json:"payEndDate";gorm:"type:timestamp"`
	AmountPaid   float32   `json:"amountPaid";gorm:"type:real"`
}
