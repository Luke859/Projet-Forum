//package BDD

package main

/*
msg d'erreur print dans le terminal
*/

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	status, db := GestionData()
	fmt.Println(status)
	//index := NewUser("test1", "test1", db)
	//fmt.Println(index)
	//statusUser, tab := CheckUser("test1", db)
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
	//fmt.Println(NewCmt(1, 1, "lorem ipsum", db))
	fmt.Println(CreateLike(1, 1, true, db))
	fmt.Println(UpdateLikeCMT(db, 1, 1, 1))
}

/*///////////////////////////////////recuperation de la base de donnée ///////////////////////////*/

func GestionData() (int, *sql.DB) {
	db, err := sql.Open("sqlite3", "../../BDD/ProjetForum.db") //lancer depuis : (../../bdd.go) lancer depuis serveur.go : (./BDD/ProjetForum.db) le chemin du projet devra changer dependant de l'endroit exectution
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

func GetAllPost(db *sql.DB) (int, []string) {
	var tabAllPost []string

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

////////////////////////creation post///////////////////////////////

func NewCmt(Id_user int, Id_post int, contenu string, db *sql.DB) int {
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
