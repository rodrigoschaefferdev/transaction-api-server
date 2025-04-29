package service

import (
	"transaction-api-server/internal/domain/entity"
	"transaction-api-server/repository"
)

type TransactionUseCase struct {
	repository repository.TransactionRepository
}

func NewTransactionUseCase(repo repository.TransactionRepository) TransactionUseCase {
	return TransactionUseCase{
		repository: repo,
	}
}

func (au *TransactionUseCase) ListTransactions() ([]entity.Transaction, error) {
	return au.repository.ListTransactions()
}

func (au *TransactionUseCase) ListTransactionsAmount(document string) ([]entity.TransactionAmount, error) {
	return au.repository.ListTransactionsAmount(document)
}

func (au TransactionUseCase) CreateTransaction(transaction entity.Transaction) (entity.Transaction, error) {
	return au.repository.CreateTransaction(&transaction)
}

func (au TransactionUseCase) ListTransactionsByAccountId(accountID int64) ([]entity.TransactionJoinAccount, error) {
	return au.repository.ListTransactionsByAccountId(accountID)
}

func (au TransactionUseCase) ListTransactionsByDocument(documento string) ([]entity.TransactionJoinAccount, error) {
	return au.repository.ListTransactionsByDocument(documento)
}
