package http_service

import (
	"github.com/gin-gonic/gin"
	"github.com/mayerkv/go-auth/domain"
	"net/http"
)

func CreateRouter(authController *AuthController, service *domain.AuthService, jwksController *JWKSController) *gin.Engine {
	r := gin.Default()

	mediaTypeMiddleware := MediaTypeMiddleware()
	authMiddleware := AuthMiddleware(service)

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/sign-in", mediaTypeMiddleware, authController.SignIn)
		authGroup.POST("/refresh", mediaTypeMiddleware, authController.Refresh)
		authGroup.GET("/profile", authMiddleware, authController.Profile)
	}

	r.GET("/health", func(context *gin.Context) {
		context.Status(http.StatusOK)
	})

	r.GET(".well-known/jwks.json", jwksController.Keys)

	return r
}
