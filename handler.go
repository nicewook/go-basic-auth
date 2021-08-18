package main

import (
	"database/sql"
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
	var creds Credential

	// NewDecoder vs Unmarshal: https://stackoverflow.com/a/21198571/3382699
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		log.Printf("failed to Decode: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 10) // 10 means difficulty
	if err != nil {
		log.Printf("failed to GenerateFromPassword: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	SaveQuery := fmt.Sprintf(`insert into %s(username, password) values ("%s", "%s")`, userAccountTable, creds.Username, string(hashedPassword))
	if _, err := db.Exec(SaveQuery); err != nil {
		log.Printf("failed to save user account: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func Signin(w http.ResponseWriter, r *http.Request) {
	log.Println("Sign in")

	var creds Credential
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		log.Printf("failed to Decode: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	SignInSQL := fmt.Sprintf(`SELECT password FROM %s WHERE username = "%s";`, userAccountTable, creds.Username)
	var storedPassword string

	if err := db.QueryRow(SignInSQL).Scan(&storedPassword); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("no user account found: %v", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Printf("failed to get stored password: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(creds.Password)); err != nil {
		log.Printf("incorrect password, %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	log.Println("correct password. signed in")
}
