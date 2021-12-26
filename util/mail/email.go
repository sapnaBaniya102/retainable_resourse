package mail

import (
	"errors"
	"net"
	"regexp"
	"strings"
)

type Email struct {
	Email      string `json:"email"`
	domain     string
	Disposable bool      `json:"is_disposable"`
	MxError    string    `json:"mx_error,omitempty"`
	NsError    string    `json:"ns_error,omitempty"`
	HostError  string    `json:"host_error,omitempty"`
	IpError    string    `json:"ip_error,omitempty"`
	Error      string    `json:"error,omitempty"`
	Valid      bool      `json:"is_valid"`
	mx         []*net.MX `json:"-"`
}

const (
	emptyString string = ""
)

var (
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func (e *Email) IsValid() {
	e.Valid = true
	e.ValidateFormat()

	// e.ValidateDomainRecords()
	// e.ValidateHostAndUser("smtp-relay.sendinblue.com", "info@verishore.com", e.mx)
}

//ValidateFormat - validates an email address meets rfc 822 format via a regex
func (e *Email) ValidateFormat() {
	_, domain, err := validateFormatAndSplit(e.Email)
	if err != nil {
		e.Valid = false
		e.Error = err.Error()
	}
	e.domain = domain
}
func GetDomainOfEmail(email string) string {
	i := strings.LastIndexByte(email, '@')
	return email[i+1:]
}

func validateFormatAndSplit(email string) (username string, domain string, err error) {
	if len(email) < 6 || len(email) > 254 {
		return emptyString, emptyString, errors.New("Invalid Email Format")
	}

	// Regex matches as per rfc 822 https://tools.ietf.org/html/rfc822
	if !emailRegexp.MatchString(email) {
		return emptyString, emptyString, errors.New("Invalid Email Format")
	}

	i := strings.LastIndexByte(email, '@')
	username = email[:i]
	domain = email[i+1:]

	if len(username) > 64 {
		return emptyString, emptyString, errors.New("Invalid Email Format")
	}

	return username, domain, nil
}

func ValidateEmail(email string) Email {
	e := Email{Email: email}
	e.IsValid()
	return e
}
