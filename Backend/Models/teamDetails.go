package models

type TeamDetails struct {
	Manager      string       `json:"manager"`
	BusinessUnit string       `json:"businessUnit"`
	TeamMembers  []TeamMember `json:"teamMembers"`
}
