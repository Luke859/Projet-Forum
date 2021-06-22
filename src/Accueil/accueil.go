package Accueil

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	BDD "../BDD"
)

type PageAccueil struct {
	Post string
	Cmt  string
}

func AccueilPage(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser

	var postsDouble [][]string
	var postOne []PageAccueil
	_, db := BDD.GestionData()

	postsDouble = BDD.GetAllPost(db)
	_, cmtsDouble := BDD.GetAllCmt(db, 1)

	for _, postSync := range postsDouble {
		for _, cmtSync := range cmtsDouble {
			p := PageAccueil{
				Post: postSync[1],
				Cmt:  cmtSync[2],
			}
			postOne = append(postOne, p)
		}

	}

	t, err := template.ParseFiles("static/HTML/layout.html", "static/HTML/Accueil.html", "static/HTML/navbar.html")
	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	t.Execute(w, postOne)
	fmt.Println("Page Accueil ✔️")

}
