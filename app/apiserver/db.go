package apiserver

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("mysql",dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}