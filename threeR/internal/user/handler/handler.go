package handler

import (
	"threeR/internal/user/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(c *gin.Context) {
	// Регистрация пользователя
}

func (h *UserHandler) Login(c *gin.Context) {
	// Аутентификация
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	// Получение профиля
}
