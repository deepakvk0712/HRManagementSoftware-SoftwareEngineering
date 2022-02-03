package models

type Employee struct {
	//IDs
	ValidID string `json:"ValidID"`
	//Contact details
	Address string `json:"Address"`
	Phone   string `json:"Phone"`
	Email   string `json:"Email"`
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
	RoutingNumber int    `json:"RoutingNumber"`
	AccountNumber int    `json:"AccountNumber"`
	Bank          string `json:"Bank"`
}
