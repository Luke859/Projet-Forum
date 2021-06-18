func ReadCookieServer(w http.ResponseWriter, req *http.Request) {

    DrawMenu(w)
    var cookie,err = req.Cookie("CookieName")
    if err == nil {
        var cookievalue = cookie.Value
        io.WriteString(w, "<b>get cookie value is " + cookievalue + "</b>\n")
    }
}