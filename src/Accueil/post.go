package Accueil

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	BDD "../BDD"
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
			fmt.Println("Walla")
		} else {
			fmt.Println("WAlla il y avait plus de poulet curry")
		}
		http.Redirect(w, r, "/accueil", http.StatusSeeOther)
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

		CmtArea := r.FormValue("cmt")
		fmt.Println(" Voici le commentaire écrit :", CmtArea)
		statusCmt := BDD.MakeCmt(IdUser, 1, CmtArea, db)
		if statusCmt == 300 {
			fmt.Println("Walla")
		} else {
			fmt.Println("WAlla il y avait plus de poulet curry")
		}
		http.Redirect(w, r, "/accueil", http.StatusSeeOther)
	}
}

func GetLikesInformation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal()
	}
	_, db := BDD.GestionData()

	LikeArea, _ := strconv.Atoi(r.FormValue("like"))
	fmt.Println(" Voici le nombre de likes écrit :", LikeArea)
	statusLikes := BDD.CreateLike(1, 1, LikeArea, db)
	if statusLikes == 300 {
		fmt.Println("Walla")
	} else {
		fmt.Println("WAlla il y avait plus de poulet curry")
	}
	http.Redirect(w, r, "/accueil", http.StatusSeeOther)

}
