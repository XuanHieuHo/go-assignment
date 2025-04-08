package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/XuanHieuHo/go-assignment/responses"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) == 0 {
			return
		}

		err := ctx.Errors.Last().Err
		statusCode := inferStatusCode(err)

		responses.NewResponseBuilder().
			WithCode(statusCode).
			WithErrors(err).
			RespondWithJSON(ctx)
		ctx.Abort()
	}
}

func inferStatusCode(err error) int {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return http.StatusNotFound
	case containsAny(err.Error(), "binding", "validation", "unmarshal", "BR"):
		return http.StatusBadRequest
	case err.Error() == "user doesn't belong to the authenticated user":
		return http.StatusUnauthorized
	default:
		if pqErr, ok := err.(*pgconn.PgError); ok {
			return parsePGError(string(pqErr.Code))
		}
		return http.StatusInternalServerError
	}
}

func parsePGError(code string) int {
	switch code {
	case "23505":
		return http.StatusConflict
	case "23503":
		return http.StatusForbidden
	case "23502":
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func containsAny(s string, subs ...string) bool {
	for _, sub := range subs {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}
