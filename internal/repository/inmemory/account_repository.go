package inmemory

import (
	"sync"

	domain2 "github.com/mayerkv/go-auth/internal/domain"
)

type AccountRepository struct {
	sync.Mutex
	items map[string]domain2.Account
}

func NewAccountRepository() domain2.AccountRepository {
	return &AccountRepository{
		items: map[string]domain2.Account{},
	}
}

func (r *AccountRepository) Save(account *domain2.Account) error {
	r.Lock()
	defer r.Unlock()

	r.items[account.Login] = *account

	return nil
}

func (r *AccountRepository) FindByLogin(login string) (*domain2.Account, error) {
	r.Lock()
	defer r.Unlock()

	if account, ok := r.items[login]; ok {
		return &account, nil
	}

	return nil, nil
}

func (r *AccountRepository) FindByUserId(id string) (*domain2.Account, error) {
	r.Lock()
	defer r.Unlock()

	for _, item := range r.items {
		if item.UserId == id {
			return &item, nil
		}
	}

	return nil, nil
}
