package main

import (
	"net/http"
	"fmt"
	"time"
	"github.com/satori/go.uuid"
)

type Credentials struct{
	Password string `json:"password"`
	Username string`json:"username"`
}
func Signin(w http.ResponseWriter, r *http.Request){
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err!= nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := uuid.NewV4().String()
	http.SetCookie(w,&http.Cookie{
		Name: "session_token",
		Value: sessionToken,
		Expires: time.Now().Add(120*time.Second),
	})
}

func WelcomePageHandler(w http.ResponseWriter, r *http.Request){
	cookieContent, err := r.Cookie("session_token")
	if err != nil{
		if err == http.ErrNoCookie{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := cookieContent.Value
	w.Write([]byte(fmt.Sprintf("Welcome%s!", response)))
}