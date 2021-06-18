package main

import (
	"net/http"

	Accueil "./src/Accueil"
)

func main() {

    http.HandleFunc("/", IndexServer)
    http.HandleFunc("/readcookie", ReadCookieServer)
    http.HandleFunc("/writecookie", WriteCookieServer)
    http.HandleFunc("/deletecookie", DeleteCookieServer)

	http.ListenAndServe(":8080", nil)
}