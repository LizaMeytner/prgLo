package main

import (
	"github.com/LizaMeytner/prgLo/auth-service/internal/core"
	"github.com/LizaMeytner/prgLo/auth-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация чистого ядра (без БД)
	authCore := core.NewAuthCore()

	// Настройка HTTP-роутера
	r := gin.Default()
	handlers.SetupAuthRoutes(r, authCore)

	// Запуск сервера
	r.Run(":8080")
}
