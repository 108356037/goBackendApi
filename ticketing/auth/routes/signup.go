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

func SignUpHandlerFunc(w http.ResponseWriter, r *http.Request) {

	// use http.MaxBytesReader to enforce maximum read of 1MB from the response body
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

	userId, err := models.CreateUser(&user)
	if err != nil {
		middleware.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.UserId = userId

	signedToken, err := jwtutils.CreateUserTokenString(strconv.Itoa(user.UserId), user.Email)
	if err != nil {
		middleware.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// claims := jwtutils.UserJwt{
	// 	UserId: strconv.Itoa(userId),
	// 	Email:  user.Email,
	// 	StandardClaims: jwt.StandardClaims{
	// 		IssuedAt:  time.Now().Unix(),
	// 		ExpiresAt: time.Now().Add(30 * 24 * time.Hour).Unix(),
	// 	},
	// }

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// signedToken, err := token.SignedString([]byte(jwtutils.JwtSignKey))
	// if err != nil {
	// 	middleware.JSONError(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	cookie := http.Cookie{
		Name:     "session-cookie",
		Value:    signedToken,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(user)
}
