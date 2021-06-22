package Accueil

import (
	"fmt"
	"net/http"
	guuid"github.com/google/uuid"
	"time"
)

func CreateCookie(w http.ResponseWriter, r *http.Request){
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
	http.Redirect(w, r, "/accueil", http.StatusSeeOther)
}