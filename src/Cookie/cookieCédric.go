package main

import (
	"net/http"
	"fmt"
)

type Credentials struct{
	Password string `json:"password"`
	Username string`json:"username"`
}
func Signin(w http.ResponseWriter, r *http.Request){
	var creds Credentials
	//GettheJSONbodyanddecodeintocredentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err!= nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}//GettheexpectedpasswordhashfromDBandcomparehashvalues
	//Createanewrandomsessiontoken
	sessionToken := uuid.NewV4().String()
	//Keepactivesessioninmemory=cacheorDBalongwithusername
	//Finally,wesettheclientcookiefor"session_token"asthesessiontokenwejustgenerated
	//wealsosetanexpirytimeof120seconds,thesameasthecache
	http.SetCookie(w,&http.Cookie{
		Name: "session_token",
		Value: sessionToken,
		Expires: time.Now().Add(120*time.Second),
	})
}

func WelcomePageHandler(w http.ResponseWriter, r *http.Request){
	//Wecanobtainthesessiontokenfromtherequestscookies,whichcomewitheveryrequest
	cookieContent, err := r.Cookie("session_token")
	if err != nil{
		if err == http.ErrNoCookie{
			//Ifthecookieisnotset,returnanunauthorizedstatus
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		//Foranyothertypeoferror,returnabadrequeststatus
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := cookieContent.Value
	//Testifwehaveavalidsessioninservermemory(CacheorDB)
	//IFwehaveavalidsessionactivated
	w.Write([]byte(fmt.Sprintf("Welcome%s!", response)))
	//ELSEdisplayHTTP401UnauthorizedAccessorDefaultPage
}