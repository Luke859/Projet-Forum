package Accueil

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	BDD "../BDD"
)

type PageAccueil struct {
	Id_Post int
	Post    string
	Cmt     []string
	Like    int
}

func AccueilPage(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser

	var postsDouble [][]string
	var postOne []PageAccueil
	_, db := BDD.GestionData()

	_, postsDouble = BDD.GetAllPost(db)
	//_, postsDouble = BDD.IsLikedPOST(db, 1)

	for _, postSync := range postsDouble {

		id_post, _ := strconv.Atoi(postSync[0])
		_, cmtsDouble := BDD.GetAllCmt(db, id_post)

		p := PageAccueil{
			Post:    postSync[1],
			Cmt:     make([]string, 0),
			Id_Post: id_post,
			// Like: postSync[1],
		}
		fmt.Println(p.Id_Post)
		for _, cmtSync := range cmtsDouble {
			p.Cmt = append(p.Cmt, cmtSync[2])
		}

		postOne = append(postOne, p)

	}

	t, err := template.ParseFiles("static/HTML/layout.html", "static/HTML/Accueil.html", "static/HTML/navbar.html")
	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	t.Execute(w, postOne)
	fmt.Println("Page Accueil ✔️")

}
