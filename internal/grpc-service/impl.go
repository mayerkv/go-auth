package grpc_service

import (
	"context"

	domain2 "github.com/mayerkv/go-auth/internal/domain"
)

type AuthServiceServerImpl struct {
	accountService *domain2.AccountService
}

func NewAuthServiceServerImpl(accountService *domain2.AccountService) *AuthServiceServerImpl {
	return &AuthServiceServerImpl{accountService: accountService}
}

func (s *AuthServiceServerImpl) CreateAccount(
	ctx context.Context,
	request *CreateAccountRequest,
) (*CreateAccountResponse, error) {
	err := s.accountService.CreateAccount(request.Email, request.Password, request.UserId, mapAccountRole(request.Role))
	if err != nil {
		return nil, err
	}

	return &CreateAccountResponse{}, nil
}

func mapAccountRole(role AccountRole) domain2.AccountRole {
	switch role {
	case AccountRole_ADMIN:
		return domain2.AccountRoleAdmin
	}

	return domain2.AccountRoleUser
}

func (s *AuthServiceServerImpl) mustEmbedUnimplementedAuthServiceServer() {
	panic("implement me")
}
