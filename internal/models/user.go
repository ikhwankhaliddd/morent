package models

import (
	"database/sql"
	"time"
)

type User struct {
	Name           string         `db:"name"`
	Email          string         `db:"email"`
	PasswordHash   string         `db:"password_hash"`
	PhoneNumber    sql.NullString `db:"phone_number"`
	ProfilePicture sql.NullString `db:"profile_picture"`
	CreatedAt      time.Time      `db:"created_at"`
	UpdatedAt      sql.NullTime   `db:"update_at"`
}
