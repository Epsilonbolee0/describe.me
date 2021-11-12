package transport

import (
	"describe.me/internal/objects/entity"
)

// UserRegister -.

type UserRegisterRequest struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserRegisterResponse struct{}

// UserLogin -.

type UserLoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type UserLoginResponse struct {
	*entity.User
}
