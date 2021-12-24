package crypt

import "unicode"

var DefaultPasswordValidator = &PasswordValidator{
	AtLeastOneLowerCase: true,
	AtLeastOneUpperCase: true,
	AtLeastOneDigit:     true,
	AtLeastOneSymbol:    true,
	MinLength:           6,
	MaxLength:           32,
}

type PasswordValidator struct {
	AtLeastOneLowerCase bool
	AtLeastOneUpperCase bool
	AtLeastOneDigit     bool
	AtLeastOneSymbol    bool
	MinLength           uint8
	MaxLength           uint8
}

// Validate validates plain password against the rules defined below.
func (p *PasswordValidator) Validate(pass string) bool {
	var (
		upp, low, num, sym bool
		tot                uint8
	)

	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false
		}
	}
	if (p.AtLeastOneLowerCase && !low) ||
		(p.AtLeastOneUpperCase && !upp) ||
		(p.AtLeastOneDigit && !num) ||
		(p.AtLeastOneSymbol && !sym) ||
		(p.MinLength != 0 && tot < p.MinLength) ||
		p.MaxLength != 0 && tot > p.MaxLength {
		return false
	}

	return true
}

func ValidatePassword(pass string) bool {
	return DefaultPasswordValidator.Validate(pass)
}
