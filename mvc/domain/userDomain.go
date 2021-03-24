package domain

import (
	"github.com/108356037/goBackendMvc/database"
)

var (
	tempusers = map[string]*User{
		"123": {9527, "Wei Che", "Tsai", "108356037@nccu.edu.tw", "LA Lakers"},
	}
)

func GetUserById(userId string) (*User, error) {
	sqlStatement := `Select * From tbl_userinfo WHERE user_id=$1`
	user := User{}
	if err := database.DB.QueryRow(sqlStatement, userId).Scan(
		&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Team); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAllUsers() ([]*User, error) {
	rows, err := database.DB.Query("SELECT * FROM tbl_userinfo")
	if err != nil {
		return nil, err
	}

	users := make([]*User, 0)
	user := User{}

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Team); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
