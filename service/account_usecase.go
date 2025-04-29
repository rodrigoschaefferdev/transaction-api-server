package service

import (
	"transaction-api-server/internal/domain/entity"
	"transaction-api-server/repository"
)

type AccountUseCase struct {
	repository repository.AccountRepository
}

func NewAccountUseCase(repo repository.AccountRepository) AccountUseCase {
	return AccountUseCase{
		repository: repo,
	}
}

func (au *AccountUseCase) ListAccounts() ([]entity.Account, error) {
	return au.repository.ListAccounts()
}

func (au AccountUseCase) GetAccountById(accountID int64) (entity.Account, error) {
	return au.repository.GetAccountById(accountID)
}

func (au AccountUseCase) CreateAccount(account entity.Account) (entity.Account, error) {
	return au.repository.CreateAccount(&account)
}
