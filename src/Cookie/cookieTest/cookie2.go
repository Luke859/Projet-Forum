package main

import (
    "io"
    "net/http"
    "log"
    "fmt"
    "time"
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

    var cookie,err = req.Cookie("testcookiename")
    if err == nil {
        var cookievalue = cookie.Value
        io.WriteString(w, "<b>get cookie value is " + cookievalue + "</b>\n")
    }

}

func WriteCookieServer(w http.ResponseWriter, req *http.Request) {

    myuuid, _ := uuid.NewV4()
    expire := time.Now().AddDate(0, 0, 1)
    cookie := http.Cookie{Name: "testcookiename", Value: myuuid.String(), Path: "/", Expires: expire, MaxAge: 86400}

    http.SetCookie(w, &cookie)

    DrawMenu(w)

}


func DeleteCookieServer(w http.ResponseWriter, req *http.Request) {
    cookie := http.Cookie{Name: "testcookiename", Path: "/", MaxAge: -1}
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