package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password string) string {
	// Generate a bcrypt hash of the password with a default cost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error : %v", err)
		return ""
	}
	return string(hash)
}

// CheckPassword compares a plain text password with a hashed password.
func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil // Return true if passwords match, false otherwise
}

func GenerateSecureToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		_ = fmt.Errorf("failed to generate secure token: %w", err)
		return ""
	}
	return hex.EncodeToString(bytes)
}
