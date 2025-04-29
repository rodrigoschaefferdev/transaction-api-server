package postgres

import (
	"database/sql"
	"strconv"
	"time"
	"transaction-api-server/internal/domain/entity"

	"github.com/sirupsen/logrus"
)

type AccountRepository struct {
	connection *sql.DB
}

// CreateAccount implements repository.AccountRepository.
func (ac AccountRepository) CreateAccount(account *entity.Account) (entity.Account, error) {
	sqlStatement := `
		INSERT INTO accounts (
			"name",
			"document",
			"created_at"
		)
		VALUES (
			$1,
			$2,
			$3
		)
		RETURNING id
	`
	var newID int64
	err := ac.connection.QueryRow(
		sqlStatement,
		account.Name,
		account.Document,
		time.Now(),
	).Scan(&newID)

	if err != nil {
		logrus.Error("Error inserting account into database: " + err.Error())
		return entity.Account{}, err
	}

	account.ID = newID

	logrus.Info("Account inserted successfully with ID: " + strconv.FormatInt(newID, 10))
	return *account, nil
}

// GetAccountById implements repository.AccountRepository.
func (ac AccountRepository) GetAccountById(accountID int64) (entity.Account, error) {
	query := "SELECT id, name, document FROM ACCOUNTS WHERE id = $1"

	logrus.Info("Executing query: " + query + " with accountId: " + strconv.FormatInt(accountID, 10))
	row := ac.connection.QueryRow(query, accountID)

	var accountObj entity.Account

	if err := row.Scan(&accountObj.ID, &accountObj.Name, &accountObj.Document); err != nil {
		if err == sql.ErrNoRows {
			logrus.Info("No account found with accountId: " + strconv.FormatInt(accountID, 10))
			return entity.Account{}, nil
		}
		logrus.Error("Error scanning row: " + err.Error())
		return entity.Account{}, err
	}

	logrus.WithFields(logrus.Fields{
		"account": accountObj,
	}).Info("Account found")

	return accountObj, nil
}

// ListAccounts implements repository.AccountRepository.
func (ac AccountRepository) ListAccounts() ([]entity.Account, error) {
	query := "SELECT id, name, document FROM ACCOUNTS"
	rows, err := ac.connection.Query(query)

	if err != nil {
		logrus.Error(err)
		return []entity.Account{}, err
	}
	var accountList []entity.Account
	var accountObj entity.Account

	for rows.Next() {
		err = rows.Scan(
			&accountObj.ID,
			&accountObj.Name,
			&accountObj.Document,
		)
		if err != nil {
			logrus.Error(err)
			return []entity.Account{}, err
		}
		accountList = append(accountList, accountObj)
	}
	rows.Close()

	return accountList, nil
}

func NewAccountRepository(connection *sql.DB) AccountRepository {
	return AccountRepository{
		connection: connection,
	}
}
