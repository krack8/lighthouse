package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"strings"
	"time"
)

type Crypto interface {
	GenerateSecureToken(length int) (string, error)
	CreateCombinedToken(rawToken string, clusterID primitive.ObjectID) (string, error)
	ParseCombinedToken(combinedToken string) (primitive.ObjectID, string, error)
}

type CryptoImpl struct {
	encryptionKey []byte
	signingKey    []byte
}

// Key generation function
func generateKey(keyName string, length int) ([]byte, error) {
	// Check if the key is already set as an environment variable
	key := os.Getenv(keyName)
	if key != "" {
		// If the key is already set, decode it from base64 to []byte
		decodedKey, err := base64.URLEncoding.DecodeString(key)
		if err != nil {
			return nil, fmt.Errorf("failed to decode key: %w", err)
		}
		return decodedKey, nil
	}

	// Otherwise, generate a new random key
	newKey := make([]byte, length)
	if _, err := rand.Read(newKey); err != nil {
		return nil, fmt.Errorf("failed to generate key: %w", err)
	}

	// Store the new key as an environment variable (encoded in base64)
	err := os.Setenv(keyName, base64.URLEncoding.EncodeToString(newKey))
	if err != nil {
		return nil, fmt.Errorf("failed to set environment variable: %w", err)
	}

	return newKey, nil
}

// Function to generate a random 32-byte encryption key and a signing key
func NewCryptoImpl() (*CryptoImpl, error) {
	// Generate encryption and signing keys if not already present
	encryptionKey, err := generateKey("ENCRYPTION_KEY", 32) // 32 bytes for AES-256
	if err != nil {
		return nil, err
	}

	signingKey, err := generateKey("SIGNING_KEY", 32) // 32 bytes for HMAC SHA-256
	if err != nil {
		return nil, err
	}

	return &CryptoImpl{
		encryptionKey: encryptionKey,
		signingKey:    signingKey,
	}, nil
}

// GenerateSecureToken generates a random token of specified length
func (c *CryptoImpl) GenerateSecureToken(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}
	// Generate timestamp
	timestamp := time.Now().Unix()

	// Combine token and timestamp
	tokenData := fmt.Sprintf("%s|%d", base64.URLEncoding.EncodeToString(bytes), timestamp)
	return tokenData, nil
}

// CreateCombinedToken creates a combined token format: encryptedToken.clusterID.signature
func (c *CryptoImpl) CreateCombinedToken(rawToken string, clusterID primitive.ObjectID) (string, error) {
	// Encrypt the raw token
	encrypted, err := c.encrypt([]byte(rawToken))
	if err != nil {
		return "", fmt.Errorf("encryption error: %w", err)
	}

	// Convert to base64
	encryptedBase64 := base64.URLEncoding.EncodeToString(encrypted)
	clusterIDHex := clusterID.Hex()

	// Create payload for signature
	payload := fmt.Sprintf("%s.%s", encryptedBase64, clusterIDHex)

	// Generate signature
	signature := c.generateSignature([]byte(payload))

	// Combine all parts
	return fmt.Sprintf("%s.%s.%s", encryptedBase64, clusterIDHex, signature), nil
}

// ParseCombinedToken parses and validates a combined token
func (c *CryptoImpl) ParseCombinedToken(combinedToken string) (primitive.ObjectID, string, error) {
	parts := strings.Split(combinedToken, ".")
	if len(parts) != 3 {
		return primitive.NilObjectID, "", fmt.Errorf("invalid token format")
	}

	encryptedBase64, clusterIDHex, signature := parts[0], parts[1], parts[2]

	// Verify signature
	payload := fmt.Sprintf("%s.%s", encryptedBase64, clusterIDHex)
	expectedSignature := c.generateSignature([]byte(payload))

	if signature != expectedSignature {
		return primitive.NilObjectID, "", fmt.Errorf("invalid signature")
	}

	// Convert clusterID
	clusterID, err := primitive.ObjectIDFromHex(clusterIDHex)
	if err != nil {
		return primitive.NilObjectID, "", fmt.Errorf("invalid cluster ID: %w", err)
	}

	// Decrypt token
	encryptedBytes, err := base64.URLEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return primitive.NilObjectID, "", fmt.Errorf("invalid encrypted token: %w", err)
	}

	decrypted, err := c.decrypt(encryptedBytes)
	if err != nil {
		return primitive.NilObjectID, "", fmt.Errorf("decryption error: %w", err)
	}

	return clusterID, string(decrypted), nil
}

// Helper functions for encryption/decryption
func (c *CryptoImpl) encrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.encryptionKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}

func (c *CryptoImpl) decrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.encryptionKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

// signs token using HMAC
func (c *CryptoImpl) generateSignature(data []byte) string {
	h := hmac.New(sha256.New, c.signingKey)
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
