package main

import (
	"github.com/LizaMeytner/prgLo/forum-service/internal/core"
	"github.com/LizaMeytner/prgLo/forum-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	forumCore := core.NewForumCore()
	r := gin.Default()
	handlers.SetupForumRoutes(r, forumCore)
	r.Run(":8081")
}
