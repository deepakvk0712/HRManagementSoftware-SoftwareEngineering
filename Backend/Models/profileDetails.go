package models

type ProfileDetails struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Title     string `json:"title"`
	AboutMe   string `json:"aboutMe"`
}
