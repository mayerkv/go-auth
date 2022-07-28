package domain

import (
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthConfig struct {
	RsaPrivateKey   *rsa.PrivateKey
	Method          jwt.SigningMethod
	Issuer          string
	AccessDuration  time.Duration
	RefreshDuration time.Duration
}

func NewAuthConfig(
	rsaPrivateKey *rsa.PrivateKey,
	method jwt.SigningMethod,
	issuer string,
	accessDuration time.Duration,
	refreshDuration time.Duration,
) *AuthConfig {
	return &AuthConfig{RsaPrivateKey: rsaPrivateKey, Method: method, Issuer: issuer, AccessDuration: accessDuration, RefreshDuration: refreshDuration}
}
