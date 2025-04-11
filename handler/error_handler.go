package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/XuanHieuHo/go-assignment/responses"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(handler func(ctx *gin.Context) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := handler(ctx); err != nil {
			var errApi *responses.ErrApi
			if errors.As(err, &errApi) {
				ctx.JSON(errApi.Code, responses.CommonResponse{
					Code:    errApi.Code,
					Message: errApi.Message,
					Errors:  errApi.Errors,
				})
				return
			}
			
			ctx.JSON(http.StatusInternalServerError, responses.CommonResponse{
				Code:    http.StatusInternalServerError,
				Message: "Internal Server Error",
			})
			log.Printf("%+v", err)
			return
		}
	}
}
