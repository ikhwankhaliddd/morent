package repositories

import (
	"context"
	"github.com/ikhwankhaleed/morent/internal/models"
	"github.com/jmoiron/sqlx"
	"log"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

const (
	registerNewUser = "INSERT INTO users (name, email, password_hash, phone_number, profile_picture, created_at) VALUES ($1, $2, $3, $4, $5,$6)"
	getUserData     = "SELECT name, email, password_hash, profile_picture, created_at, updated_at FROM users WHERE email = $1"
)

func (ur *UserRepository) InsertUser(ctx context.Context, user models.User) error {
	_, err := ur.db.ExecContext(ctx, registerNewUser,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.PhoneNumber,
		user.ProfilePicture,
		user.CreatedAt)
	if err != nil {
		log.Printf("[User Repository] failed insert user data with error: %v", err)
		return err
	}
	return nil
}

func (ur *UserRepository) GetUser(ctx context.Context, email string) (models.User, error) {
	userData := models.User{}

	err := ur.db.GetContext(ctx, &userData, getUserData, email)
	if err != nil {
		log.Printf("[User Repository] failed get user data with error : %v", err)
		return userData, err
	}
	return userData, nil
}
