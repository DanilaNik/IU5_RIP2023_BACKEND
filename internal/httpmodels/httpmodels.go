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

type TestingGetItemsRequest struct {
	SearchText string `json:"SearchText"`
}

type TestingGetItemsResponse struct {
	Items []*Item `json:"items"`
}

type TestingGetItemByIDRequest struct {
	ID int64 `json:"id"`
}

type TestingGetItemByIDResponse struct {
	Item Item `json:"item"`
}

type TestingDeleteItemRequest struct {
	ID int64 `json:"id"`
}

type TestingPostItemRequest struct {
	Item Item `json:"item"`
}

type TestingPostItemResponse struct {
	Item Item `json:"item"`
}

type TestingPutItemRequset struct {
	Item Item `json:"item"`
}

type TestingGetUserByIDRequest struct {
	ID int64 `json:"id"`
}

type TestingGetUserByIDResponse struct {
	User User `json:"user"`
}
