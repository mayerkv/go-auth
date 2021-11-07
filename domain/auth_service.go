package domain

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

var (
	ErrAccountNotExists = errors.New("account not exists")
	ErrInvalidPassword  = errors.New("invalid password")
	ErrInvalidToken     = errors.New("invalid token")
	ErrInvalidTokenType = errors.New("invalid token type")
)

type AuthService struct {
	accountRepository AccountRepository
	passwordEncoder   PasswordEncoder
	authConfig        AuthConfig
}

func NewAuthService(accountRepository AccountRepository, passwordEncoder PasswordEncoder, authConfig AuthConfig) *AuthService {
	return &AuthService{accountRepository: accountRepository, passwordEncoder: passwordEncoder, authConfig: authConfig}
}

func (s *AuthService) SignIn(dto SignInDto) (string, string, error) {
	account, err := s.getAccount(dto.Login)
	if err != nil {
		return "", "", err
	}

	if !s.passwordEncoder.Compare(dto.Password, account.PasswordHash) {
		return "", "", ErrInvalidPassword
	}

	now := time.Now()
	accessToken := s.createToken(now, account, TokenTypeAccess, s.authConfig.AccessDuration)
	refreshToken := s.createToken(now, account, TokenTypeRefresh, s.authConfig.RefreshDuration)

	access, err := accessToken.SignedString(s.authConfig.RsaPrivateKey)
	if err != nil {
		return "", "", err
	}
	refresh, err := refreshToken.SignedString(s.authConfig.RsaPrivateKey)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}

func (s *AuthService) Refresh(dto RefreshDto) (string, error) {
	claims, err := s.Parse(dto.RefreshToken)
	if err != nil {
		return "", err
	}

	if claims.Type != TokenTypeRefresh {
		return "", ErrInvalidTokenType
	}

	account, err := s.getAccount(claims.Login)
	if err != nil {
		return "", err
	}

	accessToken := s.createToken(time.Now(), account, TokenTypeAccess, s.authConfig.AccessDuration)

	return accessToken.SignedString(s.authConfig.RsaPrivateKey)
}

func (s *AuthService) Parse(token string) (*AuthClaims, error) {
	t, err := jwt.ParseWithClaims(token, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return s.authConfig.RsaPrivateKey.Public(), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(*AuthClaims); ok && t.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

func (s *AuthService) getAccount(login string) (*Account, error) {
	account, err := s.accountRepository.FindByLogin(login)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, ErrAccountNotExists
	}

	return account, nil
}

func (s *AuthService) createToken(t time.Time, a *Account, tokenType TokenType, d time.Duration) *jwt.Token {
	return jwt.NewWithClaims(s.authConfig.Method, AuthClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.NewString(),
			IssuedAt:  jwt.NewNumericDate(t),
			ExpiresAt: jwt.NewNumericDate(t.Add(d)),
			Issuer:    s.authConfig.Issuer,
			Subject:   a.Login,
		},
		UserId: a.UserId,
		Login:  a.Login,
		Type:   tokenType,
		Role:   a.Role,
	})
}
