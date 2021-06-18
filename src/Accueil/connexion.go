package Accueil

import (
	"log"
	"net/http"
	"text/template"
    "time"
	uuid "github.com/satori/go.uuid"
)

func WriteCookieServer(w http.ResponseWriter, req *http.Request) {
    myuuid, err := uuid.NewV4()
    if err != nil {
        http.SetCookie(w,&http.Cookie{
            Name:       "CookieUUID",
            Value:      myuuid.String(),
            Domain:     "ProjetForum",
            Secure:     true,
            HttpOnly:   true,
        })
    }
}

func ConnexionPage(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("static/HTML/layout.html", "static/HTML/connexion.html", "static/HTML/navbar.html")
	if err != nil {
		log.Fatalf("Template execution: %s", err)
		myuuid, err2 := uuid.NewV4()
		if err2 != nil {
			http.SetCookie(w,&http.Cookie{
				Name:       "CookieUUID",
				Value:      myuuid.String(),
				Domain:     "ProjetForum",
				Secure:     true,
				HttpOnly:   true,
			})
		return
	}
	t.Execute(w, nil)
}
