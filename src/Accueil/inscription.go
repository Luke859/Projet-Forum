package Accueil

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	BDD "../BDD"

	"golang.org/x/crypto/bcrypt"
)

func InscriptionPage(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("static/HTML/layout.html", "static/HTML/inscription.html", "static/HTML/navbar.html")
	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	t.Execute(w, nil)
	fmt.Println("Page inscription ")
}

// Fonction qui récupère le PSEUDO et le MDP du formulaire "inscription"

func GetSign(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	pseudo := r.FormValue("Pseudo")     // pseudo de l'inscription
	password := r.FormValue("Password") // mdp de l'inscription

	fmt.Println(" Identidiant d'Inscription : ", pseudo, "/", password)
	http.Redirect(w, r, "/connexion", http.StatusSeeOther)
	var HashPass = hashPassword(password) // password de l'incription Hashé (Hashpass)
	fmt.Println("Mot de passe Hashé ⌛ :", HashPass)
	statusBDD, db := BDD.GestionData()
	status := BDD.NewUser(pseudo, HashPass, db)
	if status == 0 && statusBDD == 0 {
		fmt.Println("walla")
	} else {
		fmt.Println("Walla il y avait plus de poulet curry")
	}

}

// Hash du mot de passe puis l'afficher dans le terminal
func hashPassword(password string) string {
	var passByte = []byte(password)

	hash, err := bcrypt.GenerateFromPassword(passByte, bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}

// Fonction qui récupère le PSEUDO et le MDP du formulaire "connexion"

func GetSignConnect(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal()
	}
	pseudoconnect := r.FormValue("PseudoConnect")     // pseudo de la connexion
	passwordconnect := r.FormValue("PasswordConnect") // mdp de la connexion

	fmt.Println(" Identifiant de connexion : ", pseudoconnect, "/", passwordconnect)
	http.Redirect(w, r, "/accueil", http.StatusSeeOther)
	fmt.Println()

}

// Verif du mot de passe
func comparePasswords(HashPass string, passwordconnect []byte) bool {

	byteHash := []byte(HashPass)
	err := bcrypt.CompareHashAndPassword(byteHash, passwordconnect)
	if err != nil {
		log.Println(err)
		return false
	}

	return true

}
