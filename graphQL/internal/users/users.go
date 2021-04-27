package users

import (
	"golang.org/x/crypto/bcrypt"

	database "graphql/example.com/internal/pkg/db/mysql"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func (user *User) Create() error {
	stmt, err := database.Db.Prepare("INSERT INTO Users(Username,Password) VALUES(?,?)")
	if err != nil {
		return err
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Username, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUserIdByUsername(username string) (int, error) {
	stmt, err := database.Db.Prepare("SELECT ID FROM Users WHERE Username = ?")
	if err != nil {
		return -1, err
	}

	Id := new(int)
	err = stmt.QueryRow(username).Scan(Id)
	if err != nil {
		//log.Fatal(err)
		return -1, err
	}
	return *Id, nil
}
