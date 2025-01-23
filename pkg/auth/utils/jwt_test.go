package utils

import (
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	secret := "test-secret"
	username := "testuser"
	expiry := 15 * time.Minute

	token, err := GenerateToken(username, secret, expiry)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	if token == "" {
		t.Error("Generated token is empty")
	}
}

func TestValidateToken(t *testing.T) {
	secret := "test-secret"
	username := "testuser"
	expiry := 15 * time.Minute

	token, _ := GenerateToken(username, secret, expiry)
	claims, err := ValidateToken(token, secret)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	if claims.Username != username {
		t.Errorf("Expected username %s, got %s", username, claims.Username)
	}
}
