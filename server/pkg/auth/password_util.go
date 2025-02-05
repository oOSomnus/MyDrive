package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrHashPassword = "failed to hash password"
)

func HashPassword(password string) (string, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), errors.New(ErrHashPassword)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
