package models

import "github.com/sujit-baniya/db"

type Users struct {
	Email  string `validate:"required|uniqueEmail" json:"email" gorm:"type:varchar(255);column:email;unique"`
	Status string `json:"status" gorm:"status"`
	UserId uint   `json:"userid" gorm:"userid"`
	db.BaseModel
}

func (u Users) UniqueEmail(val string) bool {
	var user Users
	if err := db.DB.Model(&Users{}).Where(Users{Email: val}).First(&user).Error; err != nil {
		return true
	}
	if user.ID != 0 {
		return false
	}
	return true
}

func (u *Users) Store() error {

	return db.DB.Create(u).Error

}
