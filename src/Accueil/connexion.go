package Accueil

import (
	"log"
	"net/http"
	"text/template"
    "time"
	uuid "github.com/satori/go.uuid"
)

func ConnexionPage(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("static/HTML/layout.html", "static/HTML/connexion.html", "static/HTML/navbar.html")
	if err != nil {
		log.Fatalf("Template execution: %s", err)
		/*myuuid, _ := uuid.NewV4()
		expire := time.Now().AddDate(0, 0, 1)
		cookie := http.Cookie{
			Name: "testcookiename",
			Value: myuuid.String(),
			Path: "/",
			Expires: expire,
			MaxAge: 86400,
			Secure:     true,
            HttpOnly:   true,
		}*/

    	http.SetCookie(w, &cookie)
		t.Execute(w, nil)
	}
}
