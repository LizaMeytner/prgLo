package handlers

import (
	"github.com/LizaMeytner/prgLo/auth-service/internal/core"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.Engine, core *core.AuthCore) {
	r.POST("/register", func(c *gin.Context) {
		var user core.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}
		
		if err := core.Register(user); err != nil {
			c.JSON(409, gin.H{"error": err.Error()})
			return
		}
		
		c.JSON(200, gin.H{"status": "User created"})
	})