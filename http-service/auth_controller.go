package http_service

import (
	"github.com/gin-gonic/gin"
	"github.com/mayerkv/go-auth/domain"
	"net/http"
)

type AuthController struct {
	authService *domain.AuthService
}

func NewAuthController(authService *domain.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) SignIn(ctx *gin.Context) {
	var dto domain.SignInDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	accessToken, refreshToken, err := c.authService.SignIn(dto)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (c *AuthController) Refresh(ctx *gin.Context) {
	var dto domain.RefreshDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	accessToken, err := c.authService.Refresh(dto)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}

func (c *AuthController) Profile(ctx *gin.Context) {
	value, exists := ctx.Get("authClaims")
	if !exists {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	claims, ok := value.(*domain.AuthClaims)
	if !ok {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	ctx.Header("x-user-id", claims.UserId)
	ctx.Header("x-user-login", claims.Login)
	ctx.JSON(http.StatusOK, gin.H{
		"user_id": claims.UserId,
		"login":   claims.Login,
	})
}
