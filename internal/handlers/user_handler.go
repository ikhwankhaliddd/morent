package handlers

import (
	"github.com/ikhwankhaleed/morent/internal/services"
	"github.com/ikhwankhaleed/morent/internal/utils"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type UserHandler struct {
	service     *services.UserService
	authService *services.AuthService
}

func NewUserHandler(service *services.UserService, authService *services.AuthService) UserHandler {
	return UserHandler{
		service:     service,
		authService: authService,
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

func (h *UserHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	req := request{}

	if err := c.Bind(&req); err != nil {
		log.Printf("[User Handler] error login user with error : %v", err)
		response := utils.CreateResponse(err.Error(), http.StatusBadRequest, nil, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	result, err := h.service.LoginUser(ctx, req.Email, req.Password)
	if err != nil {
		log.Printf("[User Handler] error login user with error : %v", err)
		response := utils.CreateResponse(err.Error(), http.StatusInternalServerError, nil, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	authToken, err := h.authService.GenerateToken(result.Email)
	if err != nil {
		log.Printf("[User Handler] error login user with error : %v", err)
		response := utils.CreateResponse(err.Error(), http.StatusInternalServerError, nil, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := utils.CreateResponse("SUCCESS", http.StatusOK, authToken, nil)
	return c.JSON(http.StatusOK, response)
}
