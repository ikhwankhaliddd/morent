package handlers

import (
	"github.com/ikhwankhaleed/morent/internal/services"
	"net/http"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {

}
