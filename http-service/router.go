package http_service

import (
	"github.com/gin-gonic/gin"
	"github.com/mayerkv/go-auth/domain"
	"net/http"
)

func CreateRouter(controller *AuthController, service *domain.AuthService) *gin.Engine {
	r := gin.Default()

	mediaTypeMiddleware := MediaTypeMiddleware()
	authMiddleware := AuthMiddleware(service)

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/sign-in", mediaTypeMiddleware, controller.SignIn)
		authGroup.POST("/refresh", mediaTypeMiddleware, controller.Refresh)
		authGroup.GET("/profile", authMiddleware, controller.Profile)
	}

	r.GET("/health", func(context *gin.Context) {
		context.Status(http.StatusOK)
	})

	return r
}
