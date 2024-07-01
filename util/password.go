package util

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	encodedHash := base64.StdEncoding.EncodeToString(hashedPassword)
	return encodedHash, nil
}

func CheckPassword(password string, hashPassword string) error {
	decodedHash, err := base64.StdEncoding.DecodeString(hashPassword)
	if err != nil {
		return fmt.Errorf("failed to decode hash: %w", err)
	}
	return bcrypt.CompareHashAndPassword(decodedHash, []byte(password))
}
