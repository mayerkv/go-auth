package domain

type AccountRepository interface {
	Save(account *Account) error
	FindByLogin(login string) (*Account, error)
	FindByUserId(id string) (*Account, error)
}
