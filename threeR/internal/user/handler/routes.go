package handler

import (
	"threeR/internal/user/service"
	"threeR/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, userService service.UserService) {
	handler := NewUserHandler(userService)

	authGroup := router.Group("/api/v1")
	{
		authGroup.POST("/register", handler.Register)
		authGroup.POST("/login", handler.Login)
	}

	authProtected := router.Group("/api/v1")
	authProtected.Use(middleware.AuthMiddleware())
	{
		authProtected.GET("/profile", handler.GetProfile)
	}
}
