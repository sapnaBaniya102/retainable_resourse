package models

import "github.com/sujit-baniya/db"

type Profile struct {
	FirstName string `json:"firstname" gorm:"type:varchar(255);column:firstname;"` //nolint:gofmt
	LastName  string `json:"lastname" gorm:"type:varchar(255);column:lastname;"`   //nolint:gofmt
	UserName  string `json:"username" gorm:"type:varchar(255);column:username;"`
	Email     string `validate:"required|uniqueEmail" json:"email" gorm:"type:varchar(255);column:email;unique"`
	Status    string `json:"status" gorm:"status"`
	BaseModel
}

func (u *Profile) Store() error {
	return db.DB.Create(u).Error
}
