package BDD

import (
	"database/sql"
	"fmt"
)

/////////////////////////////////////////////// like /////////////////////////////////////////////

func CreateLike(ID_User int, Id_cmt int, likebool int, db *sql.DB) int {
	statm, err := db.Prepare("INSERT INTO Likes (Id_post, Id_user, like_button) VALUES (?,?,?)")
	if err != nil {
		fmt.Println(err)
		fmt.Println("error Prepare new like")
		return (500)
	}
	statm.Exec(ID_User, Id_cmt, likebool)
	//db.Close()
	return (0)
}

//////////////////////////////////////////// update like /////////////////////////////////////////////////

func IsLikedPOST(db *sql.DB, ID_like int) (int, int) {
	var IsLike int = 0

	tsql, err := db.Query("SELECT like_button FROM Likes WHERE Id_likes = (?)", ID_like) // check for UUID name in database
	if err != nil {
		fmt.Println(err)
		return 500, IsLike
	}

	for tsql.Next() {
		tsql.Scan(&IsLike)
	}
	return 0, IsLike
}

/*
La fonction a besoins d'avoir le pointeur de la base de donné et des identifiant de like, user et post
La fonction va renvoyer un int qui vaut soit : 500 soit : 0.
500 etant un code d'erreur http error du a un probleme interne du serveur
les erreurs peuvent venir du fait qu'on a pas pu recupper l'ettat du button like dans la BDD ou que la Querry n'as pas fonctionnait
*/
func UpdateLikePOST(db *sql.DB, ID_like int, ID_User int, ID_Post int) int {
	status, IsLike := IsLikedPOST(db, ID_like)
	if status == 500 {
		fmt.Println("ISlike a bugger")
		return 500
	} else {
		if IsLike == 0 {
			statement, err := db.Prepare("UPDATE Likes SET like_button = true WHERE Id_likes = ? AND Id_post =? AND Id_user =?")
			if err != nil {
				fmt.Println("A")
				fmt.Println(err)
				return 500
			}
			statement.Exec(ID_like, ID_User, ID_Post)
			fmt.Println("B")
			db.Close()
			return (0)
		} else {
			statement, err := db.Prepare("UPDATE Likes SET like_button = 0 WHERE Id_likes = ? AND Id_post =? AND Id_user =?")
			if err != nil {
				fmt.Println(err)
				return 500
			}
			fmt.Println("C")
			statement.Exec(ID_like, ID_User, ID_Post)
			db.Close()
			return (0)
		}
	}
}

////////////////////////////////// like cmt ////////////////////////////////////////////////

func IsLikedCMT(db *sql.DB, ID_like int) (int, int) {
	var IsLike int = 0

	tsql, err := db.Query("SELECT like_button FROM Likes WHERE Id_likes = (?)", ID_like) // check for UUID name in database
	if err != nil {
		fmt.Println(err)
		return 500, IsLike
	}

	for tsql.Next() {
		tsql.Scan(&IsLike)
	}
	return 0, IsLike
}

/*
La fonction a besoins d'avoir le pointeur de la base de donné et des identifiant de like, user et commentaire
La fonction va renvoyer un int qui vaut soit : 500 soit : 0.
500 etant un code d'erreur http error du a un probleme interne du serveur
les erreurs peuvent venir du fait qu'on a pas pu recupper l'ettat du button like dans la BDD ou que la Querry n'as pas fonctionnait
*/

func UpdateLikeCMT(db *sql.DB, ID_like int, ID_User int, ID_cmt int) int {
	status, IsLike := IsLikedCMT(db, ID_like)
	if status == 500 {
		return 500
	} else {
		if IsLike == 0 {
			statement, err := db.Prepare("UPDATE Likes SET like_button = 1 WHERE (Id_likes = ?, Id_cmt =?, Id_user =?)")
			if err != nil {
				fmt.Println(err)
				return 500
			}
			statement.Exec(ID_like, ID_cmt, ID_User)
			db.Close()
			return (0)
		} else {
			statement, err := db.Prepare("UPDATE Likes SET like_button = 0 WHERE Id_likes = ?, Id_cmt =? , Id_post =?, Id_user =?)")
			if err != nil {
				fmt.Println("B")
				fmt.Println(err)
				return 500
			}
			statement.Exec(ID_like, ID_cmt, ID_User)
			db.Close()
			return (0)
		}
	}
}
