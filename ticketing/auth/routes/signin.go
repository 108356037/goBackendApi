package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/108356037/goticketapp/auth/jwtutils"
	"github.com/108356037/goticketapp/auth/middleware"
	"github.com/108356037/goticketapp/auth/models"
)

func SignInHandlerFunc(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	user := models.User{}
	err := dec.Decode(&user)

	if err != nil {
		middleware.RequestBodyError(w, err)
		return
	}

	existingUser, err := models.GetUser(user.Email, user.Password)
	if err != nil {
		middleware.JSONError(w, err.Error(), http.StatusForbidden)
		return
	}

	signedToken, err := jwtutils.CreateUserTokenString(strconv.Itoa(existingUser.UserId), existingUser.Email)
	if err != nil {
		middleware.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:     "session-cookie",
		Value:    signedToken,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&models.ResponseUser{
		UserId: existingUser.UserId,
		Email:  existingUser.Email,
	})
}
