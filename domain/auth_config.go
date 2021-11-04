package domain

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type AuthConfig struct {
	rsaPrivateKey   *rsa.PrivateKey
	method          jwt.SigningMethod
	issuer          string
	accessDuration  time.Duration
	refreshDuration time.Duration
}

func NewAuthConfig(rsaPrivateKey *rsa.PrivateKey, method jwt.SigningMethod, issuer string, accessDuration time.Duration, refreshDuration time.Duration) *AuthConfig {
	return &AuthConfig{rsaPrivateKey: rsaPrivateKey, method: method, issuer: issuer, accessDuration: accessDuration, refreshDuration: refreshDuration}
}
