package BDD

//package main

/*
msg d'erreur print dans le terminal
*/

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

//func main() {
//status, db := GestionData()
//fmt.Println("statu de l'ouverture de la BDD : ", status)
//index := NewUser("PIERRIC", "TESTPASS", db)
//fmt.Println(index)
//status, db = GestionData()
//fmt.Println(GetUUID_User("PIERRIC", db))
//fmt.Println("statu get uuid", PutUUID("testUUID456", "PIERRIC", db))
//fmt.Println(GetAllUsername(db))
//statusUser, tab := CheckUser("PIERRIC", db)
//fmt.Println(statusUser)
//fmt.Println(tab)
//statusPost := MakePost("Lorem IPSUM", 3)
//fmt.Println(statusPost)
//statusGetPost, tabPost := GetPost(db, 4)
//fmt.Println(statusGetPost)
//fmt.Println(tabPost)
//statusAllPost, tabAllPost := GetAllPost(db)
//fmt.Println(statusAllPost)
//fmt.Println(tabAllPost)
//fmt.Println("status de la creation d'un nouveau cmt :", NewCmt(1, 1, "lorem ipsum", db))
//status, db = GestionData()
//fmt.Println(GetAllCmt(db, 2))
//fmt.Println(CreateLike(1, 1, true, db))
//fmt.Println(UpdateLikePOST(db, 3, 1, 1))
//}

/*///////////////////////////////////recuperation de la base de donnée ///////////////////////////*/

/*
ouverture de la BDD
renvois un int qui sert de status d'erreur :
	500 en cas d'erreur
	0 en cas de reussite
renvois un pointeur *sql.DB qui sert dans toute les autre fonctions
*/
func GestionData() (int, *sql.DB) {
	db, err := sql.Open("sqlite3", "./BDD/ProjetForum_5.db") //lancer depuis : (../../bdd.go) lancer depuis serveur.go : (./BDD/ProjetForum.db) le chemin du projet devra changer dependant de l'endroit exectution
	if err != nil {
		fmt.Println(err)
		fmt.Print("error ouvertur base")
		return 500, db
	}
	return 0, db
}
