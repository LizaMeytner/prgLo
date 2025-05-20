package core

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserExists = errors.New("user already exists")
)

// User структура должна быть публичной (с большой буквы)
type User struct {
	Email    string `json:"email"`
	Password string `json:"-"` // Исключаем пароль из JSON
}

// AuthCore - основная структура сервиса
type AuthCore struct {
	users map[string]User
}

// NewAuthCore - конструктор (должен быть публичным)
func NewAuthCore() *AuthCore {
	return &AuthCore{
		users: make(map[string]User),
	}
}

func (c *AuthCore) Register(email, password string) error {
	if _, exists := c.users[email]; exists {
		return ErrUserExists
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	c.users[email] = User{
		Email:    email,
		Password: string(hashedPass),
	}
	return nil
}
