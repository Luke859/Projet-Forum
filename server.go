package main

import (
	"net/http"

	accueil "./src/handlers"
)

func main() {

	http.HandleFunc("/", accueil.AccueilPage)
	// http.HandleFunc("/groupes", APIartist.ArtistPage)
	// http.HandleFunc("/carte", APIcarte.CartePage)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}
