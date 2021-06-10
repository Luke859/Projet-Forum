package Accueil

import (
	"log"
	"net/http"
	"text/template"
)

func ConnexionPage(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("static/HTML/layout.html", "static/HTML/connexion.html", "static/HTML/navbar.html")
	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	t.Execute(w, nil)

}
