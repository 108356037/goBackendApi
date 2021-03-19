package domain

import (
	"errors"
	"fmt"
)

var (
	users = map[int64]*User{
		123: {9527, "Tsai", "wei", "108356037@nccu.edu.tw"},
	}
)

func GetUser(userId int64) (*User, error) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, errors.New(fmt.Sprintf("user_id %d is not found\n", userId))
}
