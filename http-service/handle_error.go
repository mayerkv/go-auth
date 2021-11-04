package http_service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mayerkv/go-auth/domain"
	"io"
	"net/http"
)

func handleError(ctx *gin.Context, err error) {
	message := err.Error()
	code := http.StatusInternalServerError

	switch err {
	case io.EOF:
		code = http.StatusBadRequest
		message = "bad request"
	case ErrInvalidAuthHeader, domain.ErrInvalidPassword, domain.ErrAccountNotExists, domain.ErrInvalidToken, domain.ErrInvalidTokenType:
		message = err.Error()
		code = http.StatusUnauthorized
	case ErrUnsupportedContentType:
		message = err.Error()
		code = http.StatusUnsupportedMediaType
	}

	switch err.(type) {
	case *validator.ValidationErrors:
		message = err.Error()
		code = http.StatusBadRequest
	case *jwt.ValidationError:
		message = err.Error()
		code = http.StatusUnauthorized
	}

	ctx.JSON(code, gin.H{
		"error": message,
		"code":  code,
	})
}
