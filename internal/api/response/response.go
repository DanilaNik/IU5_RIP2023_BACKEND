package response

import (
	"github.com/gin-gonic/gin"
)

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func Error(ctx *gin.Context, statusCode int, err error) {
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
