package utils

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/krack8/lighthouse/pkg/config"
	"golang.org/x/crypto/bcrypt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func HashPassword(password string) string {
	// Check password length before hashing
	if len(password) == 0 {
		log.Printf("Error: Empty password")
		return ""
	}
	if len(password) > 72 {
		log.Printf("Error: Password length exceeds 72 bytes")
		return ""
	}

	// Generate a bcrypt hash of the password with a default cost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error generating hash: %v", err)
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

// Helper function to generate reset token
func GenerateResetToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", b)
}

func GetOrCreateSecret(name, namespace string, authToken string) (string, error) {
	// Prepare secret data (base64 encoding the authToken)
	encodedToken := base64.StdEncoding.EncodeToString([]byte(authToken))
	secretData := map[string][]byte{
		"AUTH_TOKEN": []byte(encodedToken), // Store the base64 encoded token
	}

	// Try to retrieve the secret
	secret, err := config.GetKubeClientSet().CoreV1().Secrets(namespace).Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		// Check if the error is a "not found" error
		if errors.IsNotFound(err) {
			// Secret does not exist, create it
			log.Printf("Secret %s not found in namespace %s. Creating a new one...", name, namespace)

			secret = &v1.Secret{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Secret",
					APIVersion: "v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: namespace,
				},
				Data: secretData,
				Type: v1.SecretTypeOpaque,
			}

			// Create the secret in Kubernetes
			secret, err = config.GetKubeClientSet().CoreV1().Secrets(namespace).Create(context.Background(), secret, metav1.CreateOptions{})
			if err != nil {
				return "", fmt.Errorf("failed to create secret: %w", err)
			}

			log.Printf("Secret %s successfully created in namespace %s.", name, namespace)
			// Return the raw token
			return authToken, nil
		}

		// For other errors, return the error
		return "", fmt.Errorf("failed to fetch secret: %w", err)
	}

	// If the secret exists, get the encoded token (base64 encoded)
	encodedTokenFromSecret, exists := secret.Data["AUTH_TOKEN"]
	if !exists {
		return "", fmt.Errorf("key 'AUTH_TOKEN' not found in the secret")
	}

	// Decode the URL-safe base64 encoded token
	rawToken, err := base64.URLEncoding.DecodeString(string(encodedTokenFromSecret))
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 value for AUTH_TOKEN: %w", err)
	}

	// Return the raw token as a string
	return string(rawToken), nil
}
