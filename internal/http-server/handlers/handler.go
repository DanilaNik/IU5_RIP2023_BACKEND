package handlers

import (
	"github.com/DanilaNik/IU5_RIP2023/internal/minio"
	"github.com/DanilaNik/IU5_RIP2023/internal/repository"
	auth "github.com/DanilaNik/IU5_RIP2023/internal/service/authorization"
	itemservice "github.com/DanilaNik/IU5_RIP2023/internal/service/itemService"
	requestitemservice "github.com/DanilaNik/IU5_RIP2023/internal/service/requestItemService"
	requestservice "github.com/DanilaNik/IU5_RIP2023/internal/service/requestService"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Logger               *logrus.Logger
	Repository           *repository.Repository
	AuthorizationService *auth.AuthorizationService
	ItemService          *itemservice.ItemService
	RequestService       *requestservice.RequestService
	RequestItemService   *requestitemservice.RequestItemService
	Minio                *minio.MinioClient
}

func NewHandler(log *logrus.Logger, r *repository.Repository, a *auth.AuthorizationService, i *itemservice.ItemService, req *requestservice.RequestService, reqItem *requestitemservice.RequestItemService, m *minio.MinioClient) *Handler {
	return &Handler{
		Logger:               log,
		Repository:           r,
		AuthorizationService: a,
		ItemService:          i,
		RequestService:       req,
		RequestItemService:   reqItem,
		Minio:                m,
	}
}

const (
	StatusError = "Error"
)

func (h *Handler) Error(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{
		"Status": StatusError,
		"Error":  err.Error(),
	})
}

//func ValidationError(errs validator.ValidationErrors) Response {
//	var errMsgs []string
//
//	for _, err := range errs {
//		switch err.ActualTag() {
//		case "required":
//			errMsgs = append(errMsgs, fmt.Sprintf("field %s is a required field", err.Field()))
//		default:
//			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not valid", err.Field()))
//		}
//	}
//
//	return Response{
//		Status: StatusError,
//		Error:  strings.Join(errMsgs, ", "),
//	}
//}
