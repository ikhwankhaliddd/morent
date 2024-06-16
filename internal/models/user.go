package models

import "time"

type User struct {
	Name           string    `db:"name"`
	Email          string    `db:"email"`
	PasswordHash   string    `db:"password_hash"`
	PhoneNumber    string    `db:"phone_number"`
	ProfilePicture string    `db:"profile_picture"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"update_at"`
}
