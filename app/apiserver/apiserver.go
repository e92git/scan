package apiserver

import (
	"database/sql"
	// "net/http"
)

// Start ...
// func Start(config *Config) error {
// 	db, err := newDB(config.DatabaseURL)
// 	if err != nil {
// 		return err
// 	}

// 	defer db.Close()
// 	// store := sqlstore.New(db)
// 	// srv := newServer(store, sessionStore)

// 	// return http.ListenAndServe(config.BindAddr, srv)
// }

func ConnectDb(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("mysql", "gen_user:0fgxqh8bc@85.193.83.246:3306/default_db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}