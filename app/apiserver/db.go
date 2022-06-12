// package apiserver

// import (
//     "database/sql"
//     "fmt"
//     "log"
// )

// var db *sql.DB

// func connectDb() *sql.DB {

//     // Get a database handle.
//     var err error
//     db, err = sql.Open("mysql", "gen_user:0fgxqh8bc@85.193.83.246:3306/default_db")
//     if err != nil {
//         log.Fatal(err)
//     }

//     pingErr := db.Ping()
//     if pingErr != nil {
//         log.Fatal(pingErr)
//     }
//     fmt.Println("Connected!")	

// 	return db
// }