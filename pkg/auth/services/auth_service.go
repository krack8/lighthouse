package services

import (
	"errors"
	"github.com/krack8/lighthouse/pkg/auth/models"
	"github.com/krack8/lighthouse/pkg/auth/utils"
	"log"
	"os"
	"time"
)

func Login(username string, password string) (string, string, error) {
	var user *models.User
	user, err := GetUserByUsername(username)
	if err != nil {
		return "", "", errors.New("user not found")
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", "", errors.New("invalid credentials")
	}

	// Load expiry durations from environment variables
	accessTokenExpiry, err := parseDurationFromEnv("ACCESS_TOKEN_EXPIRY")
	if err != nil {
		return "", "", err
	}

	refreshTokenExpiry, err := parseDurationFromEnv("REFRESH_TOKEN_EXPIRY")
	if err != nil {
		return "", "", err
	}

	// Generate JWT tokens with the username and permissions
	accessToken, err := utils.GenerateToken(username, os.Getenv("JWT_SECRET"), accessTokenExpiry)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.GenerateToken(username, os.Getenv("JWT_REFRESH_SECRET"), refreshTokenExpiry)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// Helper function to parse durations from environment variables
func parseDurationFromEnv(envKey string) (time.Duration, error) {
	value := os.Getenv(envKey)
	if value == "" {
		return 0, errors.New(envKey + " is not set in environment variables")
	}

	parsed, err := time.ParseDuration(value)
	if err != nil {
		return 0, errors.New("invalid duration format for " + envKey + ": " + value)
	}

	return parsed, nil
}

func IsAgentAuthTokenValid(authToken string) bool {
	cluster, err := ValidateAgentClusterToken(authToken)
	if cluster == nil && err != nil {
		log.Println("[Error]", err)
		return false
	}
	return true
}
