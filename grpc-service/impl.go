package grpc_service

import (
	"context"
	"github.com/mayerkv/go-auth/domain"
)

type AuthServiceServerImpl struct {
	accountService *domain.AccountService
}

func NewAuthServiceServerImpl(accountService *domain.AccountService) *AuthServiceServerImpl {
	return &AuthServiceServerImpl{accountService: accountService}
}

func (s *AuthServiceServerImpl) CreateAccount(ctx context.Context, request *CreateAccountRequest) (*CreateAccountResponse, error) {
	err := s.accountService.CreateAccount(request.Email, request.Password, request.UserId, mapAccountRole(request.Role))
	if err != nil {
		return nil, err
	}

	return &CreateAccountResponse{}, nil
}

func mapAccountRole(role AccountRole) domain.AccountRole {
	switch role {
	case AccountRole_ADMIN:
		return domain.AccountRoleAdmin
	}

	return domain.AccountRoleUser
}

func (s *AuthServiceServerImpl) mustEmbedUnimplementedAuthServiceServer() {
	panic("implement me")
}
