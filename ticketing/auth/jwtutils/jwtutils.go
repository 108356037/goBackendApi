package jwtutils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	JwtSignKey = "asdf"
)

type UserJwt struct {
	UserId             string `json:"userId"`
	Email              string `json:"email"`
	jwt.StandardClaims `json:",inline"`
}

func CreateUserTokenString(userId, email string) (string, error) {
	claims := UserJwt{
		UserId: userId,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(30 * 24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(JwtSignKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func VerifyUserToken(signedToken string) (payload *UserJwt, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&UserJwt{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(JwtSignKey), nil
		})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserJwt)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		return nil, err
	}

	return claims, nil
}
