package models

type TeamDetails struct {
	Manager      string   `json:"manager"`
	BusinessUnit string   `json:"businessUnit"`
	TeamMembers  []string `json:"teamMembers"`
}
