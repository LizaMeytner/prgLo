package handlers

import (
	"github.com/LizaMeytner/prgLo/auth-service/internal/core"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.Engine, authCore *core.AuthCore) {
	r.POST("/register", func(c *gin.Context) {
		var request struct {
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required,min=8"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		if err := authCore.Register(request.Email, request.Password); err != nil {
			c.JSON(409, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"status": "User created"})
	})
}
