package Accueil

import (
	"log"
	"net/http"
	"strconv"
	"html/template"

	"Forum/src/BDD"
)

type PageAccueil struct {
	Id_Post int
	Post    string
	Cmt     []string
	User    []string
	Like    int
}

// AccueilPage gère l'affichage de la page d'accueil avec tous les posts
func AccueilPage(w http.ResponseWriter, r *http.Request) {
	// Ouverture de la base de données
	status, db := BDD.GestionData()
	if status != 0 {
		log.Printf("Erreur ouverture BDD: status=%d", status)
		http.Error(w, "Erreur interne", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Récupération de tous les posts
	status, postsRaw := BDD.GetAllPost(db)
	if status != 0 {
		log.Printf("Erreur GetAllPost: status=%d", status)
		http.Error(w, "Erreur interne", http.StatusInternalServerError)
		return
	}
	log.Printf("Nombre de posts récupérés : %d", len(postsRaw))

	// Construction du modèle pour le template
	var posts []PageAccueil
	for _, entry := range postsRaw {
		idPost, err := strconv.Atoi(entry[0])
		if err != nil {
			log.Printf("Erreur conversion Id_post: %v", err)
			continue
		}

		// Récupérer les commentaires du post
		cStatus, cmtsRaw := BDD.GetAllCmt(db, idPost)
		if cStatus != 0 {
			log.Printf("Erreur GetAllCmt pour post %d: status=%d", idPost, cStatus)
		}
		comments := make([]string, 0, len(cmtsRaw))
		for _, c := range cmtsRaw {
			if len(c) > 2 {
				comments = append(comments, c[2])
			}
		}

		// Initialisation du slice User pour éviter l'erreur {{.User}} dans le template
		posts = append(posts, PageAccueil{
			Id_Post: idPost,
			Post:    entry[1],
			Cmt:     comments,
			User:    make([]string, 0),
			Like:    0,
		})
	}

	// Parsing des templates (layout, navbar puis body)
	tmpl, err := template.ParseFiles(
		"static/HTML/layout.html",
		"static/HTML/navbar.html",
		"static/HTML/Accueil.html",
	)
	if err != nil {
		log.Printf("Erreur ParseFiles template: %v", err)
		http.Error(w, "Erreur interne template", http.StatusInternalServerError)
		return
	}

	// Exécution du template avec les données
	if err := tmpl.Execute(w, posts); err != nil {
		log.Printf("Erreur Execute template: %v", err)
	}

	log.Println("Page Accueil ✔️")
}
