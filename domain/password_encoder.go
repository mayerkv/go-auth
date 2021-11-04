package domain

type PasswordEncoder interface {
	Encode(password string) string
	Compare(password string, passwordHash string) bool
}
