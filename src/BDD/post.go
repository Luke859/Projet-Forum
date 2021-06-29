package BDD

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

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

func GetAllPost(db *sql.DB) (int, [][]string, int) {
	var tabAllPost [][]string

	var text string
	var id int

	statement, err := db.Query("SELECT Id_post, texte FROM Post")
	if err != nil {
		fmt.Println(err)
		return 500, tabAllPost, id
	}

	for statement.Next() {
		statement.Scan(&id, &text)
		save := []string{strconv.Itoa(id), text}
		tabAllPost = append(tabAllPost, save)
	}

	return 0, tabAllPost, id
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
