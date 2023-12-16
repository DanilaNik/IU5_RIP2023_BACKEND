package httpmodels

import (
	"github.com/DanilaNik/IU5_RIP2023/internal/service/role"
)

type TestingRegisterRequest struct {
	ID       uint64    `json:"id"`
	UserName string    `json:"userName"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Role     role.Role `json:"role"`
	ImageURL string    `json:"image_url"`
}

type TestingRegisterResponse struct {
	ID       uint64    `json:"id"`
	UserName string    `json:"userName"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Role     role.Role `json:"role"`
	ImageURL string    `json:"image_url"`
}

type TestingLoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type TestingLoginResponse struct {
	Token string `json:"token"`
}
