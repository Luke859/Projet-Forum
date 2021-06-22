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
//fmt.Println("statu de l'ouverture de la BDD : ", status)
//index := NewUser("PIERRIC", "TESTPASS", db)
//fmt.Println(index)
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
//fmt.Println(UpdateLikeCMT(db, 1, 1, 1))
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
	db, err := sql.Open("sqlite3", "./BDD/ProjetForum.db") //lancer depuis : (../../bdd.go) lancer depuis serveur.go : (./BDD/ProjetForum.db) le chemin du projet devra changer dependant de l'endroit exectution
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

func CheckPassword(username string, db *sql.DB) (int, string) {
	var HashPass string
	var password string

	tsql, err := db.Query("SELECT password FROM User WHERE pseudo = (?)", username)
	if err != nil {
		fmt.Println(err)
		return 500, HashPass
	}

	for tsql.Next() {
		tsql.Scan(&password)
	}
	HashPass = password
	return 0, HashPass
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

////////////////////////////////// Get All Username /////////////////////////////////////////////////////////////////////:

func GetAllUsername(db *sql.DB) (int, []string) {
	var allUsername []string
	var username string

	tsql, err := db.Query("SELECT pseudo FROM User")
	if err != nil {
		fmt.Println(err)
		return 500, allUsername
	}

	for tsql.Next() {
		tsql.Scan(&username)
		allUsername = append(allUsername, username)
	}
	return 0, allUsername
}

/*//////////////////////////////////////////////////recupe post////////////////////////////////////////////////*/

func GetPost(db *sql.DB, id int) (int, [1]string) {
	var tabPost [1]string
	var text string

	statement, err := db.Query("SELECT  texte FROM Post WHERE Id_post = (?)", id)
	if err != nil {
		fmt.Println(err)
		return 500, tabPost
	}

	for statement.Next() {
		statement.Scan(&text)
	}
	tabPost[0] = text

	return 0, tabPost
}

////////////////////////////////Get All Post///////////////////////////////////////////////////////////////////////////////////////////////////////////

func GetAllPost(db *sql.DB) (int, [][]string) {
	var tabAllPost [][]string

	var text string
	var id int

	statement, err := db.Query("SELECT Id_post, texte FROM Post")
	if err != nil {
		fmt.Println(err)
		return 500, tabAllPost
	}

	for statement.Next() {
		statement.Scan(&id, &text)
		save := []string{strconv.Itoa(id), text}
		tabAllPost = append(tabAllPost, save)
	}

	return 0, tabAllPost
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

/////////////////////////////////////////////// like /////////////////////////////////////////////

func CreateLike(ID_User int, Id_cmt int, likebool bool, db *sql.DB) int {
	statm, err := db.Prepare("INSERT INTO Likes (Id_cmt, Id_user, like_button) VALUES (?,?,?)")
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
		return 500
	} else {
		if IsLike == 0 {
			statement, err := db.Prepare("UPDATE Likes SET like_button = true WHERE Id_likes = ?, Id_post =?, Id_user =?)")
			if err != nil {
				fmt.Println(err)
				return 500
			}
			statement.Exec(ID_like, ID_User, ID_Post)
			db.Close()
			return (0)
		} else {
			statement, err := db.Prepare("UPDATE Likes SET like_button = 0 WHERE Id_likes = ?, Id_post =?, Id_user =?)")
			if err != nil {
				fmt.Println(err)
				return 500
			}
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
				fmt.Println("A")
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
