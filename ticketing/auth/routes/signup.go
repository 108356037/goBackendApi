package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/108356037/goticketapp/auth/middleware"
	"github.com/108356037/goticketapp/auth/models"
	"github.com/108356037/goticketapp/auth/sesscookie"
	"github.com/dgrijalva/jwt-go"
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
	userJs, err := json.Marshal(user)
	if err != nil {
		middleware.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	claims := sesscookie.UserJwt{
		UserId: strconv.Itoa(userId),
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(sesscookie.JwtSignKey))
	if err != nil {
		middleware.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err := sesscookie.Store.Get(r, sesscookie.SessionName)
	if err != nil {
		middleware.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["authenticated"] = true
	session.Values["jwt"] = signedToken
	if err := session.Save(r, w); err != nil {
		middleware.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(201)
	w.Write(userJs)

}
