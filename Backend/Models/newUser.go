package models

type User struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	BusinessUnit string `json:"businessUnit"`
	ManagerId    int    `json:"managerId"`
	Grade        string `json:"grade"`
	Location     string `json:"location"`
	Title        string `json:"title"`
	Type         string `json:"type"`
	Email        string `json:"email"`
}
