package usecase

import (
	"testing"
)

var (
	password = "secret"
)

func TestHashPassword(t *testing.T) {
	t.Log("Raw Password : ", password)
	cases := accountUsecase{}
	hash, err := cases.HashPassword(password)
	t.Log("Hashed Password : ", hash)
	if err != nil {
		t.Error("Error TestHashPassword")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	t.Log("Raw Password : ", password)
	cases := accountUsecase{}
	hash, err := cases.HashPassword(password)
	if err != nil {
		t.Error("Error TestHashPassword")
	}
	match := cases.CheckPasswordHash(password, hash)
	if !match {
		t.Error("Something went wrong")
	}
}
