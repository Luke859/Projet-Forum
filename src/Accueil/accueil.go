package Accueil

import (
	"log"
	"net/http"
	"text/template"
	"strings"
	BDD "../BDD"
)

type PageAccueil struct {
	Post string
}


func AccueilPage(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	status, db := BDD.GestionData()

	p := PageAccueil{
		Post: strings.Join(BDD.GetAllPost(db), ""),
	}


	t, err := template.ParseFiles("static/HTML/layout.html", "static/HTML/Accueil.html", "static/HTML/navbar.html")
	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	t.Execute(w, p)

}

