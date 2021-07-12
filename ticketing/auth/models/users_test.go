package models_test

import (
	"testing"

	"github.com/108356037/goticketapp/auth/models"
)

var (
	passwd1 = "passwordOne"
	passwd2 = "passwordTwo"
)

func TestCheckPasswordHash1(t *testing.T) {
	hashedpasswd, err := models.HashPassword(passwd1)
	if err != nil {
		t.Errorf(err.Error())
	}

	res := models.CheckPasswordHash(passwd1, hashedpasswd)
	if !res {
		t.Errorf("Hashed pwd %s is not correct from origin passwd %s\n", hashedpasswd, passwd1)
	}
}

func TestCheckPasswordHash2(t *testing.T) {
	hashedpasswd, err := models.HashPassword(passwd1)
	if err != nil {
		t.Errorf(err.Error())
	}

	res := models.CheckPasswordHash(passwd2, hashedpasswd)
	if res {
		t.Errorf("Hashed pwd %s is not correct from origin passwd %s\n", hashedpasswd, passwd1)
	}
}
