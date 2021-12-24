package modules

import (
	"math/big"

	"github.com/dgrijalva/jwt-go"
	"github.com/sujit-baniya/db"
	"gorm.io/gorm"
)

//User represents a user schema
type User_Profiles struct {
	gorm.Model
	Base
	Firstname string `json:"firstname" gorm:"type:varchar;not null"`
	Lastname  string `json:"lastname" gorm:"type:varchar;not null"`

	Username string `json:"username" gorm:"type:varchar;unique"`
}
type Users struct {
	gorm.Model
	Base

	UserId uint   `json:"user_id" gorm:"column:user_id"`
	Email  string `json:"email" gorm:"type:varchar;unique" validate:"required,email,min=6,max=32"`
}
type Userscredentials struct {
	gorm.Model
	Base
	UserId   uint   `json:"user_id" gorm:"column:user_id"`
	Password string `json:"password" gorm:"type:varchar;not null"`
}
type Validation struct {
	Code *big.Int
}

// UserError represent the error format for user routes
type UserError struct {
	gorm.Model
	Err      bool   `json:"error"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//Claims represent the structure of the jwt authentication
type Claims struct {
	jwt.StandardClaims
	ID uint `gorm:"primaryKey"`
}

func (u *User_Profiles) Store() error {
	return db.DB.Create(u).Error
}
