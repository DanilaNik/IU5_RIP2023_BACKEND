package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/DanilaNik/IU5_RIP2023/internal/httpmodels"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// @Summary Register
// @Description Register a new user with the provided user data
// @Tags users
// @Accept  json
// @Produce  json
// @Param input body httpmodels.TestingRegisterRequest true "User data to register"
// @Success 201 {object} httpmodels.TestingRegisterResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /register [post]
func (h *Handler) SignUp(ctx *gin.Context) {
	var userJSON httpmodels.TestingRegisterRequest
	if err := ctx.ShouldBindJSON(&userJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.AuthorizationService.RegisterUser(userJSON)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

// @Summary Login
// @Description Login with the provided user credentials and receive a JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param input body httpmodels.TestingLoginRequest true "User credentials for login"
// @Success 200 {object} httpmodels.TestingLoginResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /login [post]
func (h *Handler) Login(ctx *gin.Context) {
	var userJSON httpmodels.TestingLoginRequest

	if err := ctx.ShouldBindJSON(&userJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.AuthorizationService.LoginUser(userJSON)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("auth", token.Token, 3600*24*30, "", "", false, true)

	ctx.JSON(http.StatusCreated, gin.H{"token": token})
}

func (h *Handler) getUserRole(ctx *gin.Context) (string, string, error) {
	cookie, err := ctx.Cookie("auth")
	if err != nil {
		return "", "", err
	}
	token, err := jwt.Parse(cookie, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неверный метод подписи")
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	id := ""
	role := ""
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id = strconv.Itoa(int(claims["userID"].(float64)))
		role = claims["role"].(string)
	} else {
		return "", "", err
	}

	return id, role, nil
}

func (h *Handler) Logout(ctx *gin.Context) {
	id, _, err := h.getUserRole(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.AuthorizationService.LogoutUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "cant delete from logout",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) GetUsers(ctx *gin.Context) {
	users, err := h.Repository.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (h *Handler) GetUserById(ctx *gin.Context) {
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

func (h *Handler) DeleteUser(ctx *gin.Context) {
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

func (h *Handler) GetUserRequests(ctx *gin.Context) {
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

	requests, err := h.Repository.GetUserRequests(int(id.Id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}
	ctx.JSON(http.StatusOK, requests)
}

// @Summary Protected test endpoint
// @Description Test endpoint accessible only with valid Bearer Token
// @Security ApiKeyAuth
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Router /protected/test [get]
func (h *Handler) ProtectedTest(ctx *gin.Context) {
	userID := ctx.MustGet("UserID").(int)

	ctx.JSON(http.StatusOK, gin.H{"message": "user is authorized admin and moderator", "userID": userID})
}
