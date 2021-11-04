package repository

import (
	"github.com/mayerkv/go-auth/domain"
	"sync"
)

type InMemoryAccountRepository struct {
	sync.Mutex
	items map[string]domain.Account
}

func NewInMemoryAccountRepository() domain.AccountRepository {
	return &InMemoryAccountRepository{
		items: map[string]domain.Account{},
	}
}

func (r *InMemoryAccountRepository) Save(account *domain.Account) error {
	r.Lock()
	defer r.Unlock()

	r.items[account.Login] = *account

	return nil
}

func (r *InMemoryAccountRepository) FindByLogin(login string) (*domain.Account, error) {
	r.Lock()
	defer r.Unlock()

	if account, ok := r.items[login]; ok {
		return &account, nil
	}

	return nil, nil
}

func (r *InMemoryAccountRepository) FindByUserId(id string) (*domain.Account, error) {
	r.Lock()
	defer r.Unlock()

	for _, item := range r.items {
		if item.UserId == id {
			return &item, nil
		}
	}

	return nil, nil
}
