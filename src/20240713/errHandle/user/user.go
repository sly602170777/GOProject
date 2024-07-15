package user

import (
	"errors"
)

type User struct {
	Name     string
	Password string
	Age      int
}

func QueryUser(name string) (string, error) {
	if name == "jack" {
		return "jack123", nil
	}
	return "", errors.New("user not found this user")
}
