package services

import (
	"github.com/108356037/goBackendMvc/domain"
)

func GetUserById(userId string) (*domain.User, error) {
	return domain.GetUserById(userId)
}

func GetAllUsers() ([]*domain.User, error) {
	return domain.GetAllUsers()
}
