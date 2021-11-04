package domain

type AccountRole string

const (
	AccountRoleUser  AccountRole = "USER"
	AccountRoleAdmin AccountRole = "ADMIN"
)

type Account struct {
	Login        string
	PasswordHash string
	UserId       string
	Role         AccountRole
}

func CreateAccount(login, password, userId string, role AccountRole, encoder PasswordEncoder) *Account {
	return &Account{
		Login:        login,
		PasswordHash: encoder.Encode(password),
		UserId:       userId,
		Role:         role,
	}
}
