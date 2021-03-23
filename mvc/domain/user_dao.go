package domain

import (
	"fmt"
	"net/http"

	"github.com/108356037/goBackendMvc/errors"
)

var (
	users = map[int64]*User{
		123: {9527, "Wei Che", "Tsai", "108356037@nccu.edu.tw"},
	}
)

func GetUser(userId int64) (*User, *errors.ServiceError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &errors.ServiceError{
		Message:    fmt.Sprintf("user_id %d is not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not found",
	}
}
