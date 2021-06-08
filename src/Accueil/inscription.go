package Accueil

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

func InscriptionPage(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("static/HTML/layout.html", "static/HTML/inscription.html", "static/HTML/navbar.html")
	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	t.Execute(w, nil)
	fmt.Println("Page inscription ⌛")
}

// Fonction qui récupère le PSEUDO et le MDP du formulaire "inscription"

func GetSign(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	pseudo := r.FormValue("Pseudo")
	password := r.FormValue("Password")

	fmt.Println(pseudo, password)
	http.Redirect(w, r, "/connexion", http.StatusSeeOther)
	var HashPass = hashPassword(password)
	fmt.Println(HashPass)
}

func hashPassword(password string) string {
	var passByte = []byte(password)

	hash, err := bcrypt.GenerateFromPassword(passByte, bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}
