package Accueil

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func PostPage(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("static/HTML/layout.html", "static/HTML/post.html", "static/HTML/navbar.html")
	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	t.Execute(w, nil)
	fmt.Println("Page Post ⌛")
}

// Recupération du TEXT dans le form TextArea
func GetPostInformation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal()
	}

	TextArea := r.FormValue("text")
	fmt.Println(" Voici le post écrit :", TextArea)
	http.Redirect(w, r, "/accueil", http.StatusSeeOther)

}
