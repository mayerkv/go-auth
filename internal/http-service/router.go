package http_service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mayerkv/go-auth/internal/domain"
)

func CreateRouter(
	authController *AuthController,
	service *domain.AuthService,
	jwksController *JWKSController,
	accountController *AccountController,
) *gin.Engine {
	mediaTypeMiddleware := MediaTypeMiddleware()
	authMiddleware := AuthMiddleware(service)

	r := gin.Default()
	r.NoRoute(authMiddleware, authController.Auth)

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/sign-up", mediaTypeMiddleware, accountController.SignUp)
		authGroup.POST("/sign-in", mediaTypeMiddleware, authController.SignIn)
		authGroup.POST("/refresh", mediaTypeMiddleware, authController.Refresh)
		authGroup.GET("/profile", authMiddleware, authController.Profile)
	}

	r.GET(
		"/health", func(context *gin.Context) {
			context.Status(http.StatusOK)
		},
	)

	r.GET("/.well-known/jwks.json", jwksController.Keys)

	return r
}
