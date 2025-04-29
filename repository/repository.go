package repository

import "transaction-api-server/internal/domain/entity"

type AccountRepository interface {
	CreateAccount(account *entity.Account) (entity.Account, error)
	GetAccountById(accountID int64) (entity.Account, error)
	ListAccounts() ([]entity.Account, error)
}

type TransactionRepository interface {
	ListTransactions() ([]entity.Transaction, error)
	CreateTransaction(transaction *entity.Transaction) (entity.Transaction, error)
	ListTransactionsAmount(document string) ([]entity.TransactionAmount, error)
	ListTransactionsByAccountId(id int64) ([]entity.TransactionJoinAccount, error)
	ListTransactionsByDocument(document string) ([]entity.TransactionJoinAccount, error)
}
