import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

func gestionData() {
	db, err := sql.Open("sqlitte3", "./nomBase.db")
	if err != nil {
		fmt.Print("error ouvertur base")
		return (5)
	}
}

func newUser(pseudo, Hmpd) {
	statement, err := db.Prepare("INSERT INTO User (pseudo, Hmpd) VALUES(?,?)")
	if err != nil {
		fmt.Print("error new user")
		return (5)
	}
	statement.Exec()
}