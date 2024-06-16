package utils

import "time"

type UserDTO struct {
	Name           string
	Email          string
	PasswordHash   string
	PhoneNumber    string
	ProfilePicture string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
