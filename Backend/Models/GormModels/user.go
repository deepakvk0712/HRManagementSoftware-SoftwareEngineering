package gormModels

import "time"

type User struct {
	OfficialEmail string `json:"OfficialEmail"`
	EmployeeID    int64  `json:"EmployeeID";gorm:"primaryKey"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	BusinessUnit  string `json:"businessUnit"`
	ManagerId     int    `json:"managerId"`
	Grade         string `json:"grade"`
	Location      string `json:"location"`
	Country       string `json:"country"`
	Title         string `json:"title"`
	Type          string `json:"type"`
	PersonalEmail string `json:"personalEmail"`
	//IDs
	DriversLicense string `json:"DriversLicense"`
	SSN            string `json:"SSN"`
	StateID        string `json:"StateID"`
	//Contact details
	Address         string `json:"Address"`
	Phone           string `json:"Phone"`
	AlternateEmails string `json:"AlternateEmails"`
	//Personal Info
	Gender      string `json:"Gender"`
	DateOfBirth string `json:"DateOfBirth"` /* MM/DD/YYYY */
	Race        string `json:"Race"`
	Ethnicity   string `json:"Ethnicity"`
	Citizenship string `json:"Citizenship"`
	Nationality string `json:"Nationality"`
	Pronouns    string `json:"Pronouns"`
	//Payment Method
	RoutingNumber string    `json:"RoutingNumber"`
	AccountNumber string    `json:"AccountNumber"`
	Bank          string    `json:"Bank"`
	CreatedTS     time.Time `json:"createdTs"`
	UpdatedTS     time.Time `json:"updatedTs"`
}
