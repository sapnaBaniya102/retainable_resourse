package models

import (
	"math/big"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

//User represents a user schema
type User struct {
	gorm.Model
	Base
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email" gorm:"unique" validate:"required,email,min=6,max=32"`
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password"`
	Toggle    string `json:"toggle"`
	Errors    map[string]string
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
