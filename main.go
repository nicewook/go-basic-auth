package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

const (
	userDB           = "user.db"
	userAccountTable = "userAccountTable"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", userDB)
	if err != nil {
		log.Fatalf("fail to open db file. %v", err)
	}

	createTableQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (username TEXT NOT NULL,	password TEXT NOT NULL);`, userAccountTable)
	if _, err = db.Exec(createTableQuery); err != nil {
		log.Fatalf("failed to create table %s: %v", userAccountTable, err)
	}
}

func main() {
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/signup", Signup)

	initDB()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
