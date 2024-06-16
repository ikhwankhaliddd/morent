package services

import (
	"context"
	"github.com/ikhwankhaleed/morent/internal/repositories"
	"github.com/ikhwankhaleed/morent/internal/utils"
)

type UserService struct {
	repo repositories.UserRepository
}

func RegisterUser(ctx context.Context, data utils.UserDTO) error {
	return nil
}
