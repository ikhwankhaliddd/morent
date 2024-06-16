package repositories

import (
	"context"
	"github.com/ikhwankhaleed/morent/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func (ur *UserRepository) InsertUser(ctx context.Context, user models.User) error {
	return nil
}
