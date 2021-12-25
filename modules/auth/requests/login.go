package requests

import (
	"github.com/gookit/validate"
)

type Login struct {
	Email      string `json:"email" gorm:"email" form:"email" validate:"required|validateEmail"`
	Password   string `json:"password" gorm:"password" form:"password" validate:"required"`
	RememberMe bool   `json:"remember_me" gorm:"remember_me" form:"remember_me"`
}

func (l Login) ValidateEmail(val string) bool {
	validatedEmail := mail.ValidateEmail(l.Email)
	return validatedEmail.Valid && !validatedEmail.Disposable
}

// Messages you can custom validator error messages.
func (l *Login) Messages() map[string]string {
	return validate.MS{
		"required":      "oh! the {field} is required",
		"validateEmail": "Invalid email format",
	}
}

func (Login) TableName() string {
	return "user_profiles"
}
