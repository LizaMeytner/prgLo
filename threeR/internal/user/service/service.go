package service

import (
	"errors"
	"threeR/internal/user/repository"
	"threeR/pkg/security"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(username, email, password string) (*repository.User, error) {
	// Хеширование пароля
	hashedPassword, err := security.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &repository.User{
		Username:     username,
		Email:        email,
		PasswordHash: hashedPassword,
	}

	return s.repo.Create(user)
}

func (s *UserService) Authenticate(email, password string) (*repository.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !security.CheckPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *UserService) GetUserByID(id int) (*repository.User, error) {
	return s.repo.FindByID(id)
}
