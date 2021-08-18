package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

const (
	userAccountDB = "useraccounts.db"
)

var (
	db *sql.DB
)

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", userAccountDB)
	if err != nil {
		log.Fatalf("fail to open db file. %v", err)
	}
	// defer db.Close()

	createTableQuery := `CREATE TABLE IF NOT EXISTS userAccounts (
		username TEXT NOT NULL,
		password TEXT NOT NULL
		);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("create table %v", err)
	}
}

func main() {
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/signup", Signup)

	initDB()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
