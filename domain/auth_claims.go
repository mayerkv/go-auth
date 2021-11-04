package domain

import "github.com/golang-jwt/jwt/v4"

type AuthClaims struct {
	jwt.RegisteredClaims
	UserId string      `json:"uid"`
	Login  string      `json:"lgn"`
	Type   TokenType   `json:"ttp"`
	Role   AccountRole `json:"role"`
}

type TokenType string

const (
	TokenTypeAccess  TokenType = "access"
	TokenTypeRefresh TokenType = "refresh"
)
