package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/DanilaNik/IU5_RIP2023/internal/ds"
	"github.com/gin-gonic/gin"
)

func (h *Handler) JSONGetUsers(ctx *gin.Context) {
	users, err := h.Repository.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (h *Handler) JSONGetUserById(ctx *gin.Context) {
	jsonData, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	id := struct {
		Id uint64 `json:"id"`
	}{}
	err = json.Unmarshal(jsonData, &id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	user, err := h.Repository.GetUserByID(int(id.Id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (h *Handler) JSONDeleteUser(ctx *gin.Context) {
	jsonData, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	id := struct {
		Id uint64 `json:"id"`
	}{}
	err = json.Unmarshal(jsonData, &id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	err = h.Repository.DeleteUser(int(id.Id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error: ": err.Error})
}

func (h *Handler) CreateUser(ctx *gin.Context) {
	jsonData, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	var user ds.User
	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	err = h.Repository.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error: ": err.Error})
}

func (h *Handler) RegisterUser(user ds.User) (ds.User, error) {
	candidate, err := h.Repository.GetUserByLogin(user.Login)
	if err != nil {
		return ds.User{}, err
	}
	if candidate.Email == user.Email {
		return ds.User{}, errors.New("пользователь сущетсвует")
	}

	err = h.Repository.CreateUser(user)
	if err != nil {
		return ds.User{}, err
	}

	return user, nil
}
