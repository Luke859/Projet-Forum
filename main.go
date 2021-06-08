package main

/*
import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

struct type Catery {

}

func main(){
	db, err := sql.Open("sqlite3", "./cheminDeVotreBase.db")
	if err != nil {
		fmt.Println("Could not open database")
		return
	}
	entries, err := db.Query("SELECT (Name, Color) FROM Category LIMIT 10")
	if err != nil {
		fmt.Println("Could not Query database")
		return
	}
	var name string
	var color string

	for entries.Next() {
		entries.Scan(&name, &color)
		fmt.Println(name + " " + color)
	}

	var cat Category{

	}
}

*/
