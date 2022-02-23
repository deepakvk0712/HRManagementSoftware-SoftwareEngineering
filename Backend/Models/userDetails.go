package models

type UserDetails struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     byte   `json:"role"`
}
