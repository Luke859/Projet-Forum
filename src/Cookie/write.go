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