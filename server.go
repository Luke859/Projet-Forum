package main

import (
	// APIartist "Projet-Forum/src/APIartist"
	// APIcarte "Projet-Forum/src/APIlocations"
	accueil "Projet-Forum/src/accueil"
	"net/http"
)

func main() {

	http.HandleFunc("/", accueil.AccueilPage)
	// http.HandleFunc("/groupes", APIartist.ArtistPage)
	// http.HandleFunc("/carte", APIcarte.CartePage)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}
