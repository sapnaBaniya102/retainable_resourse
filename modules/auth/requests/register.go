package requests

import (
	"awesomeProject/modules/auth"
	models "awesomeProject/modules/auth/models"
	"awesomeProject/util/crypt"
	"errors"
)

type Register struct {
	FirstName string `json:"firstname" gorm:"firstname" form:"firstname"`
	LastName  string `json:"lastname" gorm:"lastname" form:"lastname"`
	Username  string `json:"username" gorm:"username" form:"username"`
	Email     string `json:"email" gorm:"email" form:"email" validate:"required|validateEmail"`
	Password  string `json:"password" gorm:"password" form:"password" validate:"required"`
}

func (l Register) ValidatePassword(val string) bool {
	return crypt.ValidatePassword(l.Password)
}

func (Register) TableName() string {
	return "user_profiles"
}

func (l *Register) Signup() (*Register, error) {
	userExists, _ := auth.GetUserByEmail(l.Email)
	if userExists != nil {
		return nil, errors.New("User Already Exists")

	}

	profile := models.Profile{

		FirstName: l.FirstName,
		LastName:  l.LastName,
		UserName:  l.Username,
		Email:     l.Email,
	}
	user := models.Users{
		Email: l.Email,
	}
	err := auth.Signup(profile, user, l.Password)
	if err != nil {
		return nil, err
	}
	return l, nil
}
