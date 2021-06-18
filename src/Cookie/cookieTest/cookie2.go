package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	uuid "github.com/satori/go.uuid"
)

func DrawMenu(w http.ResponseWriter){
    w.Header().Set("Content-Type", "text/html")
    io.WriteString(w, "<a href='/'>HOME <ba><br/>" + "\n")
    io.WriteString(w, "<a href='/readcookie'>Read Cookie <ba><br/>" + "\n")
    io.WriteString(w, "<a href='/writecookie'>Write Cookie <ba><br/>" + "\n")
    io.WriteString(w, "<a href='/deletecookie'>Delete Cookie <ba><br/>" + "\n")

}

func IndexServer(w http.ResponseWriter, req *http.Request) {
    DrawMenu(w)
}

func ReadCookieServer(w http.ResponseWriter, req *http.Request) {

    DrawMenu(w)
    var cookie,err = req.Cookie("CookieName")
    if err == nil {
        var cookievalue = cookie.Value
        io.WriteString(w, "<b>get cookie value is " + cookievalue + "</b>\n")
    }
}

func WriteCookieServer(w http.ResponseWriter, req *http.Request) {
    myuuid, err := uuid.NewV4()
    if err != nil {
        http.SetCookie(w,&http.Cookie{
            Name:       "CookieUUID",
            Value:      myuuid.String(),
            Domain:     "ProjetForum",
            Secure:     true,
            HttpOnly:   true,
        })
    }
    DrawMenu(w)
}

func DeleteCookieServer(w http.ResponseWriter, req *http.Request) {
    cookie := http.Cookie{Name: "CookieName", Path: "/", MaxAge: -1}
    http.SetCookie(w, &cookie)
    DrawMenu(w)

}

func main() {

    http.HandleFunc("/", IndexServer)
    http.HandleFunc("/readcookie", ReadCookieServer)
    http.HandleFunc("/writecookie", WriteCookieServer)
    http.HandleFunc("/deletecookie", DeleteCookieServer)

    fmt.Println("listen on 3000")
    err := http.ListenAndServe(":3000", nil)

    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}