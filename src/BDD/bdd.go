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
	//status, tab := checkUser("test", db)
	//fmt.Println(status)
	//fmt.Println(tab)
	statusPost := makePost("Id_Test", "", "lorem ipsum", "Titre")
	fmt.Println(statusPost)
	statusGetPost, tabPost := getPost(db, "Id_Test")
	fmt.Println(statusGetPost)
	fmt.Println(tabPost)
}

/*///////////////////////////////////recuperation de la base de donnée ///////////////////////////*/

func gestionData() (int, *sql.DB) {
	db, err := sql.Open("sqlite3", "../../BDD/Projet_Forum") //le chemin du projet devra changer dependant de l'endroit exectution
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

func getPost(db *sql.DB, id string) (int, [3]string) {
	var tabPost [3]string

	var image string
	var text string
	var titre string

	statement, err := db.Query("SELECT image, texte, titre FROM Post WHERE Id_user = (?)", id)
	if err != nil {
		fmt.Println(err)
		return 500, tabPost
	}

	for statement.Next() {
		statement.Scan(&image, &text, &titre)
	}
	tabPost[0] = image
	tabPost[1] = text
	tabPost[2] = titre

	return 0, tabPost
}

/*////////////////////////////////////// create post///////////////////////////////*/

func makePost(id string, image string, text string, titre string) int {
	status, db := gestionData()
	if status == 500 {
		fmt.Println("can't open BDD")
		return 500
	}
	newPost, err := db.Prepare("INSERT INTO Post (Id_user, image, texte, titre) VALUES(?,?,?,?)")
	if err != nil {
		fmt.Println("Prepare error")
		fmt.Println(err)
		return 500
	} else {
		newPost.Exec(id, image, text, titre)
		db.Close()
		return 300
	}
}
