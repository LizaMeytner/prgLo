package main

import (
	"net"
	"threeR/api/proto/user"
	"threeR/internal/user/grpc"
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

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, grpc.NewServer(userService))

	// 2. Запуск
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go grpcServer.Serve(lis)

	// Настройка роутера
	router := gin.Default()
	handler.SetupUserRoutes(router, userService)

	// Запуск сервера
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server", "error", err)
	}
}
