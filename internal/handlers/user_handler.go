package handlers

import (
	"github.com/ikhwankhaleed/morent/internal/services"
	"github.com/ikhwankhaleed/morent/internal/utils"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) UserHandler {
	return UserHandler{
		service: service,
	}
}

func (h *UserHandler) RegisterUser(c echo.Context) error {

	ctx := c.Request().Context()

	type request struct {
		Name           string `json:"name"`
		Email          string `json:"email"`
		PasswordHash   string `json:"password_hash"`
		PhoneNumber    string `json:"phone_number"`
		ProfilePicture string `json:"profile_picture"`
	}

	req := request{}

	if err := c.Bind(&req); err != nil {
		log.Printf("[User Handler] error bind request with error : %v", err)
		response := utils.CreateResponse(err.Error(), http.StatusBadRequest, nil, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	inputUserData := utils.UserDTO{
		Name:           req.Name,
		Email:          req.Email,
		PasswordHash:   req.PasswordHash,
		PhoneNumber:    req.PhoneNumber,
		ProfilePicture: req.ProfilePicture,
	}

	err := h.service.RegisterUser(ctx, inputUserData)
	if err != nil {
		log.Printf("[User Handler] error register user with error : %v", err)
		response := utils.CreateResponse(err.Error(), http.StatusInternalServerError, nil, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := utils.CreateResponse("SUCCESS", http.StatusOK, nil, nil)
	return c.JSON(http.StatusOK, response)
}
