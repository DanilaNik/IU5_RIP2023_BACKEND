package handlers

import (
	"github.com/DanilaNik/IU5_RIP2023/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Logger     *logrus.Logger
	Repository *repository.Repository
}

func NewHandler(log *logrus.Logger, r *repository.Repository) *Handler {
	return &Handler{
		Logger:     log,
		Repository: r,
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
