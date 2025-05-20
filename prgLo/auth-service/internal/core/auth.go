package core

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthCore struct {
	users map[string]User
}

func NewAuthCore() *AuthCore {
	return &AuthCore{
		users: make(map[string]User),
	}
}

func (c *AuthCore) Register(user User) error {
	if _, exists := c.users[user.Email]; exists {
		return ErrUserExists
	}
	c.users[user.Email] = user
	return nil
}
