package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	log.Println("Sign up")
	creds := &Credential{}
	if err := json.NewDecoder(r.Body).Decode(creds); err != nil { // https://stackoverflow.com/a/21198571/3382699
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 10) // what 8 means -
	if err != nil {
		log.Printf("GenerateFromPassword: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("add %v and %v", creds.Username, string(hashedPassword))
	SaveQuery := `insert into userAccounts(username, password) values ($1, $2)`
	stmt, err := db.Prepare(SaveQuery) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Fatalf("prepare: ", err)
	}
	result, err := stmt.Exec(creds.Username, string(hashedPassword))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("result: %+v", result)

}

func Signin(w http.ResponseWriter, r *http.Request) {
	log.Println("Sign in")

	creds := &Credential{}
	if err := json.NewDecoder(r.Body).Decode(creds); err != nil { // https://stackoverflow.com/a/21198571/3382699
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("received %v", creds)
	log.Printf("Username %v", creds.Username)

	SignInSQL := fmt.Sprintf(`SELECT password FROM userAccounts WHERE username = "%s";`, creds.Username)
	var pw string

	db.QueryRow(SignInSQL).Scan(&pw)
	log.Printf("pw: %v", pw)

	// stored := &Credential{}
	// if err := rows.Scan(&stored); err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Printf("Scan: %v", err)
	// 		w.WriteHeader(http.StatusUnauthorized)
	// 		return
	// 	}
	// 	log.Printf("Scan: %v", err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	// if err := bcrypt.CompareHashAndPassword([]byte(stored.Password), []byte(creds.Password)); err != nil {
	// 	log.Printf("CompareHashAndPassword %v, %v", stored.Password, err)
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }
	// log.Printf("storedCres.Password: %v", stored.Password)
}
