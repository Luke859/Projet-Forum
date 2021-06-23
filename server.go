package main

import (
	"net/http"

	Accueil "./src/Accueil"
)

func main() {

	http.HandleFunc("/", Accueil.AccueilPage)
	http.HandleFunc("/connexion", Accueil.ConnexionPage)
	http.HandleFunc("/inscription", Accueil.InscriptionPage)
	http.HandleFunc("/post", Accueil.PostPage)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/sent", Accueil.GetSign)
	http.HandleFunc("/sentConnect", Accueil.GetSignConnect)
	http.HandleFunc("/sentText", Accueil.GetPostInformation)
	http.HandleFunc("/sentCmt", Accueil.GetCmtInformation)
	// http.HandleFunc("/sentLike", Accueil.GetLikesInformation)

	http.ListenAndServe(":8080", nil)

}
