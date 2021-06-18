package BDD

//package main

/*
msg d'erreur print dans le terminal
*/

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

//func main() {
//status, db := GestionData()
//fmt.Println(status)
//index := newUser("test1", "test1", db)
//fmt.Println(index)
//status, db = gestionData()
//statusUser, tab := checkUser("test1", db)
//fmt.Println(statusUser)
//fmt.Println(tab)
//statusPost := MakePost("lorem ipsum", 1)
//statusPost1 := MakePost("lorem ipsumE", 1)
//statusPost2 := MakePost("lorem ipsumY", 1)
//statusPost3 := MakePost("lorem ipsumI", 1)
//fmt.Println(statusPost, statusPost1,statusPost2,statusPost3)
//statusGetPost, tabPost := getPost(db, "Id_Test")
//fmt.Println(statusGetPost)
//fmt.Println(tabPost)*/
//fmt.Println(NewCmt("IDUser", "IDPOST1", "lorem ipsum", db))
//}

/*///////////////////////////////////recuperation de la base de donnée ///////////////////////////*/

func GestionData() (int, *sql.DB) {
	db, err := sql.Open("sqlite3", "./BDD/ProjetForum.db") //lancer depuis : (bdd.go) lancer depuis serveur.go : (./BDD/ProjetForum.db) le chemin du projet devra changer dependant de l'endroit exectution
	if err != nil {
		fmt.Println(err)
		fmt.Print("error ouvertur base")
		return 500, db
	}
	return 0, db
}

/*/////////////////////////////////////////////////////creation d'un nouvelle identifiant///////////////////////////////////////////////*/

func NewUser(Pseudo string, HashPass string, db *sql.DB) int {
	statement, err := db.Prepare("INSERT INTO User (pseudo, password) VALUES(?,?)")
	if err != nil {
		fmt.Println(err)
		fmt.Println("error Prepare new user")
		return (500)
	}
	statement.Exec(Pseudo, HashPass)
	db.Close()
	return (0)
}

/*///////////////////////////////////////////////////////////verification de identifiant///////////////////////////////////////////////////////////////////////////////*/

func CheckUser(username string, db *sql.DB) (int, [2]string) {
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

/////////////////////////////////////////////////////////get id_user/////////////////////////////////////////

func GetId_User(username string, db *sql.DB) (int, int) {
	var Id_user int = -1

	tsql, err := db.Query("SELECT Id_user FROM User WHERE pseudo = (?)", username)
	if err != nil {
		fmt.Println(err)
		return 500, Id_user
	}

	for tsql.Next() {
		tsql.Scan(&Id_user)
	}
	return 0, Id_user
}

/*//////////////////////////////////////////////////recupe post////////////////////////////////////////////////*/

func GetPost(db *sql.DB, id int) (int, [3]string) {
	var tabPost [3]string

	var image string
	var text string
	var titre string

	statement, err := db.Query("SELECT image, texte, titre FROM Post WHERE Id_post = (?)", id)
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

////////////////////////////////Get All Post///////////////////////////////////////////////////////////////////////////////////////////////////////////

func GetAllPost(db *sql.DB) ([][]string) {
	var tabAllPost [][]string

	var text string
	var id int

	statement, err := db.Query("SELECT Id_post, texte FROM Post")
	if err != nil {
		fmt.Println(err)
		return tabAllPost
	}

	for statement.Next() {
		statement.Scan(&id, &text)
		save := []string{strconv.Itoa(id), text}
		tabAllPost = append(tabAllPost, save)
	}

	return tabAllPost
}

/*////////////////////////////////////// create post///////////////////////////////*/

func MakePost(text string, ID_User int) int {
    status, db := GestionData()
    if status == 500 {
        fmt.Println("can't open BDD")
        return 500
    }
    newPost, err := db.Prepare("INSERT INTO Post (texte, Id_user) VALUES(?, ?)")
    if err != nil {
        fmt.Println("Prepare error")
        fmt.Println(err)
        return 500
    } else {
        newPost.Exec(text, ID_User)
        db.Close()
        return 300
    }
}
////////////////////////creation post///////////////////////////////

func NewCmt(Id_user string, Id_post string, contenu string, db *sql.DB) int {
	statement, err := db.Prepare("INSERT INTO Commentaires (Id_user, Id_post, contenu) VALUES(?,?,?)")
	if err != nil {
		fmt.Println(err)
		fmt.Println("error Prepare new comment")
		return (500)
	}
	statement.Exec(Id_user, Id_post, contenu)
	db.Close()
	return (0)
}

//////////////////////////////////////// get UUID from User //////////////////////////////////////

func GetUUID_User(username string, db *sql.DB) (int, string) {
	var UUID string = ""

	tsql, err := db.Query("SELECT UUID FROM User WHERE pseudo = (?)", username) // check for UUID name in database
	if err != nil {
		fmt.Println(err)
		return 500, UUID
	}

	for tsql.Next() {
		tsql.Scan(&UUID)
	}
	return 0, UUID
}

////////////////////////////////// put UUID in BDD //////////////////////////

func PutUUID(UUID string, db *sql.DB) int {
	statement, err := db.Prepare("INSERT INTO User UUID VALUES(?)")
	if err != nil {
		fmt.Println(err)
		fmt.Println("error Prepare new user")
		return (500)
	}
	statement.Exec(UUID)
	db.Close()
	return (0)
}
