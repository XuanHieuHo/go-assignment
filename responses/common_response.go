package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

type ErrApi struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func OK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, CommonResponse{
		Code: http.StatusOK,
		Data: data,
	})
}

func Accepted(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusAccepted, CommonResponse{
		Code: http.StatusAccepted,
		Data: data,
	})
}

func BadRequest(message string, errors interface{}) *ErrApi {
	return &ErrApi{
		Code:    http.StatusBadRequest,
		Message: message,
		Errors:  errors,
	}
}

func Unauthorized(message string) *ErrApi {
	return &ErrApi{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

func (e *ErrApi) Error() string {
	return e.Message
}

func Err(ctx *gin.Context, code int, message string, errors interface{}) *ErrApi {
	return &ErrApi{
		Code:    code,
		Message: message,
		Errors:  errors,
	}
}
