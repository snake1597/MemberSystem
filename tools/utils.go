package tools

import (
	"golang.org/x/crypto/bcrypt"
)

func HashString(s string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 12)
	return string(bytes), err
}

func VerifyHashString(s, hash string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(s))
	return err
}
