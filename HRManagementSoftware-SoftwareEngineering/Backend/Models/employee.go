package models

type Employee struct {
	//IDs
	DriversLicense string `json:"DriversLicense"`
	SSN            string `json:"SSN"`
	StateID        string `json:"StateID"`
	//Contact details
	Address         string `json:"Address"`
	Phone           string `json:"Phone"`
	AlternateEmails string `json:"AlternateEmails"`
	//Personal Info
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	Gender      string `json:"Gender"`
	DateOfBirth string `json:"DateOfBirth"` /* MM/DD/YYYY */
	Race        string `json:"Race"`
	Ethnicity   string `json:"Ethnicity"`
	Citizenship string `json:"Citizenship"`
	Nationality string `json:"Nationality"`
	Pronouns    string `json:"Pronouns"`
	//Payment Method
	RoutingNumber string `json:"RoutingNumber"`
	AccountNumber string `json:"AccountNumber"`
	Bank          string `json:"Bank"`
}
