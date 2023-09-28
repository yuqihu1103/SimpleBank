package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// returns the password's bcrypt hash
func HashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash the password: %w", err)
	}

	return string(hashedPassword), nil
}

// check if provided password matches the hashed password
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
