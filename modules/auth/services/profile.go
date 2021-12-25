package services

import (
	"awesomeProject/modules/auth/models"

	"github.com/sujit-baniya/db"
)

func GetProfileByEmail(email string) (profile models.Profile, err error) {
	err = db.DB.Preload("Accounts").Find(&profile, "email = ?", email).Error
	// @TODO - feature to allow switching between accounts
	return
}
