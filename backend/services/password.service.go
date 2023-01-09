package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", errors.New("could not hash password")
	}

	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, candidatePassword string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}
