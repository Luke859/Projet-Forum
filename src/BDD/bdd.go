package main

/*
msg d'erreur print dans le terminal
*/

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	status, db := gestionData()
	fmt.Println(status)
	//index := newUser("test", "test", db)
	//fmt.Print(index)
	status, tab := checkUser("test", db)
	fmt.Println(status)
	fmt.Println(tab)
}

/*///////////////////////////////////recuperation de la base de donnée ///////////////////////////*/

func gestionData() (int, *sql.DB) {
	db, err := sql.Open("sqlite3", "../../BDD/Projet_Forum")
	if err != nil {
		fmt.Println(err)
		fmt.Print("error ouvertur base")
		return 500, db
	}
	return 0, db
}

/*/////////////////////////////////////////////////////creation d'un nouvelle identifiant///////////////////////////////////////////////*/

func newUser(pseudo string, Hmpd string, db *sql.DB) int {
	statement, err := db.Prepare("INSERT INTO User (pseudo, password) VALUES(?,?)")
	if err != nil {
		fmt.Println(err)
		fmt.Println("error Prepare new user")
		return (500)
	}
	statement.Exec(pseudo, Hmpd)
	db.Close()
	return (0)
}

/*///////////////////////////////////////////////////////////verification de identifiant///////////////////////////////////////////////////////////////////////////////*/

func checkUser(username string, db *sql.DB) (int, [2]string) {
	var tabUser [2]string

	var pseudo string
	var password string

	tsql, err := db.Query("SELECT pseudo, password FROM User WHERE pseudo = (?)", username)
	if err != nil {
		fmt.Println(err)
		return 500, tabUser
	}

	for tsql.Next() {
		tsql.Scan(&pseudo, &password)
	}
	tabUser[0] = pseudo
	tabUser[1] = password
	return 0, tabUser
}

/*//////////////////////////////////////////////////recupe post////////////////////*/

/*func getPost(db *sql.DB, id string) (int, string) {
	statement, err := db.Query("SELECT Id_post FROM Post WHERE User")
	if err != nil {
		fmt.Println("querry didn't work, can't catch ")
		return 500, "error"
	}
	return 0, ""
}*/

/*////////////////////////////////////// create post///////////////////////////////*/

func makePost(user string, text string, categorie string) (int, int) {
	status, db := gestionData()
	if status == 500 {
		fmt.Println("can't open BDD")
		return 500, 1
	}
	newPost, err := db.Prepare("INSERT INTO Post (PseudoUser, Titre, texte) VALUES(?,?,?)")
	if err != nil {
		fmt.Println("Prepare error")
		return 500, 1
	} else {
		newPost.Exec(user, categorie, text)
		return 300, 0
	}
}
