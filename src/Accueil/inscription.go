package Accueil

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	BDD "../BDD"

	"golang.org/x/crypto/bcrypt"
	guuid"github.com/google/uuid"
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

/*func CreateCookie(w http.ResponseWriter, r *http.Request){
	myuuid := guuid.New()

	if r.Method == "GET" {
		http.SetCookie(w, &http.Cookie{
			Name: "cookieName",
			Value: myuuid.String(),
			Path: "/",
			Expires: time.Now().Add(120*time.Second),
			MaxAge: 86400,
		})
	}
	fmt.Println(myuuid)
}*/


// Fonction qui récupère le PSEUDO et le MDP du formulaire "connexion"

func GetSignConnect(w http.ResponseWriter, r *http.Request) {
	myuuid := guuid.New()
	err := r.ParseForm()
	if err != nil {
		log.Fatal()
	}
	pseudoconnect := r.FormValue("PseudoConnect")     // pseudo de la connexion
	passwordconnect := r.FormValue("PasswordConnect") // mdp de la connexion

	_, db := BDD.GestionData()
	_, recuphash := BDD.CheckPassword(pseudoconnect, db)
	//_, recupUUID := BDD.GetUUID_User(myuuid, db)
	match := comparePasswords(recuphash, []byte(passwordconnect))

	fmt.Println(" Identifiant de connexion : ", pseudoconnect, "/", passwordconnect)
	fmt.Println("Match:   ", match)
	expire := time.Now().AddDate(0, 0, 1)
	http.SetCookie(w, &http.Cookie{
		Name:       "cookieName",
		Value:      myuuid.String(),
		Path:       "/",
		Domain:     "",
		Expires:    expire,
		RawExpires: "",
		MaxAge:     86400,
		Secure:     true,
		HttpOnly:   true,
		SameSite:   0,
		Raw:        "",
		Unparsed:   []string{},
	})

	fmt.Println(myuuid)
	// http.Redirect(w, r, "/accueil", http.StatusSeeOther)
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
