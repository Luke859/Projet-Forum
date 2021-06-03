package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func gestionData() int {
	db, err := sql.Open("sqlitte3", "./nomBase.db")
	if err != nil {
		fmt.Print("error ouvertur base")
		return (500)
	}
	pseudo := ""
	Hmpd := ""
	statusNewUser := newUser(pseudo, Hmpd, db)
	if statusNewUser == 0 {
		return (0)
	} else {
		fmt.Print("probleme --> creation user")
		return (500)
	}
}

func newUser(pseudo string, Hmpd string, db *sql.DB) int {
	statement, err := db.Prepare("INSERT INTO User (pseudo, Hmpd) VALUES(?,?)")
	if err != nil {
		fmt.Print("error new user")
		return (5)
	}
	statement.Exec(pseudo, Hmpd)
	return (0)
}

func checkUser(pseudo string, Hmpd string, db *sql.DB) int {
	return (0)
}
