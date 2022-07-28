package http_service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domain2 "github.com/mayerkv/go-auth/internal/domain"
)

type AuthController struct {
	authService *domain2.AuthService
}

func NewAuthController(authService *domain2.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) SignIn(ctx *gin.Context) {
	var dto domain2.SignInDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	accessToken, refreshToken, err := c.authService.SignIn(dto)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(
		http.StatusOK, gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		},
	)
}

func (c *AuthController) Refresh(ctx *gin.Context) {
	var dto domain2.RefreshDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	accessToken, err := c.authService.Refresh(dto)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(
		http.StatusOK, gin.H{
			"access_token": accessToken,
		},
	)
}

func (c *AuthController) Profile(ctx *gin.Context) {
	value, exists := ctx.Get("authClaims")
	if !exists {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	claims, ok := value.(*domain2.AuthClaims)
	if !ok {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	ctx.Header("x-user-id", claims.UserId)
	ctx.Header("x-user-login", claims.Login)
	ctx.JSON(
		http.StatusOK, gin.H{
			"user_id": claims.UserId,
			"login":   claims.Login,
		},
	)
}

func (c *AuthController) Auth(ctx *gin.Context) {
	value, exists := ctx.Get("authClaims")
	if !exists {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	claims, ok := value.(*domain2.AuthClaims)
	if !ok {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	ctx.Header("x-user-id", claims.UserId)
	ctx.Header("x-user-login", claims.Login)
	ctx.Status(http.StatusOK)
}
