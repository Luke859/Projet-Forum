package main

import (
	guuid"github.com/google/uuid"
	"fmt"
	"time"
	"net/http"
)

func CreateCookie() string{
	myuuid := guuid.New()
	expire := time.Now().AddDate(0, 0, 1)
    cookie := http.Cookie{
        Name: "cookieName",
        Value: myuuid.String(),
        Path: "/",
        Expires: expire,
        MaxAge: 86400,
    }
	fmt.Println(myuuid)
	return ""
}
