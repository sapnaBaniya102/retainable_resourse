package models

import (
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/session"
)

type UserWithProfile struct {
	User    *Users  `json:"user"`
	Profile Profile `json:"profile"`
	Domain  string  `json:"domain"`
	Account string  `json:"account"`
}

type LoginBucket struct {
	UserBucket map[string]*UserWithProfile
	mu         sync.RWMutex
}

func (l *LoginBucket) Add(sessionID string, userWithProfile *UserWithProfile) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.UserBucket[sessionID] = userWithProfile
}
func (l *LoginBucket) Get(c *fiber.Ctx, sessionID string) *UserWithProfile {
	if userWithProfile, ok := l.UserBucket[sessionID]; ok {
		return userWithProfile
	}
	loginResponse, _ := session.Get(c, "user_profile")
	if loginResponse == nil {
		return nil
	}
	userWithProfile := loginResponse.(UserWithProfile)
	l.Add(sessionID, &userWithProfile)
	return &userWithProfile
}
func (l *LoginBucket) Remove(sessionID string) {
	if _, ok := l.UserBucket[sessionID]; ok {
		l.mu.Lock()
		defer l.mu.Unlock()
		delete(l.UserBucket, sessionID)
	}
}
