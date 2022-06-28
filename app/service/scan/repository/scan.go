package repository

import (
	"database/sql"
)

type Config struct {
	db *sql.DB
}

func New(db *sql.DB) *Config {
	return &Config{
		db: db,
	}
}

