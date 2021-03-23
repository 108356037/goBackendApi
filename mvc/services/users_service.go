package services

import (
	"github.com/108356037/goBackendMvc/errors"

	"github.com/108356037/goBackendMvc/domain"
)

func GetUser(userId int64) (*domain.User, *errors.ServiceError) {
	return domain.GetUser(userId)
}
