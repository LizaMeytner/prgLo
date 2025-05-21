package grpc

import (
	"context"
	"threeR/api/proto/user"
	"threeR/internal/user/service"
)

type Server struct {
	user.UnimplementedUserServiceServer // <- Наследуем сгенерированный интерфейс
	service                             service.UserService
}

func NewServer(s service.UserService) *Server {
	return &Server{service: s}
}

func (s *Server) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.UserResponse, error) {
	// 1. Получаем данные из service
	u, err := s.service.GetUserByID(int(req.UserId))
	if err != nil {
		return nil, err
	}

	// 2. Преобразуем в gRPC-формат
	return &user.UserResponse{
		Id:       int32(u.ID),
		Username: u.Username,
		Email:    u.Email,
	}, nil
}
