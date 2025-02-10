package utils

import (
	"crypto/rand"
	"fmt"
	"github.com/krack8/lighthouse/pkg/auth/dto"
	"golang.org/x/crypto/bcrypt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"os"
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

// Helper function to generate reset token
func GenerateResetToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", b)
}

func CreateAgentSecret(key dto.AgentKey) (v1.Secret, error) {

	directMap := map[string][]byte{
		"AUTH_TOKEN": key.AuthToken,
	}
	kubeSecretInfo := v1.Secret{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      os.Getenv("AGENT_SECRET_NAME"),
			Namespace: os.Getenv("AGENT_SECRET_NAMESPACE"),
		},
		Data: directMap,
		Type: "Opaque",
	}

	return kubeSecretInfo, nil
}
