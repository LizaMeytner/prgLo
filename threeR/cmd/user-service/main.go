package main

import (
	"threeR/internal/user/handler"
	"threeR/internal/user/repository"
	"threeR/internal/user/service"
	"threeR/pkg/database"
	"threeR/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация логгера
	log := logger.New("UserService")

	// Подключение к базе данных
	db, err := database.NewPostgresConnection("your-connection-string")
	if err != nil {
		log.Fatal("Failed to connect to database", "error", err)
	}

	// Инициализация репозиториев и сервисов
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	// Настройка роутера
	router := gin.Default()
	handler.SetupUserRoutes(router, userService)

	// Запуск сервера
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server", "error", err)
	}
}
