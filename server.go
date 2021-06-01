package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var t *template.Template

func main() {

	t = template.Must(template.ParseFiles("templates/index.html"))
	// Importation des fichiers statiques :
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	// Accès aux pages :
	http.Handle("/", http.NotFoundHandler())
	http.HandleFunc("/accueil", accueil)
	http.Handle("/accueil/", http.NotFoundHandler())
	http.HandleFunc("/accueil-connecter", accueil-connecter)
	http.Handle("/accueil-connecter/", http.NotFoundHandler())
	http.HandleFunc("/connection", connection)
	http.Handle("/connection/", http.NotFoundHandler())
	http.HandleFunc("/inscription", inscription)
	http.Handle("/inscription/", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func accueil(w http.ResponseWriter, req *http.Request) {

	tAccueil, err := template.ParseFiles("templates/Accueil.html")
	if err != nil {
		w.WriteHeader(400)
	}
	tAccueil.Execute(w, nil)
}

func accueilconnecter(w http.ResponseWriter, req *http.Request) {

	tAccueilconnecter, err := template.ParseFiles("templates/Accueil-Connecter.html")
	if err != nil {
		w.WriteHeader(400)
	}
	tAccueilconnecter.Execute(w, nil)
}

func connection(w http.ResponseWriter, req *http.Request) {

	tConnection, err := template.ParseFiles("templates/connection.html")
	if err != nil {
		w.WriteHeader(400)
	}
	tConnection.Execute(w, nil)
}

func inscription(w http.ResponseWriter, req *http.Request) {

	tInscription, err := template.ParseFiles("templates/inscription.html")
	if err != nil {
		w.WriteHeader(400)
	}
	tInscription.Execute(w, nil)
}