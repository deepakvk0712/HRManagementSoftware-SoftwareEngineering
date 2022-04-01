package models

import "time"

type PaycheckQuery struct {
	PayBeginDate time.Time `json:"payBeginDate"`
	PayEndDate   time.Time `json:"payEndDate"`
}
