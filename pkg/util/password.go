package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPassword[T comparable](providedPassword, password T) error {
	providedPasswordStr := fmt.Sprintf("%v", providedPassword)
	passwordStr := fmt.Sprintf("%v", password)

	err := bcrypt.CompareHashAndPassword([]byte(passwordStr), []byte(providedPasswordStr))
	if err != nil {
		return err
	}
	return nil
}
