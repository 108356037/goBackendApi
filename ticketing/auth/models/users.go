package models

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	// probably need id from db
}

func CreateUser(users *User) error {
	return nil
}
