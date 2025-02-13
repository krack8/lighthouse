package utils

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/krack8/lighthouse/pkg/config"
	"github.com/krack8/lighthouse/pkg/log"
	"golang.org/x/crypto/bcrypt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func HashPassword(password string) string {
	// Check password length before hashing
	if len(password) == 0 {
		log.Logger.Errorw("empty password", "err", "password")
		return ""
	}
	if len(password) > 72 {
		log.Logger.Errorw("password length exceeds 72 bytes", "err", "password")
		return ""
	}

	// Generate a bcrypt hash of the password with a default cost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Logger.Errorw("error generating hash", "err", err.Error())
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
		log.Logger.Errorw("failed to generate secure token", "err", err.Error())
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

func GetSecret(name, namespace string) (string, error) {
	// Try to retrieve the secret
	secret, err := config.GetKubeClientSet().CoreV1().Secrets(namespace).Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return "", fmt.Errorf("secret %s not found in namespace %s", name, namespace)
		}
		return "", fmt.Errorf("failed to fetch secret: %w", err)
	}

	// Extract the encoded token
	encodedToken, exists := secret.Data["AUTH_TOKEN"]
	if !exists {
		return "", fmt.Errorf("key 'AUTH_TOKEN' not found in the secret")
	}

	// Decode the base64-encoded token
	rawToken, err := base64.StdEncoding.DecodeString(string(encodedToken))
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 value for AUTH_TOKEN: %w", err)
	}

	return string(rawToken), nil
}

func CreateNamespaceIfNotExists(namespace string) error {
	clientSet := config.GetKubeClientSet()
	namespaceClient := clientSet.CoreV1().Namespaces()
	_, err := namespaceClient.Get(context.Background(), namespace, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			// Secret does not exist, create it
			log.Logger.Infow("namespace "+namespace+" not found. creating a new one...", "info", "namespace-create")

			ns := &v1.Namespace{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Namespace",
					APIVersion: "v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: namespace,
				},
			}

			_, err = namespaceClient.Create(context.Background(), ns, metav1.CreateOptions{})
			if err != nil {
				return fmt.Errorf("failed to create namespace: %w", err)
			}
			log.Logger.Infow("namespace "+namespace+" successfully created.", "info", "namespace-create")
			return nil
		}
		return fmt.Errorf("failed to fetch namespace: %w", err)
	}
	return nil
}

func CreateOrUpdateSecret(name, namespace, authToken string) (string, error) {
	// Prepare secret data (base64 encoding the authToken)
	encodedToken := base64.StdEncoding.EncodeToString([]byte(authToken))
	secretData := map[string][]byte{
		"AUTH_TOKEN": []byte(encodedToken), // Store the base64 encoded token
	}

	clientSet := config.GetKubeClientSet()
	err := CreateNamespaceIfNotExists(namespace)
	if err != nil {
		return "", err
	}
	secretClient := clientSet.CoreV1().Secrets(namespace)

	// Check if the secret exists
	_, err = secretClient.Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			// Secret does not exist, create it
			log.Logger.Infow("secret "+name+" not found in namespace "+namespace+". creating a new one...", "info", "secret-create")

			secret := &v1.Secret{
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

			secret, err = secretClient.Create(context.Background(), secret, metav1.CreateOptions{})
			if err != nil {
				return "", fmt.Errorf("failed to create secret: %w", err)
			}

			log.Logger.Infow("secret "+name+" successfully created in namespace "+namespace+".", "info", "secret-create")
			return authToken, nil
		}
		return "", fmt.Errorf("failed to fetch secret: %w", err)
	}

	// If the secret exists, update it with the new token
	log.Logger.Infow("Secret "+name+" exists in namespace "+namespace+". Updating it...", "info", "secret-update")

	_, err = secretClient.Update(context.Background(), &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Data: secretData,
		Type: v1.SecretTypeOpaque,
	}, metav1.UpdateOptions{})

	if err != nil {
		return "", fmt.Errorf("failed to update secret: %w", err)
	}

	log.Logger.Infow("secret "+name+" successfully updated in namespace"+namespace+".", "info", "secret-update")
	return authToken, nil
}
