package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "user"
	password = "pass"
	dbName   = "postgres_db"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	logrus.Info("Connection string:" + psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logrus.Error("Error connecting to the database: " + err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logrus.Error("Error verifying the connection to the database: " + err.Error())
		return nil, err
	}

	logrus.Info("Database connection successful!")
	return db, nil
}
