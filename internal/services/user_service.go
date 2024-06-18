package services

import (
	"context"
	"github.com/ikhwankhaleed/morent/internal/models"
	"github.com/ikhwankhaleed/morent/internal/repositories"
	"github.com/ikhwankhaleed/morent/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) RegisterUser(ctx context.Context, data utils.UserDTO) error {
	createdAt := time.Now()

	phoneNumber := utils.StringToNullString(data.PhoneNumber)
	profilePicture := utils.StringToNullString(data.ProfilePicture)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.PasswordHash), bcrypt.MinCost)
	if err != nil {
		log.Printf("[User Service] Error Generate Hash Password with error :%v", err)
		return err
	}

	userData := models.User{
		Name:           data.Name,
		Email:          data.Email,
		PasswordHash:   string(hashedPassword),
		PhoneNumber:    phoneNumber,
		ProfilePicture: profilePicture,
		CreatedAt:      createdAt,
	}
	err = us.repo.InsertUser(ctx, userData)
	if err != nil {
		log.Printf("[User Service] Error Register User with error : %v", err)
		return err
	}
	return nil
}

func (us *UserService) LoginUser(ctx context.Context, email, password string) (result utils.UserDTO, err error) {
	userData, err := us.repo.GetUser(ctx, email)
	if err != nil {
		log.Printf("[User Service] Error Login User with error : %v", err)
		return result, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.PasswordHash), []byte(password))
	if err != nil {
		log.Printf("[User Service] Error Login User with error : %v", err)
		return result, err
	}
	phoneNumber := utils.NullStringToString(userData.PhoneNumber)
	profilePicture := utils.NullStringToString(userData.ProfilePicture)
	updatedAt := utils.NullTimeToTime(userData.UpdatedAt)

	result = utils.UserDTO{
		Name:           userData.Name,
		Email:          userData.Email,
		PasswordHash:   userData.PasswordHash,
		PhoneNumber:    profilePicture,
		ProfilePicture: phoneNumber,
		CreatedAt:      userData.CreatedAt,
		UpdatedAt:      updatedAt,
	}
	return result, nil
}

func (us *UserService) GetUserByEmail(ctx context.Context, email string) (result utils.UserDTO, err error) {
	userData, err := us.repo.GetUser(ctx, email)
	if err != nil {
		log.Printf("[User Service] Error GetUserByEmail with error : %v", err)
		return result, err
	}

	phoneNumber := utils.NullStringToString(userData.PhoneNumber)
	profilePicture := utils.NullStringToString(userData.ProfilePicture)
	updatedAt := utils.NullTimeToTime(userData.UpdatedAt)

	result = utils.UserDTO{
		Name:           userData.Name,
		Email:          userData.Email,
		PasswordHash:   userData.PasswordHash,
		PhoneNumber:    profilePicture,
		ProfilePicture: phoneNumber,
		CreatedAt:      userData.CreatedAt,
		UpdatedAt:      updatedAt,
	}

	return result, nil
}
