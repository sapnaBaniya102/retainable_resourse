package auth

import (
	"awesomeProject/modules/auth/models"
	"awesomeProject/util/crypt"

	"github.com/sujit-baniya/db"
)

func Signup(profile models.Profile, user models.Users, password string) error {
	err := profile.Store()
	if err != nil {
		return err
	}

	err = user.Store()
	if err != nil {
		return err
	}
	err = StorePassword(user.ID, password)
	if err != nil {
		return err
	}
	return nil
}

func StorePassword(userID uint, rawPassword string) error {
	var passCredentials models.Credential
	password, _ := crypt.CreateHash(rawPassword)
	passCredentials.Password = []byte(password)
	passCredentials.UserID = userID
	return db.DB.Create(&passCredentials).Error
}
