package routes

import (
	"net/http"

	"github.com/108356037/goticketapp/auth/middleware"
)

func SignOutHandlerFunc(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session-cookie")
	if err != nil {
		middleware.JSONError(w, "no user logined", http.StatusForbidden)
		return
	}
	cookie := http.Cookie{
		Name:   "session-cookie",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(200)
	w.Write([]byte(""))
}
