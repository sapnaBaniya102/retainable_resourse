package crypt

import "github.com/alexedwards/argon2id"

func CreateHash(password string) (hash string, err error) {
	return argon2id.CreateHash(password, argon2id.DefaultParams)
}
