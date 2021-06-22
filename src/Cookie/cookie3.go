package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
)

func main(w http.ResponseWriter, r *http.Request){
	myuuid, err := uuid.NewV4()
	if err == nil{
		fmt.Println("Your UUID is:", myuuid)
	}
	http.SetCookie(w, &http.Cookie{
		Name: "cookieName",
		Value: myuuid.String(),
		Path: "/",
		Expires: time.Now().Add(120*time.Second),
		MaxAge: 86400,
	})
	fmt.Println(myuuid)
	return
}
