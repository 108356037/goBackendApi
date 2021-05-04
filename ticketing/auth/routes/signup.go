package routes

import (
	"encoding/json"
	"net/http"

	"github.com/108356037/goticketapp/auth/middleware"
	"github.com/108356037/goticketapp/auth/models"
	"github.com/golang/gddo/httputil/header"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {

	// returns error if the Content-Type is not "application/json"
	// * can be in common middleware *
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			middleware.JSONError(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	// use http.MaxBytesReader to enforce maximum read of 1MB from the response body
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	user := models.User{}
	err := dec.Decode(&user)

	// below logic is for any error that happens in request-body decode stage
	if err != nil {
		middleware.RequestBodyError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	userId, err := models.CreateUser(&user)
	user.UserId = userId
	if err != nil {
		middleware.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userJs, err := json.Marshal(user)
	if err != nil {
		middleware.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(userJs)

}
