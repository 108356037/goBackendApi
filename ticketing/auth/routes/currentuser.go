package routes

import (
	"encoding/json"
	"net/http"

	"github.com/108356037/goticketapp/auth/jwtutils"
	"github.com/108356037/goticketapp/auth/middleware"
)

func CurrentUserHandlerFunc(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("session-cookie")
	if err != nil || c.MaxAge < 0 {
		middleware.JSONError(w, "user not authenticated", http.StatusForbidden)
		return
	}
	payload, err := jwtutils.VerifyUserToken(c.Value)
	if err != nil {
		middleware.JSONError(w, err.Error(), http.StatusForbidden)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(payload)
}
