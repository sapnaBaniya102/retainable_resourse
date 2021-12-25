package auth

import (
	"awesomeProject/modules/auth/models"
	accountModel "awesomeProject/modules/auth/models"
	authModel "awesomeProject/modules/auth/models"
	"errors"

	"github.com/sujit-baniya/db"
)

type LoginResponse struct {
	User     *authModel.Users
	Profile  accountModel.Profile
	Password string
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

func GetVerifiedUserByEmail(email string) (*LoginResponse, error) {
	var loginResponse LoginResponse
	user, err := GetUserByEmail(email)
	if err != nil {
		return &loginResponse, err
	}
	profile, err := GetProfileByID(user.ID)
	if err != nil {
		return &loginResponse, err
	}
	if !profile.EmailVerified {
		return &loginResponse, errors.New("Email not verified")
	}
	password, err := GetPasswordByUserID(user.ID)
	loginResponse.User = user
	loginResponse.Password = password
	loginResponse.Profile = profile
	return &loginResponse, nil
}
func GetProfileByID(id uint) (profile models.Profile, err error) {
	err = db.DB.Find(&profile, "id = ?", id).Error

	return
}

func GetPasswordByUserID(userID uint) (string, error) {
	var credential authModel.Credential
	err := db.DB.Find(&credential, "user_id = ?", userID).Error
	if err != nil {
		return "", err
	}
	return string(credential.Password), nil
}
