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

/*///////////////////////////////////recuperation de la base de donnée ///////////////////////////*/

/*
ouverture de la BDD
renvois un int qui sert de status d'erreur :
	500 en cas d'erreur
	0 en cas de reussite
renvois un pointeur *sql.DB qui sert dans toute les autre fonctions
*/
func GestionData() (int, *sql.DB) {
	db, err := sql.Open("sqlite3", "./BDD/ProjetForum.db") //lancer depuis : (../../bdd.go) lancer depuis serveur.go : (./BDD/ProjetForum.db) le chemin du projet devra changer dependant de l'endroit exectution
	if err != nil {
		fmt.Println(err)
		fmt.Print("error ouvertur base")
		return 500, db
	}
	return 0, db
}
