func Get(w http.ResponseWriter, r *http.Request) *Session {
	var (
		err    error
		cookie *http.Cookie
		ses    *Session
		ok     bool
		ses_id string
	)

	cookie, err = r.Cookie(sid)
	if err != nil {
		ses_id, ses = registry.new()

		cookie = &http.Cookie{Name: sid, Value: ses_id, Path: "/", HttpOnly: true, MaxAge: int(maxlifetime.Seconds())}
		http.SetCookie(w, cookie)

		return ses
	}

	ses, ok = registry.get(cookie.Value)
	if !ok {
		ses = registry.create(cookie.Value)

		cookie.MaxAge = int(maxlifetime.Seconds())
		cookie.Path = "/"
		http.SetCookie(w, cookie)

		return ses
	}

	cookie.MaxAge = int(maxlifetime.Seconds())
	http.SetCookie(w, cookie)

	return ses
}