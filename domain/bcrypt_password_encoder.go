package domain

import "golang.org/x/crypto/bcrypt"

type BCryptPasswordEncoder struct {
}

func NewBCryptPasswordEncoder() PasswordEncoder {
	return &BCryptPasswordEncoder{}
}

func (e *BCryptPasswordEncoder) Encode(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hash)
}

func (e *BCryptPasswordEncoder) Compare(password string, passwordHash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)) == nil
}
