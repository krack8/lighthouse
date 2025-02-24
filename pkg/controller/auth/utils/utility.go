package utils

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/k8s"
	"github.com/krack8/lighthouse/pkg/common/log"
	"golang.org/x/crypto/bcrypt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
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

func CreateNamespaceIfNotExists(namespace string) error {
	clientSet := k8s.GetKubeClientSet()
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

// GetSecret retrieves and decodes a secret
func GetSecret(name, namespace string) (string, error) {
	if name == "" || namespace == "" {
		return "", fmt.Errorf("missing or invalid name or namespace")
	}

	secret, err := k8s.GetKubeClientSet().CoreV1().Secrets(namespace).Get(
		context.Background(),
		name,
		metav1.GetOptions{},
	)
	if err != nil {
		if errors.IsNotFound(err) {
			return "", fmt.Errorf("secret %s not found in namespace %s", name, namespace)
		}
		return "", fmt.Errorf("failed to fetch secret: %w", err)
	}

	// Get token directly from secret.Data (already decoded)
	tokenBytes, exists := secret.Data["AUTH_TOKEN"]
	if !exists {
		return "", fmt.Errorf("key 'AUTH_TOKEN' not found in the secret")
	}

	return string(tokenBytes), nil
}

// CreateOrUpdateSecret creates or updates a secret with the given auth token
func CreateOrUpdateSecret(name, namespace, authToken, clusterId string) (string, error) {
	// Prepare secret data (no need to base64 encode, Kubernetes will do it)
	secretData := map[string][]byte{
		"AUTH_TOKEN":  []byte(authToken),
		"AGENT_GROUP": []byte(clusterId),
	}

	clientSet := k8s.GetKubeClientSet()

	// Create namespace if it doesn't exist
	//err := CreateNamespaceIfNotExists(namespace)
	//if err != nil {
	//	return "", err
	//}

	secretClient := clientSet.CoreV1().Secrets(namespace)

	// Check if secret exists
	_, err := secretClient.Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			// Create new secret
			log.Logger.Infow(fmt.Sprintf("secret %s not found in namespace %s. creating a new one...", name, namespace),
				"info", "secret-create")

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

			_, err = secretClient.Create(context.Background(), secret, metav1.CreateOptions{})
			if err != nil {
				return "", fmt.Errorf("failed to create secret: %w", err)
			}

			log.Logger.Infow(fmt.Sprintf("secret %s successfully created in namespace %s", name, namespace),
				"info", "secret-create")
			return authToken, nil
		}
		return "", fmt.Errorf("failed to fetch secret: %w", err)
	}

	if config.RunMode != "PRODUCTION" {
		// Update existing secret
		log.Logger.Infow(fmt.Sprintf("Secret %s exists in namespace %s. Updating it...", name, namespace),
			"info", "secret-update")

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

		log.Logger.Infow(fmt.Sprintf("secret %s successfully updated in namespace %s", name, namespace),
			"info", "secret-update")
	}

	return authToken, nil
}

// Helper function to get agent group from environment or secret
func GetAgentGroup(secretName, namespace string) (string, error) {
	// First try environment variable
	groupName := os.Getenv("AGENT_GROUP")
	if groupName != "" {
		return groupName, nil
	}

	// If not in environment, try getting from secret
	if secretName != "" && namespace != "" {
		// Get the Kubernetes clientset
		clientSet := k8s.GetKubeClientSet()

		// Try to get the secret
		secret, err := clientSet.CoreV1().Secrets(namespace).Get(context.Background(), secretName, metav1.GetOptions{})
		if err != nil {
			return "", fmt.Errorf("failed to get secret %s in namespace %s: %w", secretName, namespace, err)
		}

		// Look for worker group in the secret
		if groupData, exists := secret.Data["AGENT_GROUP"]; exists && len(groupData) > 0 {
			return string(groupData), nil
		}
	} else {
		return "", fmt.Errorf("missing or invalid name or namespace")
	}
	return "", fmt.Errorf("AGENT_GROUP key not found in secret %s", secretName)
}
