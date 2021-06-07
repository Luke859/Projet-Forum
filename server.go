package main

import (
	"net/http"

	Accueil "./src/Accueil"
)

func main() {

	http.HandleFunc("/", Accueil.AccueilPage)
	http.HandleFunc("/connexion", Accueil.ConnexionPage)
	http.HandleFunc("/inscription", Accueil.InscriptionPage)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}
