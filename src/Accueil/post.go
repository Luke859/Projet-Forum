package Accueil

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"../BDD"
)

func PostPage(w http.ResponseWriter, r *http.Request) {

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
	statusPost := BDD.MakePost(TextArea, 1)
	if statusPost == 300 {
		fmt.Println("Inscrit en BDD")
	} else {
		fmt.Println("NON inscrit en BDD")
	}
	http.Redirect(w, r, "/accueil", http.StatusSeeOther)

}
