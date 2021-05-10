package sesscookie

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

type UserJwt struct {
	UserId             string `json:"userId"`
	Email              string `json:"email"`
	jwt.StandardClaims `json:",inline"`
}

var (
	SessionName = "Cookie-Session"
	sessionKey  = securecookie.GenerateRandomKey(32)
	JwtSignKey  = "asdf"
	Store       = sessions.NewCookieStore(sessionKey)
)
