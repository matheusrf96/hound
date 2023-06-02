package db

import (
	"database/sql"
	"sub/src/config"

	_ "github.com/lib/pq"
)

func connectDatabase(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Connect() (*sql.DB, error) {
	db, err := connectDatabase(config.DatabaseConnectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
