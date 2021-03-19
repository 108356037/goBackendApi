package services

import (
	"github.com/108356037/goBackendMvc/domain"
)

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}
