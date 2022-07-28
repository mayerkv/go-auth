package http_service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	domain2 "github.com/mayerkv/go-auth/internal/domain"
)

type AccountController struct {
	accountService *domain2.AccountService
}

func NewAccountController(accountService *domain2.AccountService) *AccountController {
	return &AccountController{accountService: accountService}
}

func (c *AccountController) SignUp(ctx *gin.Context) {
	var dto domain2.SignUpDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	err := c.accountService.CreateAccount(dto.Login, dto.Password, uuid.NewString(), domain2.AccountRoleUser)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.Status(http.StatusCreated)
}
