package services

import (
	"github.com/google/uuid"
	"github.com/kofkuiper/echo-api/repositories"
)

type (
	accountService struct {
		accountRepo repositories.AccountRepository
	}
)

func NewAccountService(accountRepo repositories.AccountRepository) AccountService {
	return accountService{accountRepo: accountRepo}
}

// SignUp implements AccountService.
func (a accountService) SignUp(request SignUpRequest) (*Account, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	account, err := a.accountRepo.Create(repositories.Account{
		ID:       id.String(),
		Username: request.Username,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}
	return &Account{
		ID:          account.ID,
		Username:    account.Username,
		CreatedAt:   account.CreatedAt,
		UpdatedAtAt: account.UpdatedAt,
	}, nil
}
