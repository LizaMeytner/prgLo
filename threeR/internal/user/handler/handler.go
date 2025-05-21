package handler

import (
	"net/http"

	"threeR/internal/user/service"
	"threeR/pkg/logger"
	"threeR/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userService service.UserService
	validate    *validator.Validate
	log         logger.Logger
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		validate:    validator.New(),
		log:         logger.New("UserHandler"),
	}
}

// RegisterRequest структура для входящих данных регистрации
type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// LoginRequest структура для входящих данных авторизации
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest

	// Парсинг и валидация входящих данных
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Invalid request body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		h.log.Warn("Validation failed", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ValidationErrorsToMap(err)})
		return
	}

	// Создание пользователя через сервис
	user, err := h.userService.CreateUser(req.Username, req.Email, req.Password)
	if err != nil {
		h.log.Error("Failed to create user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Генерация JWT токена
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		h.log.Error("Failed to generate token", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	h.log.Info("User registered successfully", "user_id", user.ID)
	c.JSON(http.StatusCreated, gin.H{
		"user":  user,
		"token": token,
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Invalid request body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		h.log.Warn("Validation failed", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ValidationErrorsToMap(err)})
		return
	}

	// Аутентификация пользователя
	user, err := h.userService.Authenticate(req.Email, req.Password)
	if err != nil {
		h.log.Warn("Authentication failed", "email", req.Email, "error", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Генерация токена
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		h.log.Error("Failed to generate token", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	h.log.Info("User logged in", "user_id", user.ID)
	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	// Получение userID из контекста (должен быть установлен middleware аутентификации)
	userID, exists := c.Get("userID")
	if !exists {
		h.log.Error("User ID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Преобразование userID в int
	uid, ok := userID.(int)
	if !ok {
		h.log.Error("Invalid user ID type in context")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	// Получение данных пользователя
	user, err := h.userService.GetUserByID(uid)
	if err != nil {
		h.log.Error("Failed to get user", "user_id", uid, "error", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	h.log.Debug("Profile retrieved", "user_id", uid)
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
