package http_service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/jwk"
)

type JWKSController struct {
	jwkKey jwk.Key
}

func NewJWKSController(jwkKey jwk.Key) *JWKSController {
	return &JWKSController{jwkKey: jwkKey}
}

func (c *JWKSController) Keys(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK, gin.H{
			"keys": []jwk.Key{
				c.jwkKey,
			},
		},
	)
}
