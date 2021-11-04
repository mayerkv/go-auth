package domain

import "errors"

var ErrAccountAlreadyExists = errors.New("account already exists")

type AccountService struct {
	passwordEncoder   PasswordEncoder
	accountRepository AccountRepository
}

func NewAccountService(passwordEncoder PasswordEncoder, accountRepository AccountRepository) *AccountService {
	return &AccountService{passwordEncoder: passwordEncoder, accountRepository: accountRepository}
}

func (s *AccountService) CreateAccount(login, password, userId string, role AccountRole) error {
	account, err := s.accountRepository.FindByLogin(login)
	if err != nil {
		return err
	}
	if account != nil {
		return ErrAccountAlreadyExists
	}

	newAccount := CreateAccount(login, password, userId, role, s.passwordEncoder)

	return s.accountRepository.Save(newAccount)
}
