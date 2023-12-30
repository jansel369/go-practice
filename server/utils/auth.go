package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(rawPassword string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(rawPassword),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

func ComparePasswords(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(plainPassword),
	)

	return err == nil
}
