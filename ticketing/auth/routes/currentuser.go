// package routes

// import (
// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/108356037/goticketapp/auth/sesscookie"
// 	"net/http"
// 	"github.com/108356037/goticketapp/auth/middleware"
// )

// func currentUserHandlerFunc(w http.ResponseWriter, r *http.Request) {
// 	session, err := sesscookie.Store.Get(r, sesscookie.SessionName)
// 	if err != nil {
// 		middleware.JSONError(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	jwt.
// }