package entity

import "time"

type Account struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Document  string    `db:"document"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
