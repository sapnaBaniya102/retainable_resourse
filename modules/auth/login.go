package auth

import (
	"awesomeProject/util/crypt"
	"errors"
)

func CheckLogin(email string, password string) (*LoginResponse, error) {
	loginResponse, err := GetVerifiedUserByEmail(email)
	if err != nil {
		return nil, err
	}
	match, err := crypt.MatchHash(password, loginResponse.Password)
	if !match {
		return nil, errors.New("invalid Username or Password")
	}
	return loginResponse, nil
}
