package entity

import "time"

type Transaction struct {
	ID                int64     `db:"id"`
	AccountId         int64     `db:"acount_id"`
	TransactionTypeId int32     `db:"transaction_type_id"`
	Amount            float64   `db:"amount"`
	TransactionDate   time.Time `db:"transaction_date"`
}

type TransactionJoinAccount struct {
	ID              int64     `db:"id"`
	Name            string    `db:"name"`
	Document        string    `db:"document"`
	Description     string    `db:"description"`
	Amount          float64   `db:"amount"`
	TransactionDate time.Time `db:"transaction_date"`
}

type TransactionAmount struct {
	Name     string  `db:"name"`
	Document string  `db:"document"`
	Amount   float64 `db:"amount"`
}
