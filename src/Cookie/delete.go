func DeleteCookieServer(w http.ResponseWriter, req *http.Request) {
    cookie := http.Cookie{Name: "CookieName", Path: "/", MaxAge: -1}
    http.SetCookie(w, &cookie)
    DrawMenu(w)

}