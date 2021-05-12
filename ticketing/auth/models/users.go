package models

import (
	"errors"

	database "github.com/108356037/goticketapp/auth/internal/pkg/db/postgres"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserId   int    `json:"userId"`
}

type ResponseUser struct {
	Email  string `json:"email"`
	UserId int    `json:"userId"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateUser(users *User) (int, error) {
	stmt, err := database.Db.Prepare("INSERT INTO users(email,password) VALUES($1,$2) RETURNING user_id")
	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	hashedpasswd, err := HashPassword(users.Password)
	if err != nil {
		return -1, err
	}

	var userId int
	err = stmt.QueryRow(users.Email, hashedpasswd).Scan(&userId)
	if err != nil {
		return -1, err
	}

	return userId, nil
}

func GetUser(email, password string) (*User, error) {
	stmt, err := database.Db.Prepare("SELECT user_id,email,password FROM users WHERE email=$1")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var user User
	err = stmt.QueryRow(email).Scan(&user.UserId, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	if CheckPasswordHash(password, user.Password) {
		return &user, nil
	}

	return nil, errors.New("incorrect password or email")

}
