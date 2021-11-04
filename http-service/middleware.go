package http_service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mayerkv/go-auth/domain"
	"net/http"
	"strings"
)

var (
	ErrUnsupportedContentType = errors.New("unsupported content type")
	ErrInvalidAuthHeader      = errors.New("invalid auth header")
)

func MediaTypeMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("content-type")
		if header != "application/json" {
			_ = ctx.AbortWithError(http.StatusUnsupportedMediaType, ErrUnsupportedContentType)
			return
		}

		ctx.Next()
	}
}

func AuthMiddleware(authService *domain.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.GetHeader("authorization")
		if !strings.HasPrefix(tokenHeader, "Bearer ") {
			handleError(ctx, ErrInvalidAuthHeader)
			return
		}

		claims, err := authService.Parse(tokenHeader[7:])
		if err != nil {
			handleError(ctx, err)
			return
		}

		if claims.Type != domain.TokenTypeAccess {
			handleError(ctx, domain.ErrInvalidTokenType)
			return
		}

		ctx.Set("authClaims", claims)
		ctx.Next()
	}
}
