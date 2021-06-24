package BDD

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

////////////////////////creation commentaire///////////////////////////////

/*
Creation d'un commentaire dans la BDD
besoins : d'un indentifiant de l'utilisateur, d'un identifiant du post que on veut commenter et du pointeur sur la BDD
revois un int fessant code d'erreur 500 pour un echec ou 0 en cas de reussite
*/

func MakeCmt(Id_user int, Id_post int, contenu string, db *sql.DB) int {
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

//////////////////////// recupere tous les commentaire d'un post ////////////////////////////////////////////////

/*
Recuper les commentaire d'un post
besoins du pointeur de la BDD et de l'identifiant du post en question
Renvoi un int de valeur : 0 et tous les commentaire du post en cas de succée
Renvois un int de valeur 500 et un double tableau vide en cas d'echec
*/

func GetAllCmt(db *sql.DB, id_post int) (int, [][]string) {
	var tabAllPost [][]string

	var text string
	var idUser int
	var idPost int

	statement, err := db.Query("SELECT Id_user, Id_post, contenu FROM Commentaires WHERE Id_post = ?", id_post)
	if err != nil {
		fmt.Println(err)
		return 500, tabAllPost
	}

	//remplie notre tableau de tout les commentaire
	for statement.Next() {
		statement.Scan(&idUser, &idPost, &text)
		save := []string{strconv.Itoa(idUser), strconv.Itoa(idPost), text}
		tabAllPost = append(tabAllPost, save)
	}
	return 0, tabAllPost
}
