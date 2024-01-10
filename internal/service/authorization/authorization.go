package authorization

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

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

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return httpmodels.TestingRegisterResponse{}, errors.New("failed to hash password")
	}

	err = a.Repository.CreateUser(ds.User{
		ID:       user.ID,
		UserName: user.UserName,
		Login:    user.Login,
		Password: string(hash),
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

	err = bcrypt.CompareHashAndPassword([]byte(candidate.Password), []byte(user.Password))
	if err != nil {
		return httpmodels.TestingLoginResponse{Token: ""}, errors.New("неправильный пароль или логин")
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

func (a *AuthorizationService) LogoutUser(id string) error {
	err := a.Repository.DeleteJWTToken(id)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthorizationService) IsLogout(id string) error {
	err := a.Repository.GetJWTToken(id)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthorizationService) GetUserByID(id int) (*httpmodels.TestingGetUserByIDResponse, error) {
	user, err := a.Repository.GetUserByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "get user by id")
	}

	resp := &httpmodels.TestingGetUserByIDResponse{
		User: httpmodels.User{
			ID:       user.ID,
			Login:    user.Login,
			Email:    user.Email,
			UserName: user.UserName,
			Role:     user.Role,
		},
	}
	return resp, nil
}
