package main

import (
	"fmt"
	"net/http"
	"time"
	uuid"github.com/satori/go.uuid"
)

func CreateUUID() string {
	myuuid, err := uuid.NewV4()0
	if err == nil{
	  fmt.Println("Your UUID is:", myuuid)
	  return ""
	}
	return myuuid.String()
}

func WriteCookieServer(w http.ResponseWriter, req *http.Request) {
	uuid := CreateUUID()
    expire := time.Now().AddDate(0, 0, 1)
    cookie := http.Cookie{Name: "testcookiename", Value: uuid, Path: "/", Expires: expire, MaxAge: 86400}
    http.SetCookie(w, &cookie)
}

func main(){
	WriteCookieServer(w http.ResponseWriter, req *http.Request)
}
