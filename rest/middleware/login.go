package middlewares

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"verify/modules/auth"
	"verify/modules/auth/models"
	"verify/modules/auth/requests"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"github.com/sujit-baniya/crypt"
	"github.com/sujit-baniya/db"
	"github.com/sujit-baniya/flash"
	"github.com/sujit-baniya/mail"
	"github.com/sujit-baniya/session"
)

func ValidateLoginPost(redirectTo string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var login requests.Login

		if err := c.BodyParser(&login); err != nil {
			mp := fiber.Map{
				"message": err.Error(),
			}
			return flash.WithError(c, mp).Redirect(redirectTo)
		}
		v := validate.Struct(login)
		if !v.Validate() {
			mp := fiber.Map{
				"message": v.Errors.One(),
			}
			return flash.WithError(c, mp).Redirect(redirectTo)
		}
		loginResponse, err := auth.CheckLogin(login.Email, login.Password)
		if err != nil {
			mp := fiber.Map{
				"message": err.Error(),
			}
			return flash.WithError(c, mp).Redirect(redirectTo)
		}
		c.Locals("login_response", loginResponse)
		return c.Next()
	}
}

func Login(c *fiber.Ctx, rememberMe bool) error {
	sessionStarted := time.Now().Format(`2006-01-02 15:04:05`)
	id, err := session.ID(c)
	if err != nil {
		fmt.Println("Unable to get session id")
		return err
	}
	loginResponse := c.Locals("login_response").(*auth.LoginResponse)
	userWithProfile := &models.UserWithProfile{
		User:    loginResponse.User,
		Profile: loginResponse.Profile,
		Domain:  mail.GetDomainOfEmail(loginResponse.User.Email),
		Account: loginResponse.Account,
	}

	values := fiber.Map{
		"user_profile": userWithProfile,
	}
	if rememberMe {
		err = session.SetKeys(c, values, session.RememberMeExpiry)
	} else {
		err = session.SetKeys(c, values)
	}
	if err != nil {
		fmt.Println("Unable to set keys")
		fmt.Println(values)
		return err
	}
	auth.LoggedInBucket.Add(id, userWithProfile)
	db.DB.Model(&models.LoginSession{}).Where("k = ?", id).UpdateColumns(map[string]interface{}{
		"user_id":     loginResponse.User.ID,
		"last_used":   time.Now().Format(`2006-01-02 15:04:05`),
		"created_at":  sessionStarted,
		"remember_me": rememberMe,
		"is_active":   true,
	})
	return nil
}

func Auth(redirectTo string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := session.ID(c)
		if err != nil {
			return c.Redirect(redirectTo)
		}
		userWithProfile := auth.LoggedInBucket.Get(c, id)
		if userWithProfile == nil {
			Logout(c)
			return c.Redirect(redirectTo)
		}
		c.Locals("user_id", userWithProfile.User.ID)
		c.Locals("email", userWithProfile.User.Email)
		c.Locals("domain", userWithProfile.Domain)
		c.Locals("account", userWithProfile.Account)
		return c.Next()
	}
}

func IsLoggedIn(c *fiber.Ctx) bool {
	userID := c.Locals("user_id")
	if userID != nil {
		return true
	}
	userID, _ = session.Get(c, "user_id")
	if userID == nil {
		return false
	}
	c.Locals("user_id", userID)
	return true
}

func AuthAPI(c *fiber.Ctx) error {
	authorization := c.Get(fiber.HeaderAuthorization)
	key := strings.ReplaceAll(authorization, "Bearer ", "")
	validAPIKey := crypt.ValidateApiKey(key)
	if validAPIKey.Error != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": validAPIKey.Error.Error(),
			"error":   true,
		})
	}
	if !validAPIKey.Valid {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid Attempt",
			"error":   true,
		})
	}
	uID, _ := strconv.Atoi(validAPIKey.UserID)
	c.Locals("user_id", uint(uID))
	c.Locals("is_api", true)
	tmp, err := auth.APILoggedInBucket.Get(validAPIKey.UserID)
	if err != nil {
		return c.Status(406).JSON(fiber.Map{
			"message": "Unable to process your request",
			"error":   true,
		})
	}
	if tmp == nil {
		userWithProfile, err := auth.GetVerifiedUserByID(validAPIKey.UserID)
		if err != nil {
			return c.Status(406).JSON(fiber.Map{
				"message": "Unable to process your request",
				"error":   true,
			})
		}

		bt, err := json.Marshal(userWithProfile)
		if err != nil {
			return c.Status(406).JSON(fiber.Map{
				"message": "Unable to process your request",
				"error":   true,
			})
		}
		err = auth.APILoggedInBucket.Set(validAPIKey.UserID, bt, time.Minute)
		if err != nil {
			return c.Status(406).JSON(fiber.Map{
				"message": "Unable to process your request",
				"error":   true,
			})
		}
	}
	return c.Next()
}
