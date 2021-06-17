package Accueil

import (
	"log"
	"net/http"
	"text/template"
	"io"
    "fmt"
    "time"
)

func WriteCookieServer(w http.ResponseWriter, req *http.Request) {
    expire := time.Now().AddDate(0, 0, 1)
    cookie := http.Cookie{Name: "testcookiename", Value: "testcookievalue", Path: "/", Expires: expire, MaxAge: 86400}
    http.SetCookie(w, &cookie)
}

func ConnexionPage(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("static/HTML/layout.html", "static/HTML/connexion.html", "static/HTML/navbar.html")
	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	t.Execute(w, nil)
	WriteCookieServer(w)

}
