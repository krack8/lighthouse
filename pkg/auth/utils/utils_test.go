package utils

import "testing"

func TestHashPassword(t *testing.T) {
	password := "password123"
	hashedPassword := HashPassword(password)

	if hashedPassword == "" {
		t.Error("HashPassword returned an empty string")
	}
}

func TestCheckPassword(t *testing.T) {
	password := "password123"
	hashedPassword := HashPassword(password)

	if !CheckPassword(password, hashedPassword) {
		t.Error("CheckPassword failed to match the password")
	}
}
