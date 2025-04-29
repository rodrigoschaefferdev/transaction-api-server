package postgres

import (
	"database/sql"
	"strconv"
	"time"
	"transaction-api-server/internal/domain/entity"

	"github.com/sirupsen/logrus"
)

type TransactionRepository struct {
	connection *sql.DB
}

func NewTransactionRepository(connection *sql.DB) TransactionRepository {
	return TransactionRepository{
		connection: connection,
	}
}

func (ac TransactionRepository) ListTransactions() (transactions []entity.Transaction, err error) {
	query := `
        SELECT * FROM transactions ts
    `

	rows, err := ac.connection.Query(query)
	if err != nil {
		logrus.Error(err)
		return []entity.Transaction{}, err
	}

	var transactionList []entity.Transaction
	var transactionObj entity.Transaction

	for rows.Next() {
		err = rows.Scan(
			&transactionObj.ID,
			&transactionObj.TransactionTypeId,
			&transactionObj.AccountId,
			&transactionObj.Amount,
			&transactionObj.TransactionDate,
		)
		if err != nil {
			logrus.Error(err)
			return []entity.Transaction{}, err
		}
		transactionList = append(transactionList, transactionObj)
	}

	rows.Close()
	return transactionList, nil
}

func (ac TransactionRepository) ListTransactionsAmount(document string) (transactions []entity.TransactionAmount, err error) {
	query := `
        select
            a."name",
            a."document",
            coalesce(sum(ts.amount * case when ts.transaction_type_id = 1 then 1 when ts.transaction_type_id = 2 then -1 else 0 end), 0) as amount
        from
            transactions ts
        left join (select distinct id, "document", "name" from accounts) a on
            a.id = ts.account_id
        where
            a."document" = $1
        group by
            a."name",
            a."document"
    `

	rows, err := ac.connection.Query(query, document)
	if err != nil {
		logrus.Error(err)
		return []entity.TransactionAmount{}, err
	}

	var transactionList []entity.TransactionAmount
	var transactionObj entity.TransactionAmount

	for rows.Next() {
		err = rows.Scan(
			&transactionObj.Name,
			&transactionObj.Document,
			&transactionObj.Amount,
		)
		if err != nil {
			logrus.Error(err)
			return []entity.TransactionAmount{}, err
		}
		logrus.Infof("Scanned row: Name=%s, Document=%s, Amount=%.2f",
			transactionObj.Name, transactionObj.Document, transactionObj.Amount)
		transactionList = append(transactionList, transactionObj)
	}

	rows.Close()
	return transactionList, nil
}

func (ac TransactionRepository) CreateTransaction(transaction *entity.Transaction) (entity.Transaction, error) {
	sqlStatement := `
		INSERT INTO transactions (
			"transaction_type_id",
			"account_id",
			"amount",
			"transaction_date"
		)
		VALUES (
			$1,
			$2,
			$3,
			$4
		)
		RETURNING id
	`
	var newID int64
	err := ac.connection.QueryRow(
		sqlStatement,
		transaction.TransactionTypeId,
		transaction.AccountId,
		transaction.Amount,
		time.Now(),
	).Scan(&newID)

	if err != nil {
		logrus.Error("Error inserting transaction into database: " + err.Error())
		return entity.Transaction{}, err
	}

	transaction.ID = newID

	logrus.Info("Transaction inserted successfully with ID: " + strconv.FormatInt(newID, 10))
	return *transaction, nil

}

func (ac TransactionRepository) findTransactions(filter string, filterValue interface{}) (transactions []entity.TransactionJoinAccount, err error) {
	query := `
        select
            ts.id,
            a."name",
            a."document",
            ts.amount,
            tt.description,
            ts.transaction_date 
        from
            transactions ts
        left join accounts a on
            a.id = ts.account_id
        left join transaction_type tt on
            tt.id = ts.transaction_type_id 
        where
            ` + filter + ` = $1
    `
	rows, err := ac.connection.Query(query, filterValue)
	if err != nil {
		logrus.Error(err)
		return []entity.TransactionJoinAccount{}, err
	}

	var transactionList []entity.TransactionJoinAccount
	var transactionObj entity.TransactionJoinAccount

	for rows.Next() {
		err = rows.Scan(
			&transactionObj.ID,
			&transactionObj.Name,
			&transactionObj.Document,
			&transactionObj.Amount,
			&transactionObj.Description,
			&transactionObj.TransactionDate,
		)
		if err != nil {
			logrus.Error(err)
			return []entity.TransactionJoinAccount{}, err
		}
		transactionList = append(transactionList, transactionObj)
	}

	rows.Close()
	return transactionList, nil
}

func (ac TransactionRepository) ListTransactionsByAccountId(accountId int64) (transactions []entity.TransactionJoinAccount, err error) {
	return ac.findTransactions("a.id", accountId)
}

func (ac TransactionRepository) ListTransactionsByDocument(document string) (transactions []entity.TransactionJoinAccount, err error) {
	return ac.findTransactions("a.\"document\"", document)
}
