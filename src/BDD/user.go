package BDD

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
)

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

//////////////////////////////////////// get UUID from User //////////////////////////////////////

func GetUUID_User(username string, db *sql.DB) (int, uuid.UUID) {
	var UUID uuid.UUID

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

func PutUUID(UUID uuid.UUID, pseudo string, db *sql.DB) int {
	statement, err := db.Prepare("UPDATE User SET uuid = ? WHERE (pseudo =?)")
	if err != nil {
		fmt.Println(err)
		fmt.Println("error Prepare new user")
		return (500)
	}
	statement.Exec(UUID, pseudo)
	db.Close()
	return (0)
}

//////////////////////////////////////// Is sesion OK ///////////////////////////////////////////

/*
la function prend la valeur des cookie le nom de l'utilisateur et le pointeur vers la base de donnée.
et retourne un boolean.
*/
func CheckSession(cookie uuid.UUID, username string, db *sql.DB) bool {
	var IsSessionOk bool
	status, uuidBDD := GetUUID_User(username, db)
	if status == 0 {
		if cookie == uuidBDD {
			IsSessionOk = true
		} else {
			IsSessionOk = false
		}
	}
	return IsSessionOk
}
