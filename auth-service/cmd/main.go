package main

import (
	"github.com/LizaMeytner/prgLo/auth-service/internal/core"
	"github.com/LizaMeytner/prgLo/auth-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализируем ядро приложения
	authCore := core.NewAuthCore()

	// Настраиваем роутер
	router := gin.Default()

	// Регистрируем обработчики
	handlers.SetupAuthRoutes(router, authCore)

	// Запускаем сервер
	router.Run(":8080")
}
