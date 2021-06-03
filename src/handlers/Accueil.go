<<<<<<< HEAD:src/handlers/Accueil.go
package Home

import (
	"log"
	"net/http"
	"text/template"
)

var t *template.Template
var tErr *template.Template

func AccueilPage(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("static/HTML/layout.html", "static/HTML/accueil.html", "static/HTML/navbar.html")
	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	t.Execute(w, nil)
}
=======
package Home

import (
	"log"
	"net/http"
	"text/template"
)

var t *template.Template
var tErr *template.Template

func AccueilPage(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("HTML/layout.html", "HTML/accueil.html", "HTML/navbar.html")
	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	t.Execute(w, nil)
}
>>>>>>> parent of deb4d99 (Update accueil.go):src/Accueil/accueil.go
