package Accueil

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	/*"time"
	guuid"github.com/google/uuid"*/
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

	fmt.Println(" Identidiant d'Inscription : ", pseudo, password)
	http.Redirect(w, r, "/connexion", http.StatusSeeOther)
	var HashPass = hashPassword(password)
	fmt.Println("Mot de passe Hashé ⌛ :", HashPass)

}



// Fonction qui récupère le PSEUDO et le MDP du formulaire "connexion"

func GetSignConnect(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	pseudoconnect := r.FormValue("PseudoConnect")
	passwordconnect := r.FormValue("PasswordConnect")

	/*myuuid := guuid.New()
	expire := time.Now().AddDate(0, 0, 1)
    cookie := http.Cookie{
        Name: "cookieName",
        Value: myuuid.String(),
        Path: "/",
        Expires: expire,
        MaxAge: 86400,
    }
	fmt.Println(myuuid)*/
	

	fmt.Println(" Identifiant de connexion : ", pseudoconnect, passwordconnect)
	//http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/accueil", http.StatusSeeOther)
	
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
