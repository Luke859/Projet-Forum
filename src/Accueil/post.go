package Accueil

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"Forum/src/BDD"
)

func PostPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/HTML/layout.html", "static/HTML/post.html", "static/HTML/navbar.html")
	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	t.Execute(w, nil)
	fmt.Println("Page Post ⌛")
}

// Recupération du TEXT dans le form TextArea
func GetPostInformation(w http.ResponseWriter, r *http.Request) {
	localUUID := RecupValueCookie(r)
	_, db := BDD.GestionData()
	statuErr, username := BDD.GetUserByUUID(localUUID, db)
	IdErr, IdUser := BDD.GetId_User(username, db)
	if statuErr == 0 && IdErr == 0 {
		err := r.ParseForm()
		if err != nil {
			log.Fatal()
		}
		TextArea := r.FormValue("text")
		statusPost := BDD.MakePost(TextArea, IdUser)
		if statusPost == 300 {
			fmt.Println("Post envoyer : ", TextArea)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			fmt.Println("Erreur envoie post")
		}
	} else {
		http.Redirect(w, r, "/inscription", http.StatusSeeOther)
	}
}

func GetCmtInformation(w http.ResponseWriter, r *http.Request) {
	localUUID := RecupValueCookie(r)
	_, db := BDD.GestionData()
	statuErr, username := BDD.GetUserByUUID(localUUID, db)
	IdErr, IdUser := BDD.GetId_User(username, db)
	if statuErr == 0 && IdErr == 0 {
		err := r.ParseForm()
		if err != nil {
			log.Fatal()
		}

		id_post, _ := strconv.Atoi(r.FormValue("idPost"))

		CmtArea := r.FormValue("cmt")
		fmt.Println(" Voici le commentaire écrit :", CmtArea)

		statusCmt := BDD.MakeCmt(IdUser, id_post, CmtArea, db)

		if statusCmt == 0 {
			fmt.Println("Commentaire envoyer")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			fmt.Println("Erreur envoie commentaire")
		}
	} else {
		http.Redirect(w, r, "/inscription", http.StatusSeeOther)
	}
}
