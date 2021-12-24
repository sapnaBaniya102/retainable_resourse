package auth

import (
	accountModel "awesomeProject/modules/auth/models"
	authModel "awesomeProject/modules/auth/models"

	"github.com/sujit-baniya/db"
)

type LoginResponse struct {
	User     *authModel.Users
	Profile  accountModel.Profile
	Password string
	Account  string
}

func GetUserWithProfileByEmail(email string) (*LoginResponse, error) {
	var loginResponse LoginResponse
	user, err := GetUserByEmail(email)
	if err != nil {
		return &loginResponse, err
	}

	loginResponse.User = user

	return &loginResponse, nil
}
func GetUserByEmail(email string) (*authModel.Users, error) {
	var user authModel.Users
	err := db.DB.Where(&authModel.Users{Email: email}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
