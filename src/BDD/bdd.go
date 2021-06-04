package main

/*
changer le nom de la base ln 17, 58
changer le nom de la variable HMpd ln 60
msg d'erreur print dans le terminal
*/

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

/*///////////////////////////////////recuperation de la base de donnée ///////////////////////////*/

func gestionData() int {
	db, err := sql.Open("sqlitte3", "./nomBase.db")
	if err != nil {
		fmt.Print("error ouvertur base")
		return (500)
	}
	pseudo := ""
	Hmpd := ""
	statusNewUser := newUser(pseudo, Hmpd, db)
	if statusNewUser == 0 {
		return (0)
	} else {
		fmt.Print("probleme --> creation user")
		return (500)
	}
}

/*/////////////////////////////////////////////////////creation d'un nouvelle identifiant///////////////////////////////////////////////*/

func newUser(pseudo string, Hmpd string, db *sql.DB) int {
	statement, err := db.Prepare("INSERT INTO User (pseudo, Hmpd) VALUES(?,?)")
	if err != nil {
		fmt.Print("error new user")
		return (5)
	}
	statement.Exec(pseudo, Hmpd)
	return (0)
}

/*///////////////////////////////////////////////////////////verification de identifiant///////////////////////////////////////////////////////////////////////////////*/

func checkUser(pseudo string, Hmpd string, db *sql.DB) int {
	ctx := context.Background()
	var tabUser []string

	tabUser = append(tabUser, pseudo)
	tabUser = append(tabUser, Hmpd)

	err := db.PingContext(ctx)
	if err != nil {
		fmt.Println("base donée non existante")
		return 500
	}

	tsql := fmt.Sprintf("SELECT name, Hmpd FROM NOMDEBASE.User") //il faudra metre le bon nom de la base &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&

	/*rows, err2 := db.QueryContext(ctx, tsql)
	if err2 != nil {
		fmt.Println("error Querry")
		return 500
	}

	defer rows.Close()

	for rows.Next() {
		/*if pseudo == rows. {
		}
	}*/
	tabbd := strings.Split(tsql, " ")
	for i := 0; i < 4; i++ {
		if strings.Compare(tabUser[i], tabbd[i]) != 0 {
			fmt.Println("invalide username or password")
			return 401
		}
	}
	return 0
}
