package authorization

import (
	"github.com/pkg/errors"

	"github.com/DanilaNik/IU5_RIP2023/internal/ds"
	"github.com/DanilaNik/IU5_RIP2023/internal/httpmodels"
	"github.com/DanilaNik/IU5_RIP2023/internal/repository"
	"github.com/DanilaNik/IU5_RIP2023/pkg"
)

type AuthorizationService struct {
	Repository *repository.Repository
}

func NewAuthorizationService(repo *repository.Repository) *AuthorizationService {
	return &AuthorizationService{
		Repository: repo,
	}
}

func (a *AuthorizationService) RegisterUser(user httpmodels.TestingRegisterRequest) (httpmodels.TestingRegisterResponse, error) {
	candidate, err := a.Repository.GetUserByLogin(user.Login)
	if err != nil {
		return httpmodels.TestingRegisterResponse{}, err
	}
	if candidate.Email == user.Email {
		return httpmodels.TestingRegisterResponse{}, errors.New("пользователь уже сущетсвует")
	}

	err = a.Repository.CreateUser(ds.User{
		ID:       user.ID,
		UserName: user.UserName,
		Login:    user.Login,
		Password: user.Password,
		Email:    user.Email,
		Role:     user.Role,
		ImageURL: user.ImageURL,
	})
	if err != nil {
		return httpmodels.TestingRegisterResponse{}, errors.Wrapf(err, "не удалось создать пльзователя")
	}

	return httpmodels.TestingRegisterResponse{
		ID:       user.ID,
		UserName: user.UserName,
		Login:    user.Login,
		Password: user.Password,
		Email:    user.Email,
		Role:     user.Role,
		ImageURL: user.ImageURL,
	}, nil
}

func (a *AuthorizationService) LoginUser(user httpmodels.TestingLoginRequest) (httpmodels.TestingLoginResponse, error) {
	candidate, err := a.Repository.GetUserByLogin(user.Login)
	if err != nil {
		return httpmodels.TestingLoginResponse{Token: ""}, err
	}

	if candidate.Password != user.Password {
		return httpmodels.TestingLoginResponse{Token: ""}, errors.New("неверный пароль")
	}

	token, err := pkg.GenerateJWTToken(uint(candidate.ID), candidate.Role)

	if err != nil {
		return httpmodels.TestingLoginResponse{Token: ""}, err
	}

	err = a.Repository.SaveJWTToken(uint(candidate.ID), token)
	if err != nil {
		return httpmodels.TestingLoginResponse{Token: ""}, err
	}

	return httpmodels.TestingLoginResponse{Token: token}, nil
}
