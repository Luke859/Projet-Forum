package main

import (
	"fmt"
	"net/http"
	uuid"github.com/satori/go.uuid"
)

func main(w http.ResponseWriter, r *http.Request) {
  myuuid, err := uuid.NewV4()
  if err == nil{
	fmt.Println("Your UUID is:", myuuid)
	http.SetCookie(w, &http.Cookie{
    	Name:       "CookieUUID",
    	Value:      myuuid.String(),
    	Domain:     "ProjetForum",
    	Secure:     true,
    	HttpOnly:   true,
    })
	w.WriteHeader(http.StatusBadRequest)
  }
}
