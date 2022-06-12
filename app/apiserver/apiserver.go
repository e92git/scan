package apiserver

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("mysql", "gen_user:0fgxqh8bc@tcp(85.193.83.246:3306)/default_db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}