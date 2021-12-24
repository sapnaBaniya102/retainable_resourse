package util

import (
	"awesomeProject/modules/auth/models"
	"crypto/rand"
	"fmt"
	"log"
	"net/smtp"
	"regexp"
	"runtime"
	"strings"
	"sync"

	valid "github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// IsEmpty checks if a string is empty
func IsEmpty(str string) (bool, string) {
	if valid.HasWhitespaceOnly(str) && str != "" {
		return true, "Must not be empty"
	}

	return false, ""
}

type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    string
}

var p, _ = rand.Prime(rand.Reader, 16)
var rxEmail = regexp.MustCompile(".+@.+\\..+")

// var userProfile models.User_Profiles
// var credential models.Userscredentials
var user models.Users

//ValidateRegister func validates the body of the user for registration

func process(wg *sync.WaitGroup, intChannel <-chan int, resultChannel chan<- int) {
	defer wg.Done()
	for j := range intChannel {
		resultChannel <- j
	}
}

// func Validation(c *fiber.Ctx) error {

// 	// user.Errors = make(map[string]string)
// 	if err := c.BodyParser(&userProfile); err != nil {

// 		log.Fatal(err)
// 		mp := fiber.Map{
// 			"message": err.Error(),
// 		}

// 		return flash.WithError(c, mp).Redirect("/register")
// 	}
// 	if err := c.BodyParser(&credential); err != nil {

// 		log.Fatal(err)
// 		mp := fiber.Map{
// 			"message": err.Error(),
// 		}

// 		return flash.WithError(c, mp).Redirect("/register")
// 	}
// 	if err := c.BodyParser(&user); err != nil {

// 		log.Fatal(err)
// 		mp := fiber.Map{
// 			"message": err.Error(),
// 		}

// 		return flash.WithError(c, mp).Redirect("/register")
// 	}
// 	match := rxEmail.Match([]byte(user.Email))
// 	if user.Email == "" || !match {

// 		mp := fiber.Map{
// 			"error":   true,
// 			"message": "Please enter a valid email address",
// 		}
// 		return flash.WithError(c, mp).Redirect("/register")

// 		// user.Errors["Email"] = "Please enter a valid email address"
// 	}
// 	if userProfile.Username == "" {
// 		mp := fiber.Map{
// 			"error":   true,
// 			"message": "Please enter a valid username",
// 		}
// 		return flash.WithError(c, mp).Redirect("/register")
// 		// user.Errors["Username"] = "please enter a valid username"
// 	}

// 	if credential.Password == "" || len(credential.Password) < 8 {
// 		mp := fiber.Map{
// 			"error":   true,
// 			"message": "Password must be greater than 8",
// 		}
// 		return flash.WithError(c, mp).Redirect("/register")
// 		// user.Errors["Password"] = "Password must be greater than 8"
// 	}
// 	// Get first matched record
// 	// err := db.DB.Where("email = ?", user.Email).First(&user)
// 	// if err != nil {
// 	// 	mp := fiber.Map{
// 	// 		"error":   true,
// 	// 		"message": "Email already exist",
// 	// 	}
// 	// 	return flash.WithError(c, mp).Redirect("/register")
// 	// }

// 	//creating hashed password
// 	bytes, err1 := HashPassword(credential.Password)
// 	if err1 != nil {
// 		fmt.Println("error with hashing")
// 	}
// 	credential.Password = bytes
// 	// SendEmail(user)
// 	// Insert Employee into database

// 	success := fiber.Map{
// 		"success": true,
// 		"message": "success",
// 	}
// 	return flash.WithSuccess(c, success).Redirect("/validation")
// }

// func CodeVerification(c *fiber.Ctx) error {
// 	code := new(models.Validation)

// 	if err := c.BodyParser(code); err != nil {
// 		panic(err)
// 	}

// 	if code.Code.String() == p.String() {
// 		mp := fiber.Map{
// 			"success": true,
// 			"message": "success",
// 		}

// 		err := db.DB.Create(&userProfile).Error
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		user.UserId= userProfile.Base.ID
// 		err1 := db.DB.Create(&user).Error
// 		if err1 != nil {
// 			return err1
// 		}
// 		err2 := db.DB.Create(&credential).Error
// 		if err2 != nil {
// 			fmt.Println(err2)
// 		}
// 		return flash.WithSuccess(c, mp).Redirect("/success")

// 	} else {
// 		success := fiber.Map{
// 			"error":   true,
// 			"message": "error",
// 		}
// 		return flash.WithError(c, success).Redirect("/validation")
// 	}

// }

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func SendEmail(c *fiber.Ctx, emailData string) error {
	//  send mail
	sender := "skyrootmam123@gmail.com"

	to := []string{
		emailData,
	}

	p, _ = rand.Prime(rand.Reader, 16)

	username := "02b77ac5d36a55"
	password := "198c9d08b45b98"

	subject := "Simple HTML mail"
	body := `<p>To verify your email, Please enter the below given 5 digit code in the</p><br>` + p.String()

	request := Mail{
		Sender:  sender,
		To:      to,
		Subject: subject,
		Body:    body,
	}

	addr := "smtp.mailtrap.io:2525"
	host := "smtp.mailtrap.io"

	var wg sync.WaitGroup
	// make our channels for communicating work and resultChannel
	intChannel := make(chan int) // 100 was chosen arbitrarily
	resultChannel := make(chan int)
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			process(&wg, intChannel, resultChannel)
		}()
	}
	go func() {
		defer close(resultChannel)
		wg.Wait()
	}()
	// start sending intChannel
	go func() {
		defer close(intChannel)
		msg := BuildMessage(request)
		auth := smtp.PlainAuth("", username, password, host)
		err := smtp.SendMail(addr, auth, sender, to, []byte(msg))

		if err != nil {
			log.Fatal(err)
		}
		log.Println("send email")
		result := 1
		intChannel <- result
	}()
	// read all the resultChannel
	for r := range resultChannel {
		if r == 1 {
			fmt.Println("email send successfull")
		}
	}
	return c.JSON("test")
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}
